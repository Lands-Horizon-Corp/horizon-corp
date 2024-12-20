import { toast } from 'sonner'

import { Button } from '@/components/ui/button'

import { CompanyResource } from '@/horizon-corp/types'
import useConfirmModalStore from '@/store/confirm-modal-store'

interface Props {
    company: CompanyResource
}

const CompanyAcceptBar = ({ company }: Props) => {
    const { onOpen } = useConfirmModalStore()

    // TODO: Add here company service accept once implemented

    return (
        <div className="flex w-full items-center justify-between gap-x-4 rounded-xl border px-4 py-2.5">
            <span className="text-xs text-amber-600 dark:text-amber-400/80">
                This company is awaiting approval. It must be approved before it
                can operate.
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
                    onClick={() =>
                        onOpen({
                            title: 'Accept Company',
                            confirmString: 'Approve',
                            content:
                                "Are you sure you want to approve this company? This will enable them to start using E-Coop and start their company's operation.",
                            onConfirm: () => {
                                toast.info('Kunwari na accept hahaha')
                            },
                        })
                    }
                    className="h-fit rounded-none py-2 text-xs last:rounded-r-lg hover:text-primary"
                >
                    Approve
                </Button>
            </div>
        </div>
    )
}

export default CompanyAcceptBar
