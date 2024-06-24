package main

import (
	"github.com/LumeWeb/tusc-s5/internal/client"
	"github.com/LumeWeb/tusc-s5/internal/server"
	"github.com/LumeWeb/tusc-s5/internal/util"
	"os"
)

const usage = `Usage:
  tusc (server|s) [options]
  tusc (client|c) <url> <file> [options]
  tusc --help`

func main() {
	if len(os.Args) < 2 {
		util.ExitWithMessages("No command", usage)
	}
	switch cmd := os.Args[1]; cmd {
	case "server", "s":
		server.Server()
	case "client", "c":
		client.Client()
	default:
		util.ExitWithMessages("Unknown command: "+cmd, usage)
	}
}
