// +build !windows

package sftp

import (
	"path"
	"path/filepath"
)

// Makes sure we have a clean POSIX (/) absolute path to work with
func cleanPath(p string) string {
	p = filepath.ToSlash(p)
	if !filepath.IsAbs(p) {
		p = "/" + p
	}
	return path.Clean(p)
}
