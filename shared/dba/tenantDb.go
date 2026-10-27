package dba

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/geiqin/gaiakit/shared/auth"
	"github.com/geiqin/gaiakit/shared/consts"
	"github.com/geiqin/gaiakit/shared/info"
	"github.com/geiqin/gaiakit/shared/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TenantDBManager struct {
	mu    sync.RWMutex
	cache map[string]*tenantDBItem
	// 配置
	maxIdle     int
	maxOpen     int
	lifeTime    time.Duration
	idleTimeout time.Duration // 多久没访问就关闭池
	mySqlInfo   info.MysqlConfig
}

type tenantDBItem struct {
	db        *gorm.DB
	sqlDB     *sql.DB
	lastVisit time.Time
}

func NewTenantDBManager(maxIdle, maxOpen int, lifeTime time.Duration, idleTimeout time.Duration, mySqlInfo info.MysqlConfig) *TenantDBManager {
	m := &TenantDBManager{
		cache:       make(map[string]*tenantDBItem),
		maxIdle:     maxIdle,
		maxOpen:     maxOpen,
		lifeTime:    lifeTime,
		idleTimeout: idleTimeout,
		mySqlInfo:   mySqlInfo,
	}
	// 后台定时清理过期租户连接池
	go m.cleanLoop()
	return m
}

// GetDB 从ctx拿到tenantId，返回租户db
func (m *TenantDBManager) GetDB(ctx context.Context) (*gorm.DB, error) {
	tenantID, ok := metainfo.GetValue(ctx, consts.StoreId)
	var storeId int64
	if !ok || tenantID == "" {
		tenantID = "zero"
	} else {
		storeId = utils.StringToInt64(tenantID)
	}

	m.mu.RLock()
	item, ok := m.cache[tenantID]
	m.mu.RUnlock()
	if ok {
		m.mu.Lock()
		item.lastVisit = time.Now()
		m.mu.Unlock()
		return item.db, nil
	}

	// 懒加载：根据tenantId生成dsn，你可以从配置中心/管理库读取dsn
	c := m.mySqlInfo
	dbname := auth.GetDbName(storeId)
	dsn := fmt.Sprintf(consts.MySqlDSN, c.User, c.Password, c.Host, c.Port, dbname)

	// 方式B【生产推荐】：不要硬拼DSN！从「租户元信息库」查询这个租户的db配置
	// tenantMeta, err := queryTenantMetaFromMasterDB(tenantID)
	// dsn = tenantMeta.DSN

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true, // 关闭默认事务，提升性能
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(m.maxIdle)
	sqlDB.SetMaxOpenConns(m.maxOpen)
	sqlDB.SetConnMaxLifetime(m.lifeTime)

	newItem := &tenantDBItem{
		db:        db,
		sqlDB:     sqlDB,
		lastVisit: time.Now(),
	}
	m.mu.Lock()
	m.cache[tenantID] = newItem
	m.mu.Unlock()
	return db, nil
}

func (m *TenantDBManager) cleanLoop() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		m.mu.Lock()
		for tid, item := range m.cache {
			if time.Since(item.lastVisit) > m.idleTimeout {
				_ = item.sqlDB.Close() // 释放连接池
				delete(m.cache, tid)
			}
		}
		m.mu.Unlock()
	}
}
