package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "geektime",
	Short: "Geektime",
	Long:  `Geektime`,
}

func init() {
	rootCmd.AddCommand(CmdLogin, CmdEbook, CmdQuery, CmdLearn)
}

var CmdLearn = &cobra.Command{
	Use:   "learn ",
	Short: "Query your learn courses",
	Run:   RunLogin,
}
var CmdQuery = &cobra.Command{
	Use:   "query ",
	Short: "Query all courses",
	Run:   RunLogin,
}
var CmdEbook = &cobra.Command{
	Use:   "ebook ",
	Short: "Make column article to markdown file",
	Run:   RunLogin,
}

var CmdLogin = &cobra.Command{
	Use:   "login ",
	Short: "Login your account",
	Long:  "Login your account",
	Run:   RunLogin,
}

func RunLogin(cmd *cobra.Command, args []string) {
	var dir string
	if len(args) > 0 {
		dir = args[0]
	}
	base, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
		return
	}
	if dir == "" {
		// find the directory containing the cmd/*
		cmdPath, err := findCMD(base)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
			return
		}
		if len(cmdPath) == 0 {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", "The cmd directory cannot be found in the current directory")
			return
		} else if len(cmdPath) == 1 {
			for k, v := range cmdPath {
				dir = path.Join(v, k)
			}
		} else {
			var cmdPaths []string
			for k, _ := range cmdPath {
				cmdPaths = append(cmdPaths, k)
			}
			prompt := &survey.Select{
				Message: "Which directory do you want to run?",
				Options: cmdPaths,
			}
			survey.AskOne(prompt, &dir)
			if dir == "" {
				return
			}
			dir = path.Join(cmdPath[dir], dir)
		}
	}

	return
}
func findCMD(base string) (map[string]string, error) {
	var root bool
	next := func(dir string) (map[string]string, error) {
		cmdPath := make(map[string]string)
		err := filepath.Walk(dir, func(walkPath string, info os.FileInfo, err error) error {
			// multi level directory is not allowed under the cmdPath directory, so it is judged that the path ends with cmdPath.
			if strings.HasSuffix(walkPath, "cmd") {
				paths, err := ioutil.ReadDir(walkPath)
				if err != nil {
					return err
				}
				for _, fileInfo := range paths {
					if fileInfo.IsDir() {
						cmdPath[path.Join("cmd", fileInfo.Name())] = filepath.Join(walkPath, "..")
					}
				}
				return nil
			}
			if info.Name() == "go.mod" {
				root = true
			}
			return nil
		})
		return cmdPath, err
	}
	for i := 0; i < 5; i++ {
		tmp := base
		cmd, err := next(tmp)
		if err != nil {
			return nil, err
		}
		if len(cmd) > 0 {
			return cmd, nil
		}
		if root {
			break
		}
		tmp = filepath.Join(base, "..")
	}
	return map[string]string{"": base}, nil
}
