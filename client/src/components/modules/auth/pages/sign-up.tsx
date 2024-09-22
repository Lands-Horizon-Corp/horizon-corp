import SignUpForm from '../components/forms/sign-up-form'

interface Props {}

const SignUpPage = (_props: Props) => {
    return (
        <div className="flex flex-1 px-4 py-4  justify-center">
            <SignUpForm />
        </div>
    )
}

export default SignUpPage
