package files

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"server/internal/storage/postgres"
	"strings"
)

func FileHandler(logger *slog.Logger, storage *postgres.Storage, pathToStaticFolder http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		serverRoutePrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(serverRoutePrefix, http.FileServer(pathToStaticFolder))
		fs.ServeHTTP(w, r)
	}
}
