package summer

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	_ = viper.ReadInConfig()
}

func main() {
	router := gin.Default()

	// Internal routing module
	router(*router)

	router.Run(":" + viper.Get("appPort").(string))
}
