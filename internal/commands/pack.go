package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/plexsystems/pacmod/pack"
	"github.com/spf13/cobra"
)

// NewPackCommand creates a new pack command which allows
// the user to package their Go modules
func NewPackCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "pack <version> <outputdirectory>",
		Short: "Package Go module in current directory",
		Args:  cobra.MinimumNArgs(2),

		RunE: func(cmd *cobra.Command, args []string) error {
			return runPackCommand(args)
		},
	}

	return &cmd
}

func runPackCommand(args []string) error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	version := args[0]

	path, err = filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("get abs path of module path: %w", err)
	}

	outputDirectory, err := filepath.Abs(args[1])
	if err != nil {
		return fmt.Errorf("get abs path of output directory: %w", err)
	}

	log.Printf("Packing module in path %s...", outputDirectory)
	if err := pack.Module(path, version, outputDirectory); err != nil {
		return fmt.Errorf("package module: %w", err)
	}

	return nil
}
