package info

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-box/app/internal/logic/user/info"
	"go-zero-box/app/internal/svc"
)

func ExitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewExitLogic(r.Context(), svcCtx)
		resp, err := l.Exit()
		if err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, err)
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
		}
	}
}
