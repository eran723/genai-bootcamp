# Frontend Technical Spec

## Pages

### Dashboard '/dashboard'
#### Purpose
The purpose of this page is to provide a summary of learning and act as the default page when a user visits the web app.

#### Componenets
This page contains the following components:
- Last Study Session
    - shows last activity used
    - shows when last activity used
    - summarizes wrong vs correct from last activity
    - has a link to the group

- Study Progress
    - total words studied
        - across all study sessions show the total words studied out of all possible words in out database
        - display a mastery progress eg. 0%.

- Quick Stats
    - success rate eg. 0%
    - total study sessions. eg. 4.
    - total active groups. eg. 3.
    - study streak eg. 4 days.

- Start Studying Button
    - goes to study activities page

We'll need the following API endpoints:

#### Needed API Endpoints
- GET /api/dashboard
- GET /api/dashboard/last_session
- GET /api/dashboard/stats
- GET /api/dashboard/progress

### Study Activities '/study-activities'
#### Purpose
The purpose of this page is to show a collection of study activities with a thumbnail and its name, to either launch or view a studying activity

#### Components
- Study Activity Card
    - Shows a thumbnail of the activity
    - the name of the activity
    - a launch button to take us to the launch page
    - the view page to view more information about past study sessions for this study activitiy

#### Needed API Endpoints
- Get /study_activity

### Study Activity Show '/study_activities/:id'

#### Purpose
The purpose of this page is to show detailed information about a specific study activity, including its past study sessions, performance metrics, and allow launching new study sessions for this activity. This page serves as both an analytics view for the activity and a gateway to start new study sessions.

#### Components
- Activity Header
    - displays activity name and type
    - shows overall success rate
    - has launch button to start new session

- Performance Summary
    - total sessions completed
    - average score
    - time spent on activity

- Session History
    - list of past study sessions
    - shows date, score, and duration for each session
    - allows clicking into individual session details

- Progress Chart
    - visual representation of scores over time
    - tracks improvement across sessions

- Study Activity Paginated List
    - id (int)
    - name (str)
    - group_id (int) # foreign key to groups
    - activity_type (str) # quiz, flashcard, etc
    - created_at (datetime)
    - total_sessions (int)
    - average_score (float)
    - last_session_date (datetime)
    - thumbnail_url (str)
    - number of review items

#### Needed API Endpoints
    - GET /api/study_activities/:id
    - GET /api/study_activities/:id/study_sessions

### Study Activities Launch '/study_activities/:id/launch'

#### Purpose
The purpose of this page is to serve as the launching point for a new study session within a specific study activity. It prepares and initializes a new study session with the selected activity's parameters and redirects the user to the appropriate study interface.

#### Components
- Name of activity
- Launch form
    - select field for group
    - launch now button

## Behaviour 
After the form is submitted a new tab opens with the study activity based on its URL provided in the database.
Also after the form is submitted the page will redirect you to the study session show page

#### Needed API Endpoints
- POST /api/study_sctivities

### Words '/words'
#### Purpose
The purpose of this page is to display and manage the vocabulary words available in the system.

#### Components
- Search bar for filtering words
- Paginated list of words showing:
    - Columns
        - Japanese text
        - Romaji
        - English translation
        - Correct Cound
        - Wrong Count
    - Pagination with 100 items per page
    - Clicking a japanes word would take us to the word show page.
- Filter options by word groups

#### Needed API Endpoints
- GET /api/words

### Word Show Page '/word/:id'
#### Purpose
The purpose of this page is to display detailed information about a specific vocabulary word.

#### Components
- Japanese text
- Romaji reading
- English translation
- Word parts/components
- Study statistics
    - Correct count
    - Wrong count
    - Success rate
- Word Groups
    - show a series of pills eg. tags
    - when group name is clicked it will take us to the group showpage.

#### Needed API Endpoints
- GET /api/words/:id

### Word Groups '/groups'
#### Purpose
The purpose of this page is to display and manage word groups that organize vocabulary words by themes or categories.

#### Components
- Paginated List of word groups showing:
    - Group name
    - Description
    - Word count
- Clicking a group name takes user to the group show page

#### Needed API Endpoints
- GET /api/groups

### Group Show '/groups/:ids'
#### Purpose
The purpose of this page is to display detailed information about a specific word group and its associated vocabulary words.

#### Components
- Group name and description
- Group Statistics
    - Total word count
- Paginated List of words in the group showing:
    - Should use the same component as the words index page
- Study Sessions (Paginated list of study sessions)
    - Should use the same component as the words index page

#### Needed API Endpoints
- GET /api/groups/:id (the name and group stats)
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions

### Study Sessions '/study_sessions'
#### Purpose
The purpose of this page is to display a list of study sessions and their results.

#### Components
- Paginated list of study sessions showing:
    - Columns
        - ID 
        - Activity name
        - Group name
        - start time
        - end time
        - number of review items
    - Clicking of the study session id will take us to the study session page

#### Needed API Endpoints
- GET /api/study_sessions


### Study Session Show 'study_sessions/:id'
#### Purpose
The purpose of this page is to display detailed information about a specific study session, including performance metrics and individual word review items.

#### Components
- Study session details showing:
    - Activity type
    - Group name
    - Start time
    - End time
    - Score
    - Status
    = Number of review items
- Paginated List of word review items showing:
    - Japanese word
    - English word
    - User's response
    - Whether it was correct
    - Time taken
- Summary statistics:
    - Total words reviewed
    - Accuracy percentage
    - Average time per word

#### Needed API Endpoints
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words


### Settings Page '/settings'
#### Purpose
The purpose of this page is to make configurations to the study portal

#### Components
- Theme selection (light/dark mode)
- Reset history button
    - this will delete all study sessions and word review items
- Full Reset Button
    - This will drop all tables and recreate seed data.

#### Needed API Endpoints
- POST /api/reset_history
- POST /api/full_reset
