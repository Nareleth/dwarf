package main

/*
TO DO:
## Right now:
- make actors (using ECS method)


- import game/engine is the name of projectmodule/package - both need to change later
- add actual frame sleeping in the engine. can probably make a gameloop function
- proably need to make a gamestate to track cursor and ui shit in the future
- declare a screen.
- clamp cursor within screen
- draw screen border
- ui
- world struct
- actors
- ECS

*/

import (
    "game/engine"
    "time"
)

// Declare global rendering
var r = engine.NewRenderer()

// THIS SHOULDNT BE GLOBAL
var cursor = NewCursor('X', 1, 1)



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
        s_Move(&world.Components)

        // Draw
        r.Clear()                   // Clear Screen
        s_Draw(&world.Components)   // Draw entities
        cursor.Draw()               // Draw cursor

        // Flush all rendering
        r.Flush()

        // Sleep
        time.Sleep(16 * time.Millisecond) // Remove later
    }

}

