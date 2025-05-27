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
- The API will be built using GIn
- The API will always return JSON
- There will be no authentication or authorization
- Everything would be treated as a single user.

## Database Schema

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

### API Endpoints
- GET /api/dashboard
- GET /api/dashboard/last_session
- GET /api/dashboard/stats
- GET /api/dashboard/progress
- GET /api/study_activities/:id
- GET /api/study_activities/:id/study_sessions

- POST /api/study_sctivities
  -required params: group_id, study_activity_id

- GET /api/words
  - pagination with 100 items per page
- GET /api/words/:id
- GET /api/groups
  - pagination with 100 items per page
- GET /api/groups/:id (the name and group stats)
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions
- GET /api/study_sessions
  - pagination with 100 items per page
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words
- POST /api/reset_history
- POST /api/full_reset
- POST /api/study_sessions/:id/word_id/review
  - required params: correct