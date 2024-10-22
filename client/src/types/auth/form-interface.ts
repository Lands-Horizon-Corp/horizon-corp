import { IBaseCompNoChild } from '../component'

export interface IAuthForm<T, D = unknown> extends IBaseCompNoChild {
    readOnly?: boolean
    defaultValues?: T
    onSuccess?: (data: D) => void
    onError?: (e: unknown) => void
    onLoading?: (loadingState: boolean) => void
}
