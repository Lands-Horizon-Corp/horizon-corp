import * as XLSX from 'xlsx'

export type FileNameWithoutExtension<T extends string> =
    T extends `${string}.${string}` ? never : T

export const exportTableData = <T>(
    data: T[],
    fileName: FileNameWithoutExtension<string>,
    fileType: 'csv' | 'xlsx' = 'csv'
) => {
    const wb = XLSX.utils.book_new()
    const ws = XLSX.utils.json_to_sheet(data)

    XLSX.utils.book_append_sheet(wb, ws, 'data')
    XLSX.writeFile(wb, `${fileName}.${fileType}`)
}
