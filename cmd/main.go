package main

import (
	"context"
	"os"
	"strings"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	env := strings.TrimSpace(os.Getenv("ENVIRONMENT"))
	connStr := strings.TrimSpace(os.Getenv("CONN_STR"))

	app := &App{
		Ctx:     ctx,
		Cancel:  cancel,
		Env:     env,
		ConnStr: connStr,
	}

	app.startup()
}
