package main

import (
    "os"
    "text/template"
)

func main() {
    temp := "a is {{ .A }}, b is {{ .B }}, c is {{ .C }}"
    data := map[string]interface{} {
        "A": "A", "B": "B", "C": "C",
    }
    t, _ := template.New("test").Parse(temp)
    t.Execute(os.Stdout, data)
}
