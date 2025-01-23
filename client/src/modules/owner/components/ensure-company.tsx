import CompanySetupBanner from './company-setup-banner'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { IBaseCompChildOnly } from '@/types'
import { useUserAuthStore } from '@/store/user-auth-store'
import { useOwnerCompany } from '@/hooks/api-hooks/use-owner'
import { IOperationCallbacks } from '@/hooks/api-hooks/types'
import { ICompanyResource } from '@/server'

interface Props
    extends IBaseCompChildOnly,
        IOperationCallbacks<ICompanyResource> {
    disabled?: boolean
}

const EnsureOwnerCompany = ({
    children,
    onSuccess,
    disabled = false,
}: Props) => {
    const { currentUser } = useUserAuthStore()

    const { data, isPending } = useOwnerCompany({
        ownerId: currentUser?.id as number,
        onSuccess,
    })

    if (isPending && !disabled) return <LoadingSpinner />

    if (!data && !disabled) return <CompanySetupBanner />

    return children
}

export default EnsureOwnerCompany
