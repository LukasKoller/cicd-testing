package main

import (
	"context"
	"fmt"
)

type BuildServer struct {
	RefType string
	RefName string
}

func main() {
	if err := builder.build(context.Background()); err != nil {
		fmt.Println(err)
	}
}
