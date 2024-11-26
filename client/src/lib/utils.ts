import { twMerge } from 'tailwind-merge'
import { clsx, type ClassValue } from 'clsx'

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
}

export const withCatchAsync = async <
    T,
    E extends new (message?: string) => Error,
>(
    promise: Promise<T>,
    knownErrors?: E[]
): Promise<[undefined, T] | [InstanceType<E>]> => {
    return promise
        .then((data) => {
            return [undefined, data] as [undefined, T]
        })
        .catch((err) => {
            if (knownErrors == undefined) return [err]

            if (knownErrors.some((e) => err instanceof e)) return [err]

            throw err
        })
}

export const formatNumber = (
    value: number,
    minimumFractionDigits = 0,
    maximumFractionDigits = 2
) => {
    if (isNaN(value)) return '...'
    return value.toLocaleString('en-US', {
        useGrouping: true,
        minimumFractionDigits,
        maximumFractionDigits,
    })
}
