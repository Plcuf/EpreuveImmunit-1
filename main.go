package main

import (
	r "code/routeur" //Route vers mes routes
	t "code/temps"   //Route vers mes templates
)

func main() {
	t.InitTemplate() //Initialise mes templates
	r.InitServe()    //Initialise mes routes / assets et lance le serveur
}
