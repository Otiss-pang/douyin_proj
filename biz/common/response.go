package common

type Response struct {
	StatusCode int         `json:"code"`
	StatusMsg  string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{
		StatusCode: 200,
		StatusMsg:  "Success",
		Data:       data,
	}
}

func Fail(code int) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  "System error",
		Data:       nil,
	}
}

func FailWithMsg(code int, msg string) Response {
	return Response{
		StatusCode: code,
		StatusMsg:  msg,
		Data:       nil,
	}
}
