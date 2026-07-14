package main

import (
    "game/engine"
)


// Cursor struct
type Cursor struct {
    Sprite  Sprite
    X, Y    int
    Panel   engine.Panel
}


// Generate player cursor
func NewCursor(p engine.Panel, x, y int) *Cursor {
    AbsX := p.X + x
    AbsY := p.Y + y
    char := sprite_cursor_x

    return &Cursor{
        Sprite: char, 
        X:      AbsX,
        Y:      AbsY,    
        Panel:  p,
    }
}


// Draw player cursor
func (c *Cursor) Draw() {
    r.Move(c.X, c.Y)
    r.Text(string(c.Sprite))
}


// Move player cursor
func (c *Cursor) Move(dx, dy int) {
    c.X = c.X + dx
    c.Y = c.Y + dy

    // Clamp cursor in bounds
    if c.X <= c.Panel.X {c.X = c.Panel.X + 1}
    if c.Y <= c.Panel.Y {c.Y = c.Panel.Y + 1}
    if c.X >= c.Panel.Width  { c.X = c.Panel.Width - 1 }
    if c.Y >= c.Panel.Height { c.Y = c.Panel.Height - 1 }
}
