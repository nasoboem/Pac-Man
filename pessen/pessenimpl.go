package pessen

import ( //"fmt"
		"gfx2")

type data struct {
	x,y uint16
	size uint16
	art int
	}
	
func New() *data {
	var e *data
	e = new(data)
	return e
}

func (e *data) SetzeKoordinaten (x,y uint16) {
	e.x = x
	e.y = y
}

func (e *data) SetzeGröße (size uint16) {
	e.size = size
}

func (e *data) SetzeArt (art int) {
	e.art = art
}

func (e *data) GibKoordinaten() (uint16, uint16) {
	return e.x, e.y
}

func (e *data) GibGröße () uint16 {
	return e.size
}

func (e *data) GibArt () int {
	return e.art
}

func (e *data) Draw () {
	if e.art == 1 {
		gfx2.Stiftfarbe(255,0,0)
	}else{
		gfx2.Stiftfarbe(255,255,0)
	}
	gfx2.Vollkreis(e.x,e.y,e.size)
}
