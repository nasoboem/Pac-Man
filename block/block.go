package block

type Block interface{

	Draw (x,y,größe uint16,r,g,b uint8)
	
	String () string
	
	GibGröße () uint16
	
	GibKoordinaten () (uint16,uint16)
}
