package constants

const (
	// ErrSucceed - Succeed
	ErrSucceed = 0

	// ErrPermission - Permission Denied
	ErrPermission = 401
	ErrForbidden  = 438

	// ErrToken - Invalid Token
	ErrToken = 420

	// ErrInvalidParam - Invalid Parameter
	ErrInvalidParam = 421

	// ErrAccount - No This User or Password Error
	ErrAccount = 422

	// ErrSubNats - Subscribe to nats error
	ErrSubNats = 423

	// ErrInternalServerError - Internal error.
	ErrInternalServerError = 500

	// ErrWechatPay - Wechat Pay error.
	ErrWechatPay = 520

	// ErrWechatAuth - Wechat Auth error.
	ErrWechatAuth = 521

	// ErrMongoDB - MongoDB operations error.
	ErrMongoDB = 600

	// ErrMysql - Mysql operations error.
	ErrMysql = 700
)

var errText = map[int]string{
	ErrSucceed: "Success",

	ErrPermission: "Without Permission",
	ErrForbidden:  "Forbidden",

	ErrToken: "Invalid Token",

	ErrInvalidParam: "Invalide Parameter",

	ErrAccount: "Incorrect Account or Password",

	ErrSubNats: "Subscribe to nats error",

	ErrInternalServerError: "Internal error",

	ErrWechatPay: "Wechat Pay error",

	ErrWechatAuth: "Wechat Auth error",

	ErrMongoDB: "MongoDB operations error",

	ErrMysql: "Mysql operations error",
}

func ErrorText(code int) string {
	return errText[code]
}
