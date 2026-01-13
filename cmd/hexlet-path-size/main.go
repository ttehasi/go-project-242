package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "human", Aliases: []string{"H"}, Value: true, Usage: "human-readable sizes (auto-select unit)", Required: false},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Value: false, Usage: "include hidden files and directories", Required: false},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			size, err := code.GetSize(cmd.Args().Get(0), cmd.Bool("all"))
			if err != nil {
				return err
			}
			res := code.FormatSize(size, cmd.Bool("human"))
			fmt.Printf("%s\t%s", res, cmd.Args().Get(0))
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
