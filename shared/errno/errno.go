package errno

import (
	"errors"
	"fmt"

	"github.com/geiqin/gaiakit/kitex_gen/errno"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// NewErrNo return ErrNo
func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(int64(errno.Err_Success), "success")
	NoRoute                = NewErrNo(int64(errno.Err_NoRoute), "no route")
	NoMethod               = NewErrNo(int64(errno.Err_NoMethod), "no method")
	TokenErr               = NewErrNo(int64(errno.Err_TokenErr), "token error")
	BadRequest             = NewErrNo(int64(errno.Err_BadRequest), "bad request")
	ParamsErr              = NewErrNo(int64(errno.Err_ParamsErr), "params error")
	AuthorizeFail          = NewErrNo(int64(errno.Err_AuthorizeFail), "authorize failed")
	TooManyRequest         = NewErrNo(int64(errno.Err_TooManyRequest), "too many requests")
	ServiceErr             = NewErrNo(int64(errno.Err_ServiceErr), "service error")
	RecordNotFound         = NewErrNo(int64(errno.Err_RecordNotFound), "record not found")
	RecordAlreadyExist     = NewErrNo(int64(errno.Err_RecordAlreadyExist), "record already exist")
	DirtyData              = NewErrNo(int64(errno.Err_DirtyData), "dirty data")
	ApiServiceErr          = NewErrNo(int64(errno.Err_ApiServiceErr), "api service error")
	GlobalSrvClientErr     = NewErrNo(int64(errno.Err_GlobalSrvClientErr), "global srv client error")
	ComTenantSrvErr        = NewErrNo(int64(errno.Err_ComTenantSrvErr), "com.tenant.srv error")
	ComTenantSrvClientErr  = NewErrNo(int64(errno.Err_ComTenantSrvClientErr), "com.tenant.srv client error")
	ComSettingSrvErr       = NewErrNo(int64(errno.Err_ComSettingSrvErr), "com.setting.srv error")
	ComSettingSrvClientErr = NewErrNo(int64(errno.Err_ComSettingSrvClientErr), "com.setting.srv client error")
	SysPowerSrvErr         = NewErrNo(int64(errno.Err_SysPowerSrvErr), "sys.power.srv error")
	SysPowerSrvClientErr   = NewErrNo(int64(errno.Err_SysPowerSrvClientErr), "sys.power.srv client error")
	SysBlobSrvErr          = NewErrNo(int64(errno.Err_SysBlobSrvErr), "sys.blob.srv error")
	SysBlobSrvClientErr    = NewErrNo(int64(errno.Err_SysBlobSrvClientErr), "sys.blob.srv client error")
	SysSiteSrvErr          = NewErrNo(int64(errno.Err_SysSiteSrvErr), "sys.site.srv error")
	SysSiteSrvClientErr    = NewErrNo(int64(errno.Err_SysSiteSrvClientErr), "sys.site.srv client error")
	BuyCrmSrvErr           = NewErrNo(int64(errno.Err_BuyCrmSrvErr), "buy.crm.srv error")
	BuyCrmSrvClientErr     = NewErrNo(int64(errno.Err_BuyCrmSrvClientErr), "buy.crm.srv client error")
	BuyOmsSrvErr           = NewErrNo(int64(errno.Err_BuyOmsSrvErr), "buy.oms.srv error")
	BuyOmsSrvClientErr     = NewErrNo(int64(errno.Err_BuyOmsSrvClientErr), "buy.oms.srv client error")
	BuyPaySrvErr           = NewErrNo(int64(errno.Err_BuyPaySrvErr), "buy.pay.srv error")
	BuyPaySrvClientErr     = NewErrNo(int64(errno.Err_BuyPaySrvClientErr), "buy.pay.srv client error")
	BuyPdmSrvErr           = NewErrNo(int64(errno.Err_BuyPdmSrvErr), "buy.pdm.srv error")
	BuyPdmSrvClientErr     = NewErrNo(int64(errno.Err_BuyPdmSrvClientErr), "buy.pdm.srv client error")
	KitAigcSrvErr          = NewErrNo(int64(errno.Err_KitAigcSrvErr), "kit.aigc.srv error")
	KitAigcSrvClientErr    = NewErrNo(int64(errno.Err_KitAigcSrvClientErr), "kit.aigc.srv client error")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
