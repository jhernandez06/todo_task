package main

import (
	"context"
	"fmt"

	//"log"
	"os"

	"TodoList"
	_ "TodoList/app/models"
	_ "TodoList/app/tasks"

	"github.com/wawandco/oxpecker/cli"
	"github.com/wawandco/oxpecker/tools/soda"
)

// main function for the tooling cli, will be invoked by Oxpecker
// when found in the source code. In here you can add/remove plugins that
// your app will use as part of its lifecycle.
func main() {
	cli.Use(soda.Plugins(TodoList.Migrations)...)
	err := cli.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("[error] %v \n", err.Error())

		os.Exit(1)
	}
}
