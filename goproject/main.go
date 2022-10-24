package main

import (
	"fmt"
	"goproject/helper"
)

func main() {
	fmt.Println(getStartMessage())
	fmt.Printf("Project Name: %v\n", helper.GetProjectName())
	fmt.Print("Version String: ", helper.GetVersionString(), "\n")
	fmt.Printf("Version: %v\n", helper.Version)
	//fmt.Printf("Version:%v\n", helper.buildNumber)	// not allow
}

func getStartMessage() string {
	return "Hello Go Project"
}
