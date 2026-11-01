package tcphandler

import (
	"errors"
	"net"
)

func Serve(listener net.Listener) error {
	defer listener.Close()

	for {
		conn, listenerErr := listener.Accept()

		if listenerErr != nil {
			if errors.Is(listenerErr, net.ErrClosed) {
				return nil
			}
			return listenerErr
		}

		go func() {
			HandleConnection(conn)
		}()
	}
}
