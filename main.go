package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Request struct {
	Model    string  `json:"model"`
	Prompt   string  `json:"prompt"`
	System   *string `json:"system,omitempty"`
	Template *string `json:"template,omitempty"`
	Stream   bool    `json:"stream"`
	Options  struct {
		Temperature float64 `json:"temperature"`
	} `json:"options"`
}
type Response struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	Context            []int     `json:"context"`
	TotalDuration      int       `json:"total_duration"`
	LoadDuration       int       `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int       `json:"eval_duration"`
}

type Config struct {
	URL         string  `json:"url"`
	System      string  `json:"system"`
	Prompt      string  `json:"prompt"`
	Template    string  `json:"template"`
	Model       string  `json:"model"`
	Temperature float64 `json:"temperature"`
}

func main() {

	pathConfig := flag.String("config", "./properties.json", "path to config file")
	prompt := flag.String("prompt", "", "prompt to use")
	model := flag.String("model", "", "model to use")
	template := flag.String("template", "@", "template: you need assign '' for put empty this field")
	temperature := flag.Float64("temperature", -1, "temperature to use")
	system := flag.String("system", "", "system to use")
	flag.Parse()
	var config Config

	contentConfigJson, err := os.ReadFile(*pathConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(contentConfigJson, &config)
	if err != nil {
		log.Fatal(err)
	}

	var requestBody Request

	requestBody.Model = config.Model
	if *model != "" {
		requestBody.Model = *model
	}

	if config.Prompt == "" {
		requestBody.Prompt = *prompt
	} else {
		requestBody.Prompt = fileToString(config.Prompt) + " " + *prompt
	}
	requestBody.System = nil
	if config.System != "" {
		tmp := fileToString(config.System)
		requestBody.System = &tmp
	}

	if *system != "" {
		requestBody.System = system
	}

	requestBody.Template = nil
	if config.Template != "" {
		tmp := fileToString(config.Template)
		requestBody.Template = &tmp
	}

	if *template != "@" {
		requestBody.Template = template
	}

	requestBody.Stream = false
	requestBody.Options.Temperature = config.Temperature
	if *temperature >= 0.0 {
		requestBody.Options.Temperature = *temperature
	}

	marshalled, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, config.URL, bytes.NewReader(marshalled))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var responseBody Response

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseBody.Response)

}

func fileToString(filename string) string {
	b, err := os.ReadFile(filename) // just pass the file name
	if err != nil {
		return ""
	}

	return string(b)
}
