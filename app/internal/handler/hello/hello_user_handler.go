package hello

import (
	"go-zero-box/app/internal/utils/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-box/app/internal/logic/hello"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"
)

func HelloUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, result.Response(r.Context(), err.Error()))
			return
		}

		l := hello.NewHelloUserLogic(r.Context(), svcCtx)
		resp, err := l.HelloUser(&req)
		if err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, err)
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
		}
	}
}
