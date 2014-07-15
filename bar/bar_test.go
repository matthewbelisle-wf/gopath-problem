package bar

import (
        "testing"
	"bar"
)

func TestComposeNewsletter(t *testing.T) {
        if bar.Bar != 2 {
                t.Errorf("Bar does not equal 2")
        }
}
