package cmd

import (
	"errors"
	"github.com/psssix/archiver/pkg/vlc"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc <path to source file> [path to packed file]",
	Short: "Pack file using variable-length code",
	RunE:  packVlc,
}

func init() {
	packCmd.AddCommand(vlcCmd)
}

var ErrEmptySourceFilePath = errors.New("path to source file is not specified")
var ErrEmptyPackedFilePath = errors.New("path to packed file is not specified")

func packVlc(_ *cobra.Command, args []string) error {
	var (
		srcFile    string
		packedFile string
	)

	switch len(args) {
	case 0:
		return ErrEmptySourceFilePath
	case 1:
		srcFile = args[0]
		packedFile = generateFileName(srcFile, "vlc")
	case 2:
	default:
		srcFile = args[0]
		packedFile = args[1]
	}

	if srcFile == "" {
		return ErrEmptySourceFilePath
	}

	if packedFile == "" {
		return ErrEmptyPackedFilePath
	}

	srcData, err := os.ReadFile(srcFile)
	if err != nil {
		return err
	}

	packedData := vlc.Encode(string(srcData))
	err = os.WriteFile(packedFile, []byte(packedData), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateFileName(file, packedExt string) string {
	name := filepath.Base(file)
	return strings.TrimSuffix(name, filepath.Ext(file)) + "." + packedExt
}
