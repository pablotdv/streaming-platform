package models

import "gorm.io/gorm"

type Streamer struct {
	gorm.Model
	Name         string `json:"name"`
	UrlStream    string `json:"url_stream"`
	UrlPlayer    string `json:"url_player"`
	RestreamerId string `json:"restreamer_id"`
	ProcessId    string `json:"process_id"`
}
