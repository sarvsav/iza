package artifactory

import "github.com/sarvsav/iza/models"

func (as *ArtifactoryService) Ls(lsOptions ...models.OptionsLsFunc) (models.ArtifactoryLsResponse, error) {
	result, err := as.artifactory.Ls(lsOptions...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
