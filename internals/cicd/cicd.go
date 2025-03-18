package cicd

import (
	"github.com/sarvsav/iza/devops"
	"github.com/sarvsav/iza/foundation/logger"
)

type CiCdService struct {
	devops devops.Client
	logger *logger.Logger
}

func NewCiCdService(devops devops.Client, logger *logger.Logger) *CiCdService {
	return &CiCdService{
		devops: devops,
		logger: logger,
	}
}

func (ci *CiCdService) Cat() error {
	err := ci.devops.Cat()
	if err != nil {
		return err
	}

	return nil
}

func (ci *CiCdService) Du() (int, error) {
	diskUsage, err := ci.devops.Du()
	if err != nil {
		return 0, err
	}

	return diskUsage, nil
}

func (ci *CiCdService) Ls() (string, error) {
	listOfFiles, err := ci.devops.Ls()
	if err != nil {
		return "", err
	}

	return listOfFiles, nil
}

func (ci *CiCdService) Touch() (string, error) {
	fileCreated, err := ci.devops.Touch()
	if err != nil {
		return "", err
	}

	return fileCreated, nil
}

func (ci *CiCdService) WhoAmI() (string, error) {
	userName, err := ci.devops.WhoAmI()
	if err != nil {
		return "", err
	}

	return userName, nil
}
