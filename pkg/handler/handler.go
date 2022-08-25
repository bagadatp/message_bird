package handler

import (
	"bufio"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	CLIENT_MAX_TIMEOUT = 15
)

func HandleUrls(conn net.Conn, file string) error {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	for {
		if !scanner.Scan() {
			return scanner.Err()
		}
		nextUrl := scanner.Text()
		client := http.Client{
			Timeout: CLIENT_MAX_TIMEOUT * time.Second,
		}
		resp, err := client.Get(nextUrl)
		var code int
		if err != nil {
			code = -1
		}
		code = resp.StatusCode
		writer.WriteString(nextUrl + "," + strconv.Itoa(code) + "\n")
	}

	return nil
}