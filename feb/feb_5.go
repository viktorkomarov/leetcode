package feb

import "strings"

func nextDir(path string, from int) int {
	for i := from; i < len(path); i++ {
		if path[i] == '/' {
			return i
		}
	}

	return -1
}

func clearPath(path string) string {
	paths := make([]rune, 0)
	start := 0
	for start < len(path) {
		slash := false
		if path[start] == '/' {
			slash = true
		}

		paths = append(paths, rune(path[start]))
		start++
		for slash && start < len(path) && path[start] == '/' {
			start++
		}
	}

	paths = append(paths, '/')

	return string(paths)
}

func simplifyPath(path string) string {
	dirs := []string{"/"}

	path = clearPath(path)
	start := 1
	for start < len(path) {
		next := nextDir(path, start)
		if next == -1 {
			break
		}

		dir := path[start:next]
		if start == next || dir == "." {
			start = next + 1
			continue
		}

		if dir == ".." {
			if len(dirs) != 1 {
				dirs = dirs[:len(dirs)-1]
			}
		} else {
			dirs = append(dirs, dir)
		}

		start = next + 1
	}

	return "/" + strings.Join(dirs[1:], "/")
}
