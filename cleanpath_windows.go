package sftp

import (
	"path/filepath"
	"strings"
)

// Makes sure we have a clean Windows absolute path to work with
// If the path is not absolute we assume it's on C:\
func cleanPath(p string) string {
	// Hack to make path sent by sftp client work in Windows
	if absPath, err := filepath.Abs("."); err == nil {
		p = strings.TrimPrefix(p, absPath+"/")
	}

	p = filepath.ToSlash(p)
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
