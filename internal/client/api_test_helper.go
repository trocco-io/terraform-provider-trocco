package client

import (
	"reflect"
	"testing"
)

type Case struct {
	name     string
	value    interface{}
	expected interface{}
}

func testCases(t *testing.T, cases []Case) {
	t.Helper()
	for _, c := range cases {
		value := c.value
		if c.expected == nil {
			if !reflect.ValueOf(value).IsNil() {
				t.Errorf("Expected %s to be nil, got %v", c.name, value)
			}
			continue
		}
		if reflect.ValueOf(value).Kind() == reflect.Ptr {
			value = reflect.ValueOf(value).Elem().Interface()
		}
		if c.expected != value {
			t.Errorf("Expected %s to be %v, got %v", c.name, c.expected, value)
		}
	}
}
