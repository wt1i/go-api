package topics

import (
	"net/http"

	"go-api/internal/application"
	"go-api/internal/infras/utils"

	"github.com/gin-gonic/gin"
)

// TopicHandler topic handler
type TopicHandler struct {
	TopicService *application.TopicService `inject:""`
}

// GetTopic godoc
// @Summary Show an topic
// @Description get topic by ID
// @Tags Topic
// @Accept  json
// @Produce  json
// @Param topic_id path uint true "topic ID"
// @Success 200 {object} model.Topic
// @Router /api/v1/topic/{topic_id} [get]
func (s *TopicHandler) GetTopic(c *gin.Context) {
	var r application.GetTopicReq

	if err := c.ShouldBindUri(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	topic, err := s.TopicService.GetTopic(c.Request.Context(), r.TopicID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, topic)
}

// GetAllTopic godoc
// @Summary Show all topic
// @Description Show all topic
// @Tags Topic
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Topic
// @Router /api/v1/topic [get]
func (s *TopicHandler) GetAllTopic(c *gin.Context) {
	topics, err := s.TopicService.GetAllTopic(c.Request.Context())
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, topics)
}

// CreateTopic godoc
// @Summary create an topic
// @Description create topic
// @Tags Topic
// @Accept  json
// @Produce  json
// @Param request_data body application.UpsertTopicReq true "topic info"
// @Success 200 {object} model.Topic
// @Router /api/v1/topic [post]
func (s *TopicHandler) CreateTopic(c *gin.Context) {
	var r application.UpsertTopicReq

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := s.TopicService.AddTopic(c.Request.Context(), r.Name, r.Slug); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusCreated, nil)
}

// RemoveTopic godoc
// @Summary remove an topic
// @Description remove topic by ID
// @Tags Topic
// @Accept  json
// @Produce  json
// @Param topic_id path uint true "topic ID"
// @Success 200 {object} model.Topic
// @Router /api/v1/topic/{topic_id} [delete]
func (s *TopicHandler) RemoveTopic(c *gin.Context) {
	var r application.RemoveTopicReq

	if err := c.ShouldBindUri(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := s.TopicService.RemoveTopic(c.Request.Context(), r.TopicID); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, nil)
}

// UpdateTopic godoc
// @Summary update an topic
// @Description update topic
// @Tags Topic
// @Accept  json
// @Produce  json
// @Param topic_id path uint true "topic ID"
// @Param request_data body application.UpsertTopicReq true "topic info"
// @Success 200 {object} nil
// @Router /api/v1/topic/{topic_id} [put]
func (s *TopicHandler) UpdateTopic(c *gin.Context) {
	var r application.UpsertTopicReq

	if err := c.ShouldBindUri(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := c.ShouldBindJSON(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, "param invaild", err)
		return
	}

	if err := s.TopicService.UpdateTopic(c.Request.Context(), r.ID, r.BuildModelTopic()); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(c, http.StatusOK, nil)
}
