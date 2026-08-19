package selecttimeout

import (
	"testing"
	"time"
)

func TestReceiveWithTimeoutMessageExpected(t *testing.T) {
	ch := make(chan int, 1)
	duration := 3 * time.Second

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
	}()

	a, b := ReceiveWithTimeout(ch, duration)

	if a != 1 || !b {
		t.Error("Expected message")
	}
}

func TestReceiveWithTimeoutNoMessage(t *testing.T) {
	ch := make(chan int, 1)
	duration := 1 * time.Second

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	a, b := ReceiveWithTimeout(ch, duration)

	if a != 0 || b {
		t.Error("No expected message")
	}
}

func TestFirstResultFirst(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	duration := 2 * time.Second

	go func() {
		ch1 <- "Hello First"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Hello Second"
	}()

	a, b := FirstResult(ch1, ch2, duration)

	if a != "Hello First" || !b {
		t.Error("Expected First")
	}
}

func TestFirstResultSecond(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	duration := 2 * time.Second

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Hello First"
	}()

	go func() {
		ch2 <- "Hello Second"
	}()

	a, b := FirstResult(ch1, ch2, duration)

	if a != "Hello Second" || !b {
		t.Error("Expected First")
	}
}

func TestFirstResultNoResult(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	duration := 2 * time.Second

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Hello First"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Hello Second"
	}()

	a, b := FirstResult(ch1, ch2, duration)

	if a != "" || b {
		t.Error("No expected")
	}
}
