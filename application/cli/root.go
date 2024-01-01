package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultWidth  = 3840
	defaultHeight = 2160
)

var rootCmd = &cobra.Command{
	Use:   "wallshrink",
	Short: "Wallshrink compresses the directory of wallpaper images.",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")

		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Output: %s\n", output)
		fmt.Printf("Width: %d\n", width)
		fmt.Printf("Height: %d\n", height)
	},
}

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
