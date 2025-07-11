package artifactstore

import "github.com/sarvsav/iza/models"

func (as *ArtifactoryService) Ls(lsOptions ...models.OptionsLsFunc) (models.ArtifactoryLsResponseData, error) {
	result, err := as.artifactory.Ls(lsOptions...)
	if err != nil {
		return models.ArtifactoryLsResponseData{}, err
	}

	resultData, _ := result.GetLsResult()

	return resultData, nil
}
