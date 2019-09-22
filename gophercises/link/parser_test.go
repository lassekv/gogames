package link

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadExercise1(t *testing.T) {
	f, err := os.Open("./tests/ex1.html")
	if err != nil {
		t.Error("Unable to read test file")
	}
	res, err := ParseHTML(f)
	if err != nil {
		t.Error("Unable to parse HTML")
	}
	if len(res) != 1 {
		t.Errorf("Returned the wrong number of entries. It returned %d", len(res))
	}
	if res[0].Href != "/other-page" {
		t.Errorf("Wrong url returned %v", res[0].Href)
	}
	if res[0].Text != "A link to another page" {
		t.Errorf("Wrong text returned %v", res[0].Text)
	}
}

func TestReadExercise2(t *testing.T) {
	f, err := os.Open("./tests/ex2.html")
	if err != nil {
		t.Error("Unable to read test file")
	}
	res, err := ParseHTML(f)
	if err != nil {
		t.Error("Unable to parse HTML")
	}
	assert.Equal(t, 2, len(res))
	assert.Equal(t, "https://www.twitter.com/joncalhoun", res[0].Href)
	assert.Equal(t, "Check me out on twitter", res[0].Text)
	assert.Equal(t, "https://github.com/gophercises", res[1].Href)
	assert.Equal(t, "Gophercises is on Github!", res[1].Text)
}
