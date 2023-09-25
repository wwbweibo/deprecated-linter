package depractedlinter

import (
	"go/token"
	"golang.org/x/tools/go/analysis"
	"regexp"
	"strings"
)

func commonLint(pos token.Pos, pass analysis.Pass, name, text, currentVersion string) {
	regex, _ := regexp.Compile("//\\s*Deprecated:\\s*")
	if regex.MatchString(text) {
		text = regex.ReplaceAllString(text, "")
		// 检查是否指定了删除的版本
		regex, _ = regexp.Compile("^(v|V)+\\d+.\\d+.\\d+")
		if regex.MatchString(text) {
			version := regex.FindString(text)
			if strings.ToLower(version) <= strings.ToLower(currentVersion) {
				pass.Reportf(pos, "%s was marked as deprecated, and will be removed at version %s, which it early than current version %s", name, version, currentVersion)
			}
		} else {
			pass.Reportf(pos, "%s was marked as deprecated, but not specified the version will delete", name)
		}
	}
}
