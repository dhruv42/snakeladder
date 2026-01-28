package cmd

import "fmt"

// https://patorjk.com/software/taag/#p=display&f=Roman&t=Let%27s+Play+&x=rainbow3&v=4&h=4&w=80&we=false
func Banner() {
	title := `

	ooooo                      .   o8o               ooooooooo.   oooo                             
 888                    .o8    YP                 888    Y88.  888
 888          .ooooo.  .o888oo  '   .oooo.o       888   .d88'  888   .oooo.   oooo    ooo      
 888         d88'  88b   888       d88(  "8       888ooo88P'   888   P  )88b    88.  .8'       
 888         888ooo888   888        "Y88b.        888          888   .oP"888     88..8'        
 888       o 888    .o   888 .     o.  )88b       888          888  d8(  888      888'         
o888ooooood8  Y8bod8P'   "888"     8""888P'      o888o        o888o  Y888""8o     .8'          
                                                                              .o..P'           
                                                                               Y8P'            

	`
	fmt.Println(title)
}
