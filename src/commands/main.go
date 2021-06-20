package commands

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	List = []*discordgo.ApplicationCommand{
		&niiloInfo,
		&addNiiloInfo,
	}

	Handlers = map[string]func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate){
		niiloInfo.Name:    niiloCmd,
		addNiiloInfo.Name: addNiiloCmd,
	}
)

var client = &http.Client{Timeout: 10 * time.Second}

const UserAgent = "Niilo22-Bot/1.0"

func getData(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", UserAgent)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&target)
}

func getDataRaw(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", UserAgent)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func postData(url string, target interface{}, body interface{}) error {
	json_body, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_body))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&target)
}
