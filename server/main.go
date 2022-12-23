package main

import (
	"github.com/ssibrahimbas/ssi-we/src/app"
)

func main() {
	app.New().Prepare().Listen()
}
