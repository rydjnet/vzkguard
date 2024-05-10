// gentleman - Исползьуя Perspective для определения токсичных комментарий, что упрощает проведение более качественных бесед в Интернете.
package housemd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Comment struct {
	Text string `json:"text"`
}

type RequestedAttributes struct {
	Toxicity map[string]interface{} `json:"TOXICITY"`
}

type RequestBody struct {
	Comment             Comment             `json:"comment"`
	Languages           string              `json:"languages"`
	RequestedAttributes RequestedAttributes `json:"requestedAttributes"`
}

type ResponseAttributes struct {
	Toxicity map[string]interface{} `json:"TOXICITY"`
}

type ResponseData struct {
	Attributes ResponseAttributes `json:"attributeScores"`
}

type Response struct {
	Data ResponseData `json:"attributeScores"`
}

func Gentleman(tmsg *TMessage) float64 {
	var toxicScore float64
	url := "https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key=AIzaSyBp4K5ywfML85nbJhmiBhyIA4Q43dHt3hM"
	request := RequestBody{
		Comment: Comment{
			Text: tmsg.Text,
		},
		Languages: "ru",
		RequestedAttributes: RequestedAttributes{
			Toxicity: map[string]interface{}{},
		},
	}
	requestBodyBytes, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		log.Println("Error creating request:", err)
		return toxicScore
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return toxicScore
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return toxicScore
	}

	toxicity := data["attributeScores"].(map[string]interface{})["TOXICITY"].(map[string]interface{})
	summaryScore := toxicity["summaryScore"].(map[string]interface{})
	toxicScore = summaryScore["value"].(float64)
	return toxicScore
}
