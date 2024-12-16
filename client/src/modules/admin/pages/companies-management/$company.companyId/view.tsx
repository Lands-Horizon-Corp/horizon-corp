import { toast } from 'sonner'
import { useQuery } from '@tanstack/react-query'
import { useParams } from '@tanstack/react-router'

import {
    StoreIcon,
    CalendarIcon,
    TelephoneIcon,
    LocationPinIcon,
    BadgeCheckFillIcon,
    BadgeQuestionFillIcon,
    QuestionCircleFillIcon,
} from '@/components/icons'
import UserAvatar from '@/components/user-avatar'
import LoadingSpinner from '@/components/spinners/loading-spinner'

import { serverRequestErrExtractor } from '@/helpers'
import { CompanyResource } from '@/horizon-corp/types'
import { toReadableDate, withCatchAsync } from '@/utils'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { cn } from '@/lib'

const CompanyViewPage = () => {
    const { companyId } = useParams({
        from: '/admin/companies-management/$companyId/view',
    })

    const { data: company, isPending } = useQuery<CompanyResource, string>({
        queryKey: ['company', companyId],
        queryFn: async () => {
            const [error, data] = await withCatchAsync(
                CompanyService.getById(companyId)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                throw errorMessage
            }

            return data
        },
    })

    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            {isPending ? <LoadingSpinner /> : null}
            {company && (
                <div className="flex w-full max-w-5xl flex-col items-center space-y-4">
                    <div className="flex w-full flex-col items-center space-y-4 overflow-clip rounded-xl bg-secondary">
                        {/* <h1 className="self-start text-xl">Company Details</h1> */}
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
                                    className="text-xs py-2 h-fit rounded-none first:rounded-l-lg"
                                >
                                    Decline
                                </Button>
                                <Button
                                    size="sm"
                                    variant="secondary"
                                    className="text-xs py-2 h-fit hover:text-primary rounded-none last:rounded-r-lg"
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
                        <p className="text-foreground/70">
                            {company.description}
                        </p>
                    </div>

                    <Separator className="w-full" />

                    <div className="w-full space-y-4">
                        <h3 className="font-medium">Company Owner</h3>
                        <div className="flex items-center gap-x-2">
                            {company.owner ? (
                                <>
                                    <UserAvatar
                                        src={
                                            company.owner?.media?.downloadURL ??
                                            ''
                                        }
                                    />
                                    <div>
                                        <p>{`${company.owner?.username}`}</p>
                                        <p>{`${company.owner?.firstName} ${company.owner?.lastName}`}</p>
                                    </div>
                                </>
                            ) : (
                                <span className="text-sm text-destructive">
                                    Owner is missing or has been deleted.
                                </span>
                            )}
                        </div>
                    </div>
                </div>
            )}
        </div>
    )
}

export default CompanyViewPage
