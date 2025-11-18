package main

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Generator handles fake data generation
type Generator struct {
	fake   *gofakeit.Faker
	db     *Database
	uniques map[string]map[interface{}]bool
}

// NewGenerator creates a new Generator
func NewGenerator(db *Database) *Generator {
	return &Generator{
		fake:   gofakeit.New(0),
		db:     db,
		uniques: make(map[string]map[interface{}]bool),
	}
}

// Generate generates a value based on the generator type
func (g *Generator) Generate(field PopulateField, commons map[string]interface{}) (interface{}, error) {
	generator := field.Generator
	args := field.Args

	// Handle special generators
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
	case "integer", "random_int":
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
	case "float", "decimal":
		min, max := 0.0, 100.0
		if argsMap, ok := args.(map[string]interface{}); ok {
			if minVal, ok := argsMap["min"].(float64); ok {
				min = minVal
			}
			if maxVal, ok := argsMap["max"].(float64); ok {
				max = maxVal
			}
		}
		return g.fake.Float64Range(min, max), nil
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
	case "db_random_item":
		return g.dbRandomItem(args)
	case "template":
		return g.generateFromTemplate(args, commons)

	default:
		return nil, fmt.Errorf("unknown generator: %s", generator)
	}
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

// dbRandomItem gets a random value from the database
func (g *Generator) dbRandomItem(args interface{}) (interface{}, error) {
	if attribute, ok := args.(string); ok {
		parts := strings.Split(attribute, ".")
		if len(parts) != 2 {
			return nil, fmt.Errorf("db_random_item requires format 'table.column'")
		}
		return g.db.GetRandomValue(parts[0], parts[1])
	}
	return nil, fmt.Errorf("db_random_item requires a string argument")
}

// ClearUniques clears the unique value cache
func (g *Generator) ClearUniques() {
	g.uniques = make(map[string]map[interface{}]bool)
}

// generateFromTemplate generates value from a template pattern
func (g *Generator) generateFromTemplate(args interface{}, commons map[string]interface{}) (interface{}, error) {
	argsMap, ok := args.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("template generator requires args object with 'pattern' field")
	}

	pattern, ok := argsMap["pattern"].(string)
	if !ok {
		return nil, fmt.Errorf("template generator requires 'pattern' string")
	}

	return g.parseTemplate(pattern, commons)
}

// parseTemplate parses and executes template pattern
func (g *Generator) parseTemplate(pattern string, commons map[string]interface{}) (string, error) {
	result := pattern
	
	// Find all {{...}} placeholders
	for {
		start := strings.Index(result, "{{")
		if start == -1 {
			break
		}
		
		end := strings.Index(result[start:], "}}")
		if end == -1{
			return "", fmt.Errorf("unclosed template placeholder in: %s", pattern)
		}
		end += start
		
		// Extract placeholder content
		placeholder := result[start+2 : end]
		
		// Parse placeholder: generator|modifier1|modifier2
		parts := strings.Split(placeholder, "|")
		generatorExpr := strings.TrimSpace(parts[0])
		modifiers := parts[1:]
		
		// Generate value
		value, err := g.executeTemplateGenerator(generatorExpr, commons)
		if err != nil {
			return "", fmt.Errorf("failed to execute generator '%s': %w", generatorExpr, err)
		}
		
		// Apply modifiers
		valueStr := fmt.Sprintf("%v", value)
		for _, modifier := range modifiers {
			valueStr = g.applyModifier(valueStr, strings.TrimSpace(modifier))
		}
		
		// Replace placeholder with generated value
		result = result[:start] + valueStr + result[end+2:]
	}
	
	return result, nil
}

// executeTemplateGenerator executes a generator expression
func (g *Generator) executeTemplateGenerator(expr string, commons map[string]interface{}) (interface{}, error) {
	// Parse generator name and arguments
	// Format: generator_name or generator_name(arg1,arg2)
	
	parenIndex := strings.Index(expr, "(")
	if parenIndex == -1 {
		// No arguments, simple generator
		return g.executeSimpleGenerator(expr)
	}
	
	// Has arguments
	generatorName := strings.TrimSpace(expr[:parenIndex])
	argsStr := expr[parenIndex+1:]
	
	// Remove closing parenthesis
	if !strings.HasSuffix(argsStr, ")") {
		return nil, fmt.Errorf("missing closing parenthesis in: %s", expr)
	}
	argsStr = argsStr[:len(argsStr)-1]
	
	// Parse arguments
	args := g.parseTemplateArgs(argsStr)
	
	return g.executeGeneratorWithArgs(generatorName, args)
}

// executeSimpleGenerator executes a generator without arguments
func (g *Generator) executeSimpleGenerator(name string) (interface{}, error) {
	switch name {
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
	case "word":
		return g.fake.Word(), nil
	case "sentence":
		return g.fake.Sentence(10), nil
	case "city":
		return g.fake.Address().City, nil
	case "state":
		return g.fake.Address().State, nil
	case "country":
		return g.fake.Address().Country, nil
	case "postcode", "zip_code":
		return g.fake.Address().Zip, nil
	case "company":
		return g.fake.Company(), nil
	case "uuid":
		return g.fake.UUID(), nil
	case "phone":
		return g.fake.Phone(), nil
	case "date":
		return g.fake.Date(), nil
	default:
		return nil, fmt.Errorf("unknown generator: %s", name)
	}
}

// executeGeneratorWithArgs executes a generator with arguments
func (g *Generator) executeGeneratorWithArgs(name string, args []string) (interface{}, error) {
	switch name {
	case "random_int":
		if len(args) != 2 {
			return nil, fmt.Errorf("random_int requires 2 arguments: min, max")
		}
		min, max := 0, 100
		fmt.Sscanf(args[0], "%d", &min)
		fmt.Sscanf(args[1], "%d", &max)
		return g.fake.IntRange(min, max), nil
	case "truncate":
		// Special case: handled as modifier
		return nil, fmt.Errorf("truncate should be used as modifier, not generator")
	default:
		return g.executeSimpleGenerator(name)
	}
}

// parseTemplateArgs parses comma-separated arguments
func (g *Generator) parseTemplateArgs(argsStr string) []string {
	if argsStr == "" {
		return []string{}
	}
	
	parts := strings.Split(argsStr, ",")
	args := make([]string, len(parts))
	for i, part := range parts {
		args[i] = strings.TrimSpace(part)
	}
	return args
}

// applyModifier applies a modifier to a string value
func (g *Generator) applyModifier(value, modifier string) string {
	// Check if modifier has arguments: modifier(arg)
	parenIndex := strings.Index(modifier, "(")
	if parenIndex != -1 {
		modName := modifier[:parenIndex]
		argStr := modifier[parenIndex+1 : len(modifier)-1]
		
		switch modName {
		case "truncate":
			var length int
			fmt.Sscanf(argStr, "%d", &length)
			if length > 0 && len(value) > length {
				return value[:length]
			}
			return value
		case "format":
			// Date formatting - simplified
			return value
		}
		return value
	}
	
	// Simple modifiers
	switch modifier {
	case "upper":
		return strings.ToUpper(value)
	case "lower":
		return strings.ToLower(value)
	case "title":
		caser := cases.Title(language.English)
		return caser.String(strings.ToLower(value))
	case "trim":
		return strings.TrimSpace(value)
	default:
		return value
	}
}
