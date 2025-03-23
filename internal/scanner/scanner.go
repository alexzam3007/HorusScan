package scanner

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
)

// GetCPUInfo obtiene la información del procesador
func GetCPUInfo() (string, error) {
	cmd := exec.Command("wmic", "cpu", "get", "name")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error obteniendo CPU: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) > 1 {
		return strings.TrimSpace(lines[1]), nil
	}
	return "", fmt.Errorf("no se encontró información del CPU")
}

// GetRAMInfo obtiene la cantidad de RAM instalada en GB
func GetRAMInfo() (float64, error) {
	cmd := exec.Command("wmic", "memorychip", "get", "capacity")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("error obteniendo RAM: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var totalRAM uint64
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			val, err := strconv.ParseUint(line, 10, 64)
			if err == nil {
				totalRAM += val
			}
		}
	}

	return float64(totalRAM) / (1024 * 1024 * 1024), nil
}

// GetDiskInfo obtiene la capacidad total y libre de los discos
func GetDiskInfo() ([]string, error) {
	cmd := exec.Command("wmic", "logicaldisk", "get", "caption,size,freespace")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo discos: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var disks []string
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) == 3 {
			// Convertir los valores a uint64
			free, _ := strconv.ParseUint(fields[1], 10, 64)
			size, _ := strconv.ParseUint(fields[2], 10, 64)
			

			// Calcular el porcentaje de espacio libre (espacio libre / tamaño total)
			percentageFree := (float64(free) / float64(size)) * 100

			// Convertir el tamaño total y el espacio libre a GB
			totalGB := float64(size) / (1024 * 1024 * 1024)
			freeGB := float64(free) / (1024 * 1024 * 1024)

			// Agregar la información con el porcentaje, con el orden correcto (Total, Libre)
			disks = append(disks, fmt.Sprintf("%s - Total: %.2f GB, Libre: %.2f GB, %.2f%% disponible", fields[0], totalGB, freeGB, percentageFree))
		}
	}
	return disks, nil
}


// GetBatteryInfo obtiene el porcentaje de batería
func GetBatteryInfo() (string, error) {
	cmd := exec.Command("wmic", "path", "win32_battery", "get", "estimatedchargeremaining")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error obteniendo batería: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) > 1 {
		// Aquí quitamos el "%" al final de la cadena si el valor es "No disponible"
		battery := strings.TrimSpace(lines[1])
		if battery == "No disponible" {
			return battery, nil
		}
		return battery + "%", nil
	}
	return "No disponible", nil
}
