package utility

import (
	"net/url"
	"regexp"
)

var gitCloneFormatRegex = regexp.MustCompile(`^(?P<scheme>git\+ssh)://((?P<user>\S+)@)?(?P<host>\S+):(?P<path>\S+)$`)

func ParseUriOrGitCloneArg(uri string) (*url.URL, error) {
	matches := gitCloneFormatRegex.FindStringSubmatch(uri)

	if matches != nil {
		return &url.URL{
			Host:   matches[4],
			Path:   matches[5],
			Scheme: matches[1],
			User:   url.User(matches[3]),
		}, nil
	}

	return url.Parse(uri)
}
