package routes

import (
	"example/band-protocol/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type broadcast struct {
	*gin.Context
}

type route struct {
	*gin.Engine
}

func NewBroadcast(c *gin.Context) *broadcast {
	return &broadcast{
		Context: c,
	}
}

func NewGin(handler func(services.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewBroadcast(c))
	}
}

func (b *broadcast) QueryParam(key string) string {
	return b.Context.Query(key)
}

func (b *broadcast) Json(statusCode int, input interface{}) {
	b.Context.JSON(statusCode, input)
}

func (b *broadcast) Bind(input interface{}) error {
	return b.Context.ShouldBindJSON(input)
}

func Router() *route {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	return &route{r}
}
