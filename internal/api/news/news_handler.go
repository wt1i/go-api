package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liwentao0503/go-api/internal/application"
	"github.com/liwentao0503/go-api/internal/infras/utils"
)

// NewsHandler news handler
type NewsHandler struct {
	NewsService *application.NewsService `inject:""`
}

// GetNews godoc
// @Summary Show an news
// @Description get news by ID
// @Tags News
// @Accept  json
// @Produce  json
// @Param news_id path uint true "News ID"
// @Success 200 {object} model.News
// @Router /api/v1/news/{news_id} [get]
func (s *NewsHandler) GetNews(c *gin.Context) {
	var r application.GetNewsReq

	if err := c.ShouldBindUri(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	// param is numeric
	news, err := s.NewsService.GetNews(r)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, news)

}

// GetAllNews godoc
// @Summary Show all news
// @Description show all news by status
// @Tags News
// @Accept  json
// @Produce  json
// @Param status query string false "news's status exist draft|deleted|publish"
// @Param page query int false "page, default is 1"
// @Param page_size query int false "page size, default is 20"
// @Success 200 {object} []model.News
// @Router /api/v1/news [get]
func (s *NewsHandler) GetAllNews(c *gin.Context) {
	var r application.GetAllNewsReq

	if err := c.ShouldBindQuery(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	// if status parameter exist draft|deleted|publish
	news, err := s.NewsService.GetAllNewsByFilter(r)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, news)
}

// CreateNews godoc
// @Summary create an news
// @Description create news
// @Tags News
// @Accept  json
// @Produce  json
// @Param request_data body application.UpsertNewsReq true "news info"
// @Success 200 {object} model.News
// @Router /api/v1/news [post]
func (s *NewsHandler) CreateNews(c *gin.Context) {
	var r application.UpsertNewsReq

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	err := s.NewsService.AddNews(r.BuildModelNews())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, nil)
}

// RemoveNews godoc
// @Summary remove an news
// @Description remove news by ID
// @Tags News
// @Accept  json
// @Produce  json
// @Param news_id path uint true "News ID"
// @Success 200 {object} model.News
// @Router /api/v1/news/{news_id} [delete]
func (s *NewsHandler) RemoveNews(c *gin.Context) {
	var r application.RemoveNewsReq

	if err := c.ShouldBindUri(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := s.NewsService.RemoveNews(r.NewsID); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, nil)
}

// UpdateNews godoc
// @Summary update an news
// @Description update news
// @Tags News
// @Accept  json
// @Produce  json
// @Param news_id path uint true "News ID"
// @Param request_data body application.UpsertNewsReq true "news info"
// @Success 200 {object} nil
// @Router /api/v1/news/{news_id} [put]
func (s *NewsHandler) UpdateNews(c *gin.Context) {
	var r application.UpsertNewsReq

	if err := c.ShouldBindUri(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := s.NewsService.UpdateNews(r.BuildModelNews(), r.ID); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, nil)
}
