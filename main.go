package main

import (
	"bufio"
	"fmt"
	mathrand "math/rand"
	"os"
	"os/exec"
	"strings"
)

func encrypt(key int, text string) string {
	shiftedChars := make([]rune, len(text))

	for index, char := range text {
		shiftedChars[index] = rune(int(char) + key)
	}
	return string(shiftedChars)
}

func decrypt(key int, encryptedText string) string {
	parts := strings.SplitN(encryptedText, ": ", 2)
	if len(parts) < 2 {
		return encryptedText
	}

	prefix, encryptedPassword := parts[0], parts[1]
	originalChars := make([]rune, len(encryptedPassword))

	for index, char := range encryptedPassword {
		originalChars[index] = rune(int(char) - key)
	}
	return prefix + ": " + string(originalChars)
}

func header() {
	fmt.Println(`
	 _____         _               _____           _     
	|_   _|       | |             |_   _|         | |    
	  | |  ___ ___| |_ ___  __ _    | | ___   ___ | |___ 
	  | | / __/ _ \ __/ _ \/ _  |   | |/ _ \ / _ \| / __|
	 _| || (_|  __/ ||  __/ (_| |   | | (_) | (_) | \__ \
	 \___/\___\___|\__\___|\__,_|   \_/\___/ \___/|_|___/																																									   
	`)
}

type PasswordSettings struct {
	Length int
	Save   bool
}

var passwordSettings PasswordSettings

var path string

func setPath() {
	path = "Icetea Tools\\"
}

func init() {
	setPath()
}

func clearCmd() {
	clearCommand := exec.Command("cmd", "/c", "cls")
	clearCommand.Stdout = os.Stdout
	clearCommand.Run()
}

func MainPage() {
	clearCmd()
	header()
	fmt.Printf("Path: %s\n\n", path)
	fmt.Print("1 - Generator\n")
	fmt.Print("2 - Manager\n")
	fmt.Print("\n99 - Exit")
	waitInput("")
}

func GeneratorMain() {
	clearCmd()
	path = "Icetea Tools\\Generator\\"
	fmt.Printf("Path: %s\n\n", path)
	fmt.Print("1 - Password\n")
	fmt.Print("\n99 - Back")
	waitInput("")
}

func ManagerMain() {
	clearCmd()
	path = "Icetea Tools\\Manager\\"
	fmt.Printf("Path: %s\n\n", path)
	fmt.Print("1 - Show Passwords\n")
	fmt.Print("\n99 - Back")
	waitInput("")
}

func PasswordGeneratorMain() {
	clearCmd()
	path = "Icetea Tools\\Generator\\Password\\"
	fmt.Printf("Path: %s\n\n", path)
	fmt.Print("1 - Password Length\n")
	fmt.Print("2 - Save Password in Manager (yes/no)\n")
	fmt.Print("3 - Generate Password\n")
	fmt.Print("\n99 - Back\n")
	PasswordSettingsInput()
}

func waitInput(input string) {
	fmt.Println(input)
	fmt.Print(">>> ")

	var responseInput string
	fmt.Scanln(&responseInput)

	if responseInput != "1" && responseInput != "2" && responseInput != "3" && responseInput != "99" {
		fmt.Print("Invalid input")
		waitInput(input)
	}
	switch responseInput {
	case "1":
		if path == "Icetea Tools\\Generator\\" {
			PasswordGeneratorMain()
		}
		if path == "Icetea Tools\\Manager\\" {
			ShowPasswordMain()
		}
		if path == "Icetea Tools\\" {
			GeneratorMain()
		}
	case "2":
		ManagerMain()
	case "99":
		if path != "Icetea Tools\\" {
			if path == "Icetea Tools\\Generator\\Password\\" {
				path = "Icetea Tools\\Generator\\"
				GeneratorMain()
			}
			if path == "Icetea Tools\\Generator\\" || path == "Icetea Tools\\Manager\\" {
				path = "Icetea Tools\\"
				MainPage()
			}
		} else {
			os.Exit(0)
		}
	}
}

func main() {
	MainPage()
}

func PasswordSettingsInput() {
	fmt.Print(">>> ")

	var responseInput string
	fmt.Scanln(&responseInput)

	if responseInput != "1" && responseInput != "2" && responseInput != "3" && responseInput != "99" {
		fmt.Print("Invalid input\n")
		PasswordSettingsInput()
	}
	switch responseInput {
	case "1":
		fmt.Print("Password Length: ")
		fmt.Scanln(&passwordSettings.Length)
		fmt.Printf("Password Length set to: %d\n", passwordSettings.Length)
		PasswordSettingsInput()
	case "2":
		fmt.Print("Save Password in Manager (yes/no): ")
		var response string
		fmt.Scanln(&response)
		passwordSettings.Save = response == "yes"
		fmt.Printf("Password will be saved: %t\n", passwordSettings.Save)
		PasswordSettingsInput()
	case "3":
		if passwordSettings.Length == 0 {
			fmt.Print("Password Length not set\n")
			PasswordSettingsInput()
		}
		if !passwordSettings.Save {
			fmt.Print("Password will not be saved\n")
			pass := generatePassword(passwordSettings.Length)
			fmt.Printf("Password Generated: %s\n", pass)
			PasswordSettingsInput()
		} else {
			generateAndSavePassword()
			PasswordSettingsInput()
		}
	}
	if responseInput == "99" {
		GeneratorMain()
	}
}

func ShowPasswordInput() {
	fmt.Print(">>> ")

	var responseInput string
	fmt.Scanln(&responseInput)

	if responseInput != "1" && responseInput != "99" {
		fmt.Print("Invalid input")
		ShowPasswordInput()
	}
	switch responseInput {
	case "1":
		ShowPasswordMain()
	case "99":
		ManagerMain()
	}
}

func ShowPasswordMain() {
	clearCmd()
	path = "Icetea Tools\\Manager\\Show Passwords\\"
	fmt.Printf("Path: %s\n\n", path)
	displayPasswords()
}

func displayPasswords() {
	var choicePasswords = 1
	file, err := os.Open("passwords.txt")
	if err != nil {
		fmt.Println("Error opening passwords file:", err)
		ShowPasswordInput()
	}

	var numberPasswords int
	scanner := bufio.NewScanner(file)
	fmt.Println("Stored Passwords:")
	fmt.Printf("\n")
	for scanner.Scan() {
		encryptedPassword := scanner.Text()
		decryptedPassword := decrypt(18, encryptedPassword)
		fmt.Printf("%d - %s\n", choicePasswords, decryptedPassword)
		numberPasswords++
		choicePasswords++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading passwords file:", err)
		ShowPasswordInput()
	}

	file.Close()
	fmt.Printf("\nTotal: %d\n", numberPasswords)
	fmt.Print("\n99 - Back\n")
	ShowPasswordInput()
}

func generateAndSavePassword() {
	fmt.Print("Password Name: ")
	var passwordName string
	fmt.Scanln(&passwordName)

	password := generatePassword(passwordSettings.Length)
	fmt.Printf("Password Generated: %s\n", password)

	if passwordSettings.Save {
		saveToFile(passwordName, password)
	}
}

func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=<>?/[]{}|"
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[mathrand.Intn(len(charset))]
	}
	return string(password)
}

func saveToFile(passwordName, password string) {
	encryptedPassword := encrypt(18, password)

	file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s: %s\n", passwordName, encryptedPassword))
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
	fmt.Printf("Password Saved\n")
}
