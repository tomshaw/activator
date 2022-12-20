package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomshaw/activator/system"
	"strings"
)

var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "commands to find or move font folders.",
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for files recursively from selected folder.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Searched: %s\n", root)
		res := system.FindFiles(root)
		fmt.Println(strings.Join(res, ",\n"))
		fmt.Printf("Found: %d files.\n", len(res))
	},
	Example: `activator fonts find --root "C:\Fonts"`,
}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy font files from source to destination.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Source: %s\n", source)
		fmt.Printf("Destination: %s\n", destination)
		res := system.CopyFiles(source, destination)
		fmt.Printf("Found: %d files.\n", res)
	},
	Example: `activator fonts copy --source "C:\Fonts" --destination "C:\Dest"`,
}

func init() {
	rootCmd.AddCommand(fontsCmd)
	fontsCmd.AddCommand(findCmd, copyCmd)

	findCmd.Flags().StringVarP(&root, "root", "r", "", "Root directory (required)")
	findCmd.MarkFlagRequired("root")

	copyCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory (required)")
	copyCmd.MarkFlagRequired("source")
	copyCmd.Flags().StringVarP(&destination, "destination", "d", "", "Destination directory (required)")
	copyCmd.MarkFlagRequired("destination")
}
