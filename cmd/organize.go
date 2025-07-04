package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

var organizeCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize receipts into monthly folders",
	RunE: func(cmd *cobra.Command, args []string) error {
		const srcDir = "unorganized"

		entries, err := os.ReadDir(srcDir)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("%s directory not found", srcDir)
			}
			return err
		}

		re := regexp.MustCompile(`^\d{4}`)
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			name := e.Name()
			match := re.FindString(name)
			if len(match) < 4 {
				continue
			}
			monthStr := match[2:4]
			month, err := strconv.Atoi(monthStr)
			if err != nil {
				continue
			}
			destDir := fmt.Sprintf("%d月", month)
			if err := os.MkdirAll(destDir, 0o755); err != nil {
				return err
			}
			srcPath := filepath.Join(srcDir, name)
			destPath := filepath.Join(destDir, name)
			if err := os.Rename(srcPath, destPath); err != nil {
				return err
			}
			fmt.Printf("Moved %s to %s\n", name, destDir)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
}
