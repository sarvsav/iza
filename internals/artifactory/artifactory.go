package artifactory

import (
	"github.com/sarvsav/iza/artifactory"
	"github.com/sarvsav/iza/foundation/logger"
)

type ArtifactoryService struct {
	artifactory artifactory.Client
	logger      *logger.Logger
}

func NewArtifactoryService(artifactory artifactory.Client, logger *logger.Logger) *ArtifactoryService {
	return &ArtifactoryService{
		artifactory: artifactory,
		logger:      logger,
	}
}
