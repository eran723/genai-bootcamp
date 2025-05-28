# Bacend Technical Specs
This file was manually written, the purpose of this file is to orderly lay the plan in mind, in this case the example of the course, but this only demonstrates the pottential in planning well.

## Business Goal: 
A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

## Technical Requirements
- The backend will be built using Go
- The database will be SQLite3
- The API will be built using Gin
- Mage is a task runner for Go.
- The API will always return JSON
- There will be no authentication or authorization
- Everything would be treated as a single user.

## Directory Structure
'''text
backend_go/
│   ├── cmd/                    # Application entrypoints
│   │   └── server/            # Main server application
│   ├── internal/              # Private application code
│   │   ├── api/              # API handlers and routes
│   │   ├── models/           # Database models
│   │   ├── database/         # Database connection and queries
│   │   └── service/          # Business logic
│   ├── db/                  # Database related files
│   │   ├── migrations/      # SQL migration files
│   │   └── seeds/          # Seed data files
│   ├── magefiles/           # Mage task definitions
│   ├── go.mod              # Go module file
│   ├── go.sum              # Go dependencies checksum
│   └── words.db
'''

## Database Schema

Our database will be a single sqlite database called 'words.db' that will be in the root of the project folder of 'backend-go'

We have the following tables:
- words = stored vocabulary words
    * id (int)
    * japanese (str)
    * romji (str)
    * english (str)
    * parts (json)

- word_groups - join table for words and groups : many-to-many
    * id (int)
    * name (str)

- groups - thematic groups of words
    * id (int) 
    * name (str)
    * description (str)
    
- study_sessions - records of study sessions grouping word_review_items
    * id (int)
    * start_time (datetime)
    * end_time (datetime)
    * score (float)
    * status (str) # active, completed, abandoned
    * study_activity_id (int)

- study_activities - a specific study activity, linking study session to group
    * id (int)
    * session_id (int) # foreign key to study_sessions
    * group_id (int) # foreign key to groups
    * activity_type (str) # quiz, flashcard, etc
    * created_at (datetime)

- word_review_items - a record of word practice, determining if the word was correct or not
    * id (int)
    * activity_id (int) # foreign key to study_activities
    * word_id (int) # foreign key to words
    * is_correct (bool)
    * response (str) # user's answer
    * created_at (datetime)

## API Endpoints
### GET /api/dashboard/last_session
Returns details about the user's most recent study session.

**Response**
```json
{
  "session_id": 123,
  "start_time": "2024-03-20T14:00:00Z",
  "end_time": "2024-03-20T14:30:00Z",
  "score": 85.5,
  "status": "completed",
  "activity_type": "flashcard",
  "group_id": 456,
  "group_name": "Basic Verbs",
  "words_reviewed": 20,
  "correct_answers": 17
}
```

### GET /api/dashboard/stats
```json
{
  "total_study_time": 3600,
  "sessions_completed": 25,
  "total_words_reviewed": 500,
  "success_rate": 85.5,
  "study_streak_days": 4
}
```

### GET /api/dashboard/progress
```json
{
  "overall_completion": 45.5,
  "total_words_studied": 150,
  "total_available_words": 1500
}
```

### GET /api/study_activities/:id
```json
{
  "id": 1,
  "name": "Basic Vocabulary Quiz",
  "activity_type": "quiz",
  "thumbnail_url": "https://example.com/thumbnail.jpg"
}
```

### GET /api/study_activities/:id/study_sessions
```json
{
  "items": [
    {
      "id": 123,
      "activity_name":"Vocabulary Quis",
      "group_name": "Basic greetings",
      "start_time": "2024-03-20T14:00:00Z",
      "end_time": "2024-03-20T14:30:00Z",
      "score": 85.5,
      "status": "completed",
      "review_items_count": 20
    }
  ],
  "pagination":{
    "current_page":1,
    "total_pages":5,
    "total_items":100,
    "items_per_page":20
  }
}
```

### POST /api/study_activities
Required params: group_id, study_activity_id

Request:
```json
{
  "group_id": 1,
  "study_activity_id": 123
}
```

Response:
```json
{
  "id": 456,
  "group_id": 1
}
```

### GET /api/words
Pagination with 100 items per page
```json
{
  "items": [
    {
      "id": 1,
      "japanese": "こんにちは",
      "romaji": "konnichiwa",
      "english": "hello",
      "correct_count": 8,
      "wrong_count": 2
    }
  ],
  "pagination":{
    "current_page":1,
    "total_pages":5,
    "total_items":1500,
    "items_per_page":100
  }
}
```

### GET /api/words/:id
```json
{
  "id": 1,
  "japanese": "こんにちは",
  "romaji": "konnichiwa",
  "english": "hello",
  "parts": {
    "type": "greeting",
    "formality": "neutral"
  },
  "correct_count": 8,
  "wrong_count": 2,
  "success_rate": 80.0,
  "groups": ["Basic Greetings", "Common Phrases"]
}
```


### GET /api/groups
Pagination with 100 items per page
```json
{
  "items": [
    {
      "id": 1,
      "name": "Basic Greetings",
      "description": "Common Japanese greetings",
      "word_count": 20
    }
  ],
  "pagination":{
    "current_page":1,
    "total_pages":5,
    "total_items":1500,
    "items_per_page":100
  }
}
```

### GET /api/groups/:id
```json
{
  "items": [
      "id": 1,
      "japanese": "こんにちは",
      "romaji": "konnichiwa",
      "english": "hello",
      "parts": {
        "type": "greeting",
        "formality": "neutral"
      }
  ],
  "description": "Common Japanese greetings",
  "total_words": 20,
  "mastered_words": 15,
  "average_accuracy": 85.5,
  "pagination":{
    "current_page":1,
    "total_pages":5,
    "total_items":1500,
    "items_per_page":100
  }
}
```

### GET /api/groups/:id/words
```json
{
  "group_id": 1,
  "total_words": 20,
  "items": [
    {
      "id": 1,
      "japanese": "こんにちは",
      "romaji": "konnichiwa",
      "english": "hello",
      "correct_count": 8,
      "wrong_count": 2
    }
  ]
}
```

### GET /api/groups/:id/study_sessions
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2024-03-20T14:00:00Z",
      "end_time": "2024-03-20T14:30:00Z",
      "activity_type": "quiz",
      "review_item_count": 20
    }
  ],
  "pagination":{
    "current_page":1,
    "total_pages":5,
    "total_items":1500,
    "items_per_page":100
  }
}
```

### GET /api/study_sessions
Pagination with 100 items per page
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Basic Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2024-03-20T14:00:00Z",
      "end_time": "2024-03-20T14:30:00Z",
      "review_items_count": 20
    }
  ],
  "pagination":{
    "current_page":1,
    "total_pages":5,
    "total_items":1500,
    "items_per_page":100
  }
}
```

### GET /api/study_sessions/:id
```json
{
  "id": 123,
  "activity_type": "quiz",
  "group_name": "Basic Greetings",
  "start_time": "2024-03-20T14:00:00Z",
  "end_time": "2024-03-20T14:30:00Z",
  "review_items_count": 20,
  "score": 85.5,
  "status": "completed"
}
```

### GET /api/study_sessions/:id/words
```json
{
  "items": [
    {
      "word_id": 1,
      "japanese": "こんにちは",
      "romaji": "konnichiwa",
      "english": "hello",
      "corrected_count":5,
      "wrong_count":2
    }
  ],
  "pagination":{
    "current_page":1,
    "total_pages":5,
    "total_items":1500,
    "items_per_page":100
  }
}
```

### POST /api/study_sessions/:id/word_id/review
Required params: correct

Request:
```json
{
  "study_session_id": 123,
  "word_id": 1,
  "correct": true,
  "response": "hello"
}
```

Response:
```json
{
  "session_id": 123,
  "word_id": 1,
  "is_correct": true,
  "response": "hello",
  "created_at": "2024-03-20T14:15:00Z"
}
```

POST /api/reset_history
```json
{
  "status": "success",
  "message": "Study history has been reset"
}
```

POST /api/full_reset
```json
{
  "status": "success",
  "message": "System has been fully reset"
}
```

## Tasks
Mage is a task runner for 'Go'.
Lets list out possible tasks we need for our lang portal.

### Initialize Database
This task will initialize the database called 'words.db'.

### Migrate Database
This task will run a series of migrations of sql files on the database

Migration live in the 'migration' folder.
The migration files will be run in the order of their file name.
The file names should look like this:
'''sql
0001_init.sql
0002_create_words_table.sql
'''

### Seed Data 
This task will import json files and transform them into target data for our database.

All seed files live in the 'seeds' folder.

In our task we should have a DSL to specific each seed file and its expected group word name.