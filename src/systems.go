package main

import (
    "math/rand/v2"
    "game/engine"
)


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
func s_Draw(c *Components, p *engine.Panel, camera *Camera) {
    for id, sprite := range c.Sprite {
        if pos, ok := c.Position[id]; ok {

            // Render only entities in view
            localX := pos.X - camera.X
            localY := pos.Y - camera.Y

            if localX <= 0 || localX >= p.Width || localY <= 0 || localY >= p.Height { continue }            

            
            screenX := p.X + localX
            screenY := p.Y + localY
            r.Move(screenX, screenY)
            //r.Text(string(sprite.Char))
            r.SetCell(screenX, screenY, rune(sprite.Char))
        }
    }
}


/*
    Idle System:
    Simulate random wandering for entities.
    Iterate through every entity and roll a number. That number determines if their velocity changes or not.
*/
func s_Idle(c *Components) {
    for _, actor := range c.Velocity {
        switch rand.IntN(simIdleRate) {
            case 0:
                actor.DX = 0
                actor.DY = -1
            case 1:
                actor.DX = 0
                actor.DY = 1
            case 2:
                actor.DX = -1
                actor.DY = 0
            case 3:
                actor.DX = 1
                actor.DY = 0
            default:
                actor.DX = 0
                actor.DY = 0
                    
        }
    }
}
