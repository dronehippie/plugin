package drone

// Git provides the git repository metadata.
type Git struct {
	HTTPURL string `envconfig:"DRONE_GIT_HTTP_URL"`
	SSHURL  string `envconfig:"DRONE_GIT_SSH_URL"`
}
