package info

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-box/app/internal/logic/user/info"
	"go-zero-box/app/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := info.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, err)
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
		}
	}
}
