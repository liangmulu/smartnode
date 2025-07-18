package migration

import (
	"fmt"
	"strconv"
)

func upgradeFromV1160(serializedConfig map[string]map[string]string) error {
	pruneMemSize, exists := serializedConfig["nethermind"]["pruneMemSize"]
	if !exists {
		return fmt.Errorf("expected a section called `nethermind` with a setting called `pruneMemSize` but it didn't exist")
	}

	if pruneMemSize != "" {
		// Parse the pruneMemSize as an integer not using Sscanf
		var size int
		// Sscanf is not the right function to use here, we should use strconv.Atoi or similar
		size, err := strconv.Atoi(pruneMemSize)
		if err != nil {
			return fmt.Errorf("error parsing pruneMemSize: %w", err)
		}

		// If the size is less than 1280, set it to blank
		if size < 1280 {
			serializedConfig["nethermind"]["pruneMemSize"] = ""
		}
	}

	pruneMemBudget, exists := serializedConfig["nethermind"]["fullPruneMemoryBudget"]
	if !exists {
		return fmt.Errorf("expected a section called `nethermind` with a setting called `fullPruneMemoryBudget` but it didn't exist")
	}

	if pruneMemBudget != "" {
		size, err := strconv.Atoi(pruneMemBudget)
		if err != nil {
			return fmt.Errorf("error parsing pruneMemBudget: %w", err)
		}

		// If the size is less than 1280, set it to blank
		if size < 1280 {
			serializedConfig["nethermind"]["fullPruneMemoryBudget"] = ""
		}
	}

	return nil
}
