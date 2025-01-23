import { BadgeCheckFillIcon, BadgeQuestionFillIcon } from '../icons'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { ICompanyResource } from '@/server'

interface Props extends IBaseCompNoChild {
    company: ICompanyResource
}

const CompanyBadge = ({ company, className }: Props) => {
    if (!company) return

    if (company.isAdminVerified)
        return <BadgeCheckFillIcon className={cn('text-primary', className)} />

    return <BadgeQuestionFillIcon className={cn('text-amber-500', className)} />
}

export default CompanyBadge
