import { useState, useEffect } from 'react'

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from '@/components/ui/dialog'
import TimeInCounter from './time-in-counter'
import UserAvatar from '@/components/user-avatar'
import LoadingCircle from '../loader/loading-circle'

import { cn } from '@/lib/utils'
import { UserBase } from '@/types'
import ActionTooltip from '../action-tooltip'

import { IBaseCompNoChild } from '@/types/component/base'
import TimeInTimeOut from './time-in-time-out'
import { Button } from '../ui/button'

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
                    className="w-full max-w-sm rounded-xl bg-gradient-to-b from-popover to-background p-0 sm:rounded-xl"
                >
                    <DialogHeader className="hidden">
                        <DialogTitle>Time In Out Form</DialogTitle>
                        <DialogDescription>Shows time in out</DialogDescription>
                    </DialogHeader>
                    <TimeInTimeOut
                        timeEntry={timeInEntry}
                        currentUser={currentUser}
                        onTimeOut={() => {
                            setTimeInEntry(undefined)
                            setShowTimeInOut(false)
                        }}
                        onTimeInEntry={(data) => {
                            setTimeInEntry(data)
                            setShowTimeInOut(false)
                        }}
                    />
                </DialogContent>
            </Dialog>
        </>
    )
}

export default TimeInBar
