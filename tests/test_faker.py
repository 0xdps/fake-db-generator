"""Tests for Faker providers."""

import pytest

from schemer.models.fake import (
    Person,
    PersonProvider,
    RandomProvider,
    User,
    UserProvider,
    faker,
)


class TestPersonProvider:
    """Test person data generation."""

    def test_person_generation(self):
        """Test generating person data."""
        person = faker.person()

        assert isinstance(person, Person)
        assert person.first_name
        assert person.last_name
        assert person.email
        assert person.gender in ["M", "F"]
        assert person.address
        assert "@" in person.email
        assert "." in person.email

    def test_person_full_name(self):
        """Test person full name property."""
        person = faker.person()
        expected_full_name = f"{person.first_name} {person.last_name}"
        assert person.full_name == expected_full_name


class TestUserProvider:
    """Test user data generation."""

    def test_user_generation(self):
        """Test generating user data."""
        user = faker.user()

        assert isinstance(user, User)
        assert user.username
        assert user.password
        assert user.first_name
        assert user.last_name
        assert user.email
        assert "_" in user.username

    def test_user_inherits_person(self):
        """Test that User inherits from Person."""
        user = faker.user()
        assert hasattr(user, "full_name")
        expected_full_name = f"{user.first_name} {user.last_name}"
        assert user.full_name == expected_full_name


class TestRandomProvider:
    """Test random selection provider."""

    def test_random_from_options(self):
        """Test random selection from options."""
        options = ["apple", "banana", "cherry"]
        result = faker.random_from(*options)
        assert result in options

    def test_random_from_single_option(self):
        """Test random selection with single option."""
        result = faker.random_from("only_option")
        assert result == "only_option"

    @pytest.mark.parametrize("count", [10, 50, 100])
    def test_random_distribution(self, count):
        """Test that random selection uses all options."""
        options = ["a", "b", "c", "d", "e"]
        results = [faker.random_from(*options) for _ in range(count)]

        # Should have at least some variety
        unique_results = set(results)
        assert len(unique_results) > 1


class TestUniqueProvider:
    """Test unique value generation."""

    def test_unique_integers(self):
        """Test unique integer generation."""
        faker.unique.clear()
        values = [faker.unique.random_int() for _ in range(10)]

        # All values should be unique
        assert len(values) == len(set(values))

    def test_unique_random_from(self):
        """Test unique selection from list."""
        faker.unique.clear()
        options = ["a", "b", "c"]

        # Should get all 3 unique values
        values = [faker.unique.random_from(*options) for _ in range(3)]
        assert len(set(values)) == 3
        assert set(values) == set(options)


class TestCustomProvider:
    """Test custom provider get method."""

    def test_simple_method_call(self):
        """Test calling a simple Faker method."""
        result, commons = faker.get("name", None, {})
        assert isinstance(result, str)
        assert len(result) > 0

    def test_method_with_args(self):
        """Test calling method with arguments."""
        result, commons = faker.get("random_int", {"min": 1, "max": 10}, {})
        assert isinstance(result, int)
        assert 1 <= result <= 10

    def test_nested_attribute_access(self):
        """Test accessing nested attributes."""
        commons = {}
        first_name, commons = faker.get("person.first_name", None, commons)

        assert isinstance(first_name, str)
        assert "person" in commons

        # Second call should reuse cached person
        last_name, commons2 = faker.get("person.last_name", None, commons)
        assert commons["person"] is commons2["person"]


class TestFakerIntegration:
    """Test Faker integration."""

    def test_standard_faker_methods(self):
        """Test that standard Faker methods still work."""
        # Test various standard Faker providers
        assert faker.name()
        assert faker.email()
        assert faker.address()
        assert isinstance(faker.random_int(), int)
        assert faker.date()

    def test_locale_support(self):
        """Test that Faker locale works."""
        # Default locale should work
        name = faker.name()
        assert isinstance(name, str)
