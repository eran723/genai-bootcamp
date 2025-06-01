import os
import sqlite3
from contextlib import closing

# Paths
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
TEST_DB_PATH = os.path.join(BASE_DIR, 'lang-portal', 'backend_go', 'test_database.db')
MIGRATIONS_PATH = os.path.join(BASE_DIR, 'lang-portal', 'backend_go', 'db', 'migrations', '0001_init.sql')
TEST_DATA_PATH = os.path.join(BASE_DIR, 'lang-portal', 'backend_go', 'db', 'seeds', 'test_data.sql')

def init_test_db():
    """Initialize the test database with schema and test data"""
    # Ensure the database directory exists
    os.makedirs(os.path.dirname(TEST_DB_PATH), exist_ok=True)
    
    # Remove existing database if it exists
    if os.path.exists(TEST_DB_PATH):
        os.remove(TEST_DB_PATH)
    
    print(f"Creating test database at: {TEST_DB_PATH}")
    
    # Create new database and apply schema
    with closing(sqlite3.connect(TEST_DB_PATH)) as db:
        print("Applying database schema...")
        with open(MIGRATIONS_PATH, 'r', encoding='utf-8') as f:
            db.executescript(f.read())
        
        print("Inserting test data...")
        with open(TEST_DATA_PATH, 'r', encoding='utf-8') as f:
            db.executescript(f.read())
        
        db.commit()
        print("Database initialization complete!")

if __name__ == '__main__':
    init_test_db() 