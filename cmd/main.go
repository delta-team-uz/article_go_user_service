package main

import (
	"article_user_service/pkg"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		pkg.Module,
	).Run()
}
