package code

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func GetSize(path string, isAll bool) (int64, error) {
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
		if isAll {
			return info.Size(), nil
		} else {
			if strings.HasPrefix(info.Name(), ".") {
				return int64(0), nil
			} else {
				return info.Size(), nil
			}
		}
	}

	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, fmt.Errorf("ошибка чтения директории %s: %w", path, err)
		}
		var totalSize int64 = 0
		if isAll {
			for _, entry := range entries {
				fullPath := path + "/" + entry.Name()

				if entry.IsDir() {
					continue
				}
				info, err := entry.Info()
				if err != nil {
					continue
				}
				if entry.Type()&os.ModeSymlink != 0 {
					linkInfo, err := os.Stat(fullPath)
					if err != nil {
						continue
					}

					if linkInfo.Mode().IsRegular() {
						totalSize += linkInfo.Size()
					}
				} else if info.Mode().IsRegular() {
					totalSize += info.Size()
				}
			}
		} else {
			if strings.HasPrefix(info.Name(), ".") {
				return int64(0), nil
			} else {
				for _, entry := range entries {
					fullPath := path + "/" + entry.Name()

					if entry.IsDir() {
						continue
					}
					info, err := entry.Info()
					if err != nil {
						continue
					}
					if entry.Type()&os.ModeSymlink != 0 {
						linkInfo, err := os.Stat(fullPath)
						if err != nil {
							continue
						}

						if linkInfo.Mode().IsRegular() {
							if strings.HasPrefix(linkInfo.Name(), ".") {
								totalSize += 0
							} else {
								totalSize += linkInfo.Size()
							}
						}
					} else if info.Mode().IsRegular() {
						if strings.HasPrefix(info.Name(), ".") {
							totalSize += 0
						} else {
							totalSize += info.Size()
						}
					}
				}
			}
		}

		return totalSize, nil
	}

	return 0, fmt.Errorf("%s - не обычный файл и не директория", path)
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
