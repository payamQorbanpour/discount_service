package pkg

import (
	"discount_service/internal/repository"
	"discount_service/internal/service"
	"discount_service/internal/webapi"

	"github.com/go-kit/kit/log"
)

type DiscountService struct {
	repository repository.Repository
	webAPI     webapi.WebAPI
	logger     log.Logger
}

func NewService(repo repository.Repository, webAPI webapi.WebAPI, logger log.Logger) service.Service {
	return &DiscountService{
		repository: repo,
		webAPI:     webAPI,
		logger:     logger,
	}
}
