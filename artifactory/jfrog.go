package artifactory

import (
	"net/http"
	"time"

	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/models"
)

type jFrogClient struct {
	jfc *http.Client
	log *logger.Logger
}

func NewJFrogClient(jfc *http.Client, log *logger.Logger) *jFrogClient {
	return &jFrogClient{
		jfc: jfc,
		log: log,
	}
}

func (j jFrogClient) Cat() error {
	return nil
}

func (j jFrogClient) Du() (int, error) {
	return 0, nil
}

func (j jFrogClient) Ls(lsOptions ...models.OptionsLsFunc) (models.ArtifactoryLsResponse, error) {
	lsCmd := &models.LsOptions{
		LongListing: false,
		Color:       false,
		Args:        []string{},
	}

	for _, opt := range lsOptions {
		if err := opt(lsCmd); err != nil {
			return models.JFrogResult{}, err
		}
	}

	return models.JFrogResult{
		JFrogResponse: models.ArtifactoryLsResponseData{
			ArtifactoryRepos: []models.ArtifactoryRepoData{
				{
					Name:         "repo1",
					Size:         123456,
					Perms:        "rwxr-xr-x",
					Owner:        "owner1",
					Group:        "group1",
					LastModified: time.Now(),
				},
			},
			ArtifactoryFiles: []models.ArtifactoryFileData{
				{
					Name:         "file1",
					Size:         123456,
					Perms:        "rwxr-xr-x",
					Owner:        "owner1",
					Group:        "group1",
					LastModified: time.Now(),
				},
			},
			ArtifactoryFolders: []models.ArtifactoryFolderData{
				{
					Name:         "folder1",
					Size:         123456,
					Perms:        "rwxr-xr-x",
					Owner:        "owner1",
					Group:        "group1",
					LastModified: time.Now(),
				},
			},
		},
	}, nil

}

func (j jFrogClient) Touch() (string, error) {
	return "iza-jfrog-touch", nil
}

func (j jFrogClient) WhoAmI() (string, error) {
	return "iza-jfrog-whoami", nil
}
