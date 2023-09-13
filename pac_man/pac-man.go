package pac_man

import ("pessen"
		"block"
		)

type Pac_Man interface {
	//Vor.:-
	//Eff.:Pac-Man hat jetzt die übergebenen Koordinaten x,y.
	SetzeKoordinaten(x,y uint16)
	
	//Vor.:-
	//Eff.:Pac-Man wird für 10 sec unverletzlich. Farbe helblau.
	Kirsche ()
	
	//Vor.:Grafikfenster (gfx-Paket) muss geöffnet sein.
	//Eff.:
	Draw()
	
	//Vor.:-
	//Erg.:Der Pac-Man wird als string zurückgegeben in der Formatierung
	// "Pac_Man: <verletzlich-bool>,x,y (Koordinaten), Zeit bis Verletzlich"
	String () string
	
	GibKoordinaten() (uint16, uint16)
	
	GibGröße () uint16
	
	SetzeGröße (size uint16)
	
	ÄnderRichtung()
	
	Kollision (bl block.Block) bool
	
	Untermove ()
	
	GibRichtung() int
	
	SetzeRichtung(richtung int)
}


func Move (p Pac_Man, b []block.Block) {
	var erg bool
	var pneu Pac_Man
	pneu = New()
	x,y:=p.GibKoordinaten()
	größe:=p.GibGröße()
	richtung:=p.GibRichtung()
	pneu.SetzeKoordinaten(x,y)
	pneu.SetzeGröße(größe)
	pneu.SetzeRichtung(richtung)
	pneu.Untermove()
	for i:=0;i<len(b);i++{
			if erg {
				break
			} else {
				if pneu.Kollision (b[i]) {
					erg = true
				}
			}
	}
	if !erg {
		p.Untermove()
	}
}




func Essen (p Pac_Man, e pessen.Essen) bool {
	px,py:=p.GibKoordinaten()
	psize:=p.GibGröße()
	ex,ey:=e.GibKoordinaten()
	return (int(px)-int(ex))*(int(px)-int(ex))+(int(py)-int(ey))*(int(py)-int(ey))<= int(psize)*int(psize)
}
