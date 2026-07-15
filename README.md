# TO DO:
## Right now:
- UI Panel
- hotkeys
- ui state
- colonist information
- cursor hover

## Features:
- cursor highlighting (needs to account for the offset of viewport bounds and tile bounds)
- Work
- Build
- Hunt
- map size
- map generation algorithm
- Hunger
- collision

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
- create a gamestate for holding cursor info. so cursor can stop being global
- also use that to change how the key input routine works

## Cleanup:
- proably need to make a gamestate to track cursor and ui shit in the future
- import game/engine is the name of projectmodule/package - both need to change later
- add way more comments and sort and organize engine

