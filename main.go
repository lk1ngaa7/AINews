/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"buzzGen/cmd"
	"buzzGen/helpers"
)

func main() {
	helpers.InitResource()
	cmd.Execute()
}
