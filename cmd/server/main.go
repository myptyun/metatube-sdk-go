package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/metatube-community/metatube-sdk-go/cmd"
	"github.com/metatube-community/metatube-sdk-go/engine"
	V "github.com/metatube-community/metatube-sdk-go/internal/version"
)

func showVersionAndExit() {
	fmt.Println(V.BuildString())
	os.Exit(0)
}

func main() {
	if _, isSet := os.LookupEnv("VERSION"); cmd.Config.VersionFlag &&
		!isSet /* NOTE: ignore this flag if ENV contains VERSION variable. */ {
		showVersionAndExit()
	}

	var (
		bindAddr = cmd.Config.Bind
	)
	// Default to 0.0.0.0 to ensure IPv4 listening in Docker containers
	if bindAddr == "" {
		bindAddr = "0.0.0.0"
	}
	var (
		addr   = net.JoinHostPort(bindAddr, cmd.Config.Port)
		router = cmd.Router(engine.DefaultEngineName)
	)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
