package util

import (
	"frontendmod/controller"
	"frontendmod/middleware"
)

func ExecTemplates() {
	controller.ExecTemplates()
	middleware.ExecTemplates()
}
