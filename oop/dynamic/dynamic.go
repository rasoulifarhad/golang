package dynamic

import (
	"errors"
	"io"
	"syscall"
)

type readfd int

func (r readfd) Read(buf []byte) (int, error) {
	return syscall.Read(int(r), buf)
}

type writefd int

func (w writefd) Write(buf []byte) (int, error) {
	return syscall.Write(int(w), buf)
}

const (
	Stdin  = readfd(0)
	Stdout = writefd(1)
	Stderr = writefd(2)
)

var myEOF = errors.New("EOF")

type Error string

func (e Error) Error() string {
	return string(e)
}

const err = Error("EOF")

var w io.Writer = Stdout
