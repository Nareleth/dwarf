package main

/*
TO DO:
## Right now:


## Features:
- Map
- Trees
- Creatures
- Build
- Hunt
- Hunger

## Engine:
- Screen + widgets
- fps + gameloop
- debug mode

## QoL:
- add actual frame sleeping in the engine. can probably make a gameloop function
- sprites file

## UI:
- declare a screen.
- adjust draw positions relative to the parent widget
- clamp cursor within screen
- draw screen border

## Cleanup:
- proably need to make a gamestate to track cursor and ui shit in the future
- import game/engine is the name of projectmodule/package - both need to change later
*/

import (
    "game/engine"
    "time"
)

// Declare global rendering
var r = engine.NewRenderer()

// THIS SHOULDNT BE GLOBAL
var cursor = NewCursor('X', 1, 1)

// Declare const vars
const (
    SetFPS      = 30
    simIdleRate = 100
)


func main() {
    // Raw Mode (for input)
    r.EnableRawMode()
    defer r.DisableRawMode()

    // Hide terminal cursor
    r.HideCursor()
    defer r.ShowCursor()

    // New Cursor

    // New World
    world := NewWorld()

    // For testing
    world.SpawnColonist("Jeff", 10, 10)
    world.SpawnColonist("Argo", 10, 5)

    // Goroutines
    go keyPress()    

    // Game loop
    ticker := time.NewTicker(time.Second / time.Duration(SetFPS))
    defer ticker.Stop()

    frames := 0
    lastFPS := time.Now()
    currentFPS := 0
    for range ticker.C {
        // Input
        frames++
        
        if time.Since(lastFPS) >= time.Second {
            currentFPS = frames
            frames = 0
            lastFPS = time.Now()
        }

        // Update
        s_Idle(&world.Components)   // Idle Sim
        s_Move(&world.Components)   // Move entities

        // Draw
        r.Clear()                   // Clear Screen
        s_Draw(&world.Components)   // Draw entities
        cursor.Draw()               // Draw cursor
        r.Move(1, 0)
        r.Text("Debug:")
        r.Move(1, 1)
        r.Text("FPS: %d", currentFPS)

        // Flush all rendering
        r.Flush()

        // Sleep
        //time.Sleep(SetFPS * time.Millisecond) // Remove later
    }

}

