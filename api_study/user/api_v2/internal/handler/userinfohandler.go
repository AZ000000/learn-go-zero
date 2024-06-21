package handler

import (
	"net/http"
	"zeroStudy/common/response"

	"zeroStudy/api_study/user/api_v2/internal/logic"
	"zeroStudy/api_study/user/api_v2/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)
	}
}
