package main

import (
	"fmt"
)

func main() {
	var lM, lC, rM, rC, userM, userC, k int
	lM, lC = 3, 3
	rM, rC = 0, 0
	k = 0

 
	fmt.Println("\nM M M C C C | --- | \n")

	for {
		for {
			fmt.Println("Left side -> right side river travel")

			fmt.Print("Enter number of Missionaries travel => ")
			fmt.Scan(&userM)
			fmt.Print("Enter number of Cannibals travel => ")
			fmt.Scan(&userC)

			if userM == 0 && userC == 0 {
				fmt.Println("Empty travel not possible")
				fmt.Println("Re-enter : ")
			} else if (userM+userC <= 2) && (lM-userM >= 0) && (lC-userC >= 0) {
				break
			} else {
				fmt.Println("Wrong input re-enter : ")
			}
		}

		lM -= userM
		lC -= userC
		rM += userM
		rC += userC

		fmt.Println()

		for i := 0; i < lM; i++ {
			fmt.Print("M ")
		}
		for i := 0; i < lC; i++ {
			fmt.Print("C ")
		}
		fmt.Print("| --> | ")
		for i := 0; i < rM; i++ {
			fmt.Print("M ")
		}
		for i := 0; i < rC; i++ {
			fmt.Print("C ")
		}
		fmt.Println()

		k++

		// Check if cannibals eat missionaries
		if (lC == 3 && lM == 1) || (lC == 3 && lM == 2) || (lC == 2 && lM == 1) ||
			(rC == 3 && rM == 1) || (rC == 3 && rM == 2) || (rC == 2 && rM == 1) {
			fmt.Println("Cannibals eat missionaries:\nYou lost the game")
			break
		}

		if rM+rC == 6 {
			fmt.Println("You won the game : \n\tCongrats")
			fmt.Println("Total attempts")
			fmt.Println(k)
			break
		}

		for {
			fmt.Println("Right side -> Left side river travel")

			fmt.Print("Enter number of Missionaries travel => ")
			fmt.Scan(&userM)
			fmt.Print("Enter number of Cannibals travel => ")
			fmt.Scan(&userC)

			if userM == 0 && userC == 0 {
				fmt.Println("Empty travel not possible")
				fmt.Println("Re-enter : ")
			} else if (userM+userC <= 2) && (rM-userM >= 0) && (rC-userC >= 0) {
				break
			} else {
				fmt.Println("Wrong input re-enter : ")
			}
		}

		lM += userM
		lC += userC
		rM -= userM
		rC -= userC

		k++

		fmt.Println()

		for i := 0; i < lM; i++ {
			fmt.Print("M ")
		}
		for i := 0; i < lC; i++ {
			fmt.Print("C ")
		}
		fmt.Print("| <-- | ")
		for i := 0; i < rM; i++ {
			fmt.Print("M ")
		}
		for i := 0; i < rC; i++ {
			fmt.Print("C ")
		}
		fmt.Println()

		// Check if cannibals eat missionaries
		if (lC == 3 && lM == 1) || (lC == 3 && lM == 2) || (lC == 2 && lM == 1) ||
			(rC == 3 && rM == 1) || (rC == 3 && rM == 2) || (rC == 2 && rM == 1) {
			fmt.Println("Cannibals eat missionaries:\nYou lost the game")
			break
		}
	}
}
