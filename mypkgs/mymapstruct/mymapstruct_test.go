package mymapstruct

import (
	"github.com/mitchellh/mapstructure"
	"testing"
	"reflect"
)

type Person struct {
	name string
	sex string
}

func test_mapstructure(t *testing.T){
	reflect.Zero(Person.type)
}