package main

import (
    "game/engine"
)


// Cursor struct
type Cursor struct {
    Sprite  Sprite
    X, Y    int
    Panel   *engine.Panel
    Camera  *Camera
}


// Generate player cursor
func NewCursor(p *engine.Panel, camera *Camera, x, y int) *Cursor {
    AbsX := p.X + x
    AbsY := p.Y + y
    char := sprite_cursor_x

    return &Cursor{
        Sprite: char, 
        X:      AbsX,
        Y:      AbsY,    
        Panel:  p,
        Camera: camera,
    }
}


// Draw player cursor
func (c *Cursor) Draw() {
    r.Move(c.X, c.Y)
    r.SetCell(c.X, c.Y, rune(c.Sprite))
}


// Move player cursor
func (c *Cursor) Move(dx, dy int) {
    c.X = c.X + dx
    c.Y = c.Y + dy

    // Clamp cursor in bounds
    if c.X <= c.Panel.X {
        c.X = c.Panel.X + 1
        c.Camera.X--
        
        // Clamp camera in bounds
        c.Camera.X = clamp(c.Camera.X, 0, mapsize - c.Panel.Width)
        c.Camera.Y = clamp(c.Camera.Y, 0, mapsize - c.Panel.Height)
    }

    if c.Y <= c.Panel.Y {
        c.Y = c.Panel.Y + 1
        c.Camera.Y--
        
        // Clamp camera in bounds
        c.Camera.X = clamp(c.Camera.X, 0, mapsize - c.Panel.Width)
        c.Camera.Y = clamp(c.Camera.Y, 0, mapsize - c.Panel.Height)
    }

    if c.X >= c.Panel.Width  {
        c.X = c.Panel.Width - 1
        c.Camera.X++
        
        // Clamp camera in bounds
        c.Camera.X = clamp(c.Camera.X, 0, mapsize - c.Panel.Width)
        c.Camera.Y = clamp(c.Camera.Y, 0, mapsize - c.Panel.Height)
        }
    if c.Y >= c.Panel.Height {
        c.Y = c.Panel.Height - 1 
        c.Camera.Y++
        
        // Clamp camera in bounds
        c.Camera.X = clamp(c.Camera.X, 0, mapsize - c.Panel.Width)
        c.Camera.Y = clamp(c.Camera.Y, 0, mapsize - c.Panel.Height)
    }
}

// Get name of hovered entity and show actions
func (c *Cursor) Hover(w *World) string {
    // Convert cursor coords to world coords
    mapX := c.X - c.Panel.X + c.Camera.X
    mapY := c.Y - c.Panel.Y + c.Camera.Y

    // Get hovered entity data
    id, hovered := w.GetEntityAt(mapX, mapY)
    if !hovered {
        return ""
    }
    if name, ok := w.Components.Name[id]; ok {
        return name.Value
    }
    return ""
}
