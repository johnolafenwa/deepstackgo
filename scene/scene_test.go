package scene 

import "testing"

func TestScene(t *testing.T){

	client := New("http://test.johnolafenwa.me","")

	expected := "conference_room"

	response := client.GetScene("scene.jpg")

	if response.Label != expected {

		t.Errorf("Scene error, expected %s got %s",expected,response.Label)

	}



}