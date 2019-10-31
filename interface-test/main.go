package main

type HelloInterface interface {
	Hello(s string)
}

type WarnInterface interface {
	Warn(s string)
}

type TestStruct struct {
	Content string
}

func (test *TestStruct) Hello(s string) {
	(*test).Content = s
	return
}

func main() {
	var _ HelloInterface = (*TestStruct)(nil)
	var _ WarnInterface = (*TestStruct)(nil)
}
