package readtext

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadPath() string {

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error al obtener la ruta del ejecutable:", err)
		return ("Error in read")
	}

	// Obtén el directorio que contiene el ejecutable
	exeDir := filepath.Dir(exePath)

	//fmt.Println("Ruta del directorio del ejecutable:", exeDir)

	return (string(exeDir))

}
