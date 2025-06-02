package doc

import (
	"net/http"

	"go-zero-box/app/internal/svc"
)

func DocHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
				  <html>
				  <head>
					<title>API Doc</title>
					<link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@3/swagger-ui.css">
				  </head>
				  <body>
					<div id="swagger-ui"></div>
					<script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
					<script>
					  SwaggerUIBundle({
						url: '/api/doc/json', // 你的 JSON 路径
						dom_id: '#swagger-ui',
					  })
					</script>
				  </body>
				  </html>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	}
}
