
import { useState, useEffect } from 'react';
import { AxiosResponse } from 'axios';

interface UseApiState<T> {
  data: T | null;
  loading: boolean;
  error: string | null;
}

export function useApi<T>(
  apiCall: () => Promise<AxiosResponse<T>>,
  dependencies: any[] = []
): UseApiState<T> & { refetch: () => void } {
  const [state, setState] = useState<UseApiState<T>>({
    data: null,
    loading: true,
    error: null,
  });

  const fetchData = async () => {
    try {
      setState(prev => ({ ...prev, loading: true, error: null }));
      const response = await apiCall();
      setState({
        data: response.data,
        loading: false,
        error: null,
      });
    } catch (error: any) {
      setState({
        data: null,
        loading: false,
        error: error.response?.data?.message || error.message || 'An error occurred',
      });
    }
  };

  useEffect(() => {
    fetchData();
  }, dependencies);

  return {
    ...state,
    refetch: fetchData,
  };
}

export function usePaginatedApi<T>(
  apiCall: (page: number, perPage: number) => Promise<AxiosResponse<any>>,
  initialPage = 1,
  perPage = 10
) {
  const [page, setPage] = useState(initialPage);
  const [data, setData] = useState<T[]>([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchData = async (pageNum: number) => {
    try {
      setLoading(true);
      setError(null);
      const response = await apiCall(pageNum, perPage);
      setData(response.data.data);
      setTotal(response.data.total);
      setPage(pageNum);
    } catch (error: any) {
      setError(error.response?.data?.message || error.message || 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData(page);
  }, []);

  return {
    data,
    total,
    loading,
    error,
    page,
    totalPages: Math.ceil(total / perPage),
    nextPage: () => fetchData(page + 1),
    prevPage: () => fetchData(page - 1),
    goToPage: (pageNum: number) => fetchData(pageNum),
    refetch: () => fetchData(page),
  };
}
