package exception

type ErrorModel struct {
	Code    int
	Message string
}

type ErrorCode int

const (
	INTERNAL_ERROR ErrorCode = 1
	NOT_LOGIN
	NOT_EXIST_ARTICLE_ERROR
	ARTICLE_TITLE_LEN_OVERFLOW
	ARTICLE_CATEGORY_LEN_OVERFLOW
	ARTICLE_TAG_LEN_OVERFLOW
	ARTICLE_CONTENT_LEN_OVERFLOW
)

var errorCodes = [...]ErrorModel{
	ErrorModel{Code:-1, Message:"系统内部错误"},
	ErrorModel{Code:1000, Message:"未登录"},
	ErrorModel{Code:1001, Message:"文章不存在"},
	ErrorModel{Code:1002, Message:"文章标题太长"},
	ErrorModel{Code:1003, Message:"文章分类太长"},
	ErrorModel{Code:1004, Message:"文章标题太长"},
	ErrorModel{Code:1005, Message:"文章内容太长"},
}

func (c ErrorCode) Code() int {
	return errorCodes[c - 1].Code
}

func (c ErrorCode) Error() string {
	return errorCodes[c - 1].Message
}


