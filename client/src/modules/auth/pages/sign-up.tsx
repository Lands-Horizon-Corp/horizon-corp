import SignUpForm from '../components/forms/sign-up-form'
import GuestGuard from '@/components/wrappers/guest-guard'
import AuthPageWrapper from '../components/auth-page-wrapper'

const SignUpPage = () => {
    return (
        <GuestGuard>
            <div className="flex min-h-full flex-col items-center justify-center">
                <AuthPageWrapper>
                    <SignUpForm />
                </AuthPageWrapper>
            </div>
        </GuestGuard>
    )
}

export default SignUpPage
