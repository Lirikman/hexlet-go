package main

import (
	"github.com/fatih/color"
	"github.com/Lirikman/hexlet-go/greeting"
	)

func main() {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	boldRed.Println(greeting.Hello())
	}
