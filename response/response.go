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

type FaceDetectionItem struct {

	Xmin int `json:"x_min"`
	Ymin int `json:"y_min"`
	Xmax int `json:"x_max"`
	Ymax int `json:"y_max"`
	Confidence float32 `json:"confidence"`
}

type FaceDetectionResponse struct {

	Success bool `json:"success"`
	Error string `json:"error"`
	Predictions []FaceDetectionItem `json:"predictions"`
	
}
