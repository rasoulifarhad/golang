package myhttp

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, GopherCon SG")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func TestHttp2(t *testing.T) {
	go serveDebug()
	go serveApp()
	select {}
}

func TestHttp3(t *testing.T) {

	done := make(chan error, 2)

	go func() {
		done <- serveApp2()
	}()

	go func() {

		done <- serveDebug2()
	}()

	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v", err)
		}
	}
}

func TestHttp4(t *testing.T) {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug4(stop)
	}()
	go func() {
		done <- serveApp4(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}

}
