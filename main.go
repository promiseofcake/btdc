package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type input struct {
	Devices []string `yaml:"devices"`
}

type SavedAddresses map[string]interface{}

func (s SavedAddresses) IsSaved(addr string) bool {
	if _, ok := s[addr]; ok {
		return true
	}
	return false
}

type PairedDevices []BluetoothDevice

type BluetoothDevice struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func CheckDependencies() error {
	_, err := exec.Command("command", "-v", "blueutil").Output()
	if err != nil {
		return fmt.Errorf("blueutil is not found, please install and try again, %w", err)
	}
	return nil
}

func GetSaved(path string) (SavedAddresses, error) {
	saved := make(SavedAddresses)

	bts, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read path %s, %w", path, err)
	}

	var s input
	err = yaml.Unmarshal(bts, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse saved devices, %w", err)
	}

	for _, address := range s.Devices {
		saved[address] = nil
	}

	return saved, nil
}

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

func Unpair(addr string) error {
	cmd := exec.Command("/usr/local/bin/blueutil", "--format", "json", "--unpair", addr)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to unpair device %s, %w", addr, err)
	}
	return nil
}

func main() {
	var err error

	err = CheckDependencies()
	if err != nil {
		panic(err)
	}

	saved, err := GetSaved("./allow.yml")
	if err != nil {
		panic(err)
	}

	paired, err := GetPaired()
	if err != nil {
		panic(err)
	}

	var pass, fail int
	for _, device := range paired {
		if !saved.IsSaved(device.Address) {
			updateErr := Unpair(device.Address)
			if updateErr != nil {
				fmt.Printf("error, failed to unpair %s, err: %v\n", device.Address, updateErr)
				fail++
			} else {
				fmt.Printf("Successfully unpaired %s [%s]!\n", device.Name, device.Address)
				pass++
			}
		}
	}

	fmt.Printf("Unpaired %d devices, failed to unpair %d, devices!\n", pass, fail)
}
