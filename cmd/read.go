package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a bookmark",
	Long:  `This command reads a bookmark and passes the URL to the 'rdr' command.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var bookmark string
		if len(args) > 0 {
			bookmark = args[0]
		}

		var matches []string
		if bookmark == "" || !isValidURL(bookmark) {
			matches = findBookmarks(bookmark)
			if len(matches) == 0 {
				return errors.New("no bookmarks found")
			} else if len(matches) > 1 {
				bookmark = chooseBookmark(matches)
				if bookmark == "" {
					return errors.New("no bookmark chosen")
				}
			} else {
				bookmark = matches[0]
			}
		}

		if _, err := exec.LookPath("rdr"); err != nil {
			return errors.New("'rdr' command not found")
		}

		rdrCmd := exec.Command("rdr", bookmark)

		rdrCmd.Stdout = os.Stdout
		rdrCmd.Stderr = os.Stderr

		if err := rdrCmd.Run(); err != nil {
			return fmt.Errorf("failed to run 'rdr' command: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func chooseBookmark(bookmarks []string) string {
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
