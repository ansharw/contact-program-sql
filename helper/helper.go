package helper

import (
	"contact-program-fundamental/model"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fungsi untuk kembali ke menu
func BackHandler() {
	fmt.Print("Tekan enter untuk kembali ke menu")
	var back int
	fmt.Scanln(&back)
}

func PhoneToString(phoneDatas []model.Phone) string {
	var stringPhone []string
	for _, v := range phoneDatas{
		_, phone := v.GetPhone()
		stringPhone = append(stringPhone, *phone)
	}
	return strings.Join(stringPhone, ", ")
}