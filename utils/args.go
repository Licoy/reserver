package utils

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
)

type CmdArgs struct {
	Port      int    `default:"8080" short:"p" long:"port" description:"listen port"`
	Root      string `short:"r" long:"root" description:"root directory"`
	Host      string `default:"0.0.0.0" short:"H" long:"host" description:"bind host address"`
	NoBrowser bool   `long:"no-browser"  description:"don't auto opening browser"`
	NoWatch   bool   `long:"no-watch" description:"don't listen for file changes"`
	Browser   string `long:"browser" description:"specify the browser you want to use"`
	Path      string `short:"P" long:"path" description:"default open link path"`
	HideLog   bool   `long:"hide-log" description:"displays the change log of the observation path"`
	CssReload bool   `long:"css-reload" description:"css file changes to reload the page"`
	Wait      int    `default:"100" short:"w" long:"wait" description:"wait for the specified time before reloading"`
	// tips: https://github.com/jessevdk/go-flags/issues/313
	Ignore    []string `short:"i" long:"ignore" description:"multiple observation paths are allowed to be ignored, ex: -i /a -i /b"`
	IgnoreMap map[string]struct{}
}

func ParseCommonArgs() *CmdArgs {
	args := &CmdArgs{}
	_, err := flags.ParseArgs(args, os.Args)
	if err != nil {
		et, y := err.(*flags.Error)
		if y {
			if errors.Is(flags.ErrHelp, et.Type) {
				os.Exit(0)
			}
		}
		panic(fmt.Sprintf("parsing parameter failed: %v", err))
	}
	if args.Root == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic("failed to get runtime directory")
		}
		args.Root = wd
	}
	args.Root = filepath.ToSlash(args.Root)
	args.IgnoreMap = make(map[string]struct{})
	if len(args.Ignore) > 0 {
		for _, v := range args.Ignore {
			abs, err := filepath.Abs(v)
			if err != nil {
				continue
			}
			args.IgnoreMap[filepath.ToSlash(abs)] = struct{}{} //:sparkles: 优化关于忽略目录的相对路径支持
		}
	}
	return args
}
