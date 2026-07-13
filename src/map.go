package main

import (
    "game/engine"
)


// Tile struct
type Tile struct {
    Id      string
    Sprite  Sprite
}

// Initialize tile types
var (
    Tile00 = Tile{ Id: "null",  Sprite: Tile00Sprite }
    Tile01 = Tile{ Id: "grass", Sprite: Tile01Sprite }
)


// Tilemap struct
type TileMap struct {
    Width, Height   int
    Tiles           [][]Tile
}

// Generate a tilemap
func GenerateTileMap(width, height int) *TileMap {
    // Init empty map to fill
    tilemap := make([][]Tile, height)

    // Iterate through the rows
    for y := range tilemap {
        // Initialize a row of tiles
        tilemap[y] = make([]Tile, width)

        // Iterate through the columns
        for x := range tilemap[y] {
            tilemap[y][x] = Tile01
        }
    }

    // Return map
    return &TileMap{ Width: width, Height: height, Tiles: tilemap }
}

// Draw tilemap
func (m *TileMap) Draw(p engine.Panel) {
    // Establish drawing area viewport
    /* -- may not be necessary if i use drawcell func from engine
    viewportWidth   := p.Width  -1
    viewportHeight  := p.Height -1
    */

    // Iterate through the tilemap
    for y := range m.Height {
        for x := range m.Width {
            // Snap to camera
            viewportX := x - 0
            viewportY := y - 0

            // Viewbound strapping (render only in view)
            

            // Establish absolute terminal printing coords
            ScreenX := p.X + 1 + viewportX
            ScreenY := p.X + 1 + viewportY

            // Draw tile
            r.Move(ScreenX, ScreenY)
            r.Text(string(m.Tiles[y][x].Sprite))
        }
    }
}
