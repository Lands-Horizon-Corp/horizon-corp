import { useMemo } from 'react'
import { useParams } from '@tanstack/react-router'
import { useSuspenseQuery } from '@tanstack/react-query'

import {
    StoreIcon,
    CalendarIcon,
    TelephoneIcon,
    LocationPinIcon,
    BadgeMinusFillIcon,
    BadgeCheckFillIcon,
    BadgeQuestionFillIcon,
    QuestionCircleFillIcon,
    BadgeExclamationFillIcon,
} from '@/components/icons'
import { Button } from '@/components/ui/button'
import UserAvatar from '@/components/user-avatar'
import { Separator } from '@/components/ui/separator'

import { cn } from '@/lib'
import { companyLoader } from './route'
import { toReadableDate } from '@/utils'
import { OwnerResource } from '@/horizon-corp/types/profile'

const CompanyOwnerSection = ({ owner }: { owner: OwnerResource }) => {
    const AccountBadge = useMemo(() => {
        switch (owner.status) {
            case 'Verified':
                return <BadgeCheckFillIcon className="text-primary" />
            case 'Pending':
                return <BadgeQuestionFillIcon className="text-amber-600" />
            case 'Not Allowed':
                return <BadgeExclamationFillIcon className="text-destructive" />
            default:
                return <BadgeMinusFillIcon className="text-foreground/80" />
        }
    }, [owner.status])

    return (
        <div className="space-y-1 rounded-xl bg-popover">
            <div className="flex items-start justify-between gap-x-2 px-4 pb-4 pt-6">
                <div className="space-y-2">
                    <h2 className="text-2xl">{`${owner.firstName} ${owner.middleName} ${owner.lastName}`}</h2>
                    <h4 className="text-sm text-foreground/80">
                        {`${owner.email}`}{' '}
                    </h4>
                    <p className="w-fit rounded-md bg-secondary px-2 py-0.5 text-sm text-foreground/40">
                        #{owner.id}
                    </p>
                </div>
                <div className="relative size-fit">
                    <UserAvatar
                        fallback="-"
                        className="size-20"
                        src={owner.media?.downloadURL ?? ''}
                    />
                    <span className="absolute bottom-0 left-0">
                        {AccountBadge}
                    </span>
                </div>
            </div>
            <p className="px-4 pb-4 text-foreground/80">{owner.description}</p>
        </div>
    )
}

const CompanyViewPage = () => {
    const { companyId } = useParams({
        from: '/admin/companies-management/$companyId/view',
    })

    const { data: company } = useSuspenseQuery(companyLoader(companyId))

    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <div className="flex w-full max-w-5xl flex-col items-center space-y-4">
                <div className="flex w-full flex-col items-center space-y-4 overflow-clip rounded-xl bg-secondary">
                    <div className="w-full flex-col items-center space-y-2 overflow-clip rounded-xl bg-popover">
                        <div className="h-[180px] w-full rounded-md bg-[url('/profile-cover.png')] bg-cover bg-center" />
                        <div className="relative w-full space-y-2.5 px-6 pb-6 pt-8">
                            <UserAvatar
                                src={company.media?.downloadURL ?? ''}
                                className={cn(
                                    'absolute -top-28 size-36 border-2 border-primary',
                                    !company.isAdminVerified &&
                                        'border-amber-600'
                                )}
                            />
                            <span className="flex items-center gap-x-2">
                                <h3 className="text-lg font-medium">
                                    {company.name}
                                </h3>
                                {company.isAdminVerified ? (
                                    <BadgeCheckFillIcon className="text-primary" />
                                ) : (
                                    <BadgeQuestionFillIcon className="text-amber-500" />
                                )}
                            </span>
                            <span className="inline-flex items-center gap-x-2 text-sm text-foreground/80">
                                <LocationPinIcon /> {company.address}
                            </span>
                            <div className="flex gap-x-4 text-sm text-foreground/50">
                                <span className="inline-flex items-center gap-x-2">
                                    <TelephoneIcon />
                                    {company.contactNumber}
                                </span>
                                <span className="inline-flex items-center gap-x-2">
                                    <CalendarIcon />
                                    {toReadableDate(company.createdAt)}
                                </span>
                                <span className="inline-flex items-center gap-x-2">
                                    <StoreIcon />
                                    {company.branches?.length ??
                                        0} Branch
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
                {!company.isAdminVerified && (
                    <div className="flex w-full items-center justify-between gap-x-4 rounded-xl border px-4 py-2.5">
                        <span className="text-xs text-amber-600 dark:text-amber-400/80">
                            By default, new companies need&apos;s approval
                            before they can operate. This company need&apos;s
                            approval.
                        </span>
                        <div className="flex">
                            <Button
                                size="sm"
                                variant="destructive"
                                className="h-fit rounded-none py-2 text-xs first:rounded-l-lg"
                            >
                                Decline
                            </Button>
                            <Button
                                size="sm"
                                variant="secondary"
                                className="h-fit rounded-none py-2 text-xs last:rounded-r-lg hover:text-primary"
                            >
                                Approve
                            </Button>
                        </div>
                    </div>
                )}
                <div className="w-full space-y-4">
                    <h3 className="flex items-center gap-x-2 font-medium">
                        Company&apos;s Description
                        <QuestionCircleFillIcon className="inline text-foreground/20" />
                    </h3>
                    <p className="text-foreground/70">{company.description}</p>
                </div>
                <Separator className="w-full" />
                <div className="w-full space-y-4">
                    <h3 className="font-medium">Company Owner</h3>
                    {company.owner ? (
                        <CompanyOwnerSection owner={company.owner} />
                    ) : (
                        <div className="flex items-center gap-x-2">
                            <span className="text-sm text-destructive">
                                Owner is missing or has been deleted.
                            </span>
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}

export default CompanyViewPage
