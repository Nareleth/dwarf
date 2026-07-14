package main

import (
    "os"
)

// Goroutine to handle keypress
func keyPress(cursor *Cursor) {
    key := make([]byte, 1)

    for {
        // Read keypress
        os.Stdin.Read(key)

        // Handle keypress
        switch key[0] {
            // Quit game
            case 'q', 3:
                r.DisableRawMode()
                r.ShowCursor()
                r.Clear()
                r.Flush()
                os.Exit(0)


            // Cursor Movement
            case 'w':
                cursor.Move(0, -1)
            case 'a':
                cursor.Move(-1, 0)
            case 's':
                cursor.Move(0, 1)
            case 'd':
                cursor.Move(1, 0)
        }
    }
}
