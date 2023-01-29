package constant

const (
	MissingParametersErr = "参数缺失"
	InvalidParametersErr = "无效参数"

	UserTokenVerificationErr = "token校验错误"
	UserTokenIsInvalidErr    = "无效token"
	UserTokenExpiredErr      = "token过期"
	UserTokenNotFoundErr     = "token不存在"
	UserTokenIsRemovedErr    = "token已被移除"
	UserNamePasswordErr      = "用户名或密码错误"
	UserLoginErr             = "用户登录异常"
	UserAlreadyExistErr      = "用户已存在"

	DingtalkLoginStateNotFoundErr   = "钉钉登录state不存在"
	DingtalkExecuteFalseErr         = "钉钉操作失败"
	DingtalkNoAuthorizationLoginErr = "钉钉用户无权限登录"

	ProcessInstanceNotFoundErr    = "Process Instance Not Found"
	ProcessMappingAlreadyExistErr = "审批映射已存在"

	DataModelAlreadyExistErr                           = "数据模型已存在"
	IndexDefinitionAlreadyExistErr                     = "指标定义已存在"
	IndexDefinitionTargetLakehouseTableAlreadyExistErr = "指标定义目标湖仓表已存在"
)

var errCode = map[string]int{
	MissingParametersErr: 10000,
	InvalidParametersErr: 10001,

	UserTokenVerificationErr: 20000,
	UserTokenIsInvalidErr:    20001,
	UserTokenExpiredErr:      20002,
	UserTokenNotFoundErr:     20004,
	UserTokenIsRemovedErr:    20005,
	UserNamePasswordErr:      20006,
	UserLoginErr:             20007,
	UserAlreadyExistErr:      20008,

	DingtalkLoginStateNotFoundErr:   30000,
	DingtalkExecuteFalseErr:         30001,
	DingtalkNoAuthorizationLoginErr: 30002,

	ProcessInstanceNotFoundErr:    40000,
	ProcessMappingAlreadyExistErr: 40001,

	DataModelAlreadyExistErr: 50000,

	IndexDefinitionAlreadyExistErr: 60000,
}

func GetErrorCode(message string) int {
	return errCode[message]
}
