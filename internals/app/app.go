package app

import (
	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/internals/artifactory"
	"github.com/sarvsav/iza/internals/cicd"
	"github.com/sarvsav/iza/internals/datastore"
)

type Application struct {
	Logger             *logger.Logger
	ArtifactoryService *artifactory.ArtifactoryService
	CiCdService        *cicd.CiCdService
	DataStoreService   *datastore.DataStoreService
}
