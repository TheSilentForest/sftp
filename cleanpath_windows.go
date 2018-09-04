package sftp

import (
	"path/filepath"
	"strings"
)

// Makes sure we have a clean Windows absolute path to work with
// If the path is not absolute we assume it's on C:\
func cleanPath(p string) string {
	// Hack to cope with sftp client paths
	parts := strings.SplitN(p, "/", 2)
	if len(parts) == 2 && len(parts[0]) > 2 && parts[0][0:2] == "C:" {
		p = parts[1]
	}

	// Hack to cope with sftp client according to ansible
	if len(p) > 1 && p[0] == '\'' && p[len(p)-1] == '\'' {
		p = p[1 : len(p)-1]
	}

	p = filepath.FromSlash(p)
	if filepath.VolumeName(p) == "" {
		if len(p) > 0 {
			if p[0] != '\\' {
				p = "\\" + p
			}
		} else {
			p = "\\"
		}
		p = "C:" + p
	}
	return filepath.Clean(p)
}
