export interface Pagination {
    limit: number
    page: number
    sortBy?: string
    sortOrder?: 'ASC' | 'DESC'
    total?: number
    totalPages?: number
    prevPage?: number
    nextPage?: number
}
