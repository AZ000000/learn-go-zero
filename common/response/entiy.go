package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// 构造一个 code,data,message
type Body struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Response(r *http.Request, w http.ResponseWriter, res any, err error) {
	if err != nil {
		//根据不同的错误码，返回不同的错误信息
		body := Body{
			Code: 500,
			Data: nil,
			Msg:  "错误",
		}
		httpx.WriteJson(w, http.StatusOK, body)
	}
	body := Body{
		Code: 200,
		Data: res,
		Msg:  "成功",
	}
	httpx.WriteJson(w, http.StatusOK, body)

}
