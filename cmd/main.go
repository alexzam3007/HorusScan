package main

import (
	"fmt"
	"HorusScan/internal/scanner"
)

func main() {

	// ----------------------- Bloque 0, Bienvenida a la app --------------------------------------- //

	// Mostrar mensaje de bienvenida
	fmt.Println("Bienvenido a HorusScan, la herramienta de diagnóstico y testeo de equipos. ¡Comencemos!")


	// ----------------------- Bloque 1, informacion del equipo ---------------------------------- //

	// Mostrar mensaje de informacion del equipo
	fmt.Println("Informacion del equipo:")

	// Obtener informacion del CPU
	cpu, err := scanner.GetCPUInfo()
	if err != nil {
		fmt.Println("Error obteniendo información del CPU:", err)
	} else {
		fmt.Printf("CPU: %s\n", cpu)
	}

	// Obtener informacion de la memoria RAM
	ram, err := scanner.GetRAMInfo()
	if err != nil {
		fmt.Println("Error obteniendo información de la RAM:", err)
	} else {
		fmt.Printf("RAM: %.2f GB\n", ram)
	}

	// Obtener informacion del disco
	disks, err := scanner.GetDiskInfo()
	if err != nil {
		fmt.Println("Error obteniendo información de los discos:", err)
	} else {
		fmt.Println("Información de los discos:")
		for _, disk := range disks {
			fmt.Println(disk)
		}
	}

	// Obtener informacion de la bateria
	battery, err := scanner.GetBatteryInfo()
	if err != nil {
		fmt.Println("Error obteniendo información de la batería:", err)
	} else {
		fmt.Printf("Batería: %s\n", battery)
	}

	// ----------------------- Bloque 2, Espera por el usuario --------------------------------- //

	// Mostrar mensaje de finalización y esperar entrada del usuario
	fmt.Println("\nPresiona Enter para empezar el test...")
	fmt.Scanln() // Espera la entrada del usuario para continuar


	// ----------------------- Bloque 3, Test del equipo ---------------------------------- //



}
