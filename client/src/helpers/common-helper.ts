export const randomChoose = <T>(data: Array<T>) => {
    return data[~~(Math.random() * data.length)]
}

export const formatBytes = (bytes:number) => {
    if (bytes === 0) return '0 Bytes';
    const sizeUnits = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    const formattedSize = (bytes / Math.pow(1024, i)).toFixed(2);
    return `${formattedSize} ${sizeUnits[i]}`;
};

export const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString();
};