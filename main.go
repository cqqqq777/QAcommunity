package main

import (
	"main/services/boot"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.DataBaseSetup()
	boot.InitRouters()
}
