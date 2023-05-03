package main

import (
	"fmt"
	"github.com/promiseofcake/btdc/bt"
	"github.com/promiseofcake/btdc/device"
)

func main() {
	var err error

	err = bt.CheckDependencies()
	if err != nil {
		panic(err)
	}

	saved, err := device.GetSaved("./allow.yml")
	if err != nil {
		panic(err)
	}

	paired, err := bt.GetPaired()
	if err != nil {
		panic(err)
	}

	var pass, fail int
	for _, device := range paired {
		if !saved.IsSaved(device.Address) {
			updateErr := bt.Unpair(device.Address)
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
