package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
    printf("%s\n", s);
}
*/
import "C"

func main() {
    C.SayHello(C.CString("Hello, cgo."))
}
