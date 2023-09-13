package pessen

type Essen interface {
	
	SetzeKoordinaten (x,y uint16)
	
	SetzeGröße (size uint16)
	
	SetzeArt (art int)
	
	GibKoordinaten() (uint16, uint16)
	
	GibGröße () uint16
	
	GibArt () int
	
	Draw ()
}
