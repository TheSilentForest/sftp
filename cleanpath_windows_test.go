// +build windows

package sftp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanPath(t *testing.T) {
	assert.Equal(t, "c:\\", cleanPath("c:\\"))
	assert.Equal(t, "C:\\", cleanPath("."))
	assert.Equal(t, "C:\\", cleanPath("C:\\."))
	assert.Equal(t, "C:\\", cleanPath("C:\\a\\.."))
	assert.Equal(t, "C:\\a\\c", cleanPath("C:\\a\\b\\..\\c"))
	assert.Equal(t, "C:\\a\\c", cleanPath("C:\\a\\b\\..\\c\\"))
	assert.Equal(t, "C:\\a", cleanPath("C:\\a\\b\\.."))
	assert.Equal(t, "D:\\a\\b\\c", cleanPath("D:\\a\\b\\c"))
	assert.Equal(t, "C:\\", cleanPath("C:\\\\"))
	assert.Equal(t, "C:\\a", cleanPath("C:\\a\\"))
	assert.Equal(t, "C:\\a", cleanPath("a\\"))
	assert.Equal(t, "C:\\a\\b\\c", cleanPath("a\\b\\c\\"))

	assert.Equal(t, "C:\\", cleanPath("/"))
	assert.Equal(t, "C:\\", cleanPath("//"))
	assert.Equal(t, "C:\\a", cleanPath("/a/"))
	assert.Equal(t, "C:\\a", cleanPath("a/"))
	assert.Equal(t, "C:\\a\\b\\c", cleanPath("/a//b//c/"))
}
