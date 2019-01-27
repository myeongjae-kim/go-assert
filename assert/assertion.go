package assert

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, expected, actual interface{}) bool {
	if expected != actual {
		t.Errorf("inputs are not equal. expect: %s, actual: %s",
			reflect.ValueOf(expected).String(), reflect.ValueOf(actual).String())
		return false
	}
	return true
}
