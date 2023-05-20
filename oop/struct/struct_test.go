package struct2

import (
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	quit := make(chan struct{})
	go hello(1, quit)
	// print hello for 10 ms
	time.Sleep(10 * time.Millisecond)
	quit <- struct{}{} // quit printing hello
}

func TestXxx2(t *testing.T) {
	quit := make(chan struct{})
	go hello(1, quit)
	go hello(2, quit)
	// print hello for 10 ms
	time.Sleep(10 * time.Millisecond)
	quit <- struct{}{} // quit printing hello
	time.Sleep(10 * time.Millisecond)
}

func TestXxx3(t *testing.T) {
	quit := make(chan struct{})
	go hello(1, quit)
	go hello(2, quit)
	// print hello for 10 ms
	time.Sleep(10 * time.Millisecond)
	close(quit)
	time.Sleep(10 * time.Millisecond)
}
