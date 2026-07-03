package engine

import (
    "bufio"
    "fmt"
    "os"
    "golang.org/x/term"
)


// Render struct
type Renderer struct {
    w       *bufio.Writer
    state   *term.State
}


// Create a new Renderer
func NewRenderer() *Renderer {
    return &Renderer{w: bufio.NewWriter(os.Stdout)}
}


// Enable Raw mode
func (r *Renderer) EnableRawMode() {
    oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        panic(err)    
    }
    r.state = oldState
}


// Disable Raw mode
func (r *Renderer) DisableRawMode() {
    term.Restore(int(os.Stdin.Fd()), r.state)
}


// Clear screen
func (r *Renderer) Clear() {
    fmt.Fprint(r.w, "\033[2J\033[H")
}


// Hide Cursor
func (r *Renderer) HideCursor() {
    fmt.Fprint(r.w, "\033[?25l")
}

// Show Cursor
func (r *Renderer) ShowCursor() {
    fmt.Fprint(r.w, "\033[?25h")
}

// Write Text
func (r *Renderer) Text(input string) {
    fmt.Fprint(r.w, input)
}


// Move terminal cursor for printing
func (r *Renderer) Move(x, y int) {
    fmt.Fprintf(r.w, "\033[%d;%dH", y+1, x+1)
}


// Flush buffer and print
func (r *Renderer) Flush() {
    r.w.Flush()
}
