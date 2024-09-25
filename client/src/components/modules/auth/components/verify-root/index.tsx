import { useEffect, useState } from 'react'

import StepIndicator from '../../../../steps-indicator'
import VerifyForm from '../forms/verify-form'

import { UserBase, UserStatus } from '@/types'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {
    readOnly?: boolean
    userData: UserBase
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
                <img src="/e-coop-logo-1.png" className="size-10" />
                <StepIndicator
                    totalSteps={2}
                    currentStep={step}
                    className="w-full"
                />
            </div>
            {step === 1 && (
                <VerifyForm
                    key="1"
                    id={userData.id}
                    readOnly={readOnly}
                    verifyMode="mobile"
                    onSuccess={(data) => handleOnSuccess(data, 2)}
                />
            )}
            {step === 2 && (
                <VerifyForm
                    key="2"
                    id={userData.id}
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
