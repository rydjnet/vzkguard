// perspective - Исползьуя Perspective для определения токсичных комментарий, что упрощает проведение более качественных бесед в Интернете.

package perspective

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"vzkguard/config"
)

func ToxicCheker(tmsg string) float64 {
	var toxicScore float64
	url := "https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze?key=" + config.PerspectiveToken
	request := RequestBody{
		Comment: Comment{
			Text: tmsg,
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
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		log.Println("Error perspectAPI: ", resp.Status, string(b))
		return 0
	}
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
