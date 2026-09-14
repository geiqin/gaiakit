package auth

import (
	"context"

	"github.com/geiqin/gaiakit/shared/consts"
	"github.com/geiqin/gaiakit/shared/utils"
)

//获得当前店铺ID
func GetStoreId(ctx context.Context) int64 {
	val := ctx.Value(consts.AccountID)
	if val != nil {
		v := utils.StringToInt64(utils.ToString(val))
		return v
	}
	return 0
}

//获得登录账号ID
func GetAccountId(ctx context.Context) int64 {
	val := ctx.Value(consts.AccountID)
	if val != nil {
		v := utils.StringToInt64(utils.ToString(val))
		return v
	}
	return 0
}

//获得登录账号类型
func GetAccountType(ctx context.Context) string {
	val := ctx.Value(consts.AccountType)
	if val != nil {
		v := utils.ToString(val)
		return v
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
