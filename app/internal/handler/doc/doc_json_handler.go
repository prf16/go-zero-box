package doc

import (
	"net/http"

	"go-zero-box/app/internal/svc"
)

func DocJsonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/api/api.json")
	}
}
