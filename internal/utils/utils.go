package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const bookmarksFile = ".bookmarks"

func GetBookmarksFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return filepath.Join(home, bookmarksFile)
}

func IsValidURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}

func ChooseBookmark(bookmarks []string) string {
	for i, bookmark := range bookmarks {
		fmt.Printf("[%d]: %s\n", i, bookmark)
	}

	fmt.Print("Enter the number of the bookmark to open: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %v\n", err)
		return ""
	}

	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 0 || index >= len(bookmarks) {
		fmt.Printf("Invalid choice: %s. Please enter a number between 0 and %d.\n", input, len(bookmarks)-1)
		return ""
	}

	return bookmarks[index]
}
