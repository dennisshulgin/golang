package tcphandler

import (
	"io"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	bufferSize := 100
	buffer := make([]byte, bufferSize)
	var requestBuilder strings.Builder

	for {
		byteCount, err := conn.Read(buffer)

		if err != nil && err != io.EOF {
			return err
		}

		commands := processBatch(buffer[:byteCount], &requestBuilder)

		for _, command := range commands {
			var writeError error

			if command == "QUIT\n" {
				_, writeError = conn.Write([]byte("BYE\n"))

				if writeError != nil {
					return writeError
				}

				return nil
			} else {
				_, writeError = conn.Write([]byte(command))

				if writeError != nil {
					return writeError
				}
			}

		}

		if err == io.EOF {
			return nil
		}
	}
}

func processBatch(batch []byte, requestBuilder *strings.Builder) []string {
	commands := make([]string, 0, 10)

	for _, b := range batch {
		if b == '\n' {
			requestBuilder.WriteByte(b)
			commands = append(commands, strings.ToUpper(requestBuilder.String()))
			requestBuilder.Reset()
		} else {
			requestBuilder.WriteByte(b)
		}
	}

	return commands
}
