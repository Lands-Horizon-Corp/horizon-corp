import { useCallback, useState } from 'react'
import useLoadingErrorState from './use-loading-error-state'

type ApiFunction<T> = (...args: any[]) => Promise<T>

interface UseApiRequestOptions<TData, TErr> {
    onError?: (error: unknown) => TErr
    onSuccess?: (data: TData) => void
}

const useApiRequest = <TData, TFunc extends ApiFunction<TData>, TErr = string>(
    requestFunction: TFunc,
    options?: UseApiRequestOptions<TData, TErr>
) => {
    const [data, setData] = useState<TData | null>(null)
    const { loading, setLoading, error, setError } =
        useLoadingErrorState<TErr>()

    const request = useCallback(
        async (...args: Parameters<TFunc>) => {
            setLoading(true)
            setError(null)
            try {
                const response = await requestFunction(...args)
                setData(response)
                options?.onSuccess?.(response)
                return response
            } catch (e) {
                if (options?.onError) {
                    return setError(options.onError(e))
                }

                if (e instanceof Error) setError(e.message as TErr)

                throw e
            } finally {
                setLoading(false)
            }
        },
        [requestFunction, options]
    )

    return { data, loading, setLoading, error, setError, request }
}

export default useApiRequest
