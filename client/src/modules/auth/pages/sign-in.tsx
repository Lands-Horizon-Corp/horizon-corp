import SignInForm from '../components/forms/sign-in-form'

interface Props {}

const SignInPage = (_props: Props) => {
    return (
        <div className="flex flex-1 justify-center px-4 py-4">
            <SignInForm />
        </div>
    )
}

export default SignInPage
