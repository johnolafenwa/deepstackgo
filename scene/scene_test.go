package scene 

import "testing"
import "github.com/johnolafenwa/deepstackgo/response"

func TestScene(t *testing.T){

	client := New("http://test.johnolafenwa.me","")

	expected := "conference_room"

	response := client.GetScene("scene.jpg")

	if response.Label != expected {

		t.Errorf("Scene error, expected %s got %s",expected,response.Label)

	}

}


func TestSceneAsync(t *testing.T){

	client := New("http://test.johnolafenwa.me","")

	expected := "conference_room"

	response_chan := make(chan response.SceneResponse)

	client.GetSceneAsync(response_chan,"scene.jpg")

	result := <-response_chan

	if result.Label != expected {

		t.Errorf("Scene error, expected %s got %s",expected,result.Label)

	}

}

