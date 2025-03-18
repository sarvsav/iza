package devops

type jenkinsClient struct {
	userName string
	apiToken string
}

func NewJenkinsClient(user, apiToken string) *jenkinsClient {
	return &jenkinsClient{
		userName: user,
		apiToken: apiToken,
	}
}

func (j jenkinsClient) Cat() error {
	return nil
}

func (j jenkinsClient) Du() (int, error) {
	return 0, nil
}

func (j jenkinsClient) Ls() (string, error) {
	return "iza-jenkins-ls", nil
}

func (j jenkinsClient) Touch() (string, error) {
	return "iza-jenkins-touch", nil
}

func (j jenkinsClient) WhoAmI() (string, error) {
	return "iza-jenkins-whoami", nil
}
