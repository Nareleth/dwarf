package engine

import (
    "bufio"
    "fmt"
    "os"
    "golang.org/x/term"
)


/* Core */

// Create a Cell that holds runes for printing
type Cell struct {
    Char    rune
}

// Render struct
/*
    The renderer type is the main centerpiece for rendering to the terminal.
    Initiate the renderer once in the global scope using the NewRenderer function.
    w represents a bufio writer for flushing text to screen.
    The state is for toggling raw mode.
    Front and back represent a 2d grid of cells (text) that are drawn on screen vs what will be drawn.
*/
type Renderer struct {
    w       *bufio.Writer
    state   *term.State
    width   int
    height  int
    front   [][]Cell
    back    [][]Cell
}


// Create a new Renderer
func NewRenderer() *Renderer {
    w, h, _ := term.GetSize(int(os.Stdout.Fd()))
    front   := make([][]Cell, h)
    back    := make([][]Cell, h)

    for i := range front {
        front[i]    = make([]Cell, w)
        back[i]     = make([]Cell, w)
    }

    return &Renderer{
        w:      bufio.NewWriter(os.Stdout), 
        width:  w, 
        height: h,
        front:  front,
        back:   back,
        }
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

// Get terminal dimensions
func (r *Renderer) GetSize() (int, int) {
    width, height, _ := term.GetSize(int(os.Stdout.Fd()))
    return width, height
}

// Write a cells contents and allow back buffer flushing
func (r *Renderer) SetCell(x, y int, char rune) {
    r.back[y][x] = Cell{ Char: char }
}


// Flush buffer and print
func (r *Renderer) Flush() {
    for y := range r.height {
        for x := range r.width {
            if r.back[y][x] != r.front[y][x] {
                r.Move(x, y)
                r.Text(string(r.back[y][x].Char))
                r.front[y][x] = r.back[y][x]
            }
        }
    }

    // Clear back buffer
    for y := range r.back {
        for x := range r.back[y] {
            //r.back[y][x] = Cell{}
            r.back[y][x] = Cell{Char: ' '}
        }
    }

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
func (p *Panel) DrawCell(r *Renderer, childX, childY int, text string, args ...any) {
    absX := p.X + childX
    absY := p.Y + childY
    r.Move(absX, absY)
    r.Text(text, args...)
}

// Draw cell Text and allow back buffer flushing
func (p *Panel) SetText(r *Renderer, childX, childY int, text string){
    absX := p.X + childX
    absY := p.Y + childY
    for i, char := range text {
        r.SetCell(absX+i, absY, char)
    }
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
        switch row {
        case p.X:       r.SetCell(row, p.Y, br.RuneBoxNW)
        case length:    r.SetCell(row, p.Y, br.RuneBoxNE)
        default:        r.SetCell(row, p.Y, br.RuneBoxRow)
        }

        // Draw the bottom row
        switch row {
        case p.X:       r.SetCell(row, size, br.RuneBoxSW)
        case length:    r.SetCell(row, size, br.RuneBoxSE)
        default:        r.SetCell(row, size, br.RuneBoxRow)
        }

    }

    // Draw the y axis (col)
    for col := p.Y; col <= size; col++ {
        // Draw the left col
        switch col {
        case p.Y:   r.SetCell(p.X, col, br.RuneBoxNW)
        case size:  r.SetCell(p.X, col, br.RuneBoxSW)
        default:    r.SetCell(p.X, col, br.RuneBoxCol)
        }

        // Draw the right col
        switch col {
        case p.Y:   r.SetCell(length, col, br.RuneBoxNE)
        case size:  r.SetCell(length, col, br.RuneBoxSE)
        default:    r.SetCell(length, col, br.RuneBoxCol)
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
