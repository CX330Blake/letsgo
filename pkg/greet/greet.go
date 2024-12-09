package greet

import (
	"github.com/fatih/color"
)

func Hello() {
	credit := `
______      _____                      
___  / _______  /______________ ______ 
__  /  _  _ \  __/_  ___/_  __ \/  __ \
_  /___/  __/ /_ _(__  )_  /_/ // /_/ /
/_____/\___/\__/ /____/ _\__, / \____/ 	  
                        /____/  
=======================================						               
Developed by CX330Blake      2024/12/09
=======================================

[~] Ready for traversal? Let's Goooooo!!!

`

	color.Cyan(credit)
	return
}

func End() {
	color.Cyan("[~] Finish!")
}
