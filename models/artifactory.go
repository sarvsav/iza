package models

import "time"

type ArtifactoryLsResponseData struct {
	ArtifactoryRepos   []ArtifactoryRepoData
	ArtifactoryFolders []ArtifactoryFolderData
	ArtifactoryFiles   []ArtifactoryFileData
}

type ArtifactoryRepoData struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type ArtifactoryFolderData struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type ArtifactoryFileData struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type ArtifactoryLsResponse interface {
	GetResult() (ArtifactoryLsResponseData, error)
	isArtifactoryLsResponse() // marker method
}
