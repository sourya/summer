package summer

import (
	"github.com/gin-gonic/gin"
	"log"
)

func delete(c *gin.Context) {
	path := c.Param("path")

	if deleteFolder(path) == true {
		readFolder(c)
	} else {
		deleteFile(c)
	}
}

func deleteFolder(c) {

}

func deleteFile(c) {

}
