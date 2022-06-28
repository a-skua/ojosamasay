package main

import (
	"fmt"
	"github.com/a-skua/ojosamasay/ojosama"
	"io"
	"os"
)

func main() {
	bin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(ojosama.Say(string(bin)))
}
