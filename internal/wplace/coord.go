package wplace

import (
	"fmt"
	"math"
)

const (
	// Web Mercator のズームレベル
	// タイルがx軸y軸ともに2048 (= 2^11) であることから
	ZoomLevel = 11
	// タイルサイズ (ピクセル単位)
	PixelsPerTile = 1000
)

type Location struct {
	TileX, TileY   int
	PixelX, PixelY int
}

func ParseLocation(s string) (Location, error) {
	// TODO: タイル座標のみの指定 (ピクセル座標は0とみなす)
	var loc Location
	n, err := fmt.Sscanf(s, "%d-%d-%d-%d", &loc.TileX, &loc.TileY, &loc.PixelX, &loc.PixelY)
	if err != nil || n != 4 {
		return Location{}, fmt.Errorf("invalid location format: %q", s)
	}
	return loc, nil
}

// Location から lat/lng を計算
// webメルカトル変換を利用 (https://en.wikipedia.org/wiki/Web_Mercator_projection)
func (l Location) LatLng() (lat, lng float64) {
	zoomTiles := math.Pow(2, ZoomLevel)

	px := float64(l.TileX) + float64(l.PixelX)/PixelsPerTile
	py := float64(l.TileY) + float64(l.PixelY)/PixelsPerTile

	// lat
	latRad := 2*math.Atan(math.Pow(math.E, math.Pi*(1-2*py/zoomTiles))) - math.Pi/2
	// lng
	lngRad := (2*px/zoomTiles - 1) * math.Pi

	return rad2Deg(latRad), rad2Deg(lngRad)
}

func rad2Deg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}
