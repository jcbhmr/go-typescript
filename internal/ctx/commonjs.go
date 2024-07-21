package ctx

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/buke/quickjs-go"
)

func LoadCommonJS(code string, moduleName string) (quickjs.Value, error) {
	requirePreimports := []string{}
	matches := regexp.MustCompile(`require\(["'](.*?)["']\)`).FindAllStringSubmatch(code, -1)
	for _, match := range matches {
		specifier, err := strconv.Unquote(match[1])
		if err != nil {
			continue
		}
		requirePreimports = append(requirePreimports, specifier)
	}

	exportNames := []string{}
	matches = regexp.MustCompile(`module\.exports\s*=\s*\{([\s\S]*?)\}`).FindAllStringSubmatch(code, -1)
	for _, match := range matches {
		propList := strings.Split(match[1], ",")
		for _, prop := range propList {
			prop = strings.TrimSpace(prop)
			if prop == "" {
				continue
			}
			exportNames = append(exportNames, prop)
		}
	}

	js := ""
	for i, requirePreimport := range requirePreimports {
		js += fmt.Sprintf("import * as __require$%d from %q;\n", i, requirePreimport)
	}
	js += "const __requireMap = { __proto__: null };\n"
	for i, requirePreimport := range requirePreimports {
		js += fmt.Sprintf("__requireMap[%q] = __require$%d;\n", requirePreimport, i)
	}
	js += "var require = (specifier) => {\n"
	js += "  if (specifier in __requireMap) {\n"
	js += "    const m = __requireMap[specifier];\n"
	js += "    if ('__esModule' in m && !m.__esModule) {\n"
	js += "      return m.default;\n"
	js += "    } else {\n"
	js += "      return m;\n"
	js += "    }\n"
	js += "  } else {\n"
	js += "    throw new Error(`Cannot find module ${specifier}`);\n"
	js += "  }\n"
	js += "};\n"
	js += "var module = { exports: {} };\n"
	js += "var exports = module.exports;\n"
	js += code + "\n"
	for i, exportName := range exportNames {
		js += fmt.Sprintf("const __exports$%d = module.exports.%s;\n", i, exportName)
	}
	for i, exportName := range exportNames {
		js += fmt.Sprintf("export { __exports$%d as %s };\n", i, exportName)
	}
	js += "export default module.exports;\n"

	return LoadModule(js, moduleName)
}

func LoadCommonJSFile(filePath string, moduleName string) (quickjs.Value, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return Null(), err
	}
	return LoadCommonJS(string(b), moduleName)
}
