//Explanation:
//Intrinsic State (Glyph): The characters themselves are stored once, shared across different usages.
//Extrinsic State (position): The position of each character in the document is unique and passed in during rendering.
//Glyph Factory: Manages the creation and storage of glyph objects, ensuring that each character is instantiated only once.

package main

import (
	"fmt"
)

// Glyph represents the intrinsic state of a character (static)
type Glyph struct {
	character string
}

// Render method simulates rendering the glyph (dynamic)
func (g *Glyph) Render(extrinsicState string) {
	fmt.Printf("Rendering character %s at position %s\n", g.character, extrinsicState)
}

// GlyphFactory manages flyweight objects to ensure characters are shared
type GlyphFactory struct {
	glyphs map[string]*Glyph
}

// NewGlyphFactory creates a new GlyphFactory
func NewGlyphFactory() *GlyphFactory {
	return &GlyphFactory{
		glyphs: make(map[string]*Glyph),
	}
}

// GetGlyph retrieves a shared glyph object or creates it if it doesn't exist
func (f *GlyphFactory) GetGlyph(character string) *Glyph {
	if glyph, exists := f.glyphs[character]; exists {
		return glyph
	}
	glyph := &Glyph{character: character}
	f.glyphs[character] = glyph
	return glyph
}

func main() {
	factory := NewGlyphFactory()

	// Example of using repeated characters as flyweights
	document := []struct {
		character string
		position  string
	}{
		{"A", "0,0"},
		{"B", "0,1"},
		{"A", "0,2"},
		{"C", "1,0"},
		{"B", "1,1"},
		{"A", "1,2"},
	}

	for _, entry := range document {
		glyph := factory.GetGlyph(entry.character)
		glyph.Render(entry.position)
	}
}
