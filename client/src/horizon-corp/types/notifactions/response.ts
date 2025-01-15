export type NotificationPriority = "low" | "medium" | "high";

export type NotificationType = "success" | "error" | "warning" | "info";


export interface NotificationsResource {
    id: number;
    type: NotificationType;
    message: string;
    title?: string;
    createdAt?: string;
    actionUrl?: string;
    priority?: NotificationPriority;
    duration?: number;
    read: boolean;
    userId?: number | null;
    category?: string;
}
