from flask import Flask, jsonify, request

def create_app():
    app = Flask(__name__)
    
    # In-memory storage for todos
    todos = []
    
    @app.route('/health')
    def health_check():
        return jsonify({"status": "healthy"})
    
    @app.route('/api/todos', methods=['GET'])
    def get_todos():
        return jsonify(todos)
    
    @app.route('/api/todos', methods=['POST'])
    def create_todo():
        if not request.json or 'title' not in request.json:
            return jsonify({"error": "Title is required"}), 400
        
        todo = {
            'id': len(todos) + 1,
            'title': request.json['title'],
            'completed': False
        }
        todos.append(todo)
        return jsonify(todo), 201
    
    @app.route('/api/todos/<int:todo_id>', methods=['GET'])
    def get_todo(todo_id):
        todo = next((todo for todo in todos if todo['id'] == todo_id), None)
        if todo is None:
            return jsonify({"error": "Todo not found"}), 404
        return jsonify(todo)
    
    return app

if __name__ == '__main__':
    app = create_app()
    app.run(debug=True) 