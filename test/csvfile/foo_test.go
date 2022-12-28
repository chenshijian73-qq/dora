package csvfile

import (
	"flag"
	"fmt"
	"testing"
)

var foo string

func init() {
	flag.StringVar(&foo, "foo", "", "the foo bar bang")
}

func TestFoo(t *testing.T) {
	fmt.Println(foo)
}
