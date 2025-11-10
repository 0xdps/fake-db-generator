# Data Generators Reference

Complete list of all data generators available in Fakestack.

## Personal Information

### `first_name`
Generate a random first name.
```json
{ "name": "first_name", "generator": "first_name" }
```
Example: "John", "Mary", "Ahmed"

### `last_name`
Generate a random last name.
```json
{ "name": "last_name", "generator": "last_name" }
```
Example: "Smith", "Johnson", "Lee"

### `name`
Generate a full name (first + last).
```json
{ "name": "full_name", "generator": "name" }
```
Example: "John Smith", "Mary Johnson"

### `email`
Generate a random email address.
```json
{ "name": "email", "generator": "email" }
```
Example: "john.smith@example.com"

### `user_name`
Generate a random username.
```json
{ "name": "username", "generator": "user_name" }
```
Example: "john_smith_42", "cool_user_99"

### `password`
Generate a random password.
```json
{ "name": "password", "generator": "password" }
```
Example: "aB3$xY9&mK2!"

### `phone_number`
Generate a phone number.
```json
{ "name": "phone", "generator": "phone_number" }
```
Example: "+1-555-123-4567"

### `ssn`
Generate a Social Security Number (US format).
```json
{ "name": "ssn", "generator": "ssn" }
```
Example: "123-45-6789"

## Address

### `address`
Generate a complete address.
```json
{ "name": "address", "generator": "address" }
```
Example: "123 Main St, New York, NY 10001"

### `street_address`
Generate a street address.
```json
{ "name": "street", "generator": "street_address" }
```
Example: "123 Main Street"

### `city`
Generate a city name.
```json
{ "name": "city", "generator": "city" }
```
Example: "New York", "Los Angeles", "Chicago"

### `state`
Generate a state name.
```json
{ "name": "state", "generator": "state" }
```
Example: "California", "Texas", "New York"

### `country`
Generate a country name.
```json
{ "name": "country", "generator": "country" }
```
Example: "United States", "Canada", "Mexico"

### `postcode`
Generate a postal/zip code.
```json
{ "name": "zip", "generator": "postcode" }
```
Example: "90210", "10001"

### `latitude`
Generate a latitude coordinate.
```json
{ "name": "lat", "generator": "latitude" }
```
Example: "40.7128"

### `longitude`
Generate a longitude coordinate.
```json
{ "name": "lng", "generator": "longitude" }
```
Example: "-74.0060"

## Company

### `company`
Generate a company name.
```json
{ "name": "company_name", "generator": "company" }
```
Example: "Acme Corporation", "TechStart Inc."

### `company_suffix`
Generate a company suffix.
```json
{ "name": "suffix", "generator": "company_suffix" }
```
Example: "Inc.", "LLC", "Corp."

### `job`
Generate a job title.
```json
{ "name": "job_title", "generator": "job" }
```
Example: "Software Engineer", "Product Manager"

### `catch_phrase`
Generate a company catch phrase.
```json
{ "name": "slogan", "generator": "catch_phrase" }
```
Example: "Innovative solutions for tomorrow"

## Internet

### `url`
Generate a URL.
```json
{ "name": "website", "generator": "url" }
```
Example: "https://example.com"

### `domain_name`
Generate a domain name.
```json
{ "name": "domain", "generator": "domain_name" }
```
Example: "example.com"

### `ipv4`
Generate an IPv4 address.
```json
{ "name": "ip", "generator": "ipv4" }
```
Example: "192.168.1.1"

### `ipv6`
Generate an IPv6 address.
```json
{ "name": "ipv6", "generator": "ipv6" }
```
Example: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"

### `mac_address`
Generate a MAC address.
```json
{ "name": "mac", "generator": "mac_address" }
```
Example: "00:1B:44:11:3A:B7"

### `user_agent`
Generate a browser user agent string.
```json
{ "name": "user_agent", "generator": "user_agent" }
```
Example: "Mozilla/5.0 (Windows NT 10.0; Win64; x64)..."

### `slug`
Generate a URL slug.
```json
{ "name": "slug", "generator": "slug" }
```
Example: "my-awesome-post"

## Dates & Times

### `date`
Generate a random date.
```json
{ "name": "birth_date", "generator": "date" }
```
Example: "2023-05-15"

### `date_time`
Generate a random datetime.
```json
{ "name": "created_at", "generator": "date_time" }
```
Example: "2023-05-15 14:30:00"

### `past_date`
Generate a date in the past.
```json
{ "name": "registered_at", "generator": "past_date" }
```
Example: "2022-03-20"

### `future_date`
Generate a date in the future.
```json
{ "name": "expires_at", "generator": "future_date" }
```
Example: "2024-12-31"

### `time`
Generate a time.
```json
{ "name": "login_time", "generator": "time" }
```
Example: "14:30:00"

### `unix_time`
Generate a Unix timestamp.
```json
{ "name": "timestamp", "generator": "unix_time" }
```
Example: "1684159200"

## Text

### `text`
Generate random text (multiple paragraphs).
```json
{
  "name": "description",
  "generator": "text",
  "args": { "max_chars": 500 }
}
```
Example: "Lorem ipsum dolor sit amet..."

### `sentence`
Generate a random sentence.
```json
{ "name": "title", "generator": "sentence" }
```
Example: "The quick brown fox jumps over the lazy dog."

### `paragraph`
Generate a random paragraph.
```json
{ "name": "bio", "generator": "paragraph" }
```
Example: "Lorem ipsum dolor sit amet, consectetur adipiscing elit..."

### `word`
Generate a single random word.
```json
{ "name": "tag", "generator": "word" }
```
Example: "technology"

### `words`
Generate multiple random words.
```json
{
  "name": "keywords",
  "generator": "words",
  "args": { "count": 5 }
}
```
Example: "technology innovation digital future trends"

## Numbers

### `random_int`
Generate a random integer within a range.
```json
{
  "name": "age",
  "generator": "random_int",
  "args": { "min": 18, "max": 65 }
}
```
Example: 42

### `random_digit`
Generate a single random digit (0-9).
```json
{ "name": "rating", "generator": "random_digit" }
```
Example: 7

### `random_number`
Generate a random number with specified digits.
```json
{
  "name": "code",
  "generator": "random_number",
  "args": { "digits": 6 }
}
```
Example: 123456

### `random_float`
Generate a random float.
```json
{
  "name": "price",
  "generator": "random_float",
  "args": { "min": 0.0, "max": 100.0, "decimals": 2 }
}
```
Example: 42.99

## Special Generators

### `uuid`
Generate a UUID (v4).
```json
{ "name": "id", "generator": "uuid" }
```
Example: "550e8400-e29b-41d4-a716-446655440000"

### `boolean`
Generate a random boolean.
```json
{ "name": "is_active", "generator": "boolean" }
```
Example: true or false

### `random_from`
Pick a random value from a list.
```json
{
  "name": "status",
  "generator": "random_from",
  "args": {
    "choices": ["pending", "active", "inactive", "suspended"]
  }
}
```
Example: "active"

### `foreign_key`
Reference a foreign key from another table.
```json
{
  "name": "user_id",
  "generator": "foreign_key",
  "args": { "table": "users" }
}
```
This will randomly select an existing ID from the specified table.

### `person`
Generate a complete person object (returns JSON).
```json
{ "name": "person_data", "generator": "person" }
```
Returns: `{"first_name": "John", "last_name": "Doe", "email": "john.doe@example.com", ...}`

### `user`
Generate a complete user object (returns JSON).
```json
{ "name": "user_data", "generator": "user" }
```
Returns: `{"username": "john_doe", "email": "john@example.com", "password": "secret123"}`

## Generator Arguments

Most generators support optional arguments:

```json
{
  "name": "field_name",
  "generator": "generator_name",
  "args": {
    "argument1": "value1",
    "argument2": "value2"
  }
}
```

Common arguments:
- `min`, `max` - For numeric ranges
- `length` - For string lengths
- `choices` - For random selection
- `count` - For multiple items
- `max_chars` - For text length limits
- `decimals` - For float precision

## Tips

1. **Consistency**: Use related generators for consistent data (e.g., use `user_name` and `email` together)
2. **Realistic Data**: Choose generators that match your domain
3. **Performance**: Simple generators are faster than complex ones
4. **Unique Values**: Some generators may produce duplicates; use `unique: true` in column options if needed
5. **Foreign Keys**: Always use `foreign_key` generator when referencing other tables

## See Also

- [Schema Reference](schema-reference.md) - How to use generators in schemas
- [Examples](examples.md) - Real-world generator usage
