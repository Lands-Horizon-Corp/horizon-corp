export interface IOperationCallbacks<TDataSuccess = unknown, TError = unknown> {
    onSuccess?: (data: TDataSuccess) => void
    onError?: (error: TError) => void
}