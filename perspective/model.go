// gentleman - Исползьуя Perspective для определения токсичных комментарий, что упрощает проведение более качественных бесед в Интернете.
package perspective

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
