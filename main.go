package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type input struct {
	Devices []string `yaml:"devices"`
}

type deviceMap map[string]interface{}

type output struct {
	Devices []device
}

type device struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func main() {
	var err error
	bts, err := os.ReadFile("./allow.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var i input
	err = yaml.Unmarshal(bts, &i)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	kept := make(deviceMap)
	for _, add := range i.Devices {
		kept[add] = nil
	}

	out, err := exec.Command("/usr/local/bin/blueutil", "--paired", "--format", "json").Output()
	if err != nil {
		log.Fatal(err)
	}

	var devices []device
	err = json.Unmarshal(out, &devices)
	if err != nil {
		log.Fatal(err)
	}

	for _, o := range devices {
		if _, ok := kept[o.Address]; ok {
			fmt.Printf("Skipping %s [%s]\n", o.Name, o.Address)
			continue
		}

		out, err := exec.Command("/usr/local/bin/blueutil", "--format", "json", "--unpair", o.Address).Output()
		if err != nil {
			fmt.Printf("error, failed to unpair %s, output: %s, err: %v\n ", o.Address, out, err)
		}

		fmt.Printf("Disconnected %s [%s]!\n", o.Name, o.Address)
	}
}
