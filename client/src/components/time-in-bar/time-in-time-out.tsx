import { useState } from 'react'

import { Button } from '@/components/ui/button'
import LoadingCircle from '@/components/loader/loading-circle'
import TimeInCounter from '@/components/time-in-bar/time-in-counter'
import RealtimeTimeText from '@/components/time-in-bar/realtime-time-text'

import { cn } from '@/lib/utils'
import { TTImeInEntry } from '.'
import { UserBase } from '@/types'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {
    currentUser: UserBase
    timeEntry?: TTImeInEntry
    onTimeInEntry: (timeEntry: TTImeInEntry) => void
    onTimeOut: (timeEntry: TTImeInEntry) => void
}

const TimeInTimeOut = ({
    className,
    timeEntry,
    onTimeOut,
    onTimeInEntry,
}: Props) => {
    const [loading, setLoading] = useState(false)

    const handleClick = () => {
        setLoading(true)

        if (!timeEntry) {
            // todo create time in request
            return setTimeout(() => {
                onTimeInEntry({
                    timeStart: new Date('Mon, 07 Oct 2024 06:50:39 GMT'),
                })
            }, 1000)
        }

        // todo time out request
        return setTimeout(() => {
            onTimeOut({
                timeStart: new Date('Mon, 07 Oct 2024 06:50:39 GMT'),
                timeEnd: new Date(),
            })
        }, 1000)
    }

    return (
        <div className={cn('space-y-4 p-4', className)}>
            <div className="mx-auto size-52 rounded-2xl bg-background"></div>
            <div className="flex flex-col items-center gap-y-6">
                <p className="text-sm italic text-foreground/60">Quotes Desu</p>
                {timeEntry && (
                    <div className="flex w-full items-center justify-between gap-x-2 rounded-xl bg-background p-2">
                        <span className="text-foreground/70">Work hours</span>{' '}
                        <TimeInCounter
                            className="text-base"
                            timeEntry={timeEntry}
                        />
                    </div>
                )}
                <span className="text-left text-foreground/60">
                    <span>Current : </span>
                    <RealtimeTimeText />
                </span>
                <Button
                    size="sm"
                    disabled={loading}
                    className="w-full"
                    onClick={handleClick}
                    variant={!timeEntry ? 'default' : 'secondary'}
                >
                    {loading ? (
                        <LoadingCircle />
                    ) : timeEntry ? (
                        'Timeout'
                    ) : (
                        'Capture'
                    )}
                </Button>
            </div>
        </div>
    )
}

export default TimeInTimeOut
