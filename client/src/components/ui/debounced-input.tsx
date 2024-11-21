import * as React from 'react'
import { Input, InputProps } from './input'
import useDebounce from '@/hooks/use-debounce'

export interface DebouncedInputProps<
    T extends string | number | readonly string[] | undefined,
> extends Omit<InputProps, 'value' | 'onChange'> {
    value: T
    onChange: (value: T) => void
    debounceTime?: number
}

const DebouncedInput = React.forwardRef<
    HTMLInputElement,
    DebouncedInputProps<any>
>(
    <T extends string | number | readonly string[] | undefined>(
        {
            value,
            onChange,
            debounceTime = 300,
            ...props
        }: DebouncedInputProps<T>,
        ref: React.Ref<HTMLInputElement>
    ) => {
        const [localValue, setLocalValue] = React.useState<T>(value)

        const debouncedValue = useDebounce(localValue, debounceTime)

        React.useEffect(() => {
            onChange(debouncedValue)
        }, [debouncedValue, onChange])

        return (
            <Input
                ref={ref}
                value={localValue}
                onChange={(e) => setLocalValue(e.target.value as T)}
                {...props}
            />
        )
    }
)

DebouncedInput.displayName = 'DebouncedInput'

export { DebouncedInput }
