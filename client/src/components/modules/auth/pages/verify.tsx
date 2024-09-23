import VerifyForm from '@/modules/auth/components/forms/verify-form'

interface Props {}

const Verify = ({}: Props) => {
    return (
        <div className="flex flex-1 justify-center px-4 py-4">
            <VerifyForm id="" verifyMode="email" />
        </div>
    )
}

export default Verify
