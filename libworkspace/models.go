package libworkspace

type Repository struct {
	ID                 string    `json:"id"`
	Path               string    `json:"path"`
	Dependencies       []string  `json:"dependencies"`
	PostSwitch         []string  `json:"postSwitch,omitEmpty"`
	PostSwitchCommands []Command `json:"postSwitchCommands,omitEmpty"`
}

type Command struct {
	Exe []string `json:"exe"`
}

type Config struct {
	Repositories []Repository `json:"repositories"`
}
