interface Pages {
  page: string
  pageIndex: string
}

export interface FilterPages<T> {
  data: T[],
  pageIndex: number
  totalPage: number
  pageSize: number
  totalSize: number // Total size of all data including all pages 
  pages: Pages[]
}