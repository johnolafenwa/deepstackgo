package scene

import "github.com/johnolafenwa/deepstackgo/connection"
import "github.com/johnolafenwa/deepstackgo/response"
import "github.com/johnolafenwa/deepstackgo/utils"

type Response struct{

	Success bool `json:"success"`
	Label string `json:"label"`
	Confidence float32 `json:"confidence"`
	Error string `json:"error"`
}

type scene struct {
	client connection.Connection
	apiKey string
}

func New(url string,apiKey string) scene {

	c := connection.New(url)
	
	return scene{client: c,apiKey: apiKey}
}

func (scene scene) GetScene(path string) response.SceneResponse{

	req_params := map[string]string{}
	req_params["api_key"] = scene.apiKey

	req_data := utils.Request{Params: req_params}

	file := utils.File{Name: "image",Path: path}
	filelist := []utils.File{file}
	file_data := utils.Files{Files: filelist}

	var result response.SceneResponse 

	scene.client.GetResponse("v1/vision/scene",&result,req_data,file_data)

	return result

}