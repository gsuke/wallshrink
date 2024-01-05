package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wompressor",
	Short: "Wompressor compresses the directory of wallpaper images.",
	Run:   run,
}

// Flags
func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().StringP("input", "i", "", "Source directory")
	rootCmd.Flags().StringP("output", "o", "", "Destination directory")
	rootCmd.Flags().IntP(
		"width",
		"x",
		defaultWidth,
		fmt.Sprintf("Width to scale down the image (default %d)", defaultWidth),
	)
	rootCmd.Flags().IntP(
		"height",
		"y",
		defaultHeight,
		fmt.Sprintf("Height to scale down the image (default %d)", defaultHeight),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
