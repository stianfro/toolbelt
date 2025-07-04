package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var joinCSVCmd = &cobra.Command{
	Use:   "joincsv [directory] [output]",
	Short: "Join all CSV files in a directory into a single file",
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

		var csvFiles []string
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			if strings.HasSuffix(e.Name(), ".csv") {
				csvFiles = append(csvFiles, filepath.Join(dir, e.Name()))
			}
		}

		if len(csvFiles) == 0 {
			return errors.New("no csv files found")
		}

		sort.Strings(csvFiles)

		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		writer := bufio.NewWriter(outFile)
		defer writer.Flush()

		for i, file := range csvFiles {
			f, err := os.Open(file)
			if err != nil {
				return err
			}

			reader := bufio.NewReader(f)
			firstLine := true
			for {
				line, err := reader.ReadString('\n')
				if err != nil && err != io.EOF {
					f.Close()
					return err
				}
				if line != "" {
					if i > 0 && firstLine {
						firstLine = false
						continue
					}
					_, werr := writer.WriteString(line)
					if werr != nil {
						f.Close()
						return werr
					}
				}
				if err == io.EOF {
					break
				}
				firstLine = false
			}
			f.Close()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(joinCSVCmd)
}
