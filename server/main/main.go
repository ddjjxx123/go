package main

import "github.com/ddjjxx123/go/server/impl"

func main() {
	impl.CreateServer("X-Server").Serve()
}
