package component

import "encoding/json"

// Grid contains other Components
type Grid struct {
	base
	Config GridConfig `json:"config"`
}

// GridConfig is the contents of a Grid
type GridConfig struct {
	Panels []Panel `json:"panels"`
}

// NewGrid creates a grid component
func NewGrid(title string, panels ...Panel) *Grid {
	p := append([]Panel(nil), panels...) // Make a copy
	return &Grid{
		base: newBase(typeGrid, TitleFromString(title)),
		Config: GridConfig{
			Panels: p,
		},
	}
}

// GetMetadata accesses the components metadata. Implements Component.
func (t *Grid) GetMetadata() Metadata {
	return t.Metadata
}

// Add adds additional panels to the grid
func (t *Grid) Add(panels ...Panel) {
	t.Config.Panels = append(t.Config.Panels, panels...)
}

type gridMarshal Grid

// MarshalJSON implements json.Marshaler
func (t *Grid) MarshalJSON() ([]byte, error) {
	m := gridMarshal(*t)
	m.Metadata.Type = typeGrid
	return json.Marshal(&m)
}