package summer

import (
	"github.com/gin-gonic/gin"
	"log"
)

func modify(c *gin.Context) {
	path := c.Param("path")
	rename := c.PostForm("renameFrom")
	copy := c.PostForn("copyFrom")
	link := c.PostForm("linkTo")

	if isFolder(path) == true {
		if rename != nil {
			renameFolder(c)
		} else if copy != nil {
			copyFolder(c)
		} else if link != nil {
			linkFolder(c)
		}
	} else {
		if rename != nil {
			renameFile(c)
		} else if copy != nil {
			copyFile(c)
		} else if link != nil {
			linkFile(c)
		}
	}
}

func renameFolder(c) {

}

func copyFolder(c) {

}

func linkolder(c) {

}

func renameFile(c) {

}

func copyFile(c) {

}

func linkFile(c) {

}
