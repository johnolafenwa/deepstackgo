package utils

type Request struct {

	Params map[string]string

}

type Files struct {

	Files []File

}

type File struct {

	Path string
	Name string

}

