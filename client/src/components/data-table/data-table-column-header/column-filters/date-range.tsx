import { format } from 'date-fns'
import { CalendarIcon } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/calendar'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { cn } from '@/lib/utils'
import { CaptionLayout, type DateRange } from 'react-day-picker'

type DateRangePicker = {
    value: DateRange
    disabled?: (date: Date) => boolean
    captionLayout?: CaptionLayout
    fromYear?: number
    toYear?: number
    onChange: (range: { from: Date; to?: Date }) => void
}

const DateRange = ({
    value,
    onChange,
    disabled,
    fromYear = 1930,
    toYear = new Date().getFullYear(),
    ...other
}: DateRangePicker) => {
    return (
        <Popover>
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
                        if (
                            !range ||
                            !(range.from instanceof Date) ||
                            (range.to !== undefined &&
                                !(range.to instanceof Date))
                        )
                            return

                        onChange({ from: range.from, to: range.to })
                    }}
                    disabled={disabled}
                />
            </PopoverContent>
        </Popover>
    )
}

export default DateRange
