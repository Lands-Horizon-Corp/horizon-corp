export interface IOperationCallbacks<TDataSuccess = unknown, TError = unknown> {
    onSuccess?: (data: TDataSuccess) => void
    onError?: (error: TError) => void
}

export interface IFilterPaginatedHookProps {
    filterPayload?: Record<string, unknown>
    pagination?: { pageIndex: number; pageSize: number }
    preloads?: string[]
}
