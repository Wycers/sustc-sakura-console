package util

import (
	"strings"
)

// Path prefixes.
const (
	PathRoot        = "/"
	PathAdmin       = "/admin"
	PathAPI         = "/api"
	PathFavicon     = "/favicon.ico"
	PathFetchUpload = "/fetch-upload"
	PathChangelogs  = "/changelogs"
	PathRobots      = "/robots.txt"
)

var reservedPaths = []string{
	PathAdmin, PathAPI, PathFavicon, PathFetchUpload, PathChangelogs, PathRobots,
}

func IsReservedPath(path string) bool {
	path = strings.TrimSpace(path)
	if PathRoot == path {
		return true
	}

	for _, reservedPath := range reservedPaths {
		if strings.HasPrefix(path, reservedPath) {
			return true
		}
	}

	return false
}
