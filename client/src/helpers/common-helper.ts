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

export const getFileType = (file: File): FileCategory => {
    const fileExtension = file.name.split('.').pop()?.toLowerCase()
    const mimeType = file.type

    const fileTypes: Partial<Record<FileCategory, string[]>> = {
        audio: ['mp3', 'wav', 'ogg', 'aac', 'flac', 'm4a'],
        video: ['mp4', 'mkv', 'avi', 'mov', 'wmv', 'flv'],
        doc: ['doc', 'docx', 'odt', 'rtf', 'wps'],
        pdf: ['pdf'],
        sheet: ['xls', 'xlsx', 'csv', 'ods'],
        text: ['txt', 'log', 'md'],
        image: ['jpg', 'jpeg', 'png', 'gif', 'bmp', 'svg', 'webp'],
    }

    // Check each category
    for (const [category, extensions] of Object.entries(fileTypes)) {
        if (fileExtension && extensions?.includes(fileExtension)) {
            return category as FileCategory
        }
    }

    // Alternatively, check for some specific MIME types
    if (mimeType.startsWith('audio/')) return 'audio'
    if (mimeType.startsWith('video/')) return 'video'
    if (mimeType === 'application/pdf') return 'pdf'
    if (mimeType.startsWith('text/')) return 'text'
    if (mimeType.startsWith('image/')) return 'image'

    return 'unknown'
}
