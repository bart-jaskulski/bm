package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/bart-jaskulski/bm/internal/bookmarks"
	"github.com/bart-jaskulski/bm/internal/utils"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:     "open [bookmark]",
	Aliases: []string{"o"},
	Short:   "Open a bookmark in the default web browser",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bookmark := args[0]
		var url string

		if !utils.IsValidURL(bookmark) {
			chosenBookmark, err := bookmarks.SearchAndChooseBookmark(bookmark)
			if err != nil {
				return err
			}
			bookmark = chosenBookmark
		}

		url = strings.Split(bookmark, " ")[0]

		// Open URL in the default browser
		err := exec.Command("open", url).Run()
		if err != nil {
			return fmt.Errorf("failed to open URL: %v", err)
		}
		fmt.Println("URL opened in browser:", url)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
