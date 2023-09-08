package topics

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liwentao0503/go-api/internal/application"
	"github.com/liwentao0503/go-api/internal/infras/utils"
)

// TopicHandler topic handler
type TopicHandler struct {
	TopicService *application.TopicService `inject:""`
}

// GetTopic get topic
func (s *TopicHandler) GetTopic(c *gin.Context) {
	var r application.GetTopicReq

	if err := c.ShouldBindUri(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	topic, err := s.TopicService.GetTopic(r.TopicID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, topic)
}

// GetAllTopic get all topic
func (s *TopicHandler) GetAllTopic(c *gin.Context) {
	topics, err := s.TopicService.GetAllTopic()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, topics)
}

// CreateTopic create topic
func (s *TopicHandler) CreateTopic(c *gin.Context) {
	var r application.UpsertTopicReq

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := s.TopicService.AddTopic(r.Name, r.Slug); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, nil)
}

// RemoveTopic remove topic
func (s *TopicHandler) RemoveTopic(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Param("topic_id"))
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	err = s.TopicService.RemoveTopic(topicID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, nil)
}

// UpdateTopic update topic
func (s *TopicHandler) UpdateTopic(c *gin.Context) {
	var r application.UpsertTopicReq

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := s.TopicService.UpdateTopic(r.BuildModelTopic(), r.ID); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, nil)
}
