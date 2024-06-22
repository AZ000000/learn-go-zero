package handler

import (
	"net/http"

	"zeroStudy/api_study/user/api_jwt/internal/logic"
	"zeroStudy/api_study/user/api_jwt/internal/svc"
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
