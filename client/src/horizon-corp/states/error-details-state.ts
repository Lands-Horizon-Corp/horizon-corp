import useCrud from '@/horizon-corp/composables/useCrud';
import { ErrorDetails } from '../types/error-details';


const useErrorDetailsState = () => {
  const crud = useCrud<ErrorDetails>({
    identifier: 'id',
    path: '/errors',
  });

  return crud();
};

export default useErrorDetailsState;
