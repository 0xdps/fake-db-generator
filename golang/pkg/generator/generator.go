package generator

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

// Generator handles fake data generation
type Generator struct {
	fake    *gofakeit.Faker
	uniques map[string]map[interface{}]bool
}

// New creates a new Generator
func New() *Generator {
	return &Generator{
		fake:    gofakeit.New(0),
		uniques: make(map[string]map[interface{}]bool),
	}
}

// NewWithSeed creates a new Generator with a specific seed
func NewWithSeed(seed uint64) *Generator {
	return &Generator{
		fake:    gofakeit.New(seed),
		uniques: make(map[string]map[interface{}]bool),
	}
}

// Field represents a field configuration for generation
type Field struct {
	Name      string
	Generator string
	Args      interface{}
}

// Generate generates a value based on the generator type
func (g *Generator) Generate(field Field, commons map[string]interface{}) (interface{}, error) {
	generator := field.Generator
	args := field.Args

	// Handle special generators with property access
	if strings.Contains(generator, ".") {
		return g.generateWithProperty(generator, args, commons)
	}

	switch generator {
	// Personal
	case "name":
		return g.fake.Name(), nil
	case "first_name":
		return g.fake.FirstName(), nil
	case "last_name":
		return g.fake.LastName(), nil
	case "user_name", "username":
		return g.fake.Username(), nil
	case "email":
		return g.fake.Email(), nil
	case "password":
		return g.fake.Password(true, true, true, true, false, 10), nil
	case "gender":
		return g.fake.Gender(), nil

	// Address
	case "address":
		return g.fake.Address().Address, nil
	case "street_address":
		return g.fake.Address().Street, nil
	case "city":
		return g.fake.Address().City, nil
	case "state":
		return g.fake.Address().State, nil
	case "country":
		return g.fake.Address().Country, nil
	case "postcode", "zip_code":
		return g.fake.Address().Zip, nil
	case "latitude":
		return g.fake.Latitude(), nil
	case "longitude":
		return g.fake.Longitude(), nil

	// Company
	case "company":
		return g.fake.Company(), nil
	case "job":
		return g.fake.JobTitle(), nil
	case "catch_phrase":
		return g.fake.BuzzWord(), nil

	// Internet
	case "url":
		return g.fake.URL(), nil
	case "domain_name":
		return g.fake.DomainName(), nil
	case "ipv4":
		return g.fake.IPv4Address(), nil
	case "ipv6":
		return g.fake.IPv6Address(), nil
	case "mac_address":
		return g.fake.MacAddress(), nil

	// Dates
	case "date":
		return g.fake.Date(), nil
	case "date_time":
		return g.fake.Date(), nil
	case "past_date":
		return g.fake.DateRange(g.fake.PastDate(), g.fake.Date()), nil
	case "future_date":
		return g.fake.FutureDate(), nil

	// Text
	case "text":
		return g.fake.Paragraph(3, 5, 10, " "), nil
	case "sentence":
		return g.fake.Sentence(10), nil
	case "paragraph":
		return g.fake.Paragraph(3, 5, 10, " "), nil
	case "word":
		return g.fake.Word(), nil

	// Numbers
	case "random_int":
		min, max := 0, 100
		if argsMap, ok := args.(map[string]interface{}); ok {
			if minVal, ok := argsMap["min"].(float64); ok {
				min = int(minVal)
			}
			if maxVal, ok := argsMap["max"].(float64); ok {
				max = int(maxVal)
			}
		}
		return g.fake.IntRange(min, max), nil
	case "random_digit":
		return g.fake.Digit(), nil
	case "random_number":
		return g.fake.Number(0, 999999), nil

	// Phone
	case "phone_number", "phone":
		return g.fake.Phone(), nil

	// Boolean
	case "boolean":
		return g.fake.Bool(), nil

	// UUID
	case "uuid":
		return g.fake.UUID(), nil

	// Custom generators
	case "person":
		return g.generatePerson()
	case "user":
		return g.generateUser()
	case "random_from":
		return g.randomFrom(args)
	case "unique_item":
		return g.uniqueItem(field.Name, args)

	default:
		return nil, fmt.Errorf("unknown generator: %s", generator)
	}
}

// GenerateRecords generates multiple records based on fields configuration
func (g *Generator) GenerateRecords(fields []Field, count int) ([]map[string]interface{}, error) {
	records := make([]map[string]interface{}, 0, count)

	for i := 0; i < count; i++ {
		record := make(map[string]interface{})
		commons := make(map[string]interface{})

		for _, field := range fields {
			value, err := g.Generate(field, commons)
			if err != nil {
				return nil, fmt.Errorf("failed to generate field %s: %w", field.Name, err)
			}
			record[field.Name] = value
		}

		records = append(records, record)
	}

	return records, nil
}

// generateWithProperty handles generators with property access like "person.first_name"
func (g *Generator) generateWithProperty(generator string, args interface{}, commons map[string]interface{}) (interface{}, error) {
	parts := strings.Split(generator, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid generator format: %s", generator)
	}

	objType := parts[0]
	property := parts[1]

	// Check if we already generated this object
	if obj, exists := commons[objType]; exists {
		if objMap, ok := obj.(map[string]interface{}); ok {
			if val, ok := objMap[property]; ok {
				return val, nil
			}
		}
	}

	// Generate the object
	var obj map[string]interface{}
	switch objType {
	case "person":
		person, _ := g.generatePerson()
		obj = person.(map[string]interface{})
	case "user":
		user, _ := g.generateUser()
		obj = user.(map[string]interface{})
	default:
		return nil, fmt.Errorf("unknown object type: %s", objType)
	}

	commons[objType] = obj

	if val, ok := obj[property]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("property %s not found in %s", property, objType)
}

// generatePerson generates a Person object
func (g *Generator) generatePerson() (interface{}, error) {
	gender := g.fake.Gender()
	firstName := g.fake.FirstName()
	lastName := g.fake.LastName()
	email := strings.ToLower(fmt.Sprintf("%s.%s@%s", firstName, lastName, g.fake.DomainName()))

	return map[string]interface{}{
		"first_name": firstName,
		"last_name":  lastName,
		"email":      email,
		"gender":     gender,
		"address":    g.fake.Address().Address,
		"full_name":  fmt.Sprintf("%s %s", firstName, lastName),
	}, nil
}

// generateUser generates a User object
func (g *Generator) generateUser() (interface{}, error) {
	person, _ := g.generatePerson()
	personMap := person.(map[string]interface{})

	firstName := personMap["first_name"].(string)
	lastName := personMap["last_name"].(string)

	user := make(map[string]interface{})
	for k, v := range personMap {
		user[k] = v
	}
	user["username"] = strings.ToLower(fmt.Sprintf("%s_%s", firstName, lastName))
	user["password"] = g.fake.Password(true, true, true, true, false, 10)

	return user, nil
}

// randomFrom picks a random item from args
func (g *Generator) randomFrom(args interface{}) (interface{}, error) {
	if argsList, ok := args.([]interface{}); ok {
		if len(argsList) == 0 {
			return nil, fmt.Errorf("empty list for random_from")
		}
		idx := g.fake.Number(0, len(argsList)-1)
		return argsList[idx], nil
	}
	return nil, fmt.Errorf("random_from requires an array argument")
}

// uniqueItem generates a unique item
func (g *Generator) uniqueItem(fieldName string, args interface{}) (interface{}, error) {
	if argsList, ok := args.([]interface{}); ok {
		if g.uniques[fieldName] == nil {
			g.uniques[fieldName] = make(map[interface{}]bool)
		}

		// Try to find an unused value
		for i := 0; i < len(argsList)*10; i++ {
			val, _ := g.randomFrom(args)
			if !g.uniques[fieldName][val] {
				g.uniques[fieldName][val] = true
				return val, nil
			}
		}
		return nil, fmt.Errorf("could not generate unique value for %s", fieldName)
	}
	return nil, fmt.Errorf("unique_item requires an array argument")
}

// ClearUniques clears the unique value cache
func (g *Generator) ClearUniques() {
	g.uniques = make(map[string]map[interface{}]bool)
}
