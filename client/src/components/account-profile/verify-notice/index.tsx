import { UserData } from '@/horizon-corp/types'
import VerifyContactBar from './verify-contact-bar'

interface Props {
    currentUser: UserData
    onSuccess: (newUserData: UserData) => void
}

const VerifyNotice = ({ currentUser, onSuccess }: Props) => {

    if(!currentUser) return null

    return (
        <>
            {!currentUser.isEmailVerified && !currentUser.isEmailVerified && (
                <VerifyContactBar key="verify-bar-email" verifyMode="email" onSuccess={onSuccess} />
            )}
            {!currentUser.isContactVerified && currentUser.isEmailVerified && (
                <VerifyContactBar key="verify-bar-mobile" verifyMode="mobile" onSuccess={onSuccess}/>
            )}
        </>
    )
}

export default VerifyNotice
