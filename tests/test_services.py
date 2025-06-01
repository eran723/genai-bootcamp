import pytest
from flask.testing import FlaskClient
from datetime import datetime

class TestStudySessionService:
    """Test suite for Study Session Service endpoints."""
    
    def test_create_session(self, client: FlaskClient):
        """Test creating a new study session."""
        payload = {
            "study_activity_id": 1
        }
        response = client.post('/api/study-sessions', json=payload)
        assert response.status_code == 201
        assert response.content_type == 'application/json'
        assert 'id' in response.json
        assert response.json['status'] == 'in_progress'
        assert 'start_time' in response.json
        
    def test_get_session(self, client: FlaskClient):
        """Test retrieving a study session."""
        # First create a session
        create_response = client.post('/api/study-sessions', 
                                    json={"study_activity_id": 1})
        assert create_response.status_code == 201
        session_id = create_response.json['id']
        
        # Then retrieve it
        response = client.get(f'/api/study-sessions/{session_id}')
        assert response.status_code == 200
        assert response.content_type == 'application/json'
        assert response.json['id'] == session_id
        assert response.json['status'] == 'in_progress'
        
    def test_list_sessions(self, client: FlaskClient):
        """Test listing study sessions with pagination."""
        # Create multiple sessions
        for i in range(5):
            client.post('/api/study-sessions', json={"study_activity_id": i + 1})
        
        # Test first page
        response = client.get('/api/study-sessions')
        assert response.status_code == 200
        assert response.content_type == 'application/json'
        
        # Verify pagination structure
        assert 'items' in response.json
        assert 'pagination' in response.json
        pagination = response.json['pagination']
        assert 'current_page' in pagination
        assert 'total_pages' in pagination
        assert 'total_items' in pagination
        assert 'items_per_page' in pagination
        
        # Verify pagination values
        assert pagination['current_page'] == 1
        assert pagination['items_per_page'] == 100
        assert pagination['total_items'] >= 5
        assert isinstance(response.json['items'], list)
        
    def test_pagination_behavior(self, client: FlaskClient):
        """Test pagination behavior with multiple pages."""
        # Create enough sessions for multiple pages (assuming 100 items per page)
        for i in range(105):  # Create 105 sessions to ensure we have at least 2 pages
            client.post('/api/study-sessions', json={"study_activity_id": 1})
        
        # Test first page
        response1 = client.get('/api/study-sessions?page=1')
        assert response1.status_code == 200
        assert len(response1.json['items']) == 100  # First page should have 100 items
        assert response1.json['pagination']['current_page'] == 1
        
        # Test second page
        response2 = client.get('/api/study-sessions?page=2')
        assert response2.status_code == 200
        assert len(response2.json['items']) == 5  # Second page should have remaining 5 items
        assert response2.json['pagination']['current_page'] == 2
        
        # Verify total items and pages are consistent
        assert response1.json['pagination']['total_items'] == response2.json['pagination']['total_items']
        assert response1.json['pagination']['total_pages'] == response2.json['pagination']['total_pages']
        assert response1.json['pagination']['total_pages'] == 2
        
    def test_end_session(self, client: FlaskClient):
        """Test ending a study session."""
        # Create a session
        create_response = client.post('/api/study-sessions', 
                                    json={"study_activity_id": 1})
        assert create_response.status_code == 201
        session_id = create_response.json['id']
        
        # End the session
        end_payload = {
            "score": 85.5
        }
        response = client.put(f'/api/study-sessions/{session_id}/end', 
                            json=end_payload)
        assert response.status_code == 200
        
        # Verify the session is ended
        get_response = client.get(f'/api/study-sessions/{session_id}')
        assert get_response.status_code == 200
        assert get_response.json['status'] == 'completed'
        assert get_response.json['score'] == 85.5
        assert 'end_time' in get_response.json
        
    def test_get_session_not_found(self, client: FlaskClient):
        """Test retrieving a non-existent session."""
        response = client.get('/api/study-sessions/999999')
        assert response.status_code == 404
        assert 'error' in response.json
        
    def test_create_session_invalid_activity(self, client: FlaskClient):
        """Test creating a session with invalid study activity."""
        payload = {
            "study_activity_id": 999999  # Non-existent activity
        }
        response = client.post('/api/study-sessions', json=payload)
        assert response.status_code == 400
        assert 'error' in response.json 