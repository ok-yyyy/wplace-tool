/*
Copyright © 2025 ok-yyyy 192425431+ok-yyyy@users.noreply.github.com
*/
package cmd

import (
	"fmt"

	"github.com/ok-yyyy/wplace-tool/internal/wplace"
	"github.com/spf13/cobra"
)

// pixel2urlCmd represents the pixel2url command
var pixel2urlCmd = &cobra.Command{
	Use:   "pixel2url",
	Short: "Convert pixel coordinate range to a map URL",
	Long: `Converts a pixel coordinate (e.g., "100-200-30-40") to a corresponding map URL.

Argument format:
  <tileX>-<tileY>-<pixelX>-<pixelY>

Example:
  wplace-tool pixel2url 100-200-30-40
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Error: Missing argument. Please provide pixel coordinates in the format <tileX>-<tileY>-<pixelX>-<pixelY>.\n")
			return
		}

		loc, err := wplace.ParseLocation(args[0])
		if err != nil {
			fmt.Printf("Error: Invalid argument format. Please use <tileX>-<tileY>-<pixelX>-<pixelY>.\n")
			return
		}

		lat, lng := loc.LatLng()
		url := wplace.MapURL(lat, lng, 13.5)
		fmt.Printf("URL: %s\n", url)
	},
}

func init() {
	rootCmd.AddCommand(pixel2urlCmd)
}
