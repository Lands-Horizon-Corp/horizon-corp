import { useEffect, useState } from 'react'

import VerifyForm from '@/modules/auth/components/forms/verify-form'

import { UserBase } from '@/types'

interface Props {}

const Verify = ({}: Props) => {
    const [step, setStep] = useState<1 | 2 | 3>(1)
    const [userData, setUserData] = useState<UserBase | null>(null)

    useEffect(() => {
        if (!userData) return

        //  TODO: Check what user lacks in verification
        //  then set appropriate step
        //  1 - Verify Contact Number
        //  2 - Verify Email
        //  3 - Show users account status
        //      - Show if not yet validated
        //      - Show if account cancelled
    }, [userData])

    // TODO fetch/get user info

    return (
        <div className="flex flex-1 justify-center px-4 py-4">
            {step === 1 && <VerifyForm id="" verifyMode="email" />}
        </div>
    )
}

export default Verify
