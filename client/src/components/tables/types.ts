import { IBaseCompNoChild } from '@/types'
import { TFilterObject } from '@/contexts/filter-context'

export interface TableProps<T> extends IBaseCompNoChild {
    defaultFilter?: TFilterObject
    onSelectData?: (selectedData: T[]) => void
}
