package typescript

import (
	_ "embed"
)

//go:generate npm pack typescript@5.5.3
//go:generate mv typescript-5.5.3.tgz typescript.tgz
//go:embed typescript.tgz
var TypescriptTGZ []byte
