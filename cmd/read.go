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

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a bookmark",
	Long:  `This command reads a bookmark and passes the URL to the 'rdr' command.`,
	Args:  cobra.MaximumNArgs(1), // Expect at most one argument: the bookmark URL or pattern
	RunE: func(cmd *cobra.Command, args []string) error {
		var bookmark string
		if len(args) > 0 {
			bookmark = args[0]
		}

		var matches []string
		if bookmark == "" || !isValidURL(bookmark) {
			// If no argument or not a valid URL, treat it as a pattern and find matching bookmarks
			matches = findBookmarks(bookmark)
			if len(matches) == 0 {
				return errors.New("no bookmarks found")
			} else if len(matches) > 1 {
				// If multiple matches, let the user choose a bookmark
				bookmark = chooseBookmark(matches)
				if bookmark == "" {
					return errors.New("no bookmark chosen")
				}
			} else {
				bookmark = matches[0]
			}
		}

		// Check if 'rdr' command exists
		if _, err := exec.LookPath("rdr"); err != nil {
			return errors.New("'rdr' command not found")
		}

		// Create a new command 'rdr' with the bookmark as an argument
		rdrCmd := exec.Command("rdr", bookmark)

		// Set the output of the command to our standard output
		rdrCmd.Stdout = os.Stdout
		rdrCmd.Stderr = os.Stderr

		// Run the command and handle any errors
		if err := rdrCmd.Run(); err != nil {
			return fmt.Errorf("failed to run 'rdr' command: %v", err)
		}

		return nil
	},
}

func init() {
	// Add the 'read' command as a subcommand to the root command
	rootCmd.AddCommand(readCmd)
}

// chooseBookmark lets the user interactively choose a bookmark from a list
func chooseBookmark(bookmarks []string) string {
	// Print the bookmarks and let the user choose one
	for i, bookmark := range bookmarks {
		fmt.Printf("[%d]: %s\n", i, bookmark)
	}

	// Prompt the user to enter a choice
	fmt.Print("Enter the number of the bookmark to open: ")

	// Read the user's choice
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %v\n", err)
		return ""
	}

	// Trim the input and convert it to an integer
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 0 || index >= len(bookmarks) {
		fmt.Printf("Invalid choice: %s. Please enter a number between 0 and %d.\n", input, len(bookmarks)-1)
		return ""
	}

	return bookmarks[index]
}
