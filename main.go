package main

import (
	_ "Gf-Vben/internal/boot"

	_ "Gf-Vben/internal/logic"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {

	g.Server().Run()
}
