package face

import "github.com/johnolafenwa/deepstackgo/connection"
import "github.com/johnolafenwa/deepstackgo/utils"
import "github.com/johnolafenwa/deepstackgo/response"

type Face struct {

	client connection.Connection
	apiKey string
	adminKey string
}

