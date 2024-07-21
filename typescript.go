package typescript

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/buke/quickjs-go"
	"github.com/jcbhmr/go-typescript/internal/ctx"
)

var ts quickjs.Value

func init() {
	url, err := filepath.Abs("node_modules/typescript/lib/typescript.js")
	if err != nil {
		panic(err)
	}
	url = "file://" + strings.ReplaceAll(url, "\\", "/")
	_, err = ctx.LoadCommonJSFile("node_modules/typescript/lib/typescript.js", url)
	if err != nil {
		panic(err)
	}
	p, err := ctx.Eval(fmt.Sprintf("import(%q)", url))
	if err != nil {
		panic(err)
	}
	ts, err = ctx.Await(p)
	if err != nil {
		panic(err)
	}
}

func VersionMajorMinor() string {
	return ts.Get("versionMajorMinor").String()
}
