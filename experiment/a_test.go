package experiment

import (
	"fmt"
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	a := "use abs"
	index := strings.Index(a, "use")
	fmt.Println(index)
}
