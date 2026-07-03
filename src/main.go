package main

/*
TO DO:
- import game/engine is the name of projectmodule/package - both need to change later
- add actual frame sleeping in the engine. can probably make a gameloop function
- proably need to make a gamestate to track cursor and ui shit in the future
- declare a screen.
- clamp cursor within screen

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


    // Goroutines
    go keyPress()    

    // Game loop
    for {
        // Input

        // Update


        // Draw
        r.Clear()
        cursor.Draw()

        // Flush all rendering
        r.Flush()

        // Sleep
        time.Sleep(16 * time.Millisecond) // Remove later
    }

}

