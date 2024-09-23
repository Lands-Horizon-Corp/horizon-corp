import { useState } from 'react'
import SignUpForm from '../components/forms/sign-up-form'

interface Props {}

const SignUpPage = (_props: Props) => {
    const [loading] = useState(false)

    return (
        <div className="flex flex-1 justify-center px-4 py-4">
            <SignUpForm
                isLoading={loading}
                onSubmit={(val) => {
                    // TODO: Add functionality
                }}
            />
        </div>
    )
}

export default SignUpPage
