package service

import (
	"github.com/rs/zerolog"

	"github.com/agilefoxHQ/go-function-template/config"
)

type Service struct {
	config *config.Configuration
	logger *zerolog.Logger
	//repository *repository.Repositories
}

func NewService(
	config *config.Configuration,
	logger *zerolog.Logger,
// repository *repository.Repositories,
) *Service {
	return &Service{
		config: config,
		logger: logger,
		//repository: repository
	}
}
