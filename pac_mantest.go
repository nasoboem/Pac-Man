package main

import (. "gfx2"
		"pac_man"
		//"fmt"
		. "zufallszahlen"
		//"time"
		"pessen"
		"block"
		)

func main () {
	Fenster (600,400)
	Randomisieren()
	var p pac_man.Pac_Man
	p = pac_man.New()
	p.SetzeKoordinaten(100,100)
	p.SetzeGröße(20)
	go p.ÄnderRichtung()
	var e pessen.Essen
	e = pessen.New()
	e.SetzeGröße(5)
	x:=uint16(Zufallszahl(5,595))
	y:=uint16(Zufallszahl(5,395))
	e.SetzeKoordinaten(x,y)
	var bloecke []block.Block
	for i:=0;i<10;i++{
		var b block.Block
		b = block.New()
		b.Draw(uint16(Zufallszahl(0,560)),uint16(Zufallszahl(0,360)),2*p.GibGröße(),0,0,255)
		bloecke = append(bloecke,b)
	}
	for {
		UpdateAus()
		Stiftfarbe(255,255,255)
		Cls()
		pac_man.Move(p,bloecke)
		if pac_man.Essen(p,e) {
			x:=uint16(Zufallszahl(5,595))
			y:=uint16(Zufallszahl(5,395))
			z:=Zufallszahl(1,100)
			if e.GibArt() == 1 {
				p.Kirsche()
			}
			if z<=10 {
				e.SetzeArt(1)
			} else {
				e.SetzeArt(0)
			}
			e.SetzeKoordinaten(x,y)
			
		}
		p.Draw()
		for i:=0;i<len(bloecke);i++{
			x,y:=bloecke[i].GibKoordinaten()
			r:=bloecke[i].GibGröße()
			bloecke[i].Draw(x,y,r,0,0,255)
		}
		e.Draw()
		UpdateAn()
		//time.Sleep(1*time.Millisecond)
	}
	TastaturLesen1()
}
	
	
