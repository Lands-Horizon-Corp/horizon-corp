export type TNotificationPriority = 'low' | 'medium' | 'high'

export type TNotification = 'success' | 'error' | 'warning' | 'info'

export interface INotificationResource {
    id: number
    type: TNotification
    message: string
    title?: string
    createdAt?: string
    actionUrl?: string
    priority?: TNotificationPriority
    duration?: number
    read: boolean
    userId?: number | null
    category?: string
}
