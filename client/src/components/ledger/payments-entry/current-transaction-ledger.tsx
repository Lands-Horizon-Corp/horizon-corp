import { memo } from 'react'
import LedgerCard from '@/components/accounts-card'
import SkeletonLedgerCard from '@/components/Skeleton/transaction/skeleton-ledger-card'
import { Card } from '@/components/ui/card'
import { IAccountingLedgerResource } from '@/server/types/accounts/accounting-ledger'
import { EmptyIcon } from '@/components/icons'
import { Button } from '@/components/ui/button'

type NoCurrentPaymentProps = {
    onClick: () => void
}

const NoCurrentPayment = memo(({ onClick }: NoCurrentPaymentProps) => (
    <Card className="flex flex-col items-center justify-center gap-5 rounded-3xl p-8 shadow-md">
        <div className="flex flex-col items-center gap-y-3">
            <EmptyIcon size={48} className="text-gray-400 dark:text-gray-300" />
            <h2 className="text-lg font-semibold text-gray-700 dark:text-gray-200">
                No Payments Found
            </h2>
            <p className="text-center text-sm text-gray-500 dark:text-gray-400">
                There are currently no processed payments. Try reloading the
                page.
            </p>
            <Button
                onClick={onClick}
                variant="outline"
                className="rounded-full px-6"
            >
                Reload
            </Button>
        </div>
    </Card>
))

NoCurrentPayment.displayName = 'NoCurrentPayment'

type CurrentPaymentAccountingTransactionLedgerProps = {
    data: IAccountingLedgerResource[]
    isPending: boolean
    isRefetching: boolean
    onRefetch: () => void
}

const CurrentPaymentAccountingTransactionLedger = ({
    data,
    isPending,
    isRefetching,
    onRefetch,
}: CurrentPaymentAccountingTransactionLedgerProps) => {
    const hasAccountLedger = data.length > 0

    return (
        <div aria-live="polite" className="space-y-4">
            {isPending || isRefetching ? (
                <SkeletonLedgerCard />
            ) : hasAccountLedger ? (
                data.map((ledger) => (
                    <LedgerCard
                        key={
                            ledger.id ??
                            `${ledger.createdAt}-${ledger.or_number}`
                        }
                        ledger={ledger}
                    />
                ))
            ) : (
                <NoCurrentPayment onClick={onRefetch} />
            )}
        </div>
    )
}

export default memo(CurrentPaymentAccountingTransactionLedger)
