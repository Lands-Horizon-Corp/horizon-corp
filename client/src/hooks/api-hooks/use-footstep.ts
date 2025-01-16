import { toast } from 'sonner';
import { useQuery } from '@tanstack/react-query';

import { toBase64, withCatchAsync } from '@/utils';
import { serverRequestErrExtractor } from '@/helpers';
import { FootstepService } from '@/horizon-corp/services';

import { IFilterPaginatedHookProps } from './types';
import { FootstepPaginatedResource } from '@/horizon-corp/types';

export const useFilteredPaginatedFootsteps = ({
    sort,
    filterPayload,
    preloads = [],
    pagination = { pageSize: 10, pageIndex: 1 },
}: IFilterPaginatedHookProps = {}) => {
    return useQuery<FootstepPaginatedResource, string>({
        queryKey: ['footstep', 'resource-query', filterPayload, pagination, sort],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                FootstepService.getFootsteps({
                    preloads,
                    pagination,
                    sort: sort && toBase64(sort),
                    filters: filterPayload && toBase64(filterPayload),
                })
            );

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error });
                toast.error(errorMessage);
                throw errorMessage;
            }

            return result;
        },
        initialData: {
            data: [],
            pages: [],
            totalSize: 0,
            totalPage: 1,
            ...pagination,
        },
        retry: 1,
    });
};
