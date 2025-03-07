export const sidebarRouteMatcher = (
    selfUrl: string | undefined,
    toCompareUrl: string
) => {
    return toCompareUrl === selfUrl || toCompareUrl.startsWith(selfUrl + '/')
}
