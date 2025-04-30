package cicd

import "github.com/sarvsav/iza/models"

func (ci *CiCdService) Ls(lsOptions ...models.OptionsLsFunc) (models.JenkinsLsResponse, error) {
	result, err := ci.devops.Ls(lsOptions...)
	if err != nil {
		return models.JenkinsLsResponse{}, err
	}

	return result, nil
}
