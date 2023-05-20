package myhttp

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func serveApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serveDebug() {
	if err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func serveApp2() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return http.ListenAndServe("0.0.0.0:8080", mux)
}

func serveDebug2() error {
	return http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp4(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return serve("0.0.0.0:8080", mux, stop)
}

func serveDebug4(stop <-chan struct{}) error {
	return serve("127.0.0.1:8001", http.DefaultServeMux, stop)
}

// func main() {
// 	done := make(chan error, 2)
// 	stop := make(chan struct{})
// 	go func() {
// 		done <- serveDebug(stop)
// 	}()
// 	go func() {
// 		done <- serveApp(stop)
// 	}()

// 	var stopped bool
// 	for i := 0; i < cap(done); i++ {
// 		if err := <-done; err != nil {
// 			fmt.Println("error: %v", err)
// 		}
// 		if !stopped {
// 			stopped = true
// 			close(stop)
// 		}
// 	}
// }
