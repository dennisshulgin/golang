package selecttimeout

import (
	"time"
)

func ReceiveWithTimeout(ch <-chan int, timeout time.Duration) (int, bool) {
	select {
	case msg := <-ch:
		return msg, true
	case <-time.After(timeout):
		return 0, false
	}
}

func FirstResult(
	first <-chan string,
	second <-chan string,
	timeout time.Duration,
) (string, bool) {
	select {
	case msg1 := <-first:
		return msg1, true
	case msg2 := <-second:
		return msg2, true
	case <-time.After(timeout):
		return "", false
	}
}
