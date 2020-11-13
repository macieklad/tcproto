package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/macieklad/tcproto/pkg/proto"
	"github.com/spf13/cobra"
)

var port string

var rootCmd = &cobra.Command{
	Use:   "-p [port]",
	Short: "Simple tcp server for accepting connections and serving messages",
	Run: func(cmd *cobra.Command, args []string) {
		runHub()
	},
}

// Execute fires the commands which starts the tcproto server
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "Port on which hub should listen")
}

func runHub() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("Error enncountered while starting the hub: %v", err)
	}

	log.Printf("Starting the hub on port %s", port)
	hub := proto.NewHub()
	go hub.Run()
	log.Printf("Hub started, waiting for connections...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%v", err)
		}

		c := hub.MakeClient(conn)
		go c.Read()
	}
}
