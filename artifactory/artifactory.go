package artifactory

import "github.com/sarvsav/iza/models"

type Client interface {
	Cat() error
	Du() (int, error)
	Ls(lsOptions ...models.OptionsLsFunc) (models.JFrogResult, error)
	Touch() (string, error)
	WhoAmI() (string, error)
}
