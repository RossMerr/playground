package euclid_test

import (
	"testing"

	"github.com/RossMerr/playground/euclid"
)

func Test_Parser(t *testing.T) {
	err := euclid.Parser("hello world")

	if err != nil {
		t.Error(err)
	}
}

func Test_Parser_2(t *testing.T) {
	err := euclid.Parser("import world")

	if err != nil {
		t.Error(err)
	}
}
