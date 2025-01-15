import React from 'react'
import { CheckIcon, ErrorIcon, WarningIcon } from '../icons'
import { NotificationType } from '@/horizon-corp/types'
import { cn } from '@/lib'

interface NotificationTypeProps {
    type: NotificationType
    message?: string
    icon?: React.ReactNode
    IconClassName?: string
    className?: string
}

const NotificationTypeComponent = ({
    type,
    message,
    icon,
    IconClassName,
    className,
}: NotificationTypeProps) => {
    // Default icons for each notification type
    const defaultIcons: Record<NotificationType, React.ReactNode> = {
        success: (
            <span className="text-primary">
                <CheckIcon />
            </span>
        ),
        error: (
            <span className="text-destructive">
                <ErrorIcon />
            </span>
        ),
        warning: (
            <span className="text-yellow-500">
                <WarningIcon />
            </span>
        ),
        info: (
            <span className="text-blue-400">
                <WarningIcon />
            </span>
        ),
    }

    return (
        <div className={cn(`flex items-center`, className ?? className)}>
            <span className={cn(IconClassName ?? IconClassName)}>
                {icon || defaultIcons[type]}
            </span>
            <div className="message">{message}</div>
        </div>
    )
}

export default NotificationTypeComponent
