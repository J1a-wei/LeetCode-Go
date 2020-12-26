package template 

const MAX_LAVEL int = 8

func recursive(level int,param1,param2 int){
	// recursion terminator 

	if level >  MAX_LAVEL {
		// preocess_result
		return 
	}

	// process logic in current level 

	process(level,data)

	// drill down 
	recursive(level + 1,param1,param2)

	
	// recursive the current level status if needed  

}

func process(level,data int){

}