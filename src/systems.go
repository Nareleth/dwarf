package main


// Movement System
// Check the range of all entities with a velocity.
// If they have a position, then proceed to add their velocity.
func s_Move(c *Components) {
    for id, vx := range c.Velocity {
        if pos, ok := c.Position[id]; ok {
           pos.X += vx.DX 
           pos.Y += vx.DY 
        }
    }
}

/*
    Draw System:
    Check the range of all entities with a sprite. Assign key and value pair.
    If the entity has a position (needed to draw) then use the render engine to position cursor and draw.
*/
func s_Draw(c *Components) {
    for id, sprite := range c.Sprite {
        if pos, ok := c.Position[id]; ok {
            r.Move(pos.X, pos.Y)
            r.Text(string(sprite.Char))

        }
    }
}
