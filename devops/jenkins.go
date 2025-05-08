package devops

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/models"
)

type jenkinsClient struct {
	jc  *http.Client
	log *logger.Logger
}

func NewJenkinsClient(jc *http.Client, log *logger.Logger) *jenkinsClient {
	return &jenkinsClient{
		jc:  jc,
		log: log,
	}
}

func (j jenkinsClient) Cat() error {
	return nil
}

func (j jenkinsClient) Du() (int, error) {
	return 0, nil
}

func (j jenkinsClient) Ls(lsOptions ...models.OptionsLsFunc) (models.JenkinsLsResponse, error) {
	lsCmd := &models.LsOptions{
		LongListing: false,
		Color:       false,
		Args:        []string{},
	}

	for _, opt := range lsOptions {
		if err := opt(lsCmd); err != nil {
			return models.JenkinsLsResponse{}, err
		}
	}

	blueOceanUrl := "/blue/rest/organizations/jenkins/pipelines"
	url := os.Getenv("JENKINS_URL") + blueOceanUrl

	if len(lsCmd.Args) > 0 {
		url += "/" + lsCmd.Args[0] + "/pipelines"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// Handle error
	}

	resp, err := j.jc.Do(req)
	if err != nil {
		// Handle error
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle error
	}

	var response models.JenkinsApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		// Handle error
	}

	var jenkinsJobs []models.JenkinsJob
	var jenkinsJobFolders []models.JenkinsJobFolder
	var lastRunTime time.Time
	for _, v := range response {
		if v.Class == "io.jenkins.blueocean.service.embedded.rest.PipelineFolderImpl" {
			jenkinsJobFolders = append(jenkinsJobFolders, models.JenkinsJobFolder{
				Name:         v.Name,
				Size:         0,
				Perms:        "rwx",
				Owner:        v.Organization,
				Group:        v.Organization,
				LastModified: time.Now(),
			})
		} else {
			lastRunTime, err = time.Parse("2000-01-14T12:47:00.102+0200", v.StartTime)
			if err != nil {
				// Handle error
			}
			jenkinsJobs = append(jenkinsJobs, models.JenkinsJob{
				Name:         v.Name,
				Size:         0,
				Perms:        "rwx",
				Owner:        v.Organization,
				Group:        v.Organization,
				LastModified: lastRunTime,
			})
		}
	}

	result := models.JenkinsLsResponse{
		JenkinsJobFolders: jenkinsJobFolders,
		JenkinsJobs:       jenkinsJobs,
	}

	return result, nil

}

func (j jenkinsClient) Touch() (string, error) {
	return "iza-jenkins-touch", nil
}

func (j jenkinsClient) WhoAmI() (string, error) {
	return "iza-jenkins-whoami", nil
}
