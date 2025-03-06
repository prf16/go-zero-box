package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}if err != nil  {
            httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, err)
		} else {
		    httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
		}{{else}}httpx.Ok(w){{end}}
	}
}
