package slugify

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	s := "Test->àèâ<-Test"
	slug := Marshal(s)
	expected := "test-aea-test"
	if slug != expected {
		t.Fatal("Return string is not slugified as expected", expected, slug)
	}
}
