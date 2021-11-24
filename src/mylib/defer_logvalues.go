package mylib

import (
	"io"
	"log"
)

func Logvalues(s string) (n int, err error) {
	defer func() {
		log.Printf("Logvalues(%q)=%d,%v", s, n, err)
	}()
	return 7, io.EOF
}

func Last() {

}
