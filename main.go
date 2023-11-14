package main

import (
	"context"
	"fmt"

	"github.com/LukasKoller/cicd-testing/entities"
	"github.com/LukasKoller/cicd-testing/builder"
)

func main() {
	buildServerVariables := FromBuildServer()
	if err := builder.Build(context.Background(), buildServerVariables); err != nil {
		fmt.Println(err)
	}
}

func FromBuildServer() BuildServerVariables {
	return BuildServerVariables {
		RefType:         mustLoadEnv("REF_TYPE"),
		RefName:         mustLoadEnv("REF_NAME"),
}
