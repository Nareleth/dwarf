package engine

import (
    "bufio"
    "fmt"
    "os"
    "golang.org/x/term"
)


/* Core */

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
func (r *Renderer) Text(input string, args ...any) {
    fmt.Fprintf(r.w, input, args...)
}


// Move terminal cursor for printing
func (r *Renderer) Move(x, y int) {
    fmt.Fprintf(r.w, "\033[%d;%dH", y+1, x+1)
}


// Flush buffer and print
func (r *Renderer) Flush() {
    r.w.Flush()
}


/* Cell Layout */
// Create a drawing Panel
type Panel struct {
    X, Y            int         // Absolute terminal position (top left corner)
    Width, Height   int         // Size of panel
    Style           BorderStyle // Style of border
}

// Draw cell content
func (p *Panel) DrawCell(r *Renderer, childX, childY int, text string) {
    absX := p.X + childX
    absY := p.Y + childY
    r.Move(absX, absY)
    r.Text(text)
}

// Draw panel borders
func (p *Panel) DrawBorder(r *Renderer) {
    length  := p.X + p.Width
    size    := p.Y + p.Height

    br := p.borderRunes()
    if br.RuneBoxRow == 0 {
        return
    }

    // Draw the x axis (row)
    for row := p.X; row <= length; row++ {
        // Draw the top row
        r.Move(row, p.Y)

        switch row {
        case p.X:       r.Text(string(br.RuneBoxNW))
        case length:    r.Text(string(br.RuneBoxNE))
        default:        r.Text(string(br.RuneBoxRow))
        }

        // Draw the bottom row
        r.Move(row, size)

        switch row {
        case p.X:       r.Text(string(br.RuneBoxSW))
        case length:    r.Text(string(br.RuneBoxSE))
        default:        r.Text(string(br.RuneBoxRow))
        }

    }

    // Draw the y axis (col)
    for col := p.Y; col <= size; col++ {
        // Draw the left col
        r.Move(p.X, col)

        switch col {
        case p.Y:   r.Text(string(br.RuneBoxNW))
        case size:  r.Text(string(br.RuneBoxSW))
        default:    r.Text(string(br.RuneBoxCol))
        }

        // Draw the right col
        r.Move(length, col)

        switch col {
        case p.Y:   r.Text(string(br.RuneBoxNE))
        case size:  r.Text(string(br.RuneBoxSE))
        default:    r.Text(string(br.RuneBoxCol))
        }
    }
    
}

// Style panel borders
type BorderStyle int

const (
    None BorderStyle = iota
    Light
    Heavy
)

type BorderRunes struct {
    RuneBoxRow, RuneBoxCol, RuneBoxNW, RuneBoxNE, RuneBoxSW, RuneBoxSE rune
}

func (p *Panel) borderRunes() BorderRunes {
    switch p.Style {
    case Light:
        return BorderRunes{
            RuneBoxRow: RuneBoxLightRow,
            RuneBoxCol: RuneBoxLightCol,
            RuneBoxNW:  RuneBoxLightNW,
            RuneBoxNE:  RuneBoxLightNE,
            RuneBoxSW:  RuneBoxLightSW,
            RuneBoxSE:  RuneBoxLightSE,
        }
    case Heavy:
        return BorderRunes{
            RuneBoxRow: RuneBoxHeavyRow,
            RuneBoxCol: RuneBoxHeavyCol,
            RuneBoxNW:  RuneBoxHeavyNW,
            RuneBoxNE:  RuneBoxHeavyNE,
            RuneBoxSW:  RuneBoxHeavySW,
            RuneBoxSE:  RuneBoxHeavySE,
        }
    default:
        return BorderRunes{}
    }
}



/* Runes */
const (
    /* Box*/
    //Light
    RuneBoxLightRow  = '\u2500'
    RuneBoxLightCol  = '\u2502'
    RuneBoxLightNW   = '\u250C'
    RuneBoxLightNE   = '\u2510'
    RuneBoxLightSW   = '\u2514'
    RuneBoxLightSE   = '\u2518'

    // Heavy
    RuneBoxHeavyRow  = '\u2501'
    RuneBoxHeavyCol  = '\u2503'
    RuneBoxHeavyNW   = '\u250F'
    RuneBoxHeavyNE   = '\u2513'
    RuneBoxHeavySW   = '\u2517'
    RuneBoxHeavySE   = '\u251B'
)
