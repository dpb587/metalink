package utility

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var semisemverLetter = regexp.MustCompile(`^([\d]+)\.([\d]+)\.([\d]+)([a-z]+)$`)

// for limited purpose of internal filtering/sorting non-standard semvers
func RewriteSemiSemVer(version string) string {
	match := semisemverLetter.FindStringSubmatch(version)

	if match != nil {
		idx := strings.IndexByte("abcdefghijklmnopqrstuvwxyz", match[4][0])
		atoi, _ := strconv.Atoi(match[3])

		return fmt.Sprintf("%s.%s.%d", match[1], match[2], atoi*10000+idx)
	}

	return version
}
