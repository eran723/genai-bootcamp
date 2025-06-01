# Testing Process Documentation

## Overview
This document outlines the testing process for the language learning portal's backend services. Our testing approach uses a dedicated test database and predefined test data to ensure consistent and reliable test results.

## Test Database Setup
- Location: `lang-portral/backend_go/test_database.db`
- Purpose: Isolated SQLite database for testing purposes
- Reset Policy: Database is recreated for each test session

## Test Data Structure
Our test data includes instances for all core tables:

1. Words
   - Basic vocabulary items
   - Different word types and complexities
   - Various Japanese characters and romanizations

2. Groups
   - Thematic word groupings
   - Different group sizes
   - Various difficulty levels

3. Study Sessions
   - Different completion states
   - Various scores and durations
   - Multiple review items

4. Study Activities
   - Different activity types
   - Various group associations
   - Different completion states

5. Word Review Items
   - Correct and incorrect responses
   - Different review timestamps
   - Various word associations

## Test Data Management
1. Initial Setup
   - Test database creation
   - Schema migration
   - Seed data insertion

2. Test Isolation
   - Each test starts with a fresh database state
   - Transactions are rolled back after each test
   - No cross-test data contamination

## Testing Layers

### 1. Service Layer Tests
- Unit tests for business logic
- Database interaction verification
- Error handling scenarios

### 2. API Endpoint Tests
- Request/response format validation
- HTTP status code verification
- Pagination implementation
- Error responses

### 3. Integration Tests
- End-to-end workflow testing
- Cross-service interactions
- Data consistency checks

## Running Tests
```bash
# Run all tests
python -m pytest

# Run specific test file
python -m pytest tests/test_services.py

# Run with verbose output
python -m pytest -v

# Run with coverage report
python -m pytest --cov=app
```

## Test Data Reset
The test database can be reset using:
```bash
python -m pytest --reset-db
```

## Kill if already running
If the app is running on port 8080, or another local port, kill the process.

## Summary
This will:
1. Delete existing test database
2. Create new database
3. Run migrations
4. Insert seed data 
5. Kill previous running app if already running