package copilot

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"path"

	"github.com/jtefteller/copilot_cli/utility"
)

type CompletionConfig struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Suffix           *string `json:"suffix"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	TopP             float64 `json:"top_p"`
	N                int     `json:"n"`
	Stream           bool    `json:"stream"`
	LogProbs         int     `json:"logprobs"`
	Echo             bool    `json:"echo"`
	Stop             *string `json:"stop"`
	PresencePenalty  float64 `json:"presence_penalty"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	BestOf           int     `json:"best_of"`
}

type CompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []Choice
	Usage   Usage
}

type Choice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	LogProbs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (c *CompletionConfig) Default(prompt string) *CompletionConfig {
	c.Model = "text-davinci-003"
	c.Prompt = prompt
	c.Suffix = nil
	c.MaxTokens = 4000
	c.Temperature = 0.9
	c.TopP = 1
	c.N = 1
	c.Stream = false
	c.LogProbs = 0
	c.Echo = false
	c.Stop = nil
	c.PresencePenalty = 0
	c.FrequencyPenalty = 0
	c.BestOf = 1
	return c
}

func (c *CompletionConfig) Completion(r *utility.Request) (CompletionResponse, error) {
	cpr := CompletionResponse{}
	u, err := url.Parse(openAIUrl)
	u.Path = path.Join(u.Path, "completions")
	if err != nil {
		return cpr, err
	}
	r.Url = *u
	resp, err := r.Post(c)
	if err != nil {
		return cpr, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cpr, err
	}
	err = json.Unmarshal(body, &cpr)
	if err != nil {
		return cpr, err
	}
	return cpr, nil
}
