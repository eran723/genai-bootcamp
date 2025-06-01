
import React from 'react';
import { Link } from 'react-router-dom';
import { Plus, BookOpen, Brain, Eye, FileText } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import LoadingSpinner from '@/components/UI/LoadingSpinner';
import { useApi } from '@/hooks/useApi';
import { studyActivitiesApi } from '@/services/api';
import { StudyActivity } from '@/types';

const StudyActivities: React.FC = () => {
  const { data: activities, loading, error } = useApi<StudyActivity[]>(() => studyActivitiesApi.getAll());

  const getActivityIcon = (type: string) => {
    switch (type) {
      case 'vocabulary': return BookOpen;
      case 'grammar': return FileText;
      case 'kanji': return Brain;
      case 'reading': return Eye;
      default: return BookOpen;
    }
  };

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case 'beginner': return 'bg-green-100 text-green-800';
      case 'intermediate': return 'bg-yellow-100 text-yellow-800';
      case 'advanced': return 'bg-red-100 text-red-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  };

  if (loading) return <LoadingSpinner text="Loading study activities..." />;
  if (error) return <div className="text-red-600">Error: {error}</div>;

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Study Activities</h1>
          <p className="text-gray-600">Choose your learning activities</p>
        </div>
        <Button className="flex items-center gap-2">
          <Plus size={20} />
          New Activity
        </Button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {activities?.map((activity) => {
          const Icon = getActivityIcon(activity.type);
          
          return (
            <Card key={activity.id} className="hover:shadow-lg transition-shadow">
              <CardHeader>
                <div className="flex items-center justify-between">
                  <Icon size={24} className="text-blue-600" />
                  <Badge className={getDifficultyColor(activity.difficulty)}>
                    {activity.difficulty}
                  </Badge>
                </div>
                <CardTitle className="text-lg">{activity.name}</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600 mb-4 line-clamp-3">{activity.description}</p>
                <div className="flex items-center justify-between">
                  <Badge variant="outline" className="capitalize">
                    {activity.type}
                  </Badge>
                  <Link to={`/study_activities/${activity.id}`}>
                    <Button size="sm">View Details</Button>
                  </Link>
                </div>
              </CardContent>
            </Card>
          );
        })}
      </div>

      {(!activities || activities.length === 0) && (
        <div className="text-center py-12">
          <BookOpen size={48} className="mx-auto text-gray-400 mb-4" />
          <h3 className="text-lg font-medium text-gray-900 mb-2">No activities yet</h3>
          <p className="text-gray-600 mb-4">Create your first study activity to get started</p>
          <Button>
            <Plus size={20} className="mr-2" />
            Create Activity
          </Button>
        </div>
      )}
    </div>
  );
};

export default StudyActivities;
