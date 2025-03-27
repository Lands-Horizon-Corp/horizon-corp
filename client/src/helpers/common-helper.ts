import axios from 'axios'

export const randomChoose = <T>(data: Array<T>) => {
    return data[~~(Math.random() * data.length)]
}

export const formatBytes = (bytes: number) => {
    if (bytes === 0) return '0 Bytes'
    const sizeUnits = ['Bytes', 'KB', 'MB', 'GB', 'TB']
    const i = Math.floor(Math.log(bytes) / Math.log(1024))
    const formattedSize = (bytes / Math.pow(1024, i)).toFixed(2)
    return `${formattedSize} ${sizeUnits[i]}`
}

export const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString()
}

export const dataUrlToFile = (
    dataUrl: string,
    filename: string
): File | undefined => {
    const arr = dataUrl.split(',')
    if (arr.length < 2) {
        return undefined
    }

    const mimeArr = arr[0].match(/:(.*?);/)
    if (!mimeArr || mimeArr.length < 2) {
        return undefined
    }

    const mime = mimeArr[1]
    const byteString = atob(arr[1]) // Decode base64 string
    const byteNumbers = new Uint8Array(byteString.length)

    for (let i = 0; i < byteString.length; i++) {
        byteNumbers[i] = byteString.charCodeAt(i)
    }

    return new File([byteNumbers], filename, { type: mime })
}

type FileCategory =
    | 'audio'
    | 'video'
    | 'doc'
    | 'pdf'
    | 'sheet'
    | 'text'
    | 'image'
    | 'unknown'

export const getFileCategory = (
    fileName: string,
    mimeType: string
): FileCategory => {
    const fileExtension = fileName.split('.').pop()?.toLowerCase()

    const fileTypes: Partial<Record<FileCategory, string[]>> = {
        audio: ['mp3', 'wav', 'ogg', 'aac', 'flac', 'm4a'],
        video: ['mp4', 'mkv', 'avi', 'mov', 'wmv', 'flv'],
        doc: ['doc', 'docx', 'odt', 'rtf', 'wps'],
        pdf: ['pdf'],
        sheet: ['xls', 'xlsx', 'csv', 'ods'],
        text: ['txt', 'log', 'md'],
        image: ['jpg', 'jpeg', 'png', 'gif', 'bmp', 'svg', 'webp'],
    }

    for (const [category, extensions] of Object.entries(fileTypes)) {
        if (fileExtension && extensions?.includes(fileExtension)) {
            return category as FileCategory
        }
    }

    if (mimeType.startsWith('audio/')) return 'audio'
    if (mimeType.startsWith('video/')) return 'video'
    if (mimeType === 'application/pdf') return 'pdf'
    if (mimeType.startsWith('text/')) return 'text'
    if (mimeType.startsWith('image/')) return 'image'

    return 'unknown'
}

export const getFileType = (file: File): FileCategory => {
    return getFileCategory(file.name, file.type)
}

export const isDate = (value: unknown): boolean => {
    return value instanceof Date && value !== null && value !== undefined
}

export const isNumber = (value: unknown): boolean => {
    return (
        typeof value === 'number' &&
        !isNaN(value) &&
        value !== null &&
        value !== undefined
    )
}

export const isObject = (value: unknown): boolean => {
    return typeof value === 'object' && value !== null && value !== undefined
}

export const isString = (value: unknown): boolean => {
    return (
        typeof value === 'string' &&
        value !== '' &&
        value !== null &&
        value !== undefined
    )
}

export const isArray = (value: unknown): boolean => {
    return Array.isArray(value) && value !== null && value !== undefined
}

export const isBoolean = (value: unknown): boolean => {
    return typeof value === 'boolean' && value !== null && value !== undefined
}

export const commaSeparators = (num: number | string): string => {
    const numStr = num.toString()
    const numParts = numStr.split('.')
    numParts[0] = numParts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ',')
    return numParts.join('.')
}

export const removeCommaSeparators = (num: string): number => {
    return parseInt(num.replace(/,/g, ''))
}

export const sanitizeNumberInput = (value: string) => {
    return value.replace(/,/g, '').trim()
}

export const isValidDecimalInput = (value: string) => {
    return (
        /^-?\d*\.?\d{0,2}$/.test(value) &&
        (value.match(/\./g)?.length ?? 0) <= 1
    )
}

export const formatNumberOnBlur = (
    value: string,
    onChange: (val: number | undefined) => void
) => {
    const sanitized = sanitizeNumberInput(value)
    if (!sanitized || sanitized === '-') {
        onChange(undefined)
        return
    }
    let parsedValue = parseFloat(sanitized)
    if (!isNaN(parsedValue)) {
        parsedValue = parseFloat(parsedValue.toFixed(2))
        onChange(parsedValue)
    }
}

export const downloadFile = async (url: string, fileName: string) => {
    try {
        const response = await axios.get(url, {
            responseType: 'blob',
        })

        const blob = new Blob([response.data])
        const blobUrl = URL.createObjectURL(blob)

        const a = document.createElement('a')
        a.href = blobUrl
        a.download = fileName || 'download'
        document.body.appendChild(a)
        a.click()
        document.body.removeChild(a)

        URL.revokeObjectURL(blobUrl)
    } catch (error) {
        console.error('Download failed:', error)
    }
}

export const capitalize = (text: string) =>
    text && String(text[0]).toUpperCase() + String(text).slice(1)
