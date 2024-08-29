package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "The name to greet.")
	flag.Parse()
	fmt.Printf(Hello(*name))
}

func Hello(name string) string {
  return fmt.Sprintf("Hello, %s!", name)
}
