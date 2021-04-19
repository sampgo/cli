package sampctl

import "encoding/json"

type Runtime struct {
	Plugins []string `json:"plugins"`
}

type Config struct {
	User         string   `json:"user"`
	Repo         string   `json:"repo"`
	Entry        string   `json:"entry"`
	Output       string   `json:"output"`
	Dependencies []string `json:"dependencies"`
	Local        bool     `json:"local"`
	Runtime      Runtime  `json:"runtime"`
}

func UnmarshalConfig(data []byte) (Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Config) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
