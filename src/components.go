package main


// Declare Component struct
type Components struct {
    Name            map[EntityID]*c_Name
    Sprite          map[EntityID]*c_Sprite
    Interactable    map[EntityID]*c_Interactable
    Position        map[EntityID]*c_Position
    Velocity        map[EntityID]*c_Velocity
}


// Init components
func NewComponents() Components {
    return Components{
        Name:           make(map[EntityID]*c_Name),
        Sprite:         make(map[EntityID]*c_Sprite),
        Interactable:   make(map[EntityID]*c_Interactable),
        Position:       make(map[EntityID]*c_Position),
        Velocity:       make(map[EntityID]*c_Velocity),
    }
}


/*** Components ***/

/* Bio */
// Name
type c_Name struct { Value string }

// Sprite
type c_Sprite struct { Char Sprite }

/* Movement */
// Position
type c_Position struct { X, Y int }

// Velocity
type c_Velocity struct { DX, DY int }

/* Flags */
// Interactions
type c_Interactable struct { Commands []CommandID }
