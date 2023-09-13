package pac_man

import (. "gfx2"
		"fmt"
		"block"
		)

type data struct {
	richtung int   // 0 = oben, 1 = rechts, 2 = unten, 3 = links
	verletzlich bool
	x,y uint16
	size uint16
	timer int
	
}

func New() *data {
	var p *data
	p = new(data)
	p.verletzlich = true
	return p
}


func (p *data) SetzeKoordinaten(x,y uint16) {
	p.x = x
	p.y = y
}

func (p *data) GibKoordinaten() (uint16, uint16) {
	return p.x, p.y
}

func (p *data) SetzeGröße (size uint16) {
	p.size = size
}

func (p *data) GibGröße () uint16 {
	return p.size
}


func (p *data) Kirsche () {
	p.timer = 30
	p.verletzlich = false
}

func (p *data) String () string {
	var erg string
	erg = "Pac-Man: "
	erg = erg + fmt.Sprint(p.verletzlich,p.x,p.y,p.timer)
	return erg
}

func (p *data) Draw() {
	if p.verletzlich {
		Stiftfarbe(255,255,0)
	}else {
		Stiftfarbe(0,255,255)
	}
		Vollkreis(p.x,p.y,p.size)
	if p.timer!=0 {
		p.timer--
	}
	if p.timer == 0 && !p.verletzlich {
		p.verletzlich = true
	}
}

func (p *data) ÄnderRichtung() {
	for {
		taste,gedrückt,_:=TastaturLesen1()
		if gedrückt==1 {
			switch taste {
				case 273:	// oben
					p.richtung = 0
				case 275:
					p.richtung = 1
				case 274:
					p.richtung = 2
				case 276:
					p.richtung = 3
			}
		}
	}
}

//func (p *data) blauerpixel () bool {
	//var breite,höhe uint16
	//switch p.richtung {
		//case 0:
		//if int(p.y)-int(p.size) -10 < 0
			//r,g,b:=GibPunktFarbe(p.x,p.y-p.size-10)
		//case 1:
			//r,g,b:=GibPunktFarbe(p.x+p.size+10,p.y)
		//case 2:
			//r,g,b:=GibPunktFarbe(p.x,p.y+p.size+10)
		//case 3:
			//r,g,b:=GibPunktFarbe(p.x-p.size-10,p.y)
	//}
//}

func (p *data) GibRichtung() int{
	return p.richtung
}

func (p *data) SetzeRichtung(richtung int) {
	p.richtung = richtung
}


func (p *data) Untermove () {
	var breite,höhe uint16
	breite = Grafikspalten()
	höhe = Grafikzeilen()
	switch p.richtung {
		case 0:
			p.y = p.y - 1
			if p.y < p.size {
				p.y = höhe - p.size
			}
		case 1:
			p.x = p.x + 1
			if p.x > breite-p.size {
				p.x = p.size
			}
		case 2:
			p.y = p.y + 1
			if p.y > höhe-p.size {
				p.y = p.size
			}
		case 3:
			p.x = p.x - 1
			if p.x < p.size {
				p.x = breite - p.size
			}
	}
}

func (p *data) Kollision (bl block.Block) bool {
	px,py:=p.GibKoordinaten()
	psize:=p.GibGröße()
	blx,bly:=bl.GibKoordinaten()
	blgröße:=bl.GibGröße()
	if px>=blx && px<=blx+blgröße && py>=bly && py<=bly+blgröße {
		return true
	} else if punktetest(px,py,psize,blx,bly,blgröße) {
		return true
	} else if rechtecktest(px,py,psize,blx,bly,blgröße) { 
		return true
	}
	return false
}

func punktetest (px,py,pr,bx,by,bg uint16) bool {
	if (int(px)-int(bx))*(int(px)-int(bx))+(int(py)-int(by))*(int(py)-int(by))<=int(pr)*int(pr) {
		return true
	}else if (int(px)-int(bx+bg))*(int(px)-int(bx+bg))+(int(py)-int(by))*(int(py)-int(by))<=int(pr)*int(pr) {
		return true
	}else if (int(px)-int(bx))*(int(px)-int(bx))+(int(py)-int(by+bg))*(int(py)-int(by+bg))<=int(pr)*int(pr) {
		return true
	}else if (int(px)-int(bx+bg))*(int(px)-int(bx+bg))+(int(py)-int(by+bg))*(int(py)-int(by+bg))<=int(pr)*int(pr) {
		return true
	}
	return false
}

func rechtecktest (px,py,pr,bx,by,bg uint16) bool {
	if px>=bx && px<=bx+bg && py<=by && py>=by-pr {
		return true
	}else if px>=bx && px<=bx+bg && py>=by+bg && py<=by+bg+pr {
		return true
	}else if px>=bx-pr && px<=bx && py<=by+bg && py>=by {
		return true
	}else if px<=bx+bg+pr && px>=bx+bg && py<=by+bg && py>=by {
		return true
	}
	return false
}
			
