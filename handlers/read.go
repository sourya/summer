package summer

import (
	"github.com/gin-gonic/gin"
	"log"
)

func reader(c *gin.Context) {
	path := c.Param("path")

	if isFolder(path) == true {
		readFolder(c)
	} else {
		readFile(c)
	}
}

func readFolder(c) {

}

func readFile(c) {

}
