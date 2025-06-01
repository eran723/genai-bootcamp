
import React from 'react';
import { useParams, Link } from 'react-router-dom';
import { ArrowLeft, Play, Clock, Target, BookOpen } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import LoadingSpinner from '@/components/UI/LoadingSpinner';
import { useApi } from '@/hooks/useApi';
import { studyActivitiesApi } from '@/services/api';
import { StudyActivity } from '@/types';

const StudyActivityDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const { data: activity, loading, error } = useApi<StudyActivity>(() => 
    studyActivitiesApi.getById(id!), [id]
  );

  if (loading) return <LoadingSpinner text="Loading activity details..." />;
  if (error) return <div className="text-red-600">Error: {error}</div>;
  if (!activity) return <div>Activity not found</div>;

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case 'beginner': return 'bg-green-100 text-green-800';
      case 'intermediate': return 'bg-yellow-100 text-yellow-800';
      case 'advanced': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Link to="/study_activities">
          <Button variant="ghost" size="sm">
            <ArrowLeft size={16} className="mr-2" />
            Back to Activities
          </Button>
        </Link>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Main Content */}
        <div className="lg:col-span-2 space-y-6">
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <CardTitle className="text-2xl">{activity.name}</CardTitle>
                <Badge className={getDifficultyColor(activity.difficulty)}>
                  {activity.difficulty}
                </Badge>
              </div>
              <Badge variant="outline" className="w-fit capitalize">
                {activity.type}
              </Badge>
            </CardHeader>
            <CardContent>
              <p className="text-gray-600 text-lg leading-relaxed">
                {activity.description}
              </p>
            </CardContent>
          </Card>

          {/* Recent Sessions */}
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <CardTitle>Recent Sessions</CardTitle>
                <Link to={`/study_activities/${activity.id}/study_sessions`}>
                  <Button variant="outline" size="sm">View All</Button>
                </Link>
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-center py-8 text-gray-500">
                <Clock size={48} className="mx-auto mb-4 text-gray-400" />
                <p>No sessions yet for this activity</p>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Sidebar */}
        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle>Quick Actions</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <Button className="w-full flex items-center gap-2" size="lg">
                <Play size={20} />
                Start Session
              </Button>
              <Button variant="outline" className="w-full">
                Edit Activity
              </Button>
              <Button variant="destructive" className="w-full">
                Delete Activity
              </Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Activity Stats</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex items-center justify-between">
                <span className="text-gray-600">Total Sessions</span>
                <span className="font-medium">0</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">Average Score</span>
                <span className="font-medium">-</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">Total Time</span>
                <span className="font-medium">0h 0m</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">Last Played</span>
                <span className="font-medium">Never</span>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default StudyActivityDetail;
