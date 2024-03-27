package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for bookmarks containing the given pattern",
	Run: func(cmd *cobra.Command, args []string) {
		searchPattern := ""
		if len(args) > 0 {
			searchPattern = args[0]
		}

		for _, bookmark := range findBookmarks(searchPattern) {
			fmt.Println(bookmark)
		}
	},
}

func findBookmarks(searchPattern string) []string {
	var matchingBookmarks []string
	file, err := os.Open(getBookmarksFilePath())
	if err != nil {
		fmt.Println("Error opening .bookmark file:", err)
		return matchingBookmarks
	}
	defer file.Close()

	pattern, err := regexp.Compile(searchPattern)
	if err != nil {
		fmt.Println("Invalid regex pattern:", err)
		return matchingBookmarks
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if pattern.MatchString(scanner.Text()) {
			matchingBookmarks = append(matchingBookmarks, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning .bookmark file:", err)
	}

	return matchingBookmarks
}

func init() {
	rootCmd.AddCommand(findCmd)
}
