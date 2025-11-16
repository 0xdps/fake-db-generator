package generator

import (
	"fmt"
	"testing"
)

func TestGenerateRecords(t *testing.T) {
	gen := New()

	fields := []Field{
		{Name: "id", Generator: "uuid"},
		{Name: "name", Generator: "name"},
		{Name: "email", Generator: "email"},
		{Name: "age", Generator: "random_int", Args: map[string]interface{}{"min": 18.0, "max": 65.0}},
	}

	records, err := gen.GenerateRecords(fields, 10)
	if err != nil {
		t.Fatalf("Failed to generate records: %v", err)
	}

	if len(records) != 10 {
		t.Errorf("Expected 10 records, got %d", len(records))
	}

	for i, record := range records {
		if record["id"] == nil || record["name"] == nil || record["email"] == nil || record["age"] == nil {
			t.Errorf("Record %d missing required fields", i)
		}
		fmt.Printf("Record %d: %+v\n", i, record)
	}
}

func TestGeneratePerson(t *testing.T) {
	gen := New()

	fields := []Field{
		{Name: "first_name", Generator: "person.first_name"},
		{Name: "last_name", Generator: "person.last_name"},
		{Name: "email", Generator: "person.email"},
	}

	records, err := gen.GenerateRecords(fields, 5)
	if err != nil {
		t.Fatalf("Failed to generate records: %v", err)
	}

	for i, record := range records {
		fmt.Printf("Person %d: %+v\n", i, record)
	}
}
