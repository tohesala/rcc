package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/robocorp/rcc/common"
	"github.com/robocorp/rcc/operations"

	"github.com/spf13/cobra"
)

var (
	branch string
)

var communityPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull a robot from URL or community sources.",
	Long:  "Pull a robot from URL or community sources.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if common.Debug {
			defer common.Stopwatch("Pull lasted").Report()
		}

		zipfile := filepath.Join(os.TempDir(), fmt.Sprintf("pull%x.zip", time.Now().Unix()))
		defer os.Remove(zipfile)
		if common.Debug {
			common.Log("Using temporary zipfile at %v", zipfile)
		}

		var err error
		branches := []string{branch, "master", "trunk", "main"}

		for _, selected := range branches {
			link := operations.CommunityLocation(args[0], selected)
			err = operations.DownloadCommunityRobot(link, zipfile)
			if err == nil {
				break
			}
		}

		if err != nil {
			common.Exit(1, "Download failed: %v!", err)
		}

		err = operations.Unzip(directory, zipfile, true, false)
		if err != nil {
			common.Exit(1, "Error: %v", err)
		}

		common.Log("OK.")
	},
}

func init() {
	communityCmd.AddCommand(communityPullCmd)
	rootCmd.AddCommand(communityPullCmd)
	communityPullCmd.Flags().StringVarP(&branch, "branch", "b", "main", "Branch/tag/commitid to use as basis for robot.")
	communityPullCmd.Flags().StringVarP(&directory, "directory", "d", ".", "The root directory to extract the robot into.")
}
