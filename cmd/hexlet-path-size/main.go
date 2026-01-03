package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func GetSize(path string) (int64, error) {
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
		return info.Size(), nil
	}

	if info.IsDir() {
		fmt.Println("dir")
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, fmt.Errorf("ошибка чтения директории %s: %w", path, err)
		}
		var totalSize int64 = 0
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
		return totalSize, nil
	}

	return 0, fmt.Errorf("%s - не обычный файл и не директория", path)
}

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			res, err := GetSize(cmd.Args().Get(0))
			if err != nil {
				return err
			}
			fmt.Printf("%dB\t%s", res, cmd.Args().Get(0))
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
