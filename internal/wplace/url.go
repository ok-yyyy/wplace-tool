package wplace

import "fmt"

func MapURL(lat, lng, zoom float64) string {
	return fmt.Sprintf("https://wplace.live/?lat=%f&lng=%f&zoom=%.1f", lat, lng, zoom)
}
