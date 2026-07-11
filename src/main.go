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
    SetFPS      = 16
    simIdleRate = 7
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
    for {
        // Input

        // Update
        s_Idle(&world.Components)   // Idle Sim
        s_Move(&world.Components)   // Move entities

        // Draw
        r.Clear()                   // Clear Screen
        s_Draw(&world.Components)   // Draw entities
        cursor.Draw()               // Draw cursor

        // Flush all rendering
        r.Flush()

        // Sleep
        time.Sleep(SetFPS * time.Millisecond) // Remove later
    }

}

