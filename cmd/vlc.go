package cmd

import (
	"errors"
	"github.com/psssix/archiver/pkg/compression/vlc"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var vlcPackCmd = &cobra.Command{
	Use:   "vlc <path to source file> [path to packed file]",
	Short: "Pack file using variable-length code",
	RunE:  vlcPack,
}

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc <path to source file> [path to unpacked file]",
	Short: "Unpack file using variable-length code",
	RunE:  vlcUnpack,
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
	unpackCmd.AddCommand(vlcUnpackCmd)
}

var ErrEmptySourceFilePath = errors.New("path to source file is not specified")
var ErrEmptyPackedFilePath = errors.New("path to packed file is not specified")
var ErrEmptyUnpackedFilePath = errors.New("path to unpacked file is not specified")

func vlcPack(_ *cobra.Command, args []string) error {
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

	packedData, err := vlc.New().Pack(string(srcData))
	if err != nil {
		return err
	}

	err = os.WriteFile(packedFile, packedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func vlcUnpack(_ *cobra.Command, args []string) error {
	var (
		srcFile      string
		unpackedFile string
	)

	switch len(args) {
	case 0:
		return ErrEmptySourceFilePath
	case 1:
		srcFile = args[0]
		unpackedFile = generateFileName(srcFile, "txt")
	case 2:
	default:
		srcFile = args[0]
		unpackedFile = args[1]
	}

	if srcFile == "" {
		return ErrEmptySourceFilePath
	}

	if unpackedFile == "" {
		return ErrEmptyUnpackedFilePath
	}

	srcData, err := os.ReadFile(srcFile)
	if err != nil {
		return err
	}

	unpackedData, err := vlc.New().Unpack(srcData)
	if err != nil {
		return err
	}

	err = os.WriteFile(unpackedFile, []byte(unpackedData), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateFileName(file, ext string) string {
	name := filepath.Base(file)
	return strings.TrimSuffix(name, filepath.Ext(file)) + "." + ext
}
