import CompanySetupBanner from './company-setup-banner'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { IBaseCompChildOnly } from '@/types'
import { useUserAuthStore } from '@/store/user-auth-store'
import { useOwnerCompany } from '@/hooks/api-hooks/use-owner'

interface Props extends IBaseCompChildOnly {}

const EnsureOwnerCompany = ({ children }: Props) => {
    const { currentUser } = useUserAuthStore()

    const { data, isPending } = useOwnerCompany({
        ownerId: currentUser?.id as number,
    })

    if (isPending) return <LoadingSpinner />

    if (!data) return <CompanySetupBanner />

    return children
}

export default EnsureOwnerCompany
