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
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   packVlc,
}

func init() {
	packCmd.AddCommand(vlcCmd)
}

var ErrEmptySourceFilePath = errors.New("path to source file is not specified")
var ErrEmptyPackedFilePath = errors.New("path to packed file is not specified")

func packVlc(_ *cobra.Command, args []string) {
	var (
		srcFile    string
		packedFile string
	)

	switch len(args) {
	case 0:
		handleError(ErrEmptySourceFilePath)
	case 1:
		srcFile = args[0]
		packedFile = generateFileName(srcFile, "vlc")
	case 2:
	default:
		srcFile = args[0]
		packedFile = args[1]
	}

	if srcFile == "" {
		handleError(ErrEmptySourceFilePath)
	}

	if packedFile == "" {
		handleError(ErrEmptyPackedFilePath)
	}

	srcData, err := os.ReadFile(srcFile)
	if err != nil {
		handleError(err)
	}

	packedData := vlc.Encode(string(srcData))
	err = os.WriteFile(packedFile, []byte(packedData), 0644)
	if err != nil {
		handleError(err)
	}
}

func generateFileName(file, packedExt string) string {
	name := filepath.Base(file)
	return strings.TrimSuffix(name, filepath.Ext(file)) + "." + packedExt
}
