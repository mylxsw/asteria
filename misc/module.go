package misc

import (
	"strings"
)

// ModuleNameAbbr change a module name to abbreviation
func ModuleNameAbbr(moduleName string) string {
	segs := strings.Split(moduleName, ".")
	if len(segs) > 1 {
		ss := make([]string, 0)
		for _, s := range segs[:len(segs)-1] {
			if len(s) == 0 {
				continue
			}

			// avoid Chinese being cut into two paragraphs
			ss = append(ss, string(([]rune(s))[:1]))
		}

		moduleName = strings.Join(append(ss, segs[len(segs)-1]), ".")
	} else {
		moduleName = strings.Join(segs, ".")
	}
	return moduleName
}
