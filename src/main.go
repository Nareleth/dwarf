package main

import (
    "game/engine"
    "time"
)


// Declare global rendering
var r = engine.NewRenderer()

// Declare const vars
const (
    SetFPS      = 30
    simIdleRate = 100
    mapsize     = 50
)


func main() {
    // Raw Mode (for input)
    r.EnableRawMode()
    defer r.DisableRawMode()

    // Hide terminal cursor
    r.HideCursor()
    defer r.ShowCursor()

    // Drawing Panels
    panelCanvas := engine.Panel{
        X:      0,
        Y:      0,
        Width:  100,
        Height: 50,
    }

    panelCanvas.Style = engine.Light

    // New Cursor
    cursor := NewCursor(panelCanvas, 1, 1)

    // New World
    world := NewWorld()

    // Generate tilemap
    tilemap := GenerateTileMap(mapsize, mapsize)

    // For testing
    world.SpawnColonist("Jeff", 10, 10)
    world.SpawnColonist("Argo", 10, 5)

    // Goroutines
    go keyPress(cursor)    

    // Game loop
    ticker := time.NewTicker(time.Second / time.Duration(SetFPS))
    defer ticker.Stop()

    frames := 0
    lastFPS := time.Now()
    currentFPS := 0
    for range ticker.C {
        // Input
        frames++
        
        // Calculate FPS
        if time.Since(lastFPS) >= time.Second {
            currentFPS = frames
            frames = 0
            lastFPS = time.Now()
        }
        // so it shuts up and lets me compile
        _ = currentFPS

        // Update
        s_Idle(&world.Components)   // Idle Sim
        s_Move(&world.Components)   // Move entities

        // Draw
        r.Clear()                   // Clear Screen
        panelCanvas.DrawBorder(r)   // Panel Widget
        tilemap.Draw(panelCanvas)   // Tilemap
        s_Draw(&world.Components)   // Draw entities
        cursor.Draw()               // Draw cursor
        // Temp
        /*
        r.Move(1, 0)
        r.Text("Debug:")
        r.Move(1, 1)
        r.Text("FPS: %d", currentFPS)
        r.Move(10, 1)
        r.Text("Cursor Pos: (%d, %d)", cursor.X, cursor.Y)
        */
        // Flush all rendering
        r.Flush()

        // Sleep
        //time.Sleep(SetFPS * time.Millisecond) // Remove later
    }

}

