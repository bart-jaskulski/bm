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
	Short: "Find bookmarks based on a regex pattern",
	Long:  `Find bookmarks based on a regex pattern. If no pattern is provided, all bookmarks are listed.`,
	Run: func(cmd *cobra.Command, args []string) {
    for _, bookmark := range findBookmarks(args[0]) {
      fmt.Println(bookmark)
    }
	},
}

func findBookmarks(searchPattern string) []string {
	var matchingBookmarks []string
	file, err := os.Open(bookmarksPath)
	if err != nil {
		fmt.Println("Error opening .bookmark file:", err)
		return matchingBookmarks
	}
	defer file.Close()

	// If a pattern is provided, compile it into a regex
	var pattern *regexp.Regexp
	pattern, err = regexp.Compile(searchPattern)
	if err != nil {
		fmt.Println("Invalid regex pattern:", err)
		return matchingBookmarks
	}

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Scan through each line in the file
	for scanner.Scan() {
		// If a pattern is provided, only print lines that match the pattern
		if pattern != nil {
			if pattern.MatchString(scanner.Text()) {
				matchingBookmarks = append(matchingBookmarks, scanner.Text())
			}
		} else {
			matchingBookmarks = append(matchingBookmarks, scanner.Text())
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning .bookmark file:", err)
	}

	return matchingBookmarks
}

func init() {
	rootCmd.AddCommand(findCmd)
}
