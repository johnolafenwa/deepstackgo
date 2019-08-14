package connection

import "github.com/imroc/req"
import "github.com/johnolafenwa/deepstackgo/utils"
import "os"
import "path/filepath"

type Connection struct {

	URL string
}

func New(url string) Connection{

	c := Connection{URL: url}
	return c
	
}

func (c Connection) GetResponse(endpoint string,out interface{}, params utils.Request, files utils.Files  ) interface{} {

	 var uploads []req.FileUpload
	
	 for _,f := range files.Files {
			f_data,_ := os.Open(f.Path)
			uploads = append(uploads, req.FileUpload{

				File: f_data,
				FieldName: f.Name,
				FileName: filepath.Base(f.Path),

			})
	 }

	 req_params := req.Param{}
	 
	 for k,v := range params.Params {

		req_params[k] = v
	 }

	 output,err := req.Post(c.URL + "/" + endpoint,req_params,uploads)

	 if err != nil {
		 panic(err)
	 }

	 output.ToJSON(&out)

	 return out

}