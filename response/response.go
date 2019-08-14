package response


type RecognitionResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	Label string `json:"label"`
	Confidence float32 `json:"confidence"`
}


type SceneResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
	Label string `json:"label"`
	Confidence float32 `json:"confidence"`
}

