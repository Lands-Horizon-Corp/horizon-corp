import SignUpForm from '../components/forms/sign-up-form'
import AuthPageWrapper from '../components/auth-page-wrapper'

interface Props {}

const SignUpPage = (_props: Props) => {
    return (
        <div className="flex min-h-full flex-col items-center justify-center">
            <AuthPageWrapper>
                <SignUpForm />
            </AuthPageWrapper>
        </div>
    )
}

export default SignUpPage
