import { IBaseCompNoChild } from '@/types'
import { TFilterObject } from '../data-table/data-table-filters/data-table-filter-context'

export interface TableProps<T> extends IBaseCompNoChild {
    onSelectData?: (selectedData: T[]) => void
    defaultFilter?: TFilterObject
}
