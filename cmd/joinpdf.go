package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/spf13/cobra"
)

var joinPDFCmd = &cobra.Command{
	Use:   "joinpdf [directory] [output]",
	Short: "Join all PDF files in a directory into a single file",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := args[0]
		outPath := args[1]

		entries, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("%s directory not found", dir)
			}
			return err
		}

		var pdfFiles []string
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			if strings.HasSuffix(strings.ToLower(e.Name()), ".pdf") {
				pdfFiles = append(pdfFiles, filepath.Join(dir, e.Name()))
			}
		}

		if len(pdfFiles) == 0 {
			return fmt.Errorf("no PDF files found in directory %s", dir)
		}

		sort.Strings(pdfFiles)

		if err := api.MergeCreateFile(pdfFiles, outPath, nil); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(joinPDFCmd)
}
