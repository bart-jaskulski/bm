package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const bookmarksFile = ".bookmarks"

func getBookmarksFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return filepath.Join(home, bookmarksFile)
}

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
