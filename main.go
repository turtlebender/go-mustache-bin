package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/cbroglie/mustache"

	"io/ioutil"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("Usage:")
		fmt.Println("go-mustache-bin <contextFile> <templateFile>")
		os.Exit(1)
	}
	context, err := ioutil.ReadFile(os.Args[1])
	inputFile := os.Args[2]
	var f interface{}
	err = json.Unmarshal(context, &f)
	if err != nil {
		panic("Unable to load context: " + err.Error())
	}
	out, err := mustache.RenderFile(inputFile, f)
	if err != nil {
		panic("Unable to render template: " + err.Error())
	}
	fmt.Println(out)
}