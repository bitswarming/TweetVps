package main

import (
        "github.com/bitswarming/TweetVps/cmd"
        "fmt"
        "os"
)

func main() {
        if err := cmd.RootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(-1)
        }
}
