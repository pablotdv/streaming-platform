package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v3/process/restreamer-ui:ingest:60f63649-0528-430c-b852-5fa076135ea7", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleGkiOjYwMCwiZXhwIjoxNjg1ODI2NjAzLCJpYXQiOjE2ODU4MjYwMDMsImlzcyI6ImRhdGFyaGVpLWNvcmUiLCJqdGkiOiJmMDQ4NTc0Ny05NTMyLTQyZWUtOWVmYy1mZDIwMjZkOGU1MDciLCJzdWIiOiJhZG1pbiIsInVzZWZvciI6ImFjY2VzcyJ9.OzEbTLYyl_4DvOc3lxvOz6KQm0bwlwATvsm1QckZcdc")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
