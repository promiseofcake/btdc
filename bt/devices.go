package bt

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type PairedDevices []BluetoothDevice

type BluetoothDevice struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// CheckDependencies ensures the correct dependencies are present to perform actions
// in this package
func CheckDependencies() error {
	_, err := exec.Command("command", "-v", "blueutil").Output()
	if err != nil {
		return fmt.Errorf("blueutil is not found, please install and try again, %w", err)
	}
	return nil
}

// GetPaired queries paried bluetooth devices and returns a list of them
func GetPaired() (PairedDevices, error) {
	out, err := exec.Command("blueutil", "--paired", "--format", "json").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to read call list paired devices, %w", err)
	}

	var paired PairedDevices
	err = json.Unmarshal(out, &paired)
	if err != nil {
		return nil, fmt.Errorf("failed to read paried devices, %w", err)
	}

	return paired, nil
}

// Unpair unpairs a bluetooth device based upon its address
func Unpair(addr string) error {
	cmd := exec.Command("/usr/local/bin/blueutil", "--format", "json", "--unpair", addr)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to unpair device %s, %w", addr, err)
	}
	return nil
}
