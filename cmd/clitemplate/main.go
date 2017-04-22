package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/whyrv/clitemplate"
	"github.com/whyrv/spin"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "clitemplate"
	app.Version = version
	app.Author = "Carlos Alexandro Becker (root@carlosbecker.com)"
	app.Usage = "This is an clitemplate app written in Go"
	app.Action = func(c *cli.Context) error {
		spin := spin.New("\033[36m %s Working...\033[m")
		spin.Start()
		err := clitemplate.Foo()
		spin.Stop()
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	}
	_ = app.Run(os.Args)
}
