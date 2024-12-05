import { format } from 'date-fns'
import { CalendarIcon } from 'lucide-react'
import { CaptionLayout, type DateRange } from 'react-day-picker'

import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/calendar'

import { cn } from '@/lib/utils'
import { isDate, isObject } from '@/helpers'

type DateRangePicker = {
    toYear?: number
    value: DateRange
    fromYear?: number
    modal?: boolean
    captionLayout?: CaptionLayout
    disabled?: (date: Date) => boolean
    onChange: (range: { from: Date; to?: Date }) => void
}

const isRange = (value: unknown): value is Range => {
    if (!isObject(value)) {
        return false
    }

    const { from, to } = value as { from?: unknown; to?: unknown }

    return (
        (from === undefined || isDate(from)) && (to === undefined || isDate(to))
    )
}

const DateRange = ({
    value,
    modal = false,
    onChange,
    disabled,
    fromYear = 1930,
    toYear = new Date().getFullYear(),
    ...other
}: DateRangePicker) => {
    return (
        <Popover modal={modal}>
            <PopoverTrigger asChild>
                <Button
                    variant={'outline'}
                    className={cn(
                        'w-full pl-3 text-left font-normal',
                        !value && 'text-muted-foreground'
                    )}
                >
                    {value?.from
                        ? format(value.from, 'MMM dd yyyy')
                        : 'choose from'}{' '}
                    {' - '}
                    {value?.to
                        ? format(value.to, 'MMM dd yyyy')
                        : 'choose to'}{' '}
                    <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent
                className="w-auto rounded-2xl bg-popover/85 p-0 backdrop-blur"
                align="start"
            >
                <Calendar
                    {...other}
                    mode="range"
                    toYear={toYear}
                    showOutsideDays
                    selected={value}
                    fromYear={fromYear}
                    onSelect={(range) => {
                        if (!range?.from || !isRange(range)) return
                        onChange({ from: range.from, to: range.to })
                    }}
                    disabled={disabled}
                />
            </PopoverContent>
        </Popover>
    )
}

export default DateRange
