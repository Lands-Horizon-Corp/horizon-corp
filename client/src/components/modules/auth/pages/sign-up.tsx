import SignUpForm from '../components/forms/sign-up-form'

interface Props {}

const SignUpPage = (_props: Props) => {
    return (
        <div className="flex flex-1 items-center justify-center">
            <SignUpForm />
        </div>
    )
}

export default SignUpPage
