import pytest
from flask.testing import FlaskClient
from flask import Flask
import json

def test_health_check(client: FlaskClient):
    """Test the health check endpoint."""
    response = client.get('/health')
    assert response.status_code == 200
    assert response.json == {"status": "healthy"}

class TestTodoEndpoints:
    """Test suite for Todo API endpoints."""
    
    def test_get_todos_empty(self, client: FlaskClient):
        """Test GET /api/todos returns empty list initially."""
        response = client.get('/api/todos')
        assert response.status_code == 200
        assert response.content_type == 'application/json'
        assert response.json == []
        
    def test_create_todo(self, client: FlaskClient):
        """Test POST /api/todos creates a new todo."""
        payload = {
            "title": "Learn Flask Testing"
        }
        response = client.post('/api/todos', json=payload)
        assert response.status_code == 201
        assert response.content_type == 'application/json'
        assert response.json['id'] == 1
        assert response.json['title'] == payload['title']
        assert response.json['completed'] == False
        
    def test_create_todo_invalid(self, client: FlaskClient):
        """Test POST /api/todos with invalid data."""
        response = client.post('/api/todos', json={})
        assert response.status_code == 400
        assert 'error' in response.json
        
    def test_get_todo_not_found(self, client: FlaskClient):
        """Test GET /api/todos/<id> with non-existent ID."""
        response = client.get('/api/todos/999')
        assert response.status_code == 404
        assert 'error' in response.json
        
    def test_create_and_get_todo(self, client: FlaskClient):
        """Test creating a todo and then retrieving it."""
        # Create todo
        payload = {"title": "Test Todo"}
        create_response = client.post('/api/todos', json=payload)
        assert create_response.status_code == 201
        todo_id = create_response.json['id']
        
        # Get the created todo
        get_response = client.get(f'/api/todos/{todo_id}')
        assert get_response.status_code == 200
        assert get_response.json['title'] == payload['title']

# Example of how to test with authentication (if needed)
class TestAuthenticatedEndpoints:
    """Test suite for authenticated endpoints."""
    
    @pytest.fixture
    def auth_headers(self) -> dict:
        """Fixture for authentication headers.
        
        Returns:
            dict: Headers with authentication token
        """
        # Replace this with your actual authentication logic
        return {"Authorization": "Bearer test-token"}
    
    def test_protected_endpoint(self, client: FlaskClient, auth_headers: dict):
        """Test authenticated endpoint."""
        response = client.get('/api/protected', headers=auth_headers)
        assert response.status_code == 200
        assert response.content_type == 'application/json'

class TestWordsEndpoints:
    """Test suite for Words API endpoints."""
    
    def test_list_words(self, client: FlaskClient):
        """Test GET /api/words returns paginated list of words."""
        response = client.get('/api/words')
        assert response.status_code == 200
        assert response.content_type == 'application/json'
        
        # Check response structure
        data = response.json
        assert 'items' in data
        assert 'pagination' in data
        
        # Check pagination metadata
        pagination = data['pagination']
        assert 'page' in pagination
        assert 'per_page' in pagination
        assert 'total' in pagination
        assert pagination['page'] == 1
        assert pagination['per_page'] == 100  # As per our technical specs
        
        # Check items structure
        items = data['items']
        assert isinstance(items, list)
        assert len(items) > 0  # We should have our seed data
        
        # Check first word structure
        first_word = items[0]
        assert 'id' in first_word
        assert 'japanese' in first_word
        assert 'romanization' in first_word
        assert 'english' in first_word
        assert 'word_type' in first_word
        assert 'created_at' in first_word
        assert 'updated_at' in first_word
        
        # Verify some seed data
        assert any(word['japanese'] == 'こんにちは' for word in items)
        assert any(word['english'] == 'hello' for word in items)
    
    def test_get_word_by_id(self, client: FlaskClient):
        """Test GET /api/words/<id> returns specific word."""
        # First, get the list to find a valid ID
        list_response = client.get('/api/words')
        first_word_id = list_response.json['items'][0]['id']
        
        # Get specific word
        response = client.get(f'/api/words/{first_word_id}')
        assert response.status_code == 200
        assert response.content_type == 'application/json'
        
        word = response.json
        assert word['id'] == first_word_id
        assert all(key in word for key in ['japanese', 'romanization', 'english', 'word_type', 'created_at', 'updated_at'])
    
    def test_get_word_not_found(self, client: FlaskClient):
        """Test GET /api/words/<id> with non-existent ID."""
        response = client.get('/api/words/999999')
        assert response.status_code == 404
        assert 'error' in response.json
    
    def test_list_words_pagination(self, client: FlaskClient):
        """Test GET /api/words with pagination parameters."""
        response = client.get('/api/words?page=1&per_page=2')
        assert response.status_code == 200
        
        data = response.json
        assert len(data['items']) == 2  # Requested 2 items per page
        assert data['pagination']['page'] == 1
        assert data['pagination']['per_page'] == 2
        
        # Get second page
        response = client.get('/api/words?page=2&per_page=2')
        assert response.status_code == 200
        
        data = response.json
        assert data['pagination']['page'] == 2
        # Different items than first page
        assert data['items'][0]['id'] != response.json['items'][0]['id']

def test_get_words_list(client, init_database):
    """Test GET /api/words endpoint"""
    response = client.get('/api/words')
    assert response.status_code == 200
    
    data = json.loads(response.data)
    assert 'items' in data
    assert 'pagination' in data
    
    # Check pagination structure
    assert all(key in data['pagination'] for key in ['current_page', 'total_pages', 'total_items', 'items_per_page'])
    assert data['pagination']['items_per_page'] == 100
    
    # Check word structure in items
    for word in data['items']:
        assert all(key in word for key in ['id', 'japanese', 'romaji', 'english', 'correct_count', 'wrong_count'])

def test_get_word_detail(client, init_database):
    """Test GET /api/words/:id endpoint"""
    # Test with existing word (こんにちは)
    response = client.get('/api/words/1')
    assert response.status_code == 200
    
    data = json.loads(response.data)
    assert all(key in data for key in ['id', 'japanese', 'romaji', 'english', 'parts', 'correct_count', 'wrong_count', 'success_rate', 'groups'])
    assert data['japanese'] == 'こんにちは'
    assert data['romaji'] == 'konnichiwa'
    
    # Test with non-existent word
    response = client.get('/api/words/999')
    assert response.status_code == 404

def test_get_group_words(client, init_database):
    """Test GET /api/groups/:id/words endpoint"""
    # Test with existing group (Basic Greetings)
    response = client.get('/api/groups/1/words')
    assert response.status_code == 200
    
    data = json.loads(response.data)
    assert all(key in data for key in ['group_id', 'total_words', 'items'])
    assert data['group_id'] == 1
    
    # Verify words in group
    words = data['items']
    assert len(words) > 0
    for word in words:
        assert all(key in word for key in ['id', 'japanese', 'romaji', 'english', 'correct_count', 'wrong_count'])
    
    # Test with non-existent group
    response = client.get('/api/groups/999/words')
    assert response.status_code == 404

def test_get_session_words(client, init_database):
    """Test GET /api/study_sessions/:id/words endpoint"""
    # Test with existing session
    response = client.get('/api/study_sessions/1/words')
    assert response.status_code == 200
    
    data = json.loads(response.data)
    assert 'items' in data
    assert 'pagination' in data
    
    # Check word structure
    for word in data['items']:
        assert all(key in word for key in ['word_id', 'japanese', 'romaji', 'english', 'corrected_count', 'wrong_count'])
    
    # Test with non-existent session
    response = client.get('/api/study_sessions/999/words')
    assert response.status_code == 404

def test_post_word_review(client, init_database):
    """Test POST /api/study_sessions/:id/word_id/review endpoint"""
    # Test with valid review data
    review_data = {
        "study_session_id": 1,
        "word_id": 1,
        "correct": True,
        "response": "hello"
    }
    
    response = client.post('/api/study_sessions/1/word_id/review', 
                          data=json.dumps(review_data),
                          content_type='application/json')
    assert response.status_code == 200
    
    data = json.loads(response.data)
    assert all(key in data for key in ['session_id', 'word_id', 'is_correct', 'response', 'created_at'])
    assert data['session_id'] == 1
    assert data['word_id'] == 1
    assert data['is_correct'] == True
    
    # Test with invalid session
    review_data['study_session_id'] = 999
    response = client.post('/api/study_sessions/999/word_id/review',
                          data=json.dumps(review_data),
                          content_type='application/json')
    assert response.status_code == 404
    
    # Test with invalid word
    review_data['study_session_id'] = 1
    review_data['word_id'] = 999
    response = client.post('/api/study_sessions/1/word_id/review',
                          data=json.dumps(review_data),
                          content_type='application/json')
    assert response.status_code == 404 