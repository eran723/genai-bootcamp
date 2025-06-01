
export interface StudyActivity {
  id: string;
  name: string;
  description: string;
  type: 'vocabulary' | 'grammar' | 'kanji' | 'reading';
  thumbnail?: string;
  difficulty: 'beginner' | 'intermediate' | 'advanced';
  created_at: string;
  updated_at: string;
}

export interface StudySession {
  id: string;
  activity_id: string;
  activity_name: string;
  score: number;
  total_questions: number;
  correct_answers: number;
  duration_seconds: number;
  started_at: string;
  completed_at: string;
  words_practiced: Word[];
}

export interface Word {
  id: string;
  japanese: string;
  reading: string;
  english: string;
  part_of_speech: string;
  jlpt_level?: string;
  frequency_rank?: number;
  group_id?: string;
  created_at: string;
  mastery_level: number;
  correct_count: number;
  incorrect_count: number;
  last_practiced?: string;
}

export interface WordGroup {
  id: string;
  name: string;
  description: string;
  color: string;
  word_count: number;
  average_mastery: number;
  created_at: string;
  updated_at: string;
}

export interface DashboardStats {
  total_study_sessions: number;
  total_words_learned: number;
  current_streak: number;
  longest_streak: number;
  total_study_time_minutes: number;
  average_session_score: number;
  recent_sessions: StudySession[];
  mastery_distribution: {
    beginner: number;
    intermediate: number;
    advanced: number;
    mastered: number;
  };
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  per_page: number;
  total_pages: number;
}

export interface ApiError {
  message: string;
  code?: string;
}
