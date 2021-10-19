package utils

import "testing"

func TestGetFileContentTypeBySuffix(t *testing.T) {
	contentType, b := GetFileContentTypeBySuffix("/a/b/c.css")
	if !b {
		t.Error("not found")
		return
	}
	t.Log(contentType)
}
