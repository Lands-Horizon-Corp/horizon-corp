import { useState } from 'react'

import { KeysOfOrString } from '@/types'

export interface ISortItem<T = unknown> {
    field: KeysOfOrString<T>
    order: 'asc' | 'desc'
}

export interface ISortingState<T = unknown> {
    sortingState: ISortItem<T>[]
    setSortingState: React.Dispatch<React.SetStateAction<ISortItem<T>[]>>
}

export const useSortingState = <T = unknown>(): ISortingState<T> => {
    const [sortingState, setSortingState] = useState<ISortItem<T>[]>([])
    return { sortingState, setSortingState }
}
