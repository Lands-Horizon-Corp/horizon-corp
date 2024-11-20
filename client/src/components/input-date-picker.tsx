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
import { CaptionLayout } from 'react-day-picker'

type InputDatePickerProps = {
    id: string
    value: Date | undefined
    disabled?: (date: Date) => boolean
    captionLayout?: CaptionLayout
    fromYear?: number
    toYear?: number
    onChange: (date: Date | undefined) => void
    className?: string
}

const InputDatePicker = ({
    id,
    value,
    onChange,
    disabled,
    fromYear = 1930,
    toYear = new Date().getFullYear(),
    ...other
}: InputDatePickerProps) => {
    return (
        <Popover>
            <PopoverTrigger asChild>
                <Button
                    id={id}
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
                <Calendar
                    {...other}
                    mode="single"
                    toYear={toYear}
                    showOutsideDays
                    selected={value}
                    fromYear={fromYear}
                    onSelect={onChange}
                    disabled={disabled}
                />
            </PopoverContent>
        </Popover>
    )
}

export default InputDatePicker
