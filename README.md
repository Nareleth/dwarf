# TO DO:
## Right now:
- rework DrawCell (what im using in UI) into a SetText func so we can clear text like setcell
- cursor hover
    - Clear old text
- UI Panel

## backlog
- hotkeys
- ui state
- colonist information

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
- clean up the Ui code. a lot of it can be used in gameengine

## Cleanup:
- proably need to make a gamestate to track cursor and ui shit in the future
- import game/engine is the name of projectmodule/package - both need to change later
- add way more comments and sort and organize engine

