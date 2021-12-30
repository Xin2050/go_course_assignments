package main

import "github.com/Xin2050/go_course_assignments/s4/test/server"

func main() {
	server.NewDownStreamServer(0.2).Run(":8000")
}
