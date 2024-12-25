import { toast } from "sonner";
import { useQuery } from "@tanstack/react-query";

import { toBase64, withCatchAsync } from "@/utils";
import { serverRequestErrExtractor } from "@/helpers";
import { BranchPaginatedResource } from "@/horizon-corp/types";
import BranchService from "@/horizon-corp/server/admin/BranchService";


export const useFilteredPaginatedBranch =  (
    filterState: { finalFilters: Record<string, unknown> },
    pagination: { pageIndex: number; pageSize: number }
) => {
    return useQuery<BranchPaginatedResource, string>({
        queryKey: ['table', 'branches', filterState.finalFilters, pagination],
        queryFn: async () => {
            const [error, result] = await withCatchAsync(
                BranchService.filter(
                    toBase64({
                        preloads: ['Media', 'Owner'],
                        ...pagination,
                        ...filterState.finalFilters,
                    })
                )
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            return result
        },
        initialData: {
            data: [],
            pages: [],
            totalSize: 0,
            totalPage: 1,
            ...pagination,
        },
        retry: 1,
    })
}