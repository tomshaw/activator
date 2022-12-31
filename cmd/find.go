package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomshaw/activator/system"
	"log"
	"strings"
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

func init() {
	rootCmd.AddCommand(fontsCmd)
	fontsCmd.AddCommand(findFontsCmd)

	findFontsCmd.Flags().StringVarP(&root, "root", "r", "", "Root directory (required)")
	findFontsCmd.MarkFlagRequired("root")
}
