package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zeroStudy/api_study/user/api_prefix/internal/logic"
	"zeroStudy/api_study/user/api_prefix/internal/svc"
	"zeroStudy/api_study/user/api_prefix/internal/types"

	"zeroStudy/common/response"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		//修改为
		response.Response(r, w, resp, err)
	}
}
