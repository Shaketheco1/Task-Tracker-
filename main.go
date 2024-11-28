package main

import (
	"fmt"

	"github.com/Shaketheco1/Task-Tracker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
