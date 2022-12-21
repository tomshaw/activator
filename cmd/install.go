package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tomshaw/activator/install"
	"github.com/tomshaw/activator/utils"
)

var installFontsCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs selected fonts.",
	Run: func(cmd *cobra.Command, args []string) {
		fonts := utils.AppendExists(args)
		if err := install.Init(fonts, true, temporary); err != nil {
			fmt.Println("Font installation errors:", err)
		}
	},
	Example: `activator install ...files`,
}

var uninstallFontsCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls selected fonts.",
	Run: func(cmd *cobra.Command, args []string) {
		fonts := utils.AppendExists(args)
		if err := install.Init(fonts, false, temporary); err != nil {
			fmt.Println("Font installation errors:", err)
		}
	},
	Example: `activator uninstall ...files`,
}

func init() {
	rootCmd.AddCommand(installFontsCmd, uninstallFontsCmd)
	installFontsCmd.Flags().BoolVarP(&temporary, "temporary", "t", false, "Windows temporary font installation.")
}
