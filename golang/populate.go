package main

import (
	"fmt"
)

// PopulateData populates tables with fake data
func (db *Database) PopulateData(generator *Generator) error {
	for _, table := range db.schema.Populate {
		if err := db.populateTable(table, generator); err != nil {
			return fmt.Errorf("failed to populate table %s: %w", table.Name, err)
		}
	}
	return nil
}

// populateTable populates a single table
func (db *Database) populateTable(populate DbPopulate, generator *Generator) error {
	generator.ClearUniques()

	maxTableNameLen := len(populate.Name)
	maxCountLen := len(fmt.Sprintf("%d", populate.Count))

	for i := 1; i <= populate.Count; i++ {
		commons := make(map[string]interface{})
		data := make(map[string]interface{})

		for _, field := range populate.Fields {
			value, err := generator.Generate(field, commons)
			if err != nil {
				return fmt.Errorf("failed to generate value for field %s: %w", field.Name, err)
			}
			data[field.Name] = value
		}

		if err := db.Insert(populate.Name, data); err != nil {
			return fmt.Errorf("failed to insert row %d: %w", i, err)
		}

		// Progress logging
		progress := i * 100 / populate.Count
		fmt.Printf("\r%-*s | %*d entries | %d%%",
			maxTableNameLen, populate.Name,
			maxCountLen, i,
			progress)
	}

	fmt.Println() // New line after completion
	return nil
}
