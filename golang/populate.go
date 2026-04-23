package main

import (
	"fmt"
)

const batchSize = 500

// PopulateData populates tables with fake data
func (db *Database) PopulateData(generator *Generator) error {
	for _, table := range db.schema.Populate {
		if err := db.populateTable(table, generator); err != nil {
			return fmt.Errorf("failed to populate table %s: %w", table.Name, err)
		}
	}
	return nil
}

// populateTable populates a single table using transactions and multi-row batch inserts.
func (db *Database) populateTable(populate DbPopulate, generator *Generator) error {
	generator.ClearUniques()

	maxTableNameLen := len(populate.Name)
	maxCountLen := len(fmt.Sprintf("%d", populate.Count))

	// Determine stable column order from the first field list
	columns := make([]string, len(populate.Fields))
	for i, f := range populate.Fields {
		columns[i] = f.Name
	}

	var batch [][]interface{}

	flush := func(upTo int) error {
		if len(batch) == 0 {
			return nil
		}
		tx, err := db.conn.Begin()
		if err != nil {
			return fmt.Errorf("begin transaction: %w", err)
		}
		if err := db.InsertBatch(tx, populate.Name, columns, batch); err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit transaction: %w", err)
		}
		batch = batch[:0]

		progress := upTo * 100 / populate.Count
		fmt.Printf("\r%-*s | %*d entries | %d%%",
			maxTableNameLen, populate.Name,
			maxCountLen, upTo,
			progress)
		return nil
	}

	for i := 1; i <= populate.Count; i++ {
		commons := make(map[string]interface{})
		row := make([]interface{}, len(populate.Fields))

		for j, field := range populate.Fields {
			value, err := generator.Generate(field, commons)
			if err != nil {
				return fmt.Errorf("failed to generate value for field %s: %w", field.Name, err)
			}
			row[j] = value
		}

		batch = append(batch, row)

		if len(batch) >= batchSize {
			if err := flush(i); err != nil {
				return fmt.Errorf("failed to insert row %d: %w", i, err)
			}
		}
	}

	// Flush remaining rows
	if err := flush(populate.Count); err != nil {
		return fmt.Errorf("failed to insert final batch: %w", err)
	}

	fmt.Println()
	return nil
}
