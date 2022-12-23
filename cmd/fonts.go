package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomshaw/activator/system"
	"log"
	"strings"
	"time"
)

var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "commands to find or move font folders.",
}

var findFontsCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for files recursively from selected folder.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Searched: %s\n", root)
		res, err := system.FindFiles(root)
		if err != nil {
			log.Fatalf("Fatal error %p", err)
		}
		fmt.Println(strings.Join(res, ",\n"))
		fmt.Printf("Found: %d files.\n", len(res))
	},
	Example: `activator fonts find --root "C:\Fonts"`,
}

var copyFilesCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy font files from source to destination.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Source: %s\n", source)
		fmt.Printf("Destination: %s\n", destination)
		startTime := time.Now()
		err := system.CopyFiles(source, destination)
		if err != nil {
			log.Fatalf("Fatal error %v", err)
		}
		fmt.Printf("Execution time: %s", time.Since(startTime))
	},
	Example: `activator fonts copy --source "C:\Fonts" --destination "C:\Dest"`,
}

var copyFilesFoldersCmd = &cobra.Command{
	Use:   "copyf",
	Short: "Copy font files and source folders to destination.",
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
	Example: `activator fonts copy --source "C:\Fonts" --destination "C:\Dest"`,
}

func init() {
	rootCmd.AddCommand(fontsCmd)
	fontsCmd.AddCommand(findFontsCmd, copyFilesCmd, copyFilesFoldersCmd)

	findFontsCmd.Flags().StringVarP(&root, "root", "r", "", "Root directory (required)")
	findFontsCmd.MarkFlagRequired("root")

	copyFilesCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory (required)")
	copyFilesCmd.MarkFlagRequired("source")
	copyFilesCmd.Flags().StringVarP(&destination, "destination", "d", "", "Destination directory (required)")
	copyFilesCmd.MarkFlagRequired("destination")

	copyFilesFoldersCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory (required)")
	copyFilesFoldersCmd.MarkFlagRequired("source")
	copyFilesFoldersCmd.Flags().StringVarP(&destination, "destination", "d", "", "Destination directory (required)")
	copyFilesFoldersCmd.MarkFlagRequired("destination")
}
