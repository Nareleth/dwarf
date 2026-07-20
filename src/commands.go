package main


/*
Commands are used to issue orders to entities
*/

// CommandID creates a reference to commands to be used by systems.
type CommandID string

// CommandEvent is a message that sends to a system channel to execute an event
type CommandEvent struct {
    CommandID   CommandID
    Target      EntityID
}

// CommandDef is the information that the UI and input system use to allow player inpute
type CommandDef struct {
    Label   string
    Key     rune
}

// Initialize commands at game init
func InitCommands() map[CommandID]CommandDef {
    return map[CommandID]CommandDef{
        /* Commands */
        // Chop down an entity
        "cmd_chop": { Label: "Chop", Key: 'c' },
    }
}
