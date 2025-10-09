package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-box/app/internal/config"
	"io"
	"net/http"
	"strings"
)

type LogMiddleware struct {
	config *config.Config
}

func NewLogMiddleware(config *config.Config) *LogMiddleware {
	return &LogMiddleware{config: config}
}

func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// 合并 Headers
		headerMap := make(map[string]string)
		for k, v := range r.Header {
			headerMap[k] = strings.Join(v, ", ")
		}
		headerStr, _ := json.Marshal(headerMap)

		// 打印 Body
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)

			// 重新放回 Body，避免业务逻辑里拿不到
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 打印完整请求信息
		logx.WithContext(ctx).Infof("Request | Method=%s | Path=%s | Headers=%s | Body=%s", r.Method, r.URL.Path, headerStr, string(bodyBytes))
		next.ServeHTTP(w, r)
	}
}
