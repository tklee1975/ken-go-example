package helper

import "fmt"

func GetProjectName() string {
	return "Very Simple Go Project"
}

func GetVersionString() string {
	return fmt.Sprintf("%v (%v)", Version, buildNumber)
}
