import {
    useQuery,
    useMutation,
    queryOptions,
    useQueryClient,
} from '@tanstack/react-query'
import { toast } from 'sonner'

import { withCatchAsync } from '@/utils'
import { serverRequestErrExtractor } from '@/helpers'
import AuthService from '@/server/api-service/auth-service'

import {
    IUserData,
    ISignInRequest,
    ISignUpRequest,
    IChangePasswordRequest,
    IForgotPasswordRequest,
} from '@/server/types'
import { IOperationCallbacks } from './types'

// Loader for Auth Preloading
export const authLoader = () =>
    queryOptions<IUserData>({
        queryKey: ['auth', 'current-user'],
        queryFn: async () => {
            return await AuthService.currentUser()
        },
        retry: 0,
    })

// Hook: Get Current User
export const useCurrentUser = ({
    onError,
    onSuccess,
}: IOperationCallbacks<IUserData>) => {
    return useQuery<IUserData, string>({
        queryKey: ['auth', 'current-user'],
        queryFn: async () => {
            const [error, data] = await withCatchAsync(
                AuthService.currentUser()
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            onSuccess?.(data)
            return data
        },
        retry: 1,
    })
}

// Hook: Sign In
export const useSignIn = ({
    onError,
    onSuccess,
}: IOperationCallbacks<IUserData>) => {
    const queryClient = useQueryClient()

    return useMutation<IUserData, string, ISignInRequest>({
        mutationKey: ['auth', 'signin'],
        mutationFn: async (credentials) => {
            const [error, data] = await withCatchAsync(
                AuthService.signIn(credentials)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            toast.success('Signed in successfully')
            onSuccess?.(data)

            // Optionally prefetch or update user data
            queryClient.setQueryData(['auth', 'current-user'], data)
            return data
        },
    })
}

// Hook: Sign Up
export const useSignUp = ({
    onError,
    onSuccess,
}: IOperationCallbacks<IUserData, string> | undefined = {}) => {
    return useMutation<IUserData, string, ISignUpRequest>({
        mutationKey: ['auth', 'signup'],
        mutationFn: async (newUser) => {
            const [error, data] = await withCatchAsync(
                AuthService.signUp(newUser)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            toast.success('Signed up successfully')
            onSuccess?.(data)
            return data
        },
    })
}

// Hook: Forgot Password
export const useForgotPassword = ({
    onError,
    onSuccess,
}: IOperationCallbacks<void>) => {
    return useMutation<void, string, IForgotPasswordRequest>({
        mutationKey: ['auth', 'forgot-password'],
        mutationFn: async (data) => {
            const [error] = await withCatchAsync(
                AuthService.forgotPassword(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            toast.success('Password reset link sent')
            onSuccess?.()
        },
    })
}

// Hook: Change Password
export const useChangePassword = ({
    onError,
    onSuccess,
}: IOperationCallbacks<void>) => {
    return useMutation<void, string, IChangePasswordRequest>({
        mutationKey: ['auth', 'change-password'],
        mutationFn: async (data) => {
            const [error] = await withCatchAsync(
                AuthService.changePassword(data)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            toast.success('Password changed successfully')
            onSuccess?.()
        },
    })
}

// Hook: Sign Out
export const useSignOut = ({
    onError,
    onSuccess,
}: IOperationCallbacks<void>) => {
    const queryClient = useQueryClient()

    return useMutation<void, string>({
        mutationKey: ['auth', 'signout'],
        mutationFn: async () => {
            const [error] = await withCatchAsync(AuthService.signOut())

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                onError?.(errorMessage)
                throw errorMessage
            }

            toast.success('Signed out successfully')
            onSuccess?.()

            // Invalidate cached data
            queryClient.invalidateQueries({
                queryKey: ['auth', 'current-user'],
            })
        },
    })
}
