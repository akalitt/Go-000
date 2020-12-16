package errno

var ERROR_MESSAGE = map[string]Error{
	"ErrLicenseInvalid":       Error{Code: "ErrLicenseInvalid", StatusCode: 400, MessageEN: "Invalid license", MessageCN: "无效License"},
	"ErrInsert":               Error{Code: "ErrInsert", StatusCode: 500, MessageEN: "failed to insert", MessageCN: "保存失败"},
	"ErrUpdate":               Error{Code: "ErrUpdate", StatusCode: 500, MessageEN: "failed to update", MessageCN: "更新失败"},
	"ErrDelete":               Error{Code: "ErrDelete", StatusCode: 500, MessageEN: "failed to delete", MessageCN: "删除失败"},
	"ErrBadRequestParams":     Error{Code: "ErrBadRequestParams", StatusCode: 500, MessageEN: "bad request params", MessageCN: "请求参数错误"},
	"ErrTokenSign":            Error{Code: "ErrTokenSign", StatusCode: 500, MessageEN: "Error generating user token", MessageCN: "生成用户Token时出错"},
	"ErrRecordNotFound":       Error{Code: "ErrRecordNotFound", StatusCode: 500, MessageEN: "Related record or resource do not exist", MessageCN: "记录或资源不存在"},
	"ErrUserNotFound":         Error{Code: "ErrUserNotFound", StatusCode: 404, MessageEN: "Current user does not exist", MessageCN: "当前用户不存在"},
	"ErrBind":                 Error{Code: "ErrBind", StatusCode: 400, MessageEN: "Request parameter error", MessageCN: "请求参数错误"},
	"ErrMissingAuthorization": Error{Code: "ErrMissingAuthorization", StatusCode: 400, MessageEN: "Error missing authorization", MessageCN: "未找到认证信息"},
	"ErrValidation":           Error{Code: "ErrValidation", StatusCode: 400, MessageEN: "Request parameter verification failed", MessageCN: "请求参数校验失败"},
	"ErrProcessKillFailed":    Error{Code: "ErrProcessKillFailed", StatusCode: 500, MessageEN: "Task process stops failing", MessageCN: "任务进程停止失败"},
	"ErrTokenInvalid":         Error{Code: "ErrTokenInvalid", StatusCode: 403, MessageEN: "The token was invalid", MessageCN: "非法Token"},
	"ErrInstanceNotFound":     Error{Code: "ErrInstanceNotFound", StatusCode: 500, MessageEN: "Current instance does not exist", MessageCN: "当前实例信息不存在"},
	"ErrDatabase":             Error{Code: "ErrDatabase", StatusCode: 500, MessageEN: "Database error", MessageCN: "数据库错误"},
	"ErrRateLimit":            Error{Code: "ErrRateLimit", StatusCode: 401, MessageEN: "Exceeded please try again later", MessageCN: "已超限，请稍候重试"},
	"ErrConnectFailed":        Error{Code: "ErrConnectFailed", StatusCode: 400, MessageEN: "Connect failed", MessageCN: "连接失败"},
}
