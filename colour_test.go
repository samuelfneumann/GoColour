package gocolour

import "testing"

func TestColorize(t *testing.T) {
	if Red("red") != "\033[31mred\033[0m" {
		t.Fail()
	}

	if Green("green") != "\033[32mgreen\033[0m" {
		t.Fail()
	}

	if Blue("blue") != "\033[34mblue\033[0m" {
		t.Fail()
	}
}
