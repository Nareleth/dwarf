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

    // Will update functionality later
    mapsize     = 200
    seed uint64 = 12345
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

    panelUI     := engine.Panel{
        X:      panelCanvas.X + panelCanvas.Width + 1,
        Y:      panelCanvas.Y,
        Width:  30,
        Height: panelCanvas.Height,
    }

    panelUI.Style = engine.Light

    // UI Engine
    ui := NewUIEngine(&panelUI)

    // Camera
    camera := Camera{X: 0, Y: 0}

    // New Cursor
    cursor := NewCursor(&panelCanvas, &camera, 1, 1)

    // New World
    world := NewWorld()

    // Generate tilemap
    tilemap := GenerateTileMap(world, mapsize, mapsize)

    // For testing
    world.SpawnColonist("Jeff", 10, 10)
    world.SpawnColonist("Argo", 10, 5)
    world.SpawnSquirrel(11,11)
    world.SpawnSquirrel(20,11)
    world.SpawnSquirrel(11,20)

    // Goroutines
    go keyPress(cursor)    

    // Clear screen
    r.Clear()

    // Init UI
    ui.Init()
    

    /* Game loop */
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
        s_Idle(&world.Components)               // Idle Sim
        s_Move(&world.Components)               // Move entities
        ui.Elements[0].Set(cursor.Hover(world)) // Cursor Hover
        

        // Draw
        panelCanvas.DrawBorder(r)                           // Panel Widget Game
        panelUI.DrawBorder(r)                               // Panel Widget UI
        ui.Draw()                                           // UI
        tilemap.Draw(&panelCanvas, &camera)                 // Tilemap
        s_Draw(&world.Components, &panelCanvas, &camera)    // Draw entities
        cursor.Draw()                                       // Draw cursor

        // Flush all rendering
        r.Flush()

    }

}

