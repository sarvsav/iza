package app

import (
	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/internals/artifactstore"
	"github.com/sarvsav/iza/internals/cicd"
	"github.com/sarvsav/iza/internals/datastore"
)

type Application struct {
	Logger             *logger.Logger
	ArtifactoryService *artifactstore.ArtifactoryService
	CiCdService        *cicd.CiCdService
	DataStoreService   *datastore.DataStoreService
}
