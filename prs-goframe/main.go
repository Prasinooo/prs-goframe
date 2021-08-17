package main

import (
	_ "prs-goframe/boot"
	_ "prs-goframe/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
