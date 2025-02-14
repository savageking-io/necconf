package necconf

import (
	"fmt"
	"path/filepath"
)

func ExtractDirectoryAndFilenameFromPath(path string) (string, string, error) {
	if filepath.Ext(path) != ".yml" && filepath.Ext(path) != ".yaml" {
		return "", "", fmt.Errorf("bad config file: %s is not a .yml or .yaml", path)
	}
	path = filepath.Clean(path)
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	return dir, filename, nil
}
