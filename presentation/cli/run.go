package cli

import (
	"errors"
	"log"
	"wallshrink/application/usecase"
	"wallshrink/domain"

	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	input, _ := cmd.Flags().GetString("input")
	// output, _ := cmd.Flags().GetString("output")
	// width, _ := cmd.Flags().GetInt("width")
	// height, _ := cmd.Flags().GetInt("height")

	inject()
	err := usecase.TestUseCase(input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDirectoryLoadFailed):
			log.Println(err)
		default:
			log.Fatalln(err)
		}
	}
}
