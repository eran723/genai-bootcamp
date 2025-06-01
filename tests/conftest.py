import os
import pytest
import sqlite3
from flask import Flask
from flask.testing import FlaskClient
from typing import Generator, Any
from contextlib import closing

# Import our actual application
from app import create_app

# Get the root directory of the project
ROOT_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# Define paths relative to the root directory
TEST_DB_PATH = os.path.join(ROOT_DIR, 'lang-portal', 'backend_go', 'test_database.db')
SCHEMA_PATH = os.path.join(ROOT_DIR, 'lang-portal', 'backend_go', 'db', 'migrations', '0001_init.sql')
TEST_DATA_PATH = os.path.join(ROOT_DIR, 'lang-portal', 'backend_go', 'db', 'seeds', 'test_data.sql')

@pytest.fixture(scope='session')
def app():
    """Create and configure a test Flask application"""
    app = Flask(__name__)
    app.config['TESTING'] = True
    app.config['DATABASE'] = TEST_DB_PATH
    app.config['ROOT_PATH'] = ROOT_DIR  # Add root path to config
    
    return app

@pytest.fixture
def client(app):
    """Create a test client"""
    return app.test_client()

@pytest.fixture(scope='session')
def init_database(app):
    """Initialize the test database with schema and test data"""
    # Ensure the database directory exists
    os.makedirs(os.path.dirname(TEST_DB_PATH), exist_ok=True)
    
    # Remove existing database if it exists
    if os.path.exists(TEST_DB_PATH):
        os.remove(TEST_DB_PATH)
    
    # Create new database and apply schema
    with closing(sqlite3.connect(TEST_DB_PATH)) as db:
        print(f"Reading schema from: {SCHEMA_PATH}")
        with open(SCHEMA_PATH, 'r', encoding='utf-8') as f:
            db.executescript(f.read())
        
        print(f"Reading test data from: {TEST_DATA_PATH}")
        with open(TEST_DATA_PATH, 'r', encoding='utf-8') as f:
            db.executescript(f.read())
        
        db.commit()

@pytest.fixture(autouse=True)
def setup_test_database(app):
    """Setup a fresh test database for each test"""
    # Connect to the database
    with closing(sqlite3.connect(TEST_DB_PATH)) as db:
        # Begin a transaction
        db.execute('BEGIN')
        
        yield db
        
        # Rollback the transaction after the test
        db.rollback()

@pytest.fixture
def runner(app: Flask) -> Any:
    """Create a test CLI runner for the Flask application."""
    return app.test_cli_runner()

@pytest.fixture
def db_connection():
    """Create a new database connection for tests."""
    conn = sqlite3.connect(TEST_DB_PATH)
    conn.row_factory = sqlite3.Row
    yield conn
    conn.close() 