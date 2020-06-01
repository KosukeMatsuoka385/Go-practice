package main

import (
	"fmt"
	"time"
)

func getOS() string {
	return "sdfgsrg"
}

func main(){
	switch os := getOS(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Other")

	}
}