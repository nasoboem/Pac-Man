package block

import ("gfx2"
		"fmt"
		)
type data struct {
	x,y uint16
	größe uint16
	r,g,b uint8
}

func New() *data {
	var bl *data
	bl = new(data)
	return bl
}

func (bl *data) String () string{
	var erg string
	erg = fmt.Sprint(bl.x,bl.y,bl.größe,bl.r,bl.g,bl.b)
	return erg
}

func (bl *data) GibGröße () uint16 {
	return bl.größe
}

func (bl *data) GibKoordinaten () (uint16,uint16) {
	return bl.x,bl.y
}


func (bl *data) Draw (x,y,größe uint16,r,g,b uint8) {
	bl.x = x
	bl.y = y
	bl.größe = größe
	bl.r = r
	bl.g = g
	bl.b = b
	gfx2.Stiftfarbe(r,g,b)
	gfx2.Vollrechteck(x,y,größe,größe)
}
	
