package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zeroStudy/api_study/user/api_prefix/internal/logic"
	"zeroStudy/api_study/user/api_prefix/internal/svc"
	"zeroStudy/common/response"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		//修改为
		response.Response(r, w, resp, err)
	}
}
