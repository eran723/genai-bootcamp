
import React from 'react';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, PieChart, Pie, Cell } from 'recharts';
import { Book, Clock, Target, TrendingUp, Calendar, Award } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import StatCard from '@/components/UI/StatCard';
import LoadingSpinner from '@/components/UI/LoadingSpinner';
import { useApi } from '@/hooks/useApi';
import { dashboardApi } from '@/services/api';
import { DashboardStats } from '@/types';

const Dashboard: React.FC = () => {
  const { data: stats, loading, error } = useApi<DashboardStats>(() => dashboardApi.getStats());

  if (loading) return <LoadingSpinner text="Loading dashboard..." />;
  if (error) return <div className="text-red-600">Error: {error}</div>;
  if (!stats) return <div>No data available</div>;

  const recentSessionsData = stats.recent_sessions.slice(0, 7).map(session => ({
    date: new Date(session.completed_at).toLocaleDateString('en-US', { month: 'short', day: 'numeric' }),
    score: Math.round((session.correct_answers / session.total_questions) * 100),
    duration: Math.round(session.duration_seconds / 60),
  }));

  const masteryData = [
    { name: 'Beginner', value: stats.mastery_distribution.beginner, color: '#ef4444' },
    { name: 'Intermediate', value: stats.mastery_distribution.intermediate, color: '#f59e0b' },
    { name: 'Advanced', value: stats.mastery_distribution.advanced, color: '#3b82f6' },
    { name: 'Mastered', value: stats.mastery_distribution.mastered, color: '#10b981' },
  ];

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
        <p className="text-gray-600">Track your Japanese learning progress</p>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <StatCard
          title="Study Sessions"
          value={stats.total_study_sessions}
          icon={Book}
          color="blue"
        />
        <StatCard
          title="Words Learned"
          value={stats.total_words_learned}
          icon={Target}
          color="green"
        />
        <StatCard
          title="Current Streak"
          value={`${stats.current_streak} days`}
          icon={Calendar}
          color="orange"
        />
        <StatCard
          title="Study Time"
          value={`${Math.round(stats.total_study_time_minutes / 60)}h`}
          icon={Clock}
          color="purple"
        />
      </div>

      {/* Charts Row */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Recent Performance Chart */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <TrendingUp size={20} />
              Recent Performance
            </CardTitle>
          </CardHeader>
          <CardContent>
            <ResponsiveContainer width="100%" height={300}>
              <BarChart data={recentSessionsData}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="date" />
                <YAxis />
                <Tooltip 
                  formatter={(value, name) => [
                    name === 'score' ? `${value}%` : `${value} min`,
                    name === 'score' ? 'Score' : 'Duration'
                  ]}
                />
                <Bar dataKey="score" fill="#3b82f6" radius={[4, 4, 0, 0]} />
              </BarChart>
            </ResponsiveContainer>
          </CardContent>
        </Card>

        {/* Mastery Distribution */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Award size={20} />
              Mastery Distribution
            </CardTitle>
          </CardHeader>
          <CardContent>
            <ResponsiveContainer width="100%" height={300}>
              <PieChart>
                <Pie
                  data={masteryData}
                  cx="50%"
                  cy="50%"
                  outerRadius={80}
                  dataKey="value"
                  label={({ name, value }) => `${name}: ${value}`}
                >
                  {masteryData.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.color} />
                  ))}
                </Pie>
                <Tooltip />
              </PieChart>
            </ResponsiveContainer>
          </CardContent>
        </Card>
      </div>

      {/* Recent Sessions */}
      <Card>
        <CardHeader>
          <CardTitle>Recent Study Sessions</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            {stats.recent_sessions.slice(0, 5).map((session) => (
              <div key={session.id} className="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                <div>
                  <h4 className="font-medium text-gray-900">{session.activity_name}</h4>
                  <p className="text-sm text-gray-600">
                    {new Date(session.completed_at).toLocaleDateString()} • 
                    {Math.round(session.duration_seconds / 60)} minutes
                  </p>
                </div>
                <div className="text-right">
                  <div className="text-lg font-semibold text-gray-900">
                    {Math.round((session.correct_answers / session.total_questions) * 100)}%
                  </div>
                  <div className="text-sm text-gray-600">
                    {session.correct_answers}/{session.total_questions} correct
                  </div>
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default Dashboard;
