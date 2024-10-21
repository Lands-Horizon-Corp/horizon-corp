import { useState } from 'react'

import EcoopLogo from '@/components/ecoop-logo'
import StepIndicator from '@/components/steps-indicator'
import VerifyForm from '@/modules/auth/components/forms/verify-form'

import { IBaseComp } from '@/types/component'
import { UserData } from '@/horizon-corp/types'

interface Props extends IBaseComp {
    readOnly?: boolean

    userData: UserData
    onVerifyChange?: (newUserData: UserData) => void
    onVerifyComplete?: (newUserData: UserData) => void
}

type TStep = 'mobile' | 'email'

const countCompleted = (userData: UserData) => {
    let completed = 0

    if (userData.isEmailVerified) completed++
    if (userData.isContactVerified) completed++

    return completed
}

const VerifyRoot = ({
    userData,
    readOnly = false,
    onVerifyChange,
    onVerifyComplete,
}: Props) => {
    const [skipped, setSkipped] = useState<TStep[]>([])
    const completedCount = countCompleted(userData)

    if(completedCount === 2) {
        onVerifyComplete?.(userData)
        return
    }

    if(skipped.length + completedCount >= 2) {
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
            {!userData.isContactVerified && !skipped.includes("mobile") && (
                <VerifyForm
                    key="1"
                    readOnly={readOnly}
                    verifyMode="mobile"
                    onSuccess={(data) => onVerifyChange?.(data)}
                    onSkip={() => setSkipped((val) => [...val, 'mobile'])}
                />
            )}
            {!userData.isEmailVerified && ( userData.isContactVerified || skipped.includes("mobile")) && !skipped.includes("email") && (
                <VerifyForm
                    key="2"
                    verifyMode="email"
                    readOnly={readOnly}
                    onSuccess={(data) => {
                        onVerifyComplete?.(data)
                    }}
                    onSkip={() => setSkipped((val) => [...val, 'email'])}
                />
            )}
        </div>
    )
}

export default VerifyRoot
