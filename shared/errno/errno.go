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
	Success            = NewErrNo(int64(errno.Err_Success), "success")
	NoRoute            = NewErrNo(int64(errno.Err_NoRoute), "no route")
	NoMethod           = NewErrNo(int64(errno.Err_NoMethod), "no method")
	BadRequest         = NewErrNo(int64(errno.Err_BadRequest), "bad request")
	ParamsErr          = NewErrNo(int64(errno.Err_ParamsErr), "params error")
	AuthorizeFail      = NewErrNo(int64(errno.Err_AuthorizeFail), "authorize failed")
	TooManyRequest     = NewErrNo(int64(errno.Err_TooManyRequest), "too many requests")
	ServiceErr         = NewErrNo(int64(errno.Err_ServiceErr), "service error")
	RPCTenantSrvErr    = NewErrNo(int64(errno.Err_RPCTenantSrvErr), "rpc user service error")
	TenantSrvErr       = NewErrNo(int64(errno.Err_TenantSrvErr), "user service error")
	RPCBlobSrvErr      = NewErrNo(int64(errno.Err_RPCBlobSrvErr), "rpc blob service error")
	BlobSrvErr         = NewErrNo(int64(errno.Err_BlobSrvErr), "blob service error")
	RPCPowerSrvErr     = NewErrNo(int64(errno.Err_RPCPowerSrvErr), "rpc car service error")
	PowerSrvErr        = NewErrNo(int64(errno.Err_PowerSrvErr), "car service error")
	RPCSettingSrvErr   = NewErrNo(int64(errno.Err_RPCSettingSrvErr), "rpc profile service error")
	SettingSrvErr      = NewErrNo(int64(errno.Err_SettingSrvErr), "profile service error")
	RPCOmsSrvErr       = NewErrNo(int64(errno.Err_RPCOmsSrvErr), "rpc trip service error")
	OmsSrvErr          = NewErrNo(int64(errno.Err_OmsSrvErr), "trip service error")
	RecordNotFound     = NewErrNo(int64(errno.Err_RecordNotFound), "record not found")
	RecordAlreadyExist = NewErrNo(int64(errno.Err_RecordAlreadyExist), "record already exist")
	DirtyData          = NewErrNo(int64(errno.Err_DirtyData), "dirty data")
	TradeSrcErr        = NewErrNo(int64(errno.Err_TradeSrvErr), "rpc trade service error")
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
