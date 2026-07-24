package main


/*
Commands are used to issue orders to entities
- The command ID is a direct reference to a command to be used.
- CommandEvents are pushed to a system channel so that an entity will perform the command.
- CommandDef is what the UI uses to display commands to the player.
- Init Commands initializes the list of all commands at game start.
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
