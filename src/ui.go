package main

import (
    "game/engine"
)

var (
    ui_selected = &Element{}
)


// Create UI Elements that dynamically change.
type Element struct {
    Value string
}

// Set Element Text
func (e *Element) Set(v string) {
    e.Value = v
}

// Get Element Text
func (e *Element) Get() string {
    return e.Value
}

// Create a UI Label that is statically set
func AddLabel(p *engine.Panel, X, Y int, label string, value string) {
    p.DrawCell(r, X, Y, label, value)
}

// Wrapper for all UI drawing functions
func UI_Draw(p *engine.Panel) {
    AddLabel(p, 1, 1, "Selected: %s", ui_selected.Get())
}

