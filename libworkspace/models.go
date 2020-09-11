package libworkspace

type Repository struct {
	ID                 string    `json:"id"`
	Path               string    `json:"path"`
	Dependencies       []string  `json:"dependencies"`
	PostSwitch         []string  `json:"postSwitch,omitEmpty"`
	PostSwitchCommands []Command `json:"postSwitchCommands,omitEmpty"`
	OpenWith           string    `json:"openWith,omitEmpty"`
}

type Command struct {
	Exe []string          `json:"exe"`
	Env map[string]string `json:"env,omitEmpty"`
}

type Config struct {
	Repositories []Repository `json:"repositories"`
}
