package euclid_test

import (
	"errors"
	"testing"

	"github.com/RossMerr/playground/euclid"
)

func Test_Parser(t *testing.T) {
	err, _ := euclid.Parser("hello world")

	if err != nil {
		t.Error(err)
	}
}

func Test_Parser_2(t *testing.T) {
	err, _ := euclid.Parser("import world")

	if err == nil {
		t.Error(err)
	}
}

func Test_Parser_Operation(t *testing.T) {
	err, stack := euclid.Parser("1 + 1")

	if err != nil {
		t.Error(err)
	}

	if stack != "2" {
		t.Error(errors.New("Expected 2 found" + stack))
	}
}
