package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func translateName(apiFamilyType string) string {
	out := strings.ReplaceAll(apiFamilyType, "-", " ")
	out = strings.ReplaceAll(out, "_", " - ")
	out = strings.Replace(out, "opendata", "open data", 1)
	out = strings.Title(out)
	return out
}

func importData(url string) (ParticipantsData, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}

	defer resp.Body.Close()
	
	var data ParticipantsData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println(err)
	}

	return data, nil
}

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
