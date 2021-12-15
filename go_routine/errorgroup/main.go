package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	var a, b, c []int

	g.Go(func() error {
		a = []int{1, 2, 2}
		return errors.New("test")
	})

	g.Go(func() error {
		b = []int{1, 2, 3}
		return nil
	})

	g.Go(func() error {
		c = []int{3, 2, 1}
		return nil
	})

	err := g.Wait()
	fmt.Println(a, b, c)
	fmt.Println(err)
	fmt.Println(ctx.Err())
}
