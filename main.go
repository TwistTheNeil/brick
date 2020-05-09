package main

import (
	"brick/cmd"
	"brick/utils"
)

func main() {
	utils.GetPublicIPDetails()
	cmd.Execute()
}
