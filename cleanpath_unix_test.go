// +build !windows

package sftp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanPath(t *testing.T) {
	assert.Equal(t, "/", cleanPath("/"))
	assert.Equal(t, "/", cleanPath("."))
	assert.Equal(t, "/", cleanPath("/."))
	assert.Equal(t, "/", cleanPath("/a/.."))
	assert.Equal(t, "/a/c", cleanPath("/a/b/../c"))
	assert.Equal(t, "/a/c", cleanPath("/a/b/../c/"))
	assert.Equal(t, "/a", cleanPath("/a/b/.."))
	assert.Equal(t, "/a/b/c", cleanPath("/a/b/c"))
	assert.Equal(t, "/", cleanPath("//"))
	assert.Equal(t, "/a", cleanPath("/a/"))
	assert.Equal(t, "/a", cleanPath("a/"))
	assert.Equal(t, "/a/b/c", cleanPath("/a//b//c/"))

	// filepath.ToSlash does not touch \ as char on unix systems
	// so os.PathSeparator is used for windows compatible tests
	bslash := string(os.PathSeparator)
	assert.Equal(t, "/", cleanPath(bslash))
	assert.Equal(t, "/", cleanPath(bslash+bslash))
	assert.Equal(t, "/a", cleanPath(bslash+"a"+bslash))
	assert.Equal(t, "/a", cleanPath("a"+bslash))
	assert.Equal(t, "/a/b/c",
		cleanPath(bslash+"a"+bslash+bslash+"b"+bslash+bslash+"c"+bslash))
}
