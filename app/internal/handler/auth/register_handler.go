package auth

import (
	"go-zero-box/app/internal/utils/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-box/app/internal/logic/auth"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, result.Response(r.Context(), err.Error()))
			return
		}

		l := auth.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, err)
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
		}
	}
}
