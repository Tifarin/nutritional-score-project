package main

import "fmt"

func main(){
	ns := GetNutritionalScore(NutritionalData{
		Energy: 				EnergyFromKcal(0),
		Sugars: 				SugarGram(10),
		SaturatedFattyAcids: 	SaturatedFattyAcids(2),
		Sodium: 				SodiumMilliGram(500),
		Fruits: 				FruitsPercent(60),
		Fibre: 					FibreGram(4),
		Protein: 				ProteinGram(2),
	},Food)
	fmt.Printf("Nutritional Score:%d\n", ns.Value)
	fmt.Printf("Nutri Score: %s\n", ns.GetNutriScore())
}