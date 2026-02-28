package hello

import (
	"go-zero-box/app/internal/utils/result"
	"net/http"

	"go-zero-box/app/internal/logic/hello"
	"go-zero-box/app/internal/svc"
	"go-zero-box/app/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloUserRpcHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloUserRpcReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, result.Response(r.Context(), err.Error()))
			return
		}

		l := hello.NewHelloUserRpcLogic(r.Context(), svcCtx)
		resp, err := l.HelloUserRpc(&req)
		if err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, err)
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
		}
	}
}
