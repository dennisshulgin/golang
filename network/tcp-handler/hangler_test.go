package tcphandler

import (
	"net"
	"testing"
)

func TestTcpHandlerHello(t *testing.T) {
	bufferSize := 100
	serverConn, clientConn := net.Pipe()
	var message string
	buffer := make([]byte, bufferSize)

	go func() {
		HandleConnection(serverConn)
	}()

	message = "hello\n"
	clientConn.Write([]byte(message))
	byteCount, _ := clientConn.Read(buffer)
	response := string(buffer[:byteCount])

	if response != "HELLO\n" {
		t.Errorf("Expected HELLO, but got %s", response)
	}
}
