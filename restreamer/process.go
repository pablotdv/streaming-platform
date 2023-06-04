package restreamer

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Process struct {
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

type ProcessResponse struct{}

func CreateProcess(processId string) (Process, error) {
	login, err := Login(LoginRequest{
		Username: "admin",
		Password: "123123",
	})
	if err != nil {
		return Process{}, err
	}

	process := Process{
		Input: []InputOutput{
			{
				ID:      "input_0",
				Address: "rtmp://localhost/" + processId + ".stream",
				Options: []string{
					"-fflags",
					"+genpts",
					"-thread_queue_size",
					"512",
					"-analyzeduration",
					"3000000",
				},
			},
		},
		Output: []Output{
			{
				InputOutput: InputOutput{
					ID:      "output_0",
					Address: "{memfs}/" + processId + "_{outputid}.m3u8",
					Options: []string{
						"-dn",
						"-sn",
						"-map",
						"0:1",
						"-codec:v",
						"copy",
						"-map",
						"0:0",
						"-codec:a",
						"copy",
						"-metadata",
						"title=http://177.2.159.119:8080/" + processId + "/oembed.json",
						"-metadata",
						"service_provider=datarhei-Restreamer",
						"-f",
						"hls",
						"-start_number",
						"0",
						"-hls_time",
						"2",
						"-hls_list_size",
						"6",
						"-hls_flags",
						"append_list+delete_segments+program_date_time+temp_file",
						"-hls_delete_threshold",
						"4",
						"-hls_segment_filename",
						"{memfs}/" + processId + "_{outputid}_%04d.ts",
						"-master_pl_name",
						"" + processId + ".m3u8",
						"-master_pl_publish_rate",
						"2",
						"-method",
						"PUT",
					},
				},
				Cleanup: []Cleanup{
					{
						Pattern:           "memfs:/" + processId + "**",
						MaxFiles:          0,
						MaxFileAgeSeconds: 0,
						PurgeOnDelete:     true,
					},
					{
						Pattern:           "memfs:/" + processId + "_{outputid}.m3u8",
						MaxFiles:          0,
						MaxFileAgeSeconds: 24,
						PurgeOnDelete:     true,
					},
					{
						Pattern:           "memfs:/" + processId + "_{outputid}_**.ts",
						MaxFiles:          12,
						MaxFileAgeSeconds: 24,
						PurgeOnDelete:     true,
					},
					{
						Pattern:           "memfs:/" + processId + ".m3u8",
						MaxFiles:          0,
						MaxFileAgeSeconds: 24,
						PurgeOnDelete:     true,
					},
				},
			},
		},
		Autostart: true,
		ID:        "restreamer-ui:ingest:" + processId,
		Limits: Limits{
			CPUUsage:       0,
			MemoryMBytes:   0,
			WaitforSeconds: 0,
		},
		Options: []string{
			"-err_detect",
			"ignore_err",
			"-y",
		},
		Reconnect:             true,
		ReconnectDelaySeconds: 15,
		Reference:             processId,
		StaleTimeoutSeconds:   30,
		Type:                  "ffmpeg",
	}

	jsonData, err := json.Marshal(process)
	if err != nil {
		return Process{}, err
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v3/process", bytes.NewBuffer(jsonData))
	if err != nil {
		return Process{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+login.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Process{}, err
	}
	defer resp.Body.Close()

	var response Process
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Process{}, err
	}

	return response, nil
}
