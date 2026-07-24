# TO DO:
## Right now:
- Commands:
    - Create the hotkey interaction (can only run when hovered) (will also need to reference world and gamestate in teh keypress func...)
    - Create the system channel to listen for input
- Work System
    - create entity state component (so they arent only idling)
    - Create moving to pos state
    - Create worker queue
    - Create the chop logic
    - Create Despawn entity func
- Colony
    - Init + resources
    - add resource func - add wood to colony when tree is chopped


## Features:
- Work
- resources
- Build
- Hunt
- map size
- map generation algorithm
- Hunger
- collision
- entity inventory
- colonists should carry resources to storage for it to add to colony total

## Map:
- fixed map size types + declarations
- use drawcell for map drawing

## Colony:
- create starting scenario rather than hardcoded colonists
- actor names

## Engine:
- color
- get terminal size
- move runes to its own file
- clean up panel section
- put panel in its own sub directory engine/panel
- fps + gameloop funcs
- debug mode

## UI:
- create UI panel
    - nested panel?
- console tab for showing debug info or running "console commands"
- tileset

## QOL:
- clean up the ui updating section in main
- create a gamestate for holding cursor info. so cursor can stop being global
- also use that to change how the key input routine works
- clean up the Ui code. a lot of it can be used in gameengine
- create a "text section" thing in the engine so cursor hover actions can go there

## backlog
- ui state
- colonist information

## Cleanup:
- proably need to make a gamestate to track cursor and ui shit in the future
- import game/engine is the name of projectmodule/package - both need to change later
- add way more comments and sort and organize engine

