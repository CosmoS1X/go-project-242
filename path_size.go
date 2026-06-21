package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// threshold is the size in bytes at which to switch to the next unit (KB, MB, etc.).
// Si standard units are used, so 1 KB = 1000 bytes, 1 MB = 1000 KB, etc.
// Adjust as needed for your use case. For example, if you want to use binary units (KiB, MiB, etc.), you can set threshold to 1024.
const threshold float64 = 1000

// fmtHuman formats the size in bytes to a human-readable string.
func fmtHuman(size float64, human bool) string {
	if !human {
		return fmt.Sprintf("%.0fB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIdx := 0

	for unitIdx < len(units)-1 && size >= threshold {
		size /= threshold
		unitIdx++
	}

	if unitIdx == 0 {
		return fmt.Sprintf("%.0f%s", size, units[unitIdx])
	}

	return fmt.Sprintf("%.1f%s", size, units[unitIdx])
}

// shouldSkip determines whether a file or directory should be skipped based on its name and the includeHidden flag.
func shouldSkip(name string, includeHidden bool) bool {
	return !includeHidden && strings.HasPrefix(name, ".")
}

// walkDirSize recursively calculates the total size of a directory, optionally including hidden files and directories.
func walkDirSize(path string, recursive, includeHidden bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64

	for _, entry := range entries {
		name := entry.Name()
		if shouldSkip(name, includeHidden) {
			continue
		}

		entryPath := filepath.Join(path, name)
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}

		if info.IsDir() {
			if !recursive {
				continue
			}

			dirSize, err := walkDirSize(entryPath, recursive, includeHidden)
			if err != nil {
				return 0, err
			}

			size += dirSize
			continue
		}

		size += info.Size()
	}

	return size, nil
}

// GetPathSize returns the size of a file or directory at the given path.
// If recursive is true, the size of directories is calculated recursively.
// If human is true, the size is returned in a human-readable format.
// If all is true, hidden files and directories are included in the size calculation.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return fmtHuman(float64(info.Size()), human), nil
	}

	size, err := walkDirSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return fmtHuman(float64(size), human), nil
}
