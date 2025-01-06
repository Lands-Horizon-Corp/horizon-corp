import { TSortingState } from "@/hooks/use-sorting-state";

export interface IOperationCallbacks<TDataSuccess = unknown, TError = unknown> {
    onSuccess?: (data: TDataSuccess) => void
    onError?: (error: TError) => void
}

export interface IFilterPaginatedHookProps {
    preloads?: string[],
    sort? : TSortingState
    filterPayload?: Record<string, unknown>
    pagination?: { pageIndex: number; pageSize: number }
}
