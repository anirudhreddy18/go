package main

import "io"

import "os"

func main() {

	f, _ := os.OpenFile(os.Args[1], os.O_RDWR, 0755)
	io.Copy(os.Stdout, f)
}
