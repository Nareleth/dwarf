package main

import (
    "math/rand/v2"
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


// Camera struct
type Camera struct {
    X, Y int
}

// Generic function to clamp objects within boundaries
func clamp(val, min, max int) int {
    if val < min { return min }
    if val > max { return max}
    return val
}


// Tilemap struct
type TileMap struct {
    Width, Height   int
    Tiles           [][]Tile
}

// Generate a tilemap
func GenerateTileMap(world *World, width, height int) *TileMap {
    // Init empty map to fill
    tilemap := make([][]Tile, height)

    // Init map seed
    src := rand.NewPCG(seed, 0)
    rng := rand.New(src)

    // Iterate through the rows
    for y := range tilemap {
        // Initialize a row of tiles
        tilemap[y] = make([]Tile, width)

        // Iterate through the columns
        for x := range tilemap[y] {
            tilemap[y][x] = Tile01

            // RNG generate trees (5% chance)
            if rng.IntN(100) < 5 {
                world.SpawnTree(x, y)
            }
        }
    }

    // Return map
    return &TileMap{ Width: width, Height: height, Tiles: tilemap }
}

// Draw tilemap
func (m *TileMap) Draw(p *engine.Panel, camera *Camera) {
    // Establish drawing area viewport
    // -- may not be necessary if i use drawcell func from engine
    viewportWidth   := p.Width  
    viewportHeight  := p.Height

    // Iterate through the tilemap
    for y := range m.Height {
        for x := range m.Width {
            // Snap to camera
            viewportX := x - camera.X
            viewportY := y - camera.Y

            // Viewbound strapping (render only in view)
            if viewportX <= p.X || viewportX >= viewportWidth || viewportY <= p.Y || viewportY >= viewportHeight { continue }
            

            // Establish absolute terminal printing coords
            ScreenX := p.X + viewportX
            ScreenY := p.X + viewportY

            // Draw tile
            r.Move(ScreenX, ScreenY)
            r.Text(string(m.Tiles[y][x].Sprite))
        }
    }
}
