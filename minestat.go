package main

import (
	"flag"
	"fmt"

	"github.com/FragLand/minestat/Go/minestat"
)

var (
	serverIPorDNS string
	playersOnly   bool
	versionOnly   bool
	debug         bool
)

func main() {
	flag.StringVar(&serverIPorDNS, "serverIPorDNS", "default", "help message")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.BoolVar(&playersOnly, "playersOnly", false, "Display only players")
	flag.BoolVar(&versionOnly, "versionOnly", false, "Display only version")
	flag.Parse()

	if debug {
		fmt.Printf("Debug mode enabled\n")
		fmt.Println("Server IP or DNS:", serverIPorDNS)
		fmt.Println("Players only:", playersOnly)
		fmt.Println("Version only:", versionOnly)
	}

	minestat.Init(serverIPorDNS)

	if minestat.Online {
		if playersOnly && !versionOnly {
			fmt.Printf("%d / %d \n", minestat.Current_players, minestat.Max_players)
		}

		if versionOnly && !playersOnly {
			fmt.Printf("%s\n", minestat.Version)
		}

		if !playersOnly && !versionOnly {
			fmt.Printf("Version: %s - Players: %d / %d \n", minestat.Version, minestat.Current_players, minestat.Max_players)
			fmt.Printf("Latency: %dms\n", minestat.Latency)
		}
	} else {
		fmt.Println("Server is offline!")
	}
}
