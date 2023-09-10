package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Language string

const (
	English  Language = "EN"
	Japanese Language = "JA"
)

//go:generate mockgen -source=deepl.go -destination=../../mock/deepl.go
type DeeplAPIGateway interface {
	TranslateTexts(texts []string, into Language) ([]string, error)
	TranslateTextsIntoJapanese(texts []string) ([]string, error)
}

type deeplAPIGateway struct {
}

func NewDeeplAPIGateway() DeeplAPIGateway {
	return &deeplAPIGateway{}
}

func (d *deeplAPIGateway) TranslateTextsIntoJapanese(texts []string) ([]string, error) {
	translated, err := d.TranslateTexts(texts, Japanese)
	if err != nil {
		return []string{}, err
	}
	return translated, nil
}

type TranslationResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func (d *deeplAPIGateway) TranslateTexts(texts []string, into Language) ([]string, error) {
	result := make([]string, len(texts))
	authKey := os.Getenv("DEEPL_API_KEY")
	apiUrl := "https://api-free.deepl.com/v2/translate"
	data := map[string]interface{}{
		"text":        texts,
		"target_lang": string(into),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("DeepL-Auth-Key %s", authKey))
	if err != nil {
		return result, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	var response TranslationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return result, err
	}
	for i, v := range response.Translations {
		result[i] = v.Text
	}
	return result, nil
}
