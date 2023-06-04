package restreamer

type Job struct {
	Input                 []InputOutput `json:"input"`
	Output                []Output      `json:"output"`
	Autostart             bool          `json:"autostart"`
	ID                    string        `json:"id"`
	Limits                Limits        `json:"limits"`
	Options               []string      `json:"options"`
	Reconnect             bool          `json:"reconnect"`
	ReconnectDelaySeconds int           `json:"reconnect_delay_seconds"`
	Reference             string        `json:"reference"`
	StaleTimeoutSeconds   int           `json:"stale_timeout_seconds"`
	Type                  string        `json:"type"`
}

type InputOutput struct {
	ID      string   `json:"id"`
	Address string   `json:"address"`
	Options []string `json:"options"`
}

type Output struct {
	InputOutput
	Cleanup []Cleanup `json:"cleanup"`
}

type Cleanup struct {
	Pattern           string `json:"pattern"`
	MaxFiles          int    `json:"max_files"`
	MaxFileAgeSeconds int    `json:"max_file_age_seconds"`
	PurgeOnDelete     bool   `json:"purge_on_delete"`
}

type Limits struct {
	CPUUsage       int `json:"cpu_usage"`
	MemoryMBytes   int `json:"memory_mbytes"`
	WaitforSeconds int `json:"waitfor_seconds"`
}
