import { toast } from 'sonner'
import { useState } from 'react'

import EcoopLogo from '@/components/ecoop-logo'
import StepIndicator from '@/components/steps-indicator'
import LoadingSpinner from '@/components/spinners/loading-spinner'
import VerifyForm from '@/modules/auth/components/forms/verify-form'

import { withCatchAsync } from "@/utils"
import { IBaseComp } from '@/types/component'
import { UserData } from '@/horizon-corp/types'
import { useMutation } from '@tanstack/react-query'
import { serverRequestErrExtractor } from '@/helpers'
import { useUserAuthStore } from '@/store/user-auth-store'
import UserService from '@/horizon-corp/services/auth/UserService'

interface Props extends IBaseComp {
    readOnly?: boolean
    userData: UserData

    onSkip: () => void
    onVerifyChange?: (newUserData: UserData) => void
    onVerifyComplete?: (newUserData: UserData) => void
}

const countCompleted = (userData: UserData) => {
    let completed = 0

    if (userData.isEmailVerified) completed++
    if (userData.isContactVerified) completed++

    return completed
}

const VerifyRoot = ({
    userData,
    readOnly = false,
    onSkip,
    onVerifyChange,
    onVerifyComplete,
}: Props) => {
    const { setCurrentUser } = useUserAuthStore()
    const [userStoredData, setStoredUserData] = useState(userData)

    const completedCount = countCompleted(userStoredData)

    const { mutate: handleSkip, isPending: isSkipping } = useMutation<
        UserData,
        string
    >({
        mutationKey: ['skip-verification'],
        mutationFn: async () => {
            const [error, response] = await withCatchAsync(
                UserService.SkipVerification()
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            onSkip()
            setCurrentUser(response.data)

            return response.data
        },
    })

    if (completedCount === 2) {
        onVerifyComplete?.(userData)
        return
    }

    return (
        <div>
            <div className="my-5 flex w-full items-center gap-x-4">
                <EcoopLogo className="size-10" />
                <StepIndicator
                    totalSteps={2}
                    className="w-full"
                    currentStep={completedCount + 1}
                />
            </div>
            {!userStoredData.isContactVerified && !isSkipping && (
                <VerifyForm
                    key="1"
                    readOnly={readOnly}
                    verifyMode="mobile"
                    onSuccess={(data) => {
                        onVerifyChange?.(data)
                        setStoredUserData(data)
                    }}
                    onSkip={handleSkip}
                />
            )}
            {!userStoredData.isEmailVerified &&
                userStoredData.isContactVerified &&
                !isSkipping && (
                    <VerifyForm
                        key="2"
                        verifyMode="email"
                        readOnly={readOnly}
                        onSuccess={(data) => {
                            onVerifyChange?.(data)
                            setStoredUserData(data)
                        }}
                        onSkip={handleSkip}
                    />
                )}
            {isSkipping && <LoadingSpinner className="mx-auto" />}
        </div>
    )
}

export default VerifyRoot
