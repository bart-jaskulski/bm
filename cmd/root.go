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
			printBookmarks()
		} else {
			addBookmark(args[0])
		}
	},
}

// printBookmarks reads the content of the bookmarks file and prints it
func printBookmarks() {
	data, err := os.ReadFile(bookmarksPath)
	if err != nil {
		fmt.Println("Error reading bookmarks file:", err)
		return
	}

	fmt.Println(string(data))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
