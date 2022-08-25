package main

// Build a server that accepts incoming TCP connections.
// For each connection, read URLs, delimited by a newline character "\n".
// For each URL, request it and write both the URL and the response status code (e.g. `200`) to a new line in a file on disk.
// Example
// Feeding some URLs to a server running on 127.0.0.1 6666 with netcat
// ncat 127.0.0.1 6666 < <(echo -e  "https://www.google.com/search?ei=9vImXZvBEYzSa7DNqcAI&q=MessageBird&oq=MessageBird\nhttps://messagebird.com/en/#\nhttps://github.com/yandex/ClickHouse/blob/master/dbms/programs/server/users.xml\nhttp://www.google.com")
// Expected file results.csv
// http://www.google.com/search?ei=9vImXZvBEYzSa7DNqcAI&q=MessageBird&oq=MessageBird,200
// https://messagebird.com/en/#,200
// https://github.com/yandex/ClickHouse/blob/master/dbms/programs/server/users.xml,200

import (
	"fmt"
	"log"
	"net"
	"os"
	"github.com/bagadatp/message_bird/pkg/handler"
	"time"
)
var (
	version = "v0.0.1-dev"
	outputFile = "result.csv"
)

const (
	cfgUsage =`
		-config option or ENV vars
`
	MAX_TIMEOUT_SECS = 10
)

func main() {
	l := log.New(os.Stdout, fmt.Sprintf("url-getter/%d ", os.Getpid()), 0)
	/*
	var (
		cfg = flag.NewFlagSet("", flag.ExitOnError)
		cfgPath = cfg.String("config", "", "path to config file")
		httpAddr = cfg.String("http", ":http", "listen on this addr:port")
	)
	cfg.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Url Getter Srv %s", version)
		cfg.PrintDefaults()
		_, _ = fmt.Fprintln(os.Stderr, cfgUsage)
	}
	*/

	address := ":6666"
	ln, err := net.Listen("tcp", address)
	if err != nil {
		l.Fatalf("Could not start listening")
	}
	l.Printf("Server started on %v\n", address)

	for {
		conn, err := ln.Accept()
		conn.SetReadDeadline(time.Now().Add(MAX_TIMEOUT_SECS * time.Second))
		if err != nil {
			l.Printf("Error accepting connection %v", err)
		}
		go func() {
			err := handler.HandleUrls(conn, outputFile)
			if err != nil {
				l.Printf("Error handling connection: %v", err)
				conn.Close()
			}
		}()
	}
}
