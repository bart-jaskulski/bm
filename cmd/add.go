package cmd

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bookmark",
	Long:  `Add a new bookmark to the .bookmark file`,
	Run: func(cmd *cobra.Command, args []string) {
		var urlStr string
		if len(args) < 1 {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			urlStr = scanner.Text()
		} else {
			urlStr = args[0]
		}
		addBookmark(urlStr)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

var BookmarkFileMutex = &sync.Mutex{}

func isValidURL(urlStr string) bool {
  _, err := url.ParseRequestURI(urlStr)
  return err == nil
}

func addBookmark(urlStr string) {
  if !isValidURL(urlStr) {
    fmt.Printf("Invalid URL: %s\n", urlStr)
    return
  }

	BookmarkFileMutex.Lock()
	defer BookmarkFileMutex.Unlock()

	file, err := os.OpenFile(getBookmarksFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
    os.Exit(1)
		return
	}
	defer file.Close()

	_, err = file.WriteString(urlStr + "\n")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
    os.Exit(1)
		return
	}

	fmt.Println("Bookmark added successfully")
}
