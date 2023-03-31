package route

import (
	"eight-learn/handler"
	"eight-learn/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/books")
	{
		api.GET("", server.GetBooks)
		api.GET(":id", server.GetBookById)
		api.POST("", server.CreateBook)
		api.PUT(":id", server.UpdateBook)
		api.DELETE(":id", server.DeleteBook)
	}
}
