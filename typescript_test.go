package typescript_test

import (
	"fmt"
	"testing"

	"github.com/jcbhmr/go-typescript"
)

func TestVersionMajorMinor(t *testing.T) {
	v := typescript.VersionMajorMinor()
	fmt.Println(v)
}
