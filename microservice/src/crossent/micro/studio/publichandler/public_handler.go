package publichandler

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"

	"fmt"
	"crossent/micro/studio"
)

const yearInSeconds = 31536000

func NewHandler() (http.Handler, error) {
	publicFS := &assetfs.AssetFS{
		Asset:     studio.Asset,
		AssetDir:  studio.AssetDir,
		AssetInfo: studio.AssetInfo,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, private", yearInSeconds))
		http.FileServer(publicFS).ServeHTTP(w, r)
	}), nil
}
