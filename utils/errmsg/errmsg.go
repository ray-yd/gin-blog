package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	/* code > 1000 UserModel 錯誤 */
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008

	// code > 2000 ArticleModel 錯誤
	ErrorArticleNotExist = 2001

	// code > 3000 CategoryModel 錯誤
	ErrorCategoryUsed     = 3001
	ErrorCategoryNotExist = 3002
)

var codeMessage = map[int]string{
	SUCCESS:               "SUCCESS",
	ERROR:                 "FAIL",
	ErrorUsernameUsed:     "帳號已存在",
	ErrorPasswordWrong:    "密碼錯誤",
	ErrorUserNotExist:     "帳號不存在",
	ErrorTokenExist:       "TOKEN不存在",
	ErrorTokenRuntime:     "TOKEN過期",
	ErrorTokenWrong:       "TOKEN錯誤",
	ErrorTokenTypeWrong:   "TOKEN格式錯誤",
	ErrorCategoryUsed:     "文章分類已存在",
	ErrorCategoryNotExist: "文章分類不存在",
	ErrorArticleNotExist:  "文章不存在",
	ErrorUserNoRight:      "權限不足",
}

func GetErrorMessage(code int) string {
	return codeMessage[code]
}
