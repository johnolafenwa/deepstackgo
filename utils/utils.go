package utils

type Request struct {

	Params map[string]string

}

type Files struct {

	Files []File

}

type File struct {

	Filepath string
	Name string

}

