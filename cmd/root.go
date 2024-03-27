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
	Short: "Bookmark manager",
	Long:  `Bookmark manager for managing bookmarks in the ~/.bookmarks file.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			for _, bookmark := range findBookmarks("") {
				fmt.Println(bookmark)
			}
			return
		}

		addBookmark(args[0])
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
