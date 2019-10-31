package typedefs

var ROOT string

type File struct {
	Name    string
	Content string
}

func init() {
	ROOT = "/"
}
