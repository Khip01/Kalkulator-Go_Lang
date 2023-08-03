package main

import (
	Console "fmt"
	"os"
	System "os/exec"
	"strconv"
	"time"
)

func main() {
	// Persiapan rand
	// source := rand.NewSource(time.Now().UnixNano())
	// random := rand.New(source)

	// textLength := random.Perm(len(textIntro)) // Panjang random text

	// Clear CMD First
	clearScreen()

	// Showing Text Intro
	textIntro := "Welcome in Calculator (Go-Lang)"
	showWelcome(textIntro)

	// Show Menu
	showMenu()

	// Looping system if user want to do it again
	doAgain := true
	for doAgain {

		// Program Start
		// Input Operator Area
		var input string      // input string
		var operatorInput int // to int

		Console.Print("Input Operator = ")

		for {
			// Check Input must be Int
			operatorInput = checkInputInt(input, operatorInput)

			var ans string
			Console.Print("Are you sure? (y/n)...")
			Console.Scan(&ans)
			if ans == "y" || ans == "Y" {
				// // clear line
				// clearLine()
				// // clear line
				// clearLine()
				break
			} else {
				clearScreen()
				showMenu()
				Console.Print("Input Operator: ")
			}
		}

		// Displays the selected operator
		Console.Println("\n You Choose ", operatorName(operatorInput), "\n")

		time.Sleep(2000 * time.Millisecond)

		//// Enter Number
		// First Number
		var num1 int
		num1 = showInputNum("Enter your first number")
		// Second Number
		var num2 int
		num2 = showInputNum("Enter your second number")

		Console.Println("\nNum 1 =", num1)
		Console.Println("Num 2 =", num2)

		// Process Number
		processAritmethic(operatorInput, num1, num2)

		// Condition to do it again/not
		var doAgainString string
		Console.Print("Do you want to do the operation again?(y/n)")
		Console.Scan(&doAgainString)

		if doAgainString != "y" {
			doAgain = false
			break
		} else {
			// Clear CMD First
			clearScreen()

			// Show Menu
			showMenu()
		}
	}

	// END LINE
	// reader := bufio.NewReader(os.Stdin)
	// Console.Print("Press Enter to continue...")
	// reader.ReadString('k')
	Console.Print("\nBye")

	time.Sleep(2000 * time.Millisecond)
}

//----//// METHOD AREA ////----//

func clearScreen() {
	cmd := System.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showWelcome(textIntro string) {
	// Show Text Animation
	for i := 0; i < len(textIntro); i++ {
		Console.Print(string(textIntro[i]))
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(2000 * time.Millisecond)

	// Delete Text Animation
	for j := 0; j < len(textIntro); j++ {
		Console.Print("\b \b")
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(1000 * time.Millisecond)
}

func showMenu() {
	Console.Println(".____________________.")
	Console.Println("|___TABEL OPERATOR___|")
	Console.Println("|  1. Sum            |")
	Console.Println("|  2. Substraction   |")
	Console.Println("|  3. Division       |")
	Console.Println("|  4. Multiplication |")
	Console.Println("|  5. Modulus        |")
	Console.Println("|____________________|")
}

func checkInputInt(inputString string, inputInteger int) int {
	// Check Input must be Int
	for {
		// Scan Input and Parse to Int
		_, scannedInput := Console.Scan(&inputString)
		inputInteger, scannedInput = strconv.Atoi(inputString)

		// Check result is int or string
		if scannedInput != nil || inputInteger < 1 || inputInteger > 5 {
			clearScreen()
			showMenu()
			Console.Print("[! Input must be a Number 1 - 5] Input Operator = ")
		} else {
			break
		}
	}

	return inputInteger
}

func clearLine() {
	// Source: https://groups.google.com/g/golang-nuts/c/k6l_LhI8CO0?pli=1

	// ClearLine is the CSI sequence to clear the entire of the current line.
	const ClearLine = "\033[2K"

	// clear line
	Console.Printf(ClearLine)

	// clear line
	Console.Printf(ClearLine)

	// the line is cleared but the cursor is in the wrong place. the carriage
	// return moves the cursor to the beginning of the line.
	Console.Printf("\r")
}

func operatorName(operatorInput int) string {
	var operatorInputString string

	switch operatorInput {
	case 1:
		operatorInputString = "Sum"
		break
	case 2:
		operatorInputString = "Substraction"
		break
	case 3:
		operatorInputString = "Division"
		break
	case 4:
		operatorInputString = "Multiplication"
		break
	case 5:
		operatorInputString = "Modulus"
		break
	}

	return operatorInputString
}

func showInputNum(textInput string) int {
	var numReturn int

	Console.Print(textInput, ": ")

	for {
		var num1String string
		Console.Scan(&num1String)

		// Is Input = 0?
		if num1String == "0" {
			numReturn = 0
			break
		}

		// Whether the input includes letters or numbers
		num, errCatch := strconv.Atoi(num1String)
		if errCatch != nil {
			// // Clear Line
			// clearLine()
			Console.Print("[! Should be a Number] ", textInput, ": ")
		} else {
			numReturn = num
			break
		}
	}

	return numReturn
}

func processAritmethic(operatorInput int, num1 int, num2 int) {

	var operatorInputString string
	var numResult int

	switch operatorInput {
	case 1:
		operatorInputString = " + "
		numResult = num1 + num2
		break
	case 2:
		operatorInputString = " - "
		numResult = num1 - num2
		break
	case 3:
		operatorInputString = " : "
		numResult = num1 / num2
		break
	case 4:
		operatorInputString = " x "
		numResult = num1 * num2
		break
	case 5:
		operatorInputString = " % "
		numResult = num1 % num2
		break
	}

	// Display Aritmethic
	Console.Print("\nProcessing:\n", num1, operatorInputString, num2, "\nResult: ")

	for i := 1; i <= 2; i++ {
		Console.Print(".")
		time.Sleep(500 * time.Millisecond)
		Console.Print("\b \b")
		Console.Print("..")
		time.Sleep(500 * time.Millisecond)
		Console.Print("\b \b")
		Console.Print("\b \b")
		Console.Print("...")
		time.Sleep(500 * time.Millisecond)
		Console.Print("\b \b")
		Console.Print("\b \b")
		Console.Print("\b \b")
	}

	Console.Println(numResult)

	time.Sleep(1000 * time.Millisecond)
}
