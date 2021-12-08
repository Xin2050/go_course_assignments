package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func fn() error {
	e1 := sql.ErrNoRows
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func main() {
	err := fn()
	fmt.Printf("%+v\n", err)
}
