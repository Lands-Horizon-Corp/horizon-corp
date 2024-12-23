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



// const sampleData: FilterPages<{ id: number, name: string }> = {
//   data: [
//     { id: 1, name: 'John Doe' },
//     { id: 2, name: 'Jane Smith' },
//     { id: 3, name: 'Tom Brown' },
//     { id: 4, name: 'Lisa White' },
//     { id: 5, name: 'James Black' }
//   ],
//   pageIndex: 1,
//   totalPage: 5,
//   pageSize: 5,
//   totalSize: 25,
//   pages: [
//     { page: '/api/data?filter=base64' },
//     { page: '/api/data?filter=base64' },
//     { page: '/api/data?filter=base64' },
//     { page: '/api/data?filter=base64' },
//     { page: '/api/data?filter=base64' }
//   ]