// Package fakestack provides tools for generating realistic fake data and populating databases.
//
// Fakestack is a high-performance database generator with realistic fake data.
// It supports multiple databases (SQLite, MySQL, PostgreSQL, MariaDB, MSSQL, CockroachDB)
// and provides 116+ generators for creating test data.
//
// The main functionality is implemented in the golang subdirectory.
// For CLI usage, build and run the binary from golang/main.go
// For library usage, import github.com/0xdps/fake-stack/golang/pkg/generator
//
// Example:
//
//	import "github.com/0xdps/fake-stack/golang/pkg/generator"
//
//	gen := generator.New()
//	fields := []generator.Field{
//	    {Name: "username", Generator: "username"},
//	    {Name: "email", Generator: "email"},
//	}
//	records, err := gen.GenerateRecords(fields, 10)
//
// For more information, visit: https://github.com/0xdps/fake-stack
package fakestack
