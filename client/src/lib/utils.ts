import { twMerge } from 'tailwind-merge'
import { clsx, type ClassValue } from 'clsx'

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

/**
 * A utility function designed for inline async operations that may fail,
 * returning an array `[error, result]` to simplify error handling in single requests.
 *
 * This function is useful for reducing verbosity in scenarios where you are making
 * a single asynchronous request. It avoids the need for a try-catch block and
 * handles known error types gracefully.
 *
 * @async
 * @template T - The type of the expected successful result.
 * @template {new (message?: string) => Error} E - The type of known error constructors.
 * @param {Promise<T>} promise - The promise representing the async operation to execute.
 * @param {?E[]} [knownErrors] - Optional array of known error constructors. If the error
 * matches any of these types, it is returned in the `[error, undefined]` tuple.
 * @returns {Promise<[undefined, T] | [InstanceType<E>, undefined]>} - A promise resolving
 * to a tuple where the first element is the error (if any) and the second element is the
 * result of the operation. If successful, `[undefined, result]` is returned.
 * If an error occurs, `[error, undefined]` is returned.
 */
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
