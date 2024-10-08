import { forwardRef, useEffect, useState } from 'react'

import { getTimeDifference } from './utils'
import { TTImeInEntry } from './time-in-bar'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    timeEntry: TTImeInEntry
    interval?: number
}

const TimeInCounter = forwardRef<HTMLSpanElement, Props>(
    ({ timeEntry, interval = 60_000, className, ...other }, ref) => {
        const [totalWorkHours, setTotalWorkHours] = useState('--')

        const updateTotalWorkHours = (currentDate: Date) => {
            setTotalWorkHours(
                getTimeDifference(timeEntry.timeStart, currentDate)
            )
        }

        useEffect(() => {
            const timeCounter = setInterval(() => {
                const currentDate = new Date()
                updateTotalWorkHours(currentDate)
            }, interval)

            return () => {
                clearInterval(timeCounter)
            }
        }, [timeEntry])

        useEffect(() => {
            updateTotalWorkHours(new Date())
        }, [])

        return (
            <span className={className} {...other} ref={ref}>
                {totalWorkHours}
            </span>
        )
    }
)

TimeInCounter.displayName = 'TimeInCounter'

export default TimeInCounter
