import { useEffect, useState } from 'react'

import EcoopLogo from '@/components/ecoop-logo'
import StepIndicator from '@/components/steps-indicator'
import VerifyForm from '@/modules/auth/components/forms/verify-form'

import { UserBase, UserStatus } from '@/types'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {
    readOnly?: boolean
    userData: UserBase // TODO: Change based on auth response resource fromn @/horizon-corp/types/auth/...
    onVerifyComplete?: (newUserData: UserBase) => void
    onVerifyChange?: (newUserData: UserBase) => void
}

const VerifyRoot = ({
    readOnly = false,
    userData,
    onVerifyChange,
    onVerifyComplete,
}: Props) => {
    const [step, setStep] = useState(1)

    useEffect(() => {
        if (
            (userData.validContactNumber && userData.validEmail) ||
            userData.status === UserStatus.Verified
        )
            return onVerifyComplete?.(userData)
    }, [userData])

    const handleOnSuccess = (userData: UserBase, nextStep: number) => {
        onVerifyChange?.(userData)
        setStep(nextStep)
    }

    return (
        <div>
            <div className="my-5 flex w-full items-center gap-x-4">
                <EcoopLogo className="size-10" />
                <StepIndicator
                    totalSteps={2}
                    currentStep={step}
                    className="w-full"
                />
            </div>
            {step === 1 && (
                <VerifyForm
                    key="1"
                    id={1} // TODO: Update this after returned resource is defined
                    userType="Admin" // TODO: Update this after returned resource is defined
                    readOnly={readOnly}
                    verifyMode="mobile"
                    onSuccess={(data) => handleOnSuccess(data, 2)}
                />
            )}
            {step === 2 && (
                <VerifyForm
                    key="2"
                    id={1} // TODO: Update this after returned resource is defined
                    userType="Member" // TODO: Update this after returned resource is defined
                    verifyMode="email"
                    readOnly={readOnly}
                    onSuccess={(data) => {
                        onVerifyComplete?.(data)
                    }}
                />
            )}
        </div>
    )
}

export default VerifyRoot
