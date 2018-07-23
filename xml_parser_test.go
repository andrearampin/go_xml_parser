package XMLParser

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

// Parser should return "hello, <input>"
func TestParser(t *testing.T) {
	xmlMessage, err := ioutil.ReadFile("request.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	parsed := Parser(xmlMessage)
	expected := "Pasta price $15, Pizza price $20.95"
	if strings.Compare(parsed.ToString(), expected) != 0 {
		t.Errorf("got %q, want %s", parsed, expected)
	}
}
