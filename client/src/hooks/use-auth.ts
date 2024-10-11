import { useEffect, useState } from 'react'
import useAuthStore from '@/store/auth-store'
import { SignUpRequest } from '@/horizon-corp/types'
import UserService from '@/horizon-corp/server/auth/UserService'
import { handleAxiosError } from '@/lib/utils'

type TAccountType = 'Member' | 'Employee' | 'Admin' | 'Owner'

const useAuth = () => {
    const [loading, setLoading] = useState(false)
    const { currentUser, setCurrentUser } = useAuthStore()
    const [error, setError] = useState<string | null>(null)

    const loadCurrentUser = () => {
        // TODO: Request to /auth/current-user
        // - the server should return the current signed in user based on the request auth cookie
        // if the request doesnt have the auth cookie then return nothing, 403
        // - after getting the user data null or not, I'll set it to the currentUser
    }

    useEffect(() => {
        loadCurrentUser()
    }, [])

    const signUp = (signUpData: SignUpRequest) => {
        // TODO: Sign up and set currentUser
        // setCurrentUer(something)
        setError(null)
        setLoading(true)
        try {
            const createdUser = UserService.signUp()
            return createdUser
        } catch (e) {
            setError(handleAxiosError(e))
            return null
        } finally {
            setLoading(false)
        }
    }

    const signIn = () => {
        // TODO: Sign in and set currentUser
        // setCurrentUer(something)
    }

    const verifyEmail = (code: string | null) => {
        // if code is null, send email code to current user email /api/auth/sendCode email
        // if code has value, then verify current user email /api/auth/verify email
    }

    const verifyContact = (code: String | null) => {
        // if code is null, send email code to current user contact /api/auth/sendCode email
        // if code has value, then verify current user contact /api/auth/verify email
    }

    const forgotPassword = (email: string, accountType: TAccountType) => {
        // send a request /api/auth/forgot-password to send the user temporary password reset link to email
    }

    const passwordReset = (resetId: string) => {
        // send new password and invalidate the password reset link
        // request to /api/auth/password-reset/${resetId} idk not sure
    }

    return {
        loading,
        error,
        currentUser,
        signUp,
        signIn,
        verifyEmail,
        verifyContact,
        passwordReset,
        forgotPassword,
        loadCurrentUser,
    }
}

export default useAuth
