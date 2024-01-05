package cli

import (
	"errors"
	"log"
	"wompressor/application/usecase"
	"wompressor/domain"

	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	input, _ := cmd.Flags().GetString("input")
	output, _ := cmd.Flags().GetString("output")
	width, _ := cmd.Flags().GetInt("width")
	height, _ := cmd.Flags().GetInt("height")

	dimension := domain.Dimension{
		Width:  width,
		Height: height,
	}

	inject()
	err := usecase.CompressImageSetUseCase(input, output, dimension)

	// Main error handle
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrImageSetLoadFailed):
			log.Println(err)
		default:
			log.Fatalln(err)
		}
	}
}
