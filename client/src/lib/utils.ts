import { twMerge } from 'tailwind-merge'
import { clsx, type ClassValue } from 'clsx'

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
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