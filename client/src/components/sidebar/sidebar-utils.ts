export const concatParentUrl = ({
    baseUrl = '',
    url = '',
}: {
    baseUrl?: string
    url?: string
}) => {
    return `${baseUrl}${url}`
}

export const sidebarRouteMatcher = (
    selfUrl: string | undefined,
    toCompareUrl: string
) => {
    return toCompareUrl === selfUrl || toCompareUrl.startsWith(selfUrl + '/')
}
