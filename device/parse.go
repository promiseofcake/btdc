package device

import (
	"fmt"
	"os"

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
