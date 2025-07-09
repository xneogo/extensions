/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2025/4/27 -- 14:33
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xerror xerror/xerror.go
*/

package xerror

type ErrCode int32

type XError struct {
	err   error
	inner *XError
	msg   *string
	code  ErrCode
}

func New(code int, msg string) *XError {
	return &XError{
		code: ErrCode(code),
		msg:  &msg,
	}
}

func (x *XError) Error() string {
	if x == nil || x.err.Error() == "" {
		return "<nil>"
	}

	if x.msg != nil {
		return *x.msg
	}

	if x.chain() {
		return x.inner.Error()
	}

	return x.err.Error()
}

func (x *XError) chain() bool {
	return x.inner != nil
}

func (x *XError) Unwrap() error {
	return x.inner
}

func (x *XError) Wrap(code int, msg string) *XError {
	return &XError{
		inner: New(code, msg),
	}
}

func (x *XError) WrapErr(err error) *XError {
	return &XError{
		inner: New(-1, err.Error()),
	}
}
