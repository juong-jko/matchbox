package main

import (
	"context"
	"juong/matchbox/matchbox"
	"juong/matchbox/web"
)

func main() {
	ctx := context.Background()

	m := matchbox.NewMatchbox()
	s := web.NewServer(m)
	s.Start(ctx)
}
