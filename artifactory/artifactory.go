package artifactory

import "github.com/sarvsav/iza/models"

type Client interface {
	Cat() error
	Du() (int, error)
	Ls(lsOptions ...models.OptionsLsFunc) (models.ArtifactoryLsResponse, error)
	Touch() (string, error)
	WhoAmI() (string, error)
}
