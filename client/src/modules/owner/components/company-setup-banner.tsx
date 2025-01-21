import { Button } from '@/components/ui/button'
import { BuildingIcon } from '@/components/icons'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'

interface Props extends IBaseCompNoChild {}

const CompanySetupBanner = ({ className }: Props) => {
    return (
        <div
            className={cn(
                'flex flex-col justify-center max-w-3xl gap-y-3 rounded-xl border bg-secondary dark:bg-popover p-5',
                className
            )}
        >
            <p className="text-2xl font-medium">Setup your Company</p>
            <p className="text-sm text-foreground/75">
                Create your company to unlock the full potential of ECoop. Once
                your company is set up, you can easily manage branches,
                employees, and streamline your operations!
            </p>
            <Button className="w-fit self-center">
                Create Company <BuildingIcon className="ml-2" />
            </Button>
        </div>
    )
}

export default CompanySetupBanner
