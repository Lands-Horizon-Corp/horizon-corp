import axios from 'axios'
import { twMerge } from 'tailwind-merge'
import { clsx, type ClassValue } from 'clsx'
import { ErrorDetails } from '@/horizon-corp/types'

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

type SimpleTypes = boolean | number | string | null | undefined | symbol

export default function getSimpleProperties<T extends object>(
    objectToParse: T,
    include?: object | Array<keyof T> | boolean
): { [key in keyof T]: SimpleTypes } {
    const simpleProperties: Partial<{ [key in keyof T]: SimpleTypes }> = {}

    if (include === false)
        return simpleProperties as { [key in keyof T]: SimpleTypes }

    for (const [key, value] of Object.entries(objectToParse)) {
        const isSimpleType =
            value === null ||
            value === undefined ||
            typeof value === 'boolean' ||
            typeof value === 'number' ||
            typeof value === 'string' ||
            typeof value === 'symbol'

        if (
            isSimpleType ||
            (Array.isArray(include) && include.includes(key as keyof T))
        ) {
            simpleProperties[key as keyof T] = value as SimpleTypes
        }
    }

    return simpleProperties as { [key in keyof T]: SimpleTypes }
}

export const handleAxiosError = (error: unknown): string => {
    if (axios.isAxiosError(error)) {
        if (!error.response) {
            if (error.message === 'Network Error') {
                return 'Network error: Please check your internet connection and try again.'
            } else {
                return `Error: ${error.message}`
            }
        }

        const data = error.response.data as ErrorDetails

        return data.message
    } else {
        return 'An unexpected error occurred. Please try again.'
    }
}
