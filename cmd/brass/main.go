package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mikerybka/brass"
	"github.com/mikerybka/util"
)

func main() {
	if len(os.Args) < 2 {
		help("")
		return
	}
	cmd := os.Args[1]
	if cmd == "console" || cmd == "c" {
		console()
		return
	}
	if cmd == "server" || cmd == "s" || cmd == "start" {
		port := util.EnvVar("PORT", "3000")
		startServer(port)
		return
	}
	if cmd == "help" {
		if len(os.Args) > 2 {
			page := os.Args[2]
			help(page)
		}
		help("")
		return
	}
	help("")
}

func startServer(port string) {
	s := brass.NewServer()
	http.ListenAndServe(":"+port, s)
}

func console() {
	fmt.Println(">")
}

func help(page string) {
	if page == "" {
		msg := fmt.Sprintf(`Usage: %s [server|console]`, os.Args[0])
		fmt.Println(msg)
		return
	}
	fmt.Println("help page not found:", page)
}
