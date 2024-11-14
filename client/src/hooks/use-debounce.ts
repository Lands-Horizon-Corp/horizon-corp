import { useState, useEffect } from 'react'

function useDebounce<T>(value: T, delay: number): T {
    const [debouncedValue, setDebouncedValue] = useState<T>(value)

    useEffect(() => {
        const handler = setTimeout(() => {
            setDebouncedValue(value)
        }, delay)

        return () => {
            console.log("Clearing timeout for value:", value);
            clearTimeout(handler)
        }
    }, [value, delay])

    return debouncedValue
}

export default useDebounce
