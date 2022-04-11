package routes

import (
	"file_uploader/pkg/controller"
	"file_uploader/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func RunAllRoutes() {
	r := gin.Default()

	// Использование CORS
	r.Use(controller.CORSMiddleware())

	// Установка Logger-а
	utils.SetLogger()

	// Форматирование логов
	utils.FormatLogs(r)

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	//r.Use(limits.RequestSizeLimiter(100))

	// Запуск end-point'ов
	runAllRoutes(r)

	// Запуск сервера
	runServer(r)

}

func runAllRoutes(r *gin.Engine) {
	r.GET("/", HealthCheck)

	r.POST("/file/upload", controller.UploadFile)
	r.GET("/file/download", controller.DownloadFile)
}

func runServer(r *gin.Engine) {
	var (
		port string
		host string
	)
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3001"
		host = "localhost"
	} else {
		host = "0.0.0.0"
	}
	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Println(err)
	}
}

func HealthCheck(c *gin.Context) {
	//res := map[string]interface{}{
	//	"data": "Server is up and running",
	//}

	c.JSON(http.StatusOK, gin.H{"data": "Server is up and running"})
}
