package main

import "github.com/johnearl92/xendit-account-service/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
