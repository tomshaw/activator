package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomshaw/activator/system"
	"log"
	"time"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "commands to copy fonts and folders.",
}

var copyFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Copy font files from source to destination.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Source: %q\n", args)
		fmt.Printf("Destination: %s\n", destination)
		startTime := time.Now()
		err := system.CopyFiles(destination, args)
		if err != nil {
			log.Fatalf("Fatal error %v", err)
		}
		fmt.Printf("Execution time: %s", time.Since(startTime))
	},
	Example: `activator copy files --destination "C:\Dest" ...files`,
}

var copyFoldersCmd = &cobra.Command{
	Use:   "folders",
	Short: "Copy font files and folders to destination.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Source: %s\n", source)
		fmt.Printf("Destination: %s\n", destination)
		startTime := time.Now()
		err := system.CopyFilesFolders(source, destination)
		if err != nil {
			log.Fatalf("Fatal error %v", err)
		}
		fmt.Printf("Execution time: %s", time.Since(startTime))
	},
	Example: `activator copy folders --source "C:\Fonts" --destination "C:\Dest"`,
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.AddCommand(copyFilesCmd, copyFoldersCmd)

	copyFilesCmd.Flags().StringVarP(&destination, "destination", "d", "", "Destination directory (required)")
	copyFilesCmd.MarkFlagRequired("destination")

	copyFoldersCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory (required)")
	copyFoldersCmd.MarkFlagRequired("source")
	copyFoldersCmd.Flags().StringVarP(&destination, "destination", "d", "", "Destination directory (required)")
	copyFoldersCmd.MarkFlagRequired("destination")
}
