import { serverRequestErrExtractor } from "@/helpers";
import { NotificationsResource } from "@/horizon-corp/types";
import { Role } from "@/types";
import { withCatchAsync } from "@/utils";
import { useMutation, UseMutationResult, useQuery, useQueryClient, UseQueryResult } from "@tanstack/react-query";
import { toast } from "sonner";
import { IOperationCallbacks } from "./types";

interface INotificationsProps {
    userId: number;
    role: Role;
}

const handleError = (error: unknown, onError?: (message: string) => void) => {
    const errorMessage = serverRequestErrExtractor({ error });
    toast.error(errorMessage);
    onError?.(errorMessage);
    throw new Error(errorMessage);
};

const handleSuccess = <TDataSuccess>(
    message: string,
    onSuccess?: (data: TDataSuccess) => void,
    data?: TDataSuccess
) => {
    toast.success(message);
    onSuccess?.(data as TDataSuccess);
};

// Load/get notifications
export const useNotifications = ({ userId, role }: INotificationsProps): UseQueryResult<NotificationsResource[], string> => {
    return useQuery<NotificationsResource[], string>({
        queryKey: ['admin-notifications', userId, role],
        queryFn: async () => {
              // Replace with actual API call
            const [error, result] = await withCatchAsync(
                new Promise<NotificationsResource[]>((resolve) => {
                    resolve([
                        {
                            id: 101,
                            type: 'error',
                            message: 'Database connection failed on Server #3.',
                            title: 'Critical System Alert',
                            createdAt: '2025-01-12T09:45:00Z',
                            actionUrl: '/admin/system/logs/101',
                            read: false,
                            priority: 'high',
                            userId: null,
                            category: 'System',
                        },
                        {
                            id: 102,
                            type: 'warning',
                            message: 'There are 5 pending user account approvals.',
                            title: 'Pending Approvals',
                            createdAt: '2025-01-12T08:30:00Z',
                            actionUrl: '/admin/users/pending-approvals',
                            read: false,
                            priority: 'medium',
                            userId: null,
                            category: 'User Management',
                        },
                        {
                            id: 103,
                            type: 'success',
                            message: 'Backup process completed successfully.',
                            title: 'Backup Complete',
                            createdAt: '2025-01-12T07:15:00Z',
                            actionUrl: '/admin/system/backup-logs',
                            read: true,
                            priority: 'low',
                            userId: null,
                            category: 'System',
                        },
                        {
                            id: 104,
                            type: 'info',
                            message: 'The weekly analytics report is ready for review.',
                            title: 'Analytics Report',
                            createdAt: '2025-01-11T22:00:00Z',
                            actionUrl: '/admin/reports/weekly',
                            read: false,
                            priority: 'medium',
                            userId: null,
                            category: 'Analytics',
                        },
                        {
                            id: 105,
                            type: 'error',
                            message: 'Payment processing failed for Invoice #5678.',
                            title: 'Payment Failure',
                            createdAt: '2025-01-12T05:00:00Z',
                            actionUrl: '/admin/payments/failure/5678',
                            read: false,
                            priority: 'high',
                            userId: 42,
                            category: 'Finance',
                        },
                        {
                            id: 106,
                            type: 'success',
                            message: "New user 'John Doe' has registered.",
                            title: 'New Registration',
                            createdAt: '2025-01-12T06:00:00Z',
                            actionUrl: '/admin/users/42',
                            read: true,
                            priority: 'low',
                            userId: 42,
                            category: 'User Management',
                        },
                        {
                            id: 107,
                            type: 'info',
                            message:
                                'Version 2.5.0 of the application has been deployed successfully.',
                            title: 'Deployment Success',
                            createdAt: '2025-01-12T03:30:00Z',
                            actionUrl: '/admin/deployments/2.5.0',
                            read: true,
                            priority: 'low',
                            userId: null,
                            category: 'System',
                        },
                        {
                            id: 108,
                            type: 'warning',
                            message: 'The API usage limit is at 85%.',
                            title: 'API Usage Warning',
                            createdAt: '2025-01-11T21:15:00Z',
                            actionUrl: '/admin/api/usage',
                            read: false,
                            priority: 'medium',
                            userId: null,
                            category: 'System',
                        },
                    ]);
                })
            );

            if (error) {
                throw new Error(serverRequestErrExtractor({ error }));
            }

            return result;
        },
    });
};

// Mark as read
export const useMarkasReadNotifications = ({
    onSuccess,
    onError,
}: IOperationCallbacks<NotificationsResource, string>) => {
    const queryClient = useQueryClient();

    return useMutation<void, string, number>({
        mutationKey: ["notifications", "mark-as-read"],
        mutationFn: async () => {
            // Replace with actual API call
            const [error] = await withCatchAsync(
                new Promise<void>((resolve) => {
                    resolve();
                })
            );

            if (error) {
                throw new Error("Failed to mark notification as read");
            }
        },
        onSuccess: (_, notificationId) => {
            queryClient.invalidateQueries({ queryKey: ["notifications"] });
            handleSuccess(`Notification #${notificationId} marked as read`, onSuccess);
        },
        onError: (error) => {
            handleError(error, onError);
        },
    });
};

// Delete notification
export const useDeleteNotifications = ({
    onSuccess,
    onError,
}: IOperationCallbacks): UseMutationResult<void, string, number> => {
    const queryClient = useQueryClient();

    return useMutation<void, string, number>({
        mutationKey: ["notifications", "delete"],
        mutationFn: async () => {
            // Replace with actual API call
            const [error] = await withCatchAsync(
                new Promise<void>((resolve) => {
                    resolve();
                })
            );

            if (error) {
                throw new Error("Failed to delete notification");
            }
        },
        onSuccess: (_, notificationId) => {
            queryClient.invalidateQueries({ queryKey: ["notifications"] });
            handleSuccess(`Notification #${notificationId} deleted`, onSuccess);
        },
        onError: (error) => {
            handleError(error, onError);
        },
    });
};