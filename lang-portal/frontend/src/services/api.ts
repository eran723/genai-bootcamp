
import axios, { AxiosResponse } from 'axios';
import { 
  StudyActivity, 
  StudySession, 
  Word, 
  WordGroup, 
  DashboardStats, 
  PaginatedResponse 
} from '../types';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Dashboard API
export const dashboardApi = {
  getStats: (): Promise<AxiosResponse<DashboardStats>> =>
    api.get('/dashboard/stats'),
};

// Study Activities API
export const studyActivitiesApi = {
  getAll: (): Promise<AxiosResponse<StudyActivity[]>> =>
    api.get('/study_activities'),
  
  getById: (id: string): Promise<AxiosResponse<StudyActivity>> =>
    api.get(`/study_activities/${id}`),
  
  getSessions: (id: string, page = 1, perPage = 10): Promise<AxiosResponse<PaginatedResponse<StudySession>>> =>
    api.get(`/study_activities/${id}/study_sessions?page=${page}&per_page=${perPage}`),
  
  create: (activity: Omit<StudyActivity, 'id' | 'created_at' | 'updated_at'>): Promise<AxiosResponse<StudyActivity>> =>
    api.post('/study_activities', activity),
  
  update: (id: string, activity: Partial<StudyActivity>): Promise<AxiosResponse<StudyActivity>> =>
    api.put(`/study_activities/${id}`, activity),
  
  delete: (id: string): Promise<AxiosResponse<void>> =>
    api.delete(`/study_activities/${id}`),
};

// Study Sessions API
export const studySessionsApi = {
  getAll: (page = 1, perPage = 10): Promise<AxiosResponse<PaginatedResponse<StudySession>>> =>
    api.get(`/study_sessions?page=${page}&per_page=${perPage}`),
  
  getById: (id: string): Promise<AxiosResponse<StudySession>> =>
    api.get(`/study_sessions/${id}`),
  
  create: (session: Omit<StudySession, 'id' | 'started_at' | 'completed_at'>): Promise<AxiosResponse<StudySession>> =>
    api.post('/study_sessions', session),
  
  update: (id: string, session: Partial<StudySession>): Promise<AxiosResponse<StudySession>> =>
    api.put(`/study_sessions/${id}`, session),
  
  delete: (id: string): Promise<AxiosResponse<void>> =>
    api.delete(`/study_sessions/${id}`),
};

// Words API
export const wordsApi = {
  getAll: (page = 1, perPage = 20, filters?: {
    search?: string;
    jlpt_level?: string;
    part_of_speech?: string;
    group_id?: string;
  }): Promise<AxiosResponse<PaginatedResponse<Word>>> => {
    const params = new URLSearchParams({
      page: page.toString(),
      per_page: perPage.toString(),
      ...filters,
    });
    return api.get(`/words?${params}`);
  },
  
  getById: (id: string): Promise<AxiosResponse<Word>> =>
    api.get(`/words/${id}`),
  
  create: (word: Omit<Word, 'id' | 'created_at' | 'mastery_level' | 'correct_count' | 'incorrect_count'>): Promise<AxiosResponse<Word>> =>
    api.post('/words', word),
  
  update: (id: string, word: Partial<Word>): Promise<AxiosResponse<Word>> =>
    api.put(`/words/${id}`, word),
  
  delete: (id: string): Promise<AxiosResponse<void>> =>
    api.delete(`/words/${id}`),
};

// Groups API
export const groupsApi = {
  getAll: (): Promise<AxiosResponse<WordGroup[]>> =>
    api.get('/groups'),
  
  getById: (id: string): Promise<AxiosResponse<WordGroup>> =>
    api.get(`/groups/${id}`),
  
  getWords: (id: string, page = 1, perPage = 20): Promise<AxiosResponse<PaginatedResponse<Word>>> =>
    api.get(`/groups/${id}/words?page=${page}&per_page=${perPage}`),
  
  create: (group: Omit<WordGroup, 'id' | 'created_at' | 'updated_at' | 'word_count' | 'average_mastery'>): Promise<AxiosResponse<WordGroup>> =>
    api.post('/groups', group),
  
  update: (id: string, group: Partial<WordGroup>): Promise<AxiosResponse<WordGroup>> =>
    api.put(`/groups/${id}`, group),
  
  delete: (id: string): Promise<AxiosResponse<void>> =>
    api.delete(`/groups/${id}`),
};

export default api;
