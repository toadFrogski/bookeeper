package main

import (
	"bookeeper/domain"
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
)

// Define the models to generate migrations for.
var models = []any{
	&domain.User{},
	&domain.Book{},
	&domain.Role{},
	&domain.Permission{},
	&domain.Checkout{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
