package exception

type ErrorModel struct {
	Code    int
	Message string
}

type ErrorCode int

const (
	INTERNAL_ERROR ErrorCode = 1 + iota
	NOT_LOGIN
	NOT_EXIST_ARTICLE_ERROR
)

var errorCodes = [...]ErrorModel{
	ErrorModel{Code:-1, Message:"系统内部错误"},
	ErrorModel{Code:1000, Message:"未登录"},
	ErrorModel{Code:1001, Message:"文章不存在"},

}

func (c ErrorCode) Code() int {
	return errorCodes[c - 1].Code
}

func (c ErrorCode) Message() string {
	return errorCodes[c - 1].Message
}


