import { IBaseCompNoChild } from '../../component/index'

export interface IForm<TDefaultVals, IData = unknown, IErr = unknown>
    extends IBaseCompNoChild {
    readOnly?: boolean
    defaultValues?: TDefaultVals
    onSuccess?: (data: IData) => void
    onError?: (e: IErr) => void
    onLoading?: (loadingState: boolean) => void
}
