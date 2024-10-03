import React from 'react'

export const concatParentUrl = ({
    baseUrl = '',
    url = '',
}: {
    baseUrl?: string
    url?: string
}) => {
    return `${baseUrl}${url}`
}

// TOFIX
export const sidebarCollapsableRouteMatcher = (
    selfUrl: string = '',
    toCompareUrl: string
) => {
    return toCompareUrl.startsWith(selfUrl) && toCompareUrl !== selfUrl
}

// TOFIX
export const sidebarItemRouteMatcher = (
    selfUrl: string = '',
    toCompareUrl: string
) => {
    return toCompareUrl.startsWith(selfUrl)
}

export function isReactNode(node: any): boolean {
    return (
        React.isValidElement(node) || // Check if it's a React element
        typeof node === 'string' || // Check if it's a string
        typeof node === 'number' || // Check if it's a number
        Array.isArray(node) || // Check if it's an array (could be an array of React nodes)
        node === null || // Allow null
        node === undefined // Allow undefined
    )
}
