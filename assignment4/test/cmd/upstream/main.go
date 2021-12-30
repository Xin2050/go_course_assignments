package main

import (
	"github.com/Xin2050/go_course_assignments/s4/test/server"
	"time"
)

func main() {
	server.NewUpStreamServer(
		10,
		50,
		0.8,
		time.Second*5,
	).Run(":9000")
}
