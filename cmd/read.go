package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/bart-jaskulski/bm/internal/bookmarks"
	"github.com/bart-jaskulski/bm/internal/utils"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:     "read",
	Aliases: []string{"r"},
	Short:   "Read a bookmark",
	Long:    `This command reads a bookmark and passes the URL to the 'rdr' command.`,
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var bookmark string
		if len(args) > 0 {
			bookmark = args[0]
		}

		if !utils.IsValidURL(bookmark) {
			matches, err := bookmarks.FindBookmarks(bookmark)
			if err != nil {
				return err
			}

			if len(matches) == 0 {
				return errors.New("no bookmarks found")
			} else if len(matches) > 1 {
				bookmark = utils.ChooseBookmark(matches)
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
