package result

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	i := Just("string", nil).Map(sconvert)
	t.Log(i)
	fmt.Printf("%T", i)
}

var sconvert = func(s string) (string, error) {
	return "", nil
}
