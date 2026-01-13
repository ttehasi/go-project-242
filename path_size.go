package code

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func GetSize(path string, recur, isAll bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("Ошибка %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		targetInfo, err := os.Stat(path)
		if err != nil {
			return 0, fmt.Errorf("битая символическая ссылка %s: %w", path, err)
		}
		info = targetInfo
	}
	if info.Mode().IsRegular() {
		if !isAll && strings.HasPrefix(info.Name(), ".") {
			return 0, nil
		}
		return info.Size(), nil
	}

	if info.IsDir() {
		if !isAll && strings.HasPrefix(info.Name(), ".") {
			return 0, nil
		}
		if recur {
			return calculaterecDirSize(path, isAll)
		} else {
			var totalSize int64

			entries, err := os.ReadDir(path)
			if err != nil {
				return 0, fmt.Errorf("ошибка чтения директории %s: %w", path, err)
			}

			for _, entry := range entries {
				if !isAll && strings.HasPrefix(entry.Name(), ".") {
					continue
				}

				fullPath := filepath.Join(path, entry.Name())

				if !entry.IsDir() {
					size, err := getFileSize(fullPath, entry, isAll)
					if err != nil {
						continue
					}
					totalSize += size
					continue
				}
				continue

			}

			return totalSize, nil
		}
	}

	return 0, fmt.Errorf("%s - не обычный файл и не директория", path)
}

func calculaterecDirSize(dirPath string, isAll bool) (int64, error) {
	var totalSize int64

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения директории %s: %w", dirPath, err)
	}

	for _, entry := range entries {
		if !isAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		fullPath := filepath.Join(dirPath, entry.Name())

		if !entry.IsDir() {
			size, err := getFileSize(fullPath, entry, isAll)
			if err != nil {
				continue
			}
			totalSize += size
			continue
		}

		subDirSize, err := calculaterecDirSize(fullPath, isAll)
		if err != nil {
			continue
		}
		totalSize += subDirSize
	}

	return totalSize, nil
}

func getFileSize(fullPath string, entry os.DirEntry, isAll bool) (int64, error) {
	if entry.Type()&os.ModeSymlink != 0 {
		linkInfo, err := os.Stat(fullPath)
		if err != nil {
			return 0, err
		}

		if linkInfo.IsDir() {
			return 0, nil
		}

		if !isAll && strings.HasPrefix(linkInfo.Name(), ".") {
			return 0, nil
		}

		return linkInfo.Size(), nil
	}

	info, err := entry.Info()
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

func FormatSize(size int64, formated bool) string {
	if !formated {
		return fmt.Sprintf("%dB", size)
	}
	type Res struct {
		size float64
		ed   string
	}
	res := Res{size: float64(size), ed: "B"}
	if res.size >= float64(1024*1024*1024*1024*1024*1024) {
		res.size = math.Round((res.size/float64(1024*1024*1024*1024*1024*1024))*10) / 10
		res.ed = "EB"
	}
	if res.size >= float64(1024*1024*1024*1024*1024) {
		res.size = math.Round((res.size/float64(1024*1024*1024*1024*1024))*10) / 10
		res.ed = "PB"
	}
	if res.size >= float64(1024*1024*1024*1024) {
		res.size = math.Round((res.size/float64(1024*1024*1024*1024))*10) / 10
		res.ed = "TB"
	}
	if res.size >= float64(1024*1024*1024) {
		res.size = math.Round((res.size/float64(1024*1024*1024))*10) / 10
		res.ed = "GB"
	}
	if res.size >= float64(1024*1024) {
		res.size = math.Round((res.size/float64(1024*1024))*10) / 10
		res.ed = "MB"
	}
	if res.size >= float64(1024) {
		res.size = math.Round((res.size/float64(1024))*10) / 10
		res.ed = "KB"
	}
	return fmt.Sprintf("%.1f%s", res.size, res.ed)
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, recursive, all)
	if err != nil {
		return "Ошибка", err
	}
	return FormatSize(size, human), nil
}
