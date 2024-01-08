package application

import (
	"testing"

	"go-api/internal/domain/model"
	mock "go-api/internal/domain/repository"

	"github.com/golang/mock/gomock"
)

func Test_GetNews(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockNewsService := mock.NewMockNewsRepository(mockCtrl)
	getNews := &NewsService{
		NewsRepo: mockNewsService,
	}

	var req = GetNewsReq{
		NewsID: 123,
	}
	mockNewsService.EXPECT().Get(req.NewsID).Return(&model.News{
		TopicID: 1,
		Title:   "test",
	}, nil).Times(1)

	news, err := getNews.GetNews(req)

	if err != nil {
		t.Errorf("get err :%v ", err)
	}

	if news.Title != "test" {
		t.Errorf("get title err")
	}
}
