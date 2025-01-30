import { useState } from 'react'

import {
    PAGINATION_INITIAL_INDEX,
    PAGINATION_INITIAL_PAGE_SIZE,
} from '@/constants'

export type TPagination = { pageSize?: number; pageIndex?: number }

export const usePagination = (props: TPagination = {}) => {
    const [pagination, setPagination] = useState<TPagination>({
        pageIndex: props?.pageIndex ?? PAGINATION_INITIAL_INDEX,
        pageSize: props?.pageSize ?? PAGINATION_INITIAL_PAGE_SIZE,
    })

    return { pagination, setPagination }
}
