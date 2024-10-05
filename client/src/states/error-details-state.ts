import useCrud from '@/composables/useCrud';
import { ErrorDetails } from '@/types';

const useErrorDetailsState = () => {
  const crud = useCrud<ErrorDetails>({
    identifier: 'id',
    path: '/errors',
  });

  return crud();
};

export default useErrorDetailsState;
