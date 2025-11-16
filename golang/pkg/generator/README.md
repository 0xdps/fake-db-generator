# Fakestack Generator Library

A standalone Go library for generating realistic fake data. Can be used independently without database connections.

## Installation

```bash
go get github.com/0xdps/fake-stack/core/pkg/generator
```

## Usage

```go
import "github.com/0xdps/fake-stack/core/pkg/generator"

// Create a generator
gen := generator.New()

// Define fields
fields := []generator.Field{
    {Name: "id", Generator: "uuid"},
    {Name: "name", Generator: "name"},
    {Name: "email", Generator: "email"},
    {Name: "age", Generator: "random_int", Args: map[string]interface{}{"min": 18.0, "max": 65.0}},
}

// Generate records
records, err := gen.GenerateRecords(fields, 100)
```

## Available Generators

### Personal
- `name`, `first_name`, `last_name`, `username`, `email`, `password`, `gender`

### Address
- `address`, `street_address`, `city`, `state`, `country`, `postcode`, `zip_code`
- `latitude`, `longitude`

### Company
- `company`, `job`, `catch_phrase`

### Internet
- `url`, `domain_name`, `ipv4`, `ipv6`, `mac_address`

### Dates
- `date`, `date_time`, `past_date`, `future_date`

### Text
- `text`, `sentence`, `paragraph`, `word`

### Numbers
- `random_int` (with min/max args), `random_digit`, `random_number`

### Other
- `phone_number`, `phone`, `boolean`, `uuid`

### Complex Objects
- `person` - Returns object with first_name, last_name, email, gender, address, full_name
- `user` - Returns person object + username, password
- `person.{property}` - Access specific property (e.g., "person.first_name")
- `user.{property}` - Access specific property (e.g., "user.username")

### Custom
- `random_from` - Pick random item from array: `Args: []interface{}{"red", "blue", "green"}`
- `unique_item` - Unique item from array

## Example: API Usage

Perfect for building mock APIs like apimock.codes:

```go
func GetUsers(c *gin.Context) {
    gen := generator.New()
    
    fields := []generator.Field{
        {Name: "id", Generator: "uuid"},
        {Name: "username", Generator: "username"},
        {Name: "email", Generator: "email"},
        {Name: "first_name", Generator: "first_name"},
        {Name: "last_name", Generator: "last_name"},
    }
    
    users, _ := gen.GenerateRecords(fields, 100)
    c.JSON(200, users)
}
```

## Testing

```bash
go test ./pkg/generator -v
```
