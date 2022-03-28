package workingdirectory

import (
	"fmt"
	"os"
)

func GetWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}
