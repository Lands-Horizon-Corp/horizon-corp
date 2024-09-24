import { UserBase } from '@/types'
import { IBaseCompNoChild } from './base-interface'

export interface IAuthForm<T> extends IBaseCompNoChild {
    readOnly?: boolean
    defaultValues?: T
    onSuccess?: (responseData: UserBase) => void
    onError?: (e: unknown) => void
    onLoading?: (loadingState: boolean) => void

    // isLoading?: boolean
    // errorMessage?: string
    // onSubmit?: (val: T) => void
}
