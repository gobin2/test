package test

import (
	"fmt"

	"github.com/gobin2/test/tmptest"
)

func Mainfunc() {
	fmt.Println("test main func")
	tmptest.Remotetestfunc()
}
