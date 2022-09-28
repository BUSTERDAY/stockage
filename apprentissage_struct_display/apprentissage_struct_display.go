package main

import "fmt"

type Voiture struct {
	model string
	body  string
	power int
}

func main() {
	//Initialisation de mes voitures
	var v1 Voiture
	var v2 Voiture

	//Initialisation de la voiture 1
	v1.model = "Audi TTS"
	v1.body = "coupé"
	v1.power = 245

	//Initialisation de la voiture 1
	v2.model = "Citroen Xzara phase 1"
	v2.body = "berline"
	v2.power = 90

	v1.display()
	v2.display()
	v2.gainPower()
	v2.gainPower()
}

func (v Voiture) display() {
	fmt.Println("----------------------")
	fmt.Println("Model : ", v.model)
	fmt.Println("Body : ", v.body)
	fmt.Println("Power : ", v.power)
}
func (v *Voiture) gainPower() {
	v.power += 100
	fmt.Println("New power : ", v.power)
}
