package main

import (
    "game/engine"
)

// Elements will go in the engine, so they are built around that purpose.

// Create UI Elements that dynamically change.
type Element struct {
    Label   string
    Value   string
    X, Y    int
}

// Generate new elements
func NewElement(label, value string, x, y int) *Element {
    return &Element{
        Label:  label,
        Value:  value,
        X:      x,
        Y:      y,
    }
}

// Set Element Value
func (e *Element) Set(v string) {
    e.Value = v
}

// Get Element Value
func (e *Element) Get() string {
    return e.Value
}

// Create a wrapper for tracking the UI elements. Game logic talks here and UI logic reads here
type UIEngine struct {
    Elements    []*Element
    Panel       *engine.Panel
}

// Construct a new UIEngine
func NewUIEngine(panel *engine.Panel) *UIEngine {
    return &UIEngine {
        Elements:   make([]*Element, 0),
        Panel:      panel,
    }
}

// Create a new Element using this wrapper and draw it on screen.
func (ui *UIEngine) AddElement(label, value string, x, y int) {
    ui.Elements = append(ui.Elements, NewElement(label, value, x, y))
    
    //ui.Panel.DrawCell(r, x, y, label+value)
    ui.Panel.SetText(r, x, y, label+value)
}

// Draws all UI elements
func (ui *UIEngine) Draw() {
    for _, e := range ui.Elements{
        //ui.Panel.DrawCell(r, e.X, e.Y, e.Label+e.Value)
        ui.Panel.SetText(r, e.X, e.Y, e.Label+e.Value)
    }
}


/* Create the UI Elements */
// Initialize all UI Elements and draws them on the screen for the first time.
func (ui *UIEngine) Init() {
    ui.AddElement("Selected: ", "nil", 1, 1)
}
