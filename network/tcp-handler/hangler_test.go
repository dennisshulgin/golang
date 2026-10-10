package tcphandler

import (
	"errors"
	"io"
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

func TestTcpHandlerBye(t *testing.T) {
	bufferSize := 100
	serverConn, clientConn := net.Pipe()
	var message string
	buffer := make([]byte, bufferSize)

	go func() {
		HandleConnection(serverConn)
	}()

	message = "quit\n"
	clientConn.Write([]byte(message))
	byteCount, _ := clientConn.Read(buffer)
	response := string(buffer[:byteCount])

	if response != "BYE\n" {
		t.Errorf("Expected BYE, but got %s", response)
	}

	message = "HELLO\n"
	count, err := clientConn.Write([]byte(message))

	if count != 0 || errors.Is(err, io.EOF) {
		t.Error("Connection is not closed")
	}
}

func TestTcpHandlerHelloPartially(t *testing.T) {
	bufferSize := 100
	serverConn, clientConn := net.Pipe()
	var message string
	buffer := make([]byte, bufferSize)

	go func() {
		HandleConnection(serverConn)
	}()

	message = "hel"
	clientConn.Write([]byte(message))
	message = "lo\n"
	clientConn.Write([]byte(message))

	byteCount, _ := clientConn.Read(buffer)
	response := string(buffer[:byteCount])

	if response != "HELLO\n" {
		t.Errorf("Expected HELLO, but got %s", response)
	}
}

func TestTcpHandlerCloseConnection(t *testing.T) {
	serverConn, clientConn := net.Pipe()

	go func() {
		err := HandleConnection(serverConn)

		if !errors.Is(err, io.EOF) {
			t.Error("Connection is not closed")
		}
	}()

	clientConn.Close()
}
