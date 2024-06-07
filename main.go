// Copyright © 2024- Luka Ivanović
// This code is licensed under the terms of the MIT licence (see LICENCE for details).

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filepath, ok := os.LookupEnv("MARKTHISDATEFILE")
	if !ok {
		log.Fatalln("Environment variable $MARKTHISDATEFILE not set.")
	}
	file, err := os.OpenFile(filepath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0o600)
	if err != nil {
		log.Fatalln(err.Error())
	}
	date := time.Now().Format("Monday, 02 Jan 2006 15:04:05")
	args := os.Args[1:]
	message := strings.Join(args, " ")
	if _, err := file.WriteString(fmt.Sprintf( "%s: %s\n", date, message)); err != nil {
		log.Fatalln(err.Error())
	}
}
