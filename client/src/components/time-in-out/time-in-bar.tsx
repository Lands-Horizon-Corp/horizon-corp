import { useState, useEffect, useMemo } from 'react'

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'
import TimeInCounter from './time-in-counter'
import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import TimeInTimeOut from '@/components/time-in-out'
import ActionTooltip from '@/components/action-tooltip'
import LoadingCircle from '@/components/loader/loading-circle'

import {
    randomEndOfDayQuoute,
    randomStartOfDayQuoute,
} from '@/constants/quoutes'
import { cn } from '@/lib/utils'
import { UserBase } from '@/types'
import { IBaseCompNoChild } from '@/types/component/base'

interface Props extends IBaseCompNoChild {
    currentUser: UserBase | null
}

export type TTImeInEntry = {
    timeStart: Date
    timeEnd?: Date
}

const TimeInBar = ({ className, currentUser }: Props) => {
    const [loading, setLoading] = useState(true)
    const [showTimeInOut, setShowTimeInOut] = useState(false)
    const [timeInEntry, setTimeInEntry] = useState<TTImeInEntry | undefined>(
        undefined
    )

    const quote = useMemo(() => {
        return !timeInEntry ? randomStartOfDayQuoute() : randomEndOfDayQuoute()
    }, [timeInEntry])

    useEffect(() => {
        setTimeout(() => setLoading(false), 1000)
    }, [])

    const handleClick = () => {
        setShowTimeInOut(true)
    }

    if (!currentUser) return

    return (
        <>
            <ActionTooltip
                side="bottom"
                align="center"
                tooltipContent={timeInEntry ? 'Time out' : 'Time in now'}
            >
                <Button
                    onClick={handleClick}
                    variant="secondary"
                    size={loading ? 'icon' : 'default'}
                    className={cn(
                        'pointer-events-none relative flex cursor-not-allowed items-center gap-x-2 rounded-full border-2 border-secondary bg-popover p-1.5 text-sm text-foreground/70 delay-150 duration-300 ease-in-out',
                        className,
                        timeInEntry && 'border-primary/40',
                        !loading && 'pointer-events-auto cursor-pointer pr-3'
                    )}
                >
                    {loading && <LoadingCircle />}
                    {!loading && (
                        <>
                            <UserAvatar
                                src="https://avatars.githubusercontent.com/u/48374007?s=80&v=4"
                                fallback="JX"
                            />
                            {!timeInEntry && <span>Time-in</span>}
                            {timeInEntry && (
                                <span>
                                    Onshift :{' '}
                                    <TimeInCounter
                                        className="font-medium"
                                        timeEntry={timeInEntry}
                                    />
                                </span>
                            )}
                        </>
                    )}
                </Button>
            </ActionTooltip>
            <Dialog
                open={showTimeInOut}
                onOpenChange={(state) => setShowTimeInOut(state)}
            >
                <DialogContent
                    hideCloseButton
                    overlayClassName="backdrop-blur-sm"
                    className="w-full max-w-sm overflow-hidden !rounded-3xl border bg-popover p-0 shadow-center-md"
                >
                    <DialogHeader className="hidden">
                        <DialogTitle>Time In Out Form</DialogTitle>
                        <DialogDescription>Shows time in out</DialogDescription>
                    </DialogHeader>
                    <div className="ecoop-scroll max-h-screen overflow-y-scroll [&::-webkit-scrollbar]:size-0">
                        <TimeInTimeOut
                            timeEntry={timeInEntry}
                            currentUser={currentUser}
                            message={quote}
                            onTimeOut={() => {
                                setTimeInEntry(undefined)
                                setShowTimeInOut(false)
                            }}
                            onTimeInEntry={(data) => {
                                setTimeInEntry(data)
                                setShowTimeInOut(false)
                            }}
                        />
                    </div>
                </DialogContent>
            </Dialog>
        </>
    )
}

export default TimeInBar
