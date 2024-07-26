package main

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/world"
)

func main() {
	ctx := context.Background()
	log := logger.FromContext(ctx)

	debugPrint[world.Defs](log, "./assets/world/defs.bin")
	debugPrint[world.World](log, "./assets/world/world_office_0.bin")
	debugPrint[world.World](log, "./assets/world/world_test.bin")
}

func debugPrint[T any](log *logger.Logger, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("Error opening file: %s", err)
		os.Exit(1)
	}

	var value T
	err = gob.NewDecoder(file).Decode(&value)
	if err != nil {
		log.Errorf("Error decoding file: %s", err)
		os.Exit(1)
	}

	valueJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		log.Errorf("Error marshaling: %s", err)
		os.Exit(1)
	}

	fmt.Println("----[ " + path + " ]----")
	fmt.Println(string(valueJSON))
}
