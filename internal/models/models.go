package models

import "log"

type Line struct {
	Text   string
	Prefix string
	Number int
	PosX   float32
	PosY   float32
	Style  *Style
}

type Style struct {
	AsString string
	Font     string
	Size     int
	Margin   int
}

func (l *Line) GetChords() {
	log.Println(*l.Style)
	l.PosX = 10
	l.PosY = float32(l.Style.Size + l.Style.Margin * l.Number)
}
