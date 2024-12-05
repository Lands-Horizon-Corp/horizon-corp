import { format } from 'date-fns'
import { CalendarIcon } from 'lucide-react'
import { CaptionLayout } from 'react-day-picker'

import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { Button } from '@/components/ui/button'
import DateTimePicker from './date-time-picker'

import { cn } from '@/lib/utils'

type InputDatePickerProps = {
    fromYear?: number
    toYear?: number
    className?: string
    value: Date | undefined
    captionLayout?: CaptionLayout
    disabled?: (date: Date) => boolean
    onChange: (date: Date | undefined) => void
}

const InputDatePicker = ({
    value,
    onChange,
    disabled,
    fromYear = 1900,
    toYear = new Date().getFullYear(),
    ...other
}: InputDatePickerProps) => {
    return (
        <Popover modal>
            <PopoverTrigger asChild>
                <Button
                    variant={'outline'}
                    className={cn(
                        'w-full pl-3 text-left font-normal',
                        !value && 'text-muted-foreground'
                    )}
                >
                    {value ? format(value, 'PPP') : <span>Pick a date</span>}
                    <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent
                className="w-auto rounded-2xl bg-popover/85 p-0 backdrop-blur"
                align="start"
            >
                <DateTimePicker
                    {...other}
                    toYear={toYear}
                    value={value}
                    fromYear={fromYear}
                    onChange={onChange}
                    disabled={disabled}
                />
            </PopoverContent>
        </Popover>
    )
}

export default InputDatePicker
