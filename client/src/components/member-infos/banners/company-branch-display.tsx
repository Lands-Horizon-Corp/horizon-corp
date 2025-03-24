import ImageDisplay from '@/components/image-display'
import { cn } from '@/lib'
import { IBranchResource, ICompanyResource } from '@/server'
import { IBaseCompNoChild } from '@/types'

interface Props extends IBaseCompNoChild {
    branch?: IBranchResource
    company?: ICompanyResource
}

const CompanyBranchDisplay = ({ branch, company, className }: Props) => {
    return (
        <div className={cn('flex gap-x-2', className)}>
            <div className="flex min-w-64 items-end gap-x-2">
                <ImageDisplay
                    src={company?.media?.downloadURL}
                    className="size-16 rounded-xl"
                    fallbackClassName="rounded-xl size-16"
                />
                <div className="space-y-1">
                    <p>{company?.name ?? '-'}</p>
                    <p className="text-xs text-muted-foreground/70">Company</p>
                </div>
            </div>
            <div className="flex min-w-64 items-end gap-x-2">
                <ImageDisplay
                    src={branch?.media?.downloadURL}
                    className="size-16 rounded-xl"
                    fallbackClassName="rounded-xl size-16"
                />
                <div className="space-y-1">
                    <p>{branch?.name ?? '-'}</p>
                    <p className="text-xs text-muted-foreground/70">Branch</p>
                </div>
            </div>
        </div>
    )
}

export default CompanyBranchDisplay
