package test

import (
	"io/ioutil"
	"testing"
)

func LoadFixture(t *testing.T, name string) []byte {
	bytes, err := ioutil.ReadFile("./testdata/" + name)
	if err != nil {

		t.Fatal(err)
	}
	return bytes
}
