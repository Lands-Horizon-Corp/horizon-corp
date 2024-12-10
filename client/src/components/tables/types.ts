import { IBaseCompNoChild } from '@/types'

export interface TableProps<T> extends IBaseCompNoChild {
    onSelectData? : (selectedData: T[]) => void
}
