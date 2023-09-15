package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liwentao0503/go-api/internal/api/middleware"
	"github.com/liwentao0503/go-api/internal/api/news"
	"github.com/liwentao0503/go-api/internal/api/topics"
	"github.com/liwentao0503/go-api/internal/infras/config"
	"github.com/liwentao0503/go-api/internal/infras/migration"
	"github.com/liwentao0503/go-api/internal/infras/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterHandler api router
type RouterHandler struct {
	AppConfig     *config.AppConfig        `inject:""`
	TopicHandler  *topics.TopicHandler     `inject:""`
	NewsHandler   *news.NewsHandler        `inject:""`
	MigrateAction *migration.MigrateAction `inject:""`
}

// Router create router handler
func (s *RouterHandler) Router() http.Handler {
	// gin mode
	s.ginMode()
	router := gin.New()
	router.Use(middleware.AccessLog(), middleware.RecoverHandler())
	router.NoRoute(middleware.NotFoundHandler())
	s.webRoute(router)
	return router
}

func (s *RouterHandler) ginMode() {
	// gin mode设置
	switch s.AppConfig.AppEnv {
	case "local", "dev":
		gin.SetMode(gin.DebugMode)
	case "testing":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

func (s *RouterHandler) webRoute(router *gin.Engine) {
	// home route
	router.GET("/", s.home)
	router.GET("/api/v1", s.home)

	// News route
	router.GET("/api/v1/news", s.NewsHandler.GetAllNews)
	router.GET("/api/v1/news/:news_id", s.NewsHandler.GetNews)
	router.POST("/api/v1/news", s.NewsHandler.CreateNews)
	router.DELETE("/api/v1/news/:news_id", s.NewsHandler.RemoveNews)
	router.PUT("/api/v1/news/:news_id", s.NewsHandler.UpdateNews)

	// Topic route
	router.GET("/api/v1/topic", s.TopicHandler.GetAllTopic)
	router.GET("/api/v1/topic/:topic_id", s.TopicHandler.GetTopic)
	router.POST("/api/v1/topic", s.TopicHandler.CreateTopic)
	router.DELETE("/api/v1/topic/:topic_id", s.TopicHandler.RemoveTopic)
	router.PUT("/api/v1/topic/:topic_id", s.TopicHandler.UpdateTopic)

	// Migration Route
	router.POST("/api/v1/migrate", s.migrate)
	router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// migrate godoc
// @Summary migrate db
// @Description migrate db
// @Tags migrate
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /api/v1/migrate [post]
func (s *RouterHandler) migrate(c *gin.Context) {
	err := s.MigrateAction.DBMigrate()
	if err != nil {
		log.Println("request_uri: ", c.Request.RequestURI)
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	msg := "Success Migrate"
	utils.JSON(c, http.StatusOK, msg)
}

// home index handler
func (s *RouterHandler) home(c *gin.Context) {
	c.String(http.StatusOK, "GO DDD API")
}
