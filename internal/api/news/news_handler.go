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

// GetNews get news
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

// GetAllNews get all news
func (s *NewsHandler) GetAllNews(c *gin.Context) {
	var r application.GetAllNewsReq

	if err := c.ShouldBindUri(&r); err != nil {
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

// CreateNews create news
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

// RemoveNews remove news
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

// UpdateNews update news
func (s *NewsHandler) UpdateNews(c *gin.Context) {
	var r application.UpsertNewsReq

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
