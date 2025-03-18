package devops

type Client interface {
	Cat() error
	Du() (int, error)
	Ls() (string, error)
	Touch() (string, error)
	WhoAmI() (string, error)
}
