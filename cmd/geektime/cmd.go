package main

import (
	"fmt"
	"geektime-ebook/geek"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os"
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
	Short: "Login your account ",
	Long:  "Login your account [-h] [-a ACCOUNT] [-p PASSWORD]",
	Run:   RunLogin,
}

func RunLogin(cmd *cobra.Command, args []string) {
	var req geek.LoginReq

	var qs = []*survey.Question{
		{
			Name: "country",
			Prompt: &survey.Input{
				Message: "What is your country?",
				Default: "+86",
			},
		},
		{
			Name:     "cellphone",
			Prompt:   &survey.Input{Message: "What is your cellphone?"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Input{Message: "What is your password?"},
			Validate: survey.Required,
		},
	}
	survey.Ask(qs, &req)

	err := geek.Login(req.Cellphone, req.Password)
	if err != nil {
		printError(err)
	}

	return
}
func printError(err error) {
	fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
}
