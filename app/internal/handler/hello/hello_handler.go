package hello

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-box/app/internal/logic/hello"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := hello.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)
		if err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, err)
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
		}
	}
}
