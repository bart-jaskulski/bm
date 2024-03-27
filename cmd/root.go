/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var bookmarksPath = os.Getenv("HOME") + "/.bookmarks"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bm",
	Short: "bm is a CLI tool for managing bookmarks",
	Long:  `bm is a CLI tool for managing bookmarks. It can add, find and read bookmarks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			for _, bookmark := range findBookmarks("") {
				fmt.Println(bookmark)
			}
		} else {
			addBookmark(args[0])
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
