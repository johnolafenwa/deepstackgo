package connection

import "testing"
import "github.com/johnolafenwa/deepstackgo/utils"

type Response struct{

	Success bool `json:"success"`
	Label string `json:"label"`
	Confidence float32 `json:"confidence"`
	Error string `json:"error"`
}


func TestGetResponse(t *testing.T){

	c := New("http://localhost:82")

	req_data := utils.Request{}
	file := utils.File{Name: "image", Filepath: "scene.jpg" }
	filelist := []utils.File{file}
	files := utils.Files{Files: filelist}
	
	response := Response{}

	c.getResponse("v1/vision/scene",&response,req_data,files)

	if response.Success != true {
		t.Errorf("Success %v is not %v, error is %s",response.Success,true,response.Error)
	}

	if response.Label != "conference_room" {

		t.Errorf("Label is %s, should be %s",response.Label,"conference_room")
	}
	
}