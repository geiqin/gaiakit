package auth

import (
	"context"
	"fmt"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/geiqin/gaiakit/shared/consts"
	"github.com/geiqin/gaiakit/shared/utils"
)

//获得当前店铺ID
func GetStoreId(ctx context.Context) int64 {
	val, ok := metainfo.GetValue(ctx, consts.StoreID)
	if ok {
		v := utils.StringToInt64(val)
		return v
	}
	return 0
}

//获得登录账号ID
func GetAccountId(ctx context.Context) int64 {
	val, ok := metainfo.GetValue(ctx, consts.AccountID)
	if ok {
		v := utils.StringToInt64(val)
		return v
	}
	return 0
}

//获得登录账号类型
func GetAccountType(ctx context.Context) string {
	val, ok := metainfo.GetValue(ctx, consts.AccountType)
	if ok {
		return val
	}
	return ""
}

func IsMaster(ctx context.Context) bool {
	storeId := GetStoreId(ctx)
	if storeId == 1 {
		return true
	}
	return false
}

func IsStore(ctx context.Context) bool {
	storeId := GetStoreId(ctx)
	if storeId > 1 {
		return true
	}
	return false
}

func IsMember(ctx context.Context) bool {
	accountType := GetAccountType(ctx)
	if accountType == consts.MasterMember || accountType == consts.StoreMember {
		return true
	}
	return false
}

func GetDbName(storeId int64) string {
	flag := fmt.Sprintf("%08d", storeId)
	if storeId == 1 {
		return "gaia_master"
	}
	return "gaia_store_" + flag
}
