package main

import (
	"design-pattern/oneinstance/service"
)

func main() {
	ins := service.Instance
	ins.Driver()
}
