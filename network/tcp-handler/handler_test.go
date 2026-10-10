package tcphandler

import (
	"net"
	"sync"
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

	clientConn.Close()
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

	if count != 0 || err == nil {
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

	clientConn.Close()
}

func TestTcpHandlerCloseConnection(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	var wg sync.WaitGroup

	wg.Go(func() {
		err := HandleConnection(serverConn)

		if err != nil {
			t.Error("Connection is not closed")
		}
	})

	clientConn.Close()

	wg.Wait()
}

func TestTcpHandlerFewCommands(t *testing.T) {
	bufferSize := 100
	serverConn, clientConn := net.Pipe()
	var message string
	buffer := make([]byte, bufferSize)

	go func() {
		HandleConnection(serverConn)
	}()

	message = "hello\nbye\n"
	clientConn.Write([]byte(message))

	byteCount, _ := clientConn.Read(buffer)
	response := string(buffer[:byteCount])

	if response != "HELLO\n" {
		t.Errorf("Expected HELLO, but got %s", response)
	}

	byteCount2, _ := clientConn.Read(buffer)
	response2 := string(buffer[:byteCount2])

	if response2 != "BYE\n" {
		t.Errorf("Expected BYE, but got %s", response2)
	}

	clientConn.Close()
}
