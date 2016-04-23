package summer

import (
	"github.com/gin-gonic/gin"
	"log"
)

func writer(c *gin.Context) {
	path := c.Param("path")
	content := c.query("content")

	if isFolder(path) == true {
		newFolder(c)
	} else {
		writeFile(c)
	}
}

func newFolder(c) {

}

func writeFile(c) {

}
