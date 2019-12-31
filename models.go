package main

type Repository struct {
	ID           string   `json:"id"`
	Path         string   `json:"path"`
	Dependencies []string `json:"dependencies"`
	PostSwitch   []string `json:"postSwitch,omitEmpty"`
}

type GitConfig struct {
	PrivateKeyPath    string `json:"privateKeyPath"`
	KeyPairPassphrase string `json:"keyPairPassphrase"`
}

type Config struct {
	Repositories []Repository `json:"repositories"`
	GitConfig    GitConfig    `json:"gitConfig"`
}
