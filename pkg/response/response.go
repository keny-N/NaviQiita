package response

type AlexaResponse struct {
	Version           string   `json:"version"`
	SessionAttributes struct{} `json:"sessionAttributes"`
	Response          struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
	} `json:"response"`
}

func CreateResponse(message string) AlexaResponse {
	return AlexaResponse{
		Version: "1.0",
		Response: struct {
			OutputSpeech struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"outputSpeech"`
		}{
			OutputSpeech: struct {
				Type string `json:"type"`
				Text string `json:"text"`
			}{
				Type: "PlainText",
				Text: message,
			},
		},
	}
}
