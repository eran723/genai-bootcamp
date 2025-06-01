
import React, { useState } from 'react';
import { Search, Filter, Plus, Book } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import LoadingSpinner from '@/components/UI/LoadingSpinner';
import { usePaginatedApi } from '@/hooks/useApi';
import { wordsApi } from '@/services/api';
import { Word } from '@/types';
import { Link } from 'react-router-dom';

const Words: React.FC = () => {
  const [filters, setFilters] = useState({
    search: '',
    jlpt_level: '',
    part_of_speech: '',
  });

  const { 
    data: words, 
    loading, 
    error, 
    page, 
    totalPages, 
    nextPage, 
    prevPage,
    goToPage 
  } = usePaginatedApi<Word>(
    (page, perPage) => wordsApi.getAll(page, perPage, filters),
    1, 
    20
  );

  const getMasteryColor = (level: number) => {
    if (level >= 80) return 'bg-green-100 text-green-800';
    if (level >= 60) return 'bg-blue-100 text-blue-800';
    if (level >= 40) return 'bg-yellow-100 text-yellow-800';
    return 'bg-red-100 text-red-800';
  };

  const getMasteryLabel = (level: number) => {
    if (level >= 80) return 'Mastered';
    if (level >= 60) return 'Advanced';
    if (level >= 40) return 'Intermediate';
    return 'Beginner';
  };

  if (loading) return <LoadingSpinner text="Loading vocabulary..." />;
  if (error) return <div className="text-red-600">Error: {error}</div>;

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Vocabulary</h1>
          <p className="text-gray-600">Browse and manage your Japanese vocabulary</p>
        </div>
        <Button className="flex items-center gap-2">
          <Plus size={20} />
          Add Word
        </Button>
      </div>

      {/* Filters */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Filter size={20} />
            Filters
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
            <div className="relative">
              <Search size={16} className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
              <Input
                placeholder="Search words..."
                value={filters.search}
                onChange={(e) => setFilters({ ...filters, search: e.target.value })}
                className="pl-10"
              />
            </div>
            <Select value={filters.jlpt_level} onValueChange={(value) => setFilters({ ...filters, jlpt_level: value })}>
              <SelectTrigger>
                <SelectValue placeholder="JLPT Level" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="">All Levels</SelectItem>
                <SelectItem value="N5">N5</SelectItem>
                <SelectItem value="N4">N4</SelectItem>
                <SelectItem value="N3">N3</SelectItem>
                <SelectItem value="N2">N2</SelectItem>
                <SelectItem value="N1">N1</SelectItem>
              </SelectContent>
            </Select>
            <Select value={filters.part_of_speech} onValueChange={(value) => setFilters({ ...filters, part_of_speech: value })}>
              <SelectTrigger>
                <SelectValue placeholder="Part of Speech" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="">All Types</SelectItem>
                <SelectItem value="noun">Noun</SelectItem>
                <SelectItem value="verb">Verb</SelectItem>
                <SelectItem value="adjective">Adjective</SelectItem>
                <SelectItem value="adverb">Adverb</SelectItem>
              </SelectContent>
            </Select>
            <Button>Apply Filters</Button>
          </div>
        </CardContent>
      </Card>

      {/* Words Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {words?.map((word) => (
          <Card key={word.id} className="hover:shadow-lg transition-shadow">
            <CardHeader className="pb-3">
              <div className="flex items-center justify-between">
                <Badge className={getMasteryColor(word.mastery_level)}>
                  {getMasteryLabel(word.mastery_level)}
                </Badge>
                {word.jlpt_level && (
                  <Badge variant="outline">{word.jlpt_level}</Badge>
                )}
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-3">
                <div>
                  <div className="text-2xl font-bold text-gray-900 mb-1">
                    {word.japanese}
                  </div>
                  <div className="text-sm text-gray-600">
                    {word.reading}
                  </div>
                </div>
                <div className="text-gray-700">
                  {word.english}
                </div>
                <div className="flex items-center justify-between text-sm">
                  <Badge variant="secondary" className="capitalize">
                    {word.part_of_speech}
                  </Badge>
                  <div className="text-gray-500">
                    {word.correct_count}✓ / {word.incorrect_count}✗
                  </div>
                </div>
                <Link to={`/words/${word.id}`}>
                  <Button size="sm" className="w-full">
                    View Details
                  </Button>
                </Link>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      {/* Pagination */}
      {totalPages > 1 && (
        <div className="flex items-center justify-between">
          <Button 
            variant="outline" 
            onClick={prevPage}
            disabled={page === 1}
          >
            Previous
          </Button>
          <span className="text-sm text-gray-600">
            Page {page} of {totalPages}
          </span>
          <Button 
            variant="outline" 
            onClick={nextPage}
            disabled={page === totalPages}
          >
            Next
          </Button>
        </div>
      )}

      {(!words || words.length === 0) && (
        <div className="text-center py-12">
          <Book size={48} className="mx-auto text-gray-400 mb-4" />
          <h3 className="text-lg font-medium text-gray-900 mb-2">No words found</h3>
          <p className="text-gray-600 mb-4">Add some vocabulary to get started</p>
          <Button>
            <Plus size={20} className="mr-2" />
            Add Word
          </Button>
        </div>
      )}
    </div>
  );
};

export default Words;
