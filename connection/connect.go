package connection

import "github.com/imroc/req"
import "github.com/johnolafenwa/deepstackgo/utils"
import "os"
import "path/filepath"

type connection struct {

	URL string
}

func New(url string) connection{

	c := connection{URL: url}
	return c
	
}

func (c connection) getResponse(endpoint string,out interface{}, params utils.Request, files utils.Files  ) interface{} {

	 var uploads []req.FileUpload
	
	 for _,f := range files.Files {
			f_data,_ := os.Open(f.Filepath)
			uploads = append(uploads, req.FileUpload{

				File: f_data,
				FieldName: f.Name,
				FileName: filepath.Base(f.Filepath),

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