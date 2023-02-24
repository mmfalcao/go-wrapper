package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var client *http.Client

type Episode struct {
	Id       int
	Name     string
	Air_date string
	Episode  string
	Url      string
	Created  string
}

func GetEpisode() ([]Episode, error) {
	url := "https://rickandmortyapi.com/api/episode/12,17,18"

	var episode []Episode

	err := GetJson(url, &episode)
	if err != nil {
		fmt.Printf("error getting episode: %s \n", err.Error())
		return episode, err
	} else {
		for e, _ := range episode {
			fmt.Printf("Found this episode: %s \n", episode[e].Name)
		}
		return episode, err
	}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func CreateCSV(episode []Episode) error {
	csvfile, err := os.Create("data.csv")

	if err != nil {
		log.Fatal("Failed to create file: %s \n", err)
		return err
	}

	writer := csv.NewWriter(csvfile)
	header := []string{"Name", "Episode", "Url"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, row := range episode {
		var r []string
		fmt.Printf("save episode: %s \n", row.Name)
		r = append(r, row.Name, row.Episode, row.Url)
		if err := writer.Write(r); err != nil {
			return err
		}
	}

	writer.Flush()
	csvfile.Close()

	return csvfile.Chdir()
}

func main() {

	client = &http.Client{Timeout: 10 * time.Second}

	var eps []Episode
	eps, _ = GetEpisode()
	CreateCSV(eps)
}
