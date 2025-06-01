# Frontend Infrastructure Prompt
## Overview
We are developing a desktop-only Japanese language learning web application. The application will serve as:

A centralized portal for launching and interacting with various study activities.

A system to browse, group, and manage Japanese vocabulary and grammar.

A dashboard to review learning progress, study streaks, and performance metrics.

This frontend will interact with an already-built Go (Gin) backend over a JSON REST API, using a single-user model with no authentication.

## Technical Stack
- React.js – UI library
- TypeScript – language for strict type checking
- Vite – build tool and development server
- Tailwind CSS – utility-first CSS framework
- ShadCN/UI – prebuilt component library

## Testing:
- Unit tests (Jest + React Testing Library)
- Performance audits (Lighthouse)

## Project Structure
csharp
Copy
Edit
frontend/
├── src/
│   ├── components/        # Shared UI components
│   ├── pages/             # Route-specific views
│   ├── hooks/             # Custom React hooks
│   ├── services/          # API interaction modules
│   ├── types/             # TypeScript interfaces and types
│   └── utils/             # Utility functions
├── public/
├── index.html
├── vite.config.ts
├── tailwind.config.js
└── tsconfig.json

## Routing & Pages
### Path (Route)
"/dashboard"
### Description
Shows recent session summary, study streak, and aggregated stats.

### Path (Route)
"/study_activities"
### Description
Lists available study activities with types and thumbnails.

### Path (Route)
"/study_activities/:id"
### Description
Displays details of an activity and allows launching new sessions.

### Path (Route)
"/study_activities/:id/study_sessions"
### Description
Shows all sessions for the selected activity.

### Path (Route)
"/study_sessions"
### Description
Overview of all study sessions (history).

### Path (Route)
"/study_sessions/:id"
### Description
View details for a specific study session, including score and words.

### Path (Route)
"/groups"
### Description
Explore all word groups and associated metadata.

### Path (Route)
"/groups/:id"
### Description
View words in a group and stats like accuracy and mastery.
	
### Path (Route)
"/words"
### Description	
Vocabulary browser with filters, counts, and details.

### Path (Route)
"/words/:id"
### Description
Detail view of a vocabulary word, including correctness history.

## API Integration
Integrate each page using the backend’s REST API endpoints:

- Use "/api/dashboard/*" for dashboard stats
- Use "/api/study_activities", "/api/study_sessions", "/api/groups", and "/api/words" endpoints for CRUD operations.
- All requests and responses are JSON formatted
- No authentication required (single-user system)

## UI/UX Notes
- Only desktop resolution should be supported.
- Data visualization should be minimal but informative (e.g., simple charts for scores).
- Use Tailwind for layouts and spacing; ShadCN for inputs, buttons, tabs, modals.
- All lists (e.g. words, sessions) should support pagination (based on backend response).

## Tasks
- Scaffold frontend project using Vite, TypeScript, and Tailwind.
- Build reusable layout components (sidebar, header, loader).
- Implement routing using React Router.
- Create API layer (e.g. Axios) for backend interaction.
- Implement each page with mock data first, then connect to API.
- Write tests for components and API services.
- Perform Lighthouse audits to measure accessibility and performance.