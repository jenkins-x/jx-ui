package main

import (
	"fmt"
	"os"

	"jx-ui/internal/server"
)

func main() {
	s := &server.Server{}
	if err := s.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
