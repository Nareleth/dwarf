package main


// Cursor struct
type Cursor struct {
    Char rune
    X, Y int
}


// Generate player cursor
func NewCursor(char rune, x, y int) *Cursor {
    return &Cursor{
        Char:   char, 
        X:      x,
        Y:      y,    
    }
}


// Draw player cursor
func (c *Cursor) Draw() {
    r.Move(c.X, c.Y)
    r.Text(string(c.Char))
}


// Move player cursor
func (c *Cursor) Move(dx, dy int) {
    c.X = c.X + dx
    c.Y = c.Y + dy

    // Clamp cursor in bounds
    if c.X < 0 {c.X = 0}
    if c.Y < 0 {c.Y = 0}
    //outer bounds here!!!!!!!!!!!!!!!!!!!!!!
}
