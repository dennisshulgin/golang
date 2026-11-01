package tcphandler

import (
	"errors"
	"net"
	"sync"
	"testing"
	"time"
)

var TestErr = errors.New("Error")

type CommandReqResp struct {
	Request  string
	Response string
}

type TestListener struct{}

func (l TestListener) Accept() (net.Conn, error) {
	return nil, TestErr
}

func (l TestListener) Close() error {
	return nil
}

func (l TestListener) Addr() net.Addr {
	return nil
}

func TestServeOneClient(t *testing.T) {
	listener, listenerError := net.Listen("tcp", "127.0.0.1:0")

	if listenerError != nil {
		t.Fatalf("Unexpected error %s", listenerError)
	}

	ch := make(chan error, 1)

	go func() {
		serverErr := Serve(listener)
		ch <- serverErr
	}()

	commands := []CommandReqResp{
		{Request: "hello\n", Response: "HELLO\n"},
	}

	createClientAndSendCommands(listener, commands, t)
	listener.Close()

	select {
	case msg := <-ch:
		if msg != nil {
			t.Errorf("Unexpected error %s", msg)
		}
	case <-time.After(10 * time.Second):
		t.Error("Timeout")
	}
}

func TestServeTwoClients(t *testing.T) {
	listener, listenerError := net.Listen("tcp", "127.0.0.1:0")

	if listenerError != nil {
		t.Fatalf("Unexpected error %s", listenerError)
	}

	ch := make(chan error, 1)

	go func() {
		serverErr := Serve(listener)
		ch <- serverErr
	}()

	var wg sync.WaitGroup

	wg.Go(func() {
		commands := []CommandReqResp{
			{Request: "hello\n", Response: "HELLO\n"},
		}
		createClientAndSendCommands(listener, commands, t)
	})

	wg.Go(func() {
		commands := []CommandReqResp{
			{Request: "hello\n", Response: "HELLO\n"},
			{Request: "quit\n", Response: "BYE\n"},
		}
		createClientAndSendCommands(listener, commands, t)
	})

	wg.Wait()

	listener.Close()

	select {
	case msg := <-ch:
		if msg != nil {
			t.Errorf("Unexpected error %s", msg)
		}
	case <-time.After(10 * time.Second):
		t.Error("Timeout")
	}
}

func TestServeStopsAfterListenerClose(t *testing.T) {
	listener, listenerErr := net.Listen("tcp", "127.0.0.1:0")

	if listenerErr != nil {
		t.Fatalf("Unexpected error %s", listenerErr)
	}

	ch := make(chan error, 1)

	go func() {
		ch <- Serve(listener)
	}()

	listener.Close()

	select {
	case msg := <-ch:
		if msg != nil {
			t.Errorf("Unexpected error %s", msg)
		}
	case <-time.After(10 * time.Second):
		t.Error("Timeout")
	}
}

func TestServeReturnsAcceptError(t *testing.T) {
	listener := TestListener{}
	ch := make(chan error, 1)

	go func() {
		ch <- Serve(listener)
	}()

	listener.Close()

	select {
	case msg := <-ch:
		if !errors.Is(msg, TestErr) {
			t.Errorf("Expected error %s, but got %s", TestErr, msg)
		}
	case <-time.After(10 * time.Second):
		t.Error("Timeout")
	}
}

func createClientAndSendCommands(listener net.Listener, commands []CommandReqResp, t *testing.T) {
	bufferSize := 100
	buffer := make([]byte, bufferSize)
	conn, dialError := net.Dial("tcp", listener.Addr().String())

	if dialError != nil {
		t.Errorf("Unexpected error %s", dialError)
	}

	for _, command := range commands {
		_, writeError := conn.Write([]byte(command.Request))

		if writeError != nil {
			t.Errorf("Unexpected error %s", writeError)
		}

		readCount, readError := conn.Read(buffer)

		if readError != nil {
			t.Errorf("Unexpected error %s", readError)
		}

		response := string(buffer[:readCount])

		if response != command.Response {
			t.Errorf("Expected %s, but got %s", command.Response, response)
		}
	}

	conn.Close()
}
