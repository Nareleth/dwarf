package main


/*  
    ECS utilizes entities, components, and systems.
    An entity is just an int of an ID. Thats it. No object.
    The World holds the state of entities in it, and a global nextID.
    By initializing the components with maps, each id has a unique set of components.
    Then adding new components just calls a NewEntity func and populates what it needs.
*/

// Entity
type EntityID uint64

// World struct
type World struct {
    nextID      EntityID
    Components  Components
}

// Generate the world
func NewWorld() *World {
    return &World{ Components: NewComponents() }
}

// Create new entity
func (w *World) NewEntity() EntityID {
    id := w.nextID
    w.nextID++
    return id
}

// Get the entity from a location
func (w *World) GetEntityAt(X, Y int) (EntityID, bool) {
    for id, pos := range w.Components.Position {
        if pos.X == X && pos.Y == Y {
            return id, true
        }
    }

    return 0, false
}

/* Spawn functions - entity templates */
// Colonist
func (w *World) SpawnColonist(name string, x, y int) {
    id := w.NewEntity()
    c := &w.Components
    c.Name[id]      = &c_Name{ Value: name }
    c.Sprite[id]    = &c_Sprite{ Char: sprite_colonist }
    c.Position[id]  = &c_Position{ X: x, Y: y }
    c.Velocity[id]  = &c_Velocity{ DX: 0, DY: 0 }
}

// Tree
func (w *World) SpawnTree(x, y int) {
    id              := w.NewEntity()
    e               := &w.Components
    e.Name[id]      = &c_Name{ Value: "tree" }
    e.Sprite[id]    = &c_Sprite{ Char: sprite_tree }
    e.Position[id]  = &c_Position{ X: x, Y: y }
}

// Squirrel
func (w *World) SpawnSquirrel(x, y int) {
    id              := w.NewEntity()
    e               := &w.Components
    e.Name[id]      = &c_Name{ Value: "squirrel" }
    e.Sprite[id]    = &c_Sprite{ Char: sprite_squirrel }
    e.Position[id]  = &c_Position{ X: x, Y: y}
    e.Velocity[id]  = &c_Velocity{ DX: 0, DY: 0} 
}
