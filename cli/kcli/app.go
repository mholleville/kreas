package kcli

import (
	"strings"
)

func GenerateAlias(module string) string {
	aliasMap := map[string]string{
		"terraform": "tf",
		"ansible": "as",
	}
	splittedModule := strings.SplitN(module, "-", 2)
	if val, ok := aliasMap[splittedModule[0]]; ok {
		suffix := ""
		if len(splittedModule) > 1 {
			suffix = splittedModule[1]
		}
		return val + suffix
	} else {
		return module
	}
}
