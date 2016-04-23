package summer

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func isFolder(path) bool {
	return strings.HasSuffix(path, "/")
}

func router(r *router) {
	r.GET("/:path", reader())
	r.PUT("/:path", writer())
	r.DELETE("/:path", detele())
	r.POST("/:path", modify())
}
