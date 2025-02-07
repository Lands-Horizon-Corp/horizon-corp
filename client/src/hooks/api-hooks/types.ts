import { TSortingState } from '@/hooks/use-sorting-state'

export interface IOperationCallbacks<TDataSuccess = unknown, TError = unknown> {
    onSuccess?: (data: TDataSuccess) => void
    onError?: (error: TError) => void
}

export interface IFilterPaginatedHookProps extends IAPIPreloads {
    sort?: TSortingState
    filterPayload?: Record<string, unknown>
    pagination?: { pageIndex: number; pageSize: number }
}

export interface IAPIPreloads {
    preloads?: string[]
}

export interface IQueryProps {
    enabled?: boolean
}
