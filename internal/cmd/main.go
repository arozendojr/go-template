package main

import (
	"github.com/arozendojr/go-rest/internal/application/rest"
	"github.com/arozendojr/go-rest/internal/infra/conf"
)

func main() {
	conf.Load()
	rest.Start()
}
