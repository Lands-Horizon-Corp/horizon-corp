import React from 'react'
import { Card, CardContent } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import {
    IAccountingLedgerResource,
    TransactionSource,
} from '@/server/types/accounts/accounting-ledger'
import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from '@/components/ui/accordion'
// import { Tabs, TabsContent, TabsList, TabsTrigger } from '../ui/tabs'
import { CheckCircle } from 'lucide-react'
import { cn } from '@/lib'

type TransactionSourceBadgeProps = {
    source: TransactionSource
    className?: string
}

export const TransactionSourceBadge: React.FC<TransactionSourceBadgeProps> = ({
    source,
    className,
}) => {
    const colorMap: Record<TransactionSource, string> = {
        withdrawal: 'bg-red-100 text-red-700',
        deposit: 'bg-blue-100 text-blue-700',
        payment: 'bg-green-100 text-green-700',
    }

    return (
        <Badge
            className={cn(
                'rounded-full px-2 text-[10px]',
                colorMap[source],
                className
            )}
        >
            {source.toUpperCase()}
        </Badge>
    )
}

type LedgerCardProps = {
    ledger: IAccountingLedgerResource
}

const LedgerCard: React.FC<LedgerCardProps> = ({ ledger }) => {
    return (
        <Card className="flex flex-col gap-4 rounded-2xl p-6 text-gray-800 shadow-lg dark:text-white">
            <CardContent className="relative flex flex-col p-0">
                <div className="flex w-full justify-end">
                    <TransactionSourceBadge
                        className=""
                        source={ledger.transaction_source}
                    />
                </div>
                <div className="flex items-center justify-between">
                    <div>
                        <p className="text-sm text-gray-500">Description</p>
                        <h2 className="text-lg font-bold">
                            {ledger.description}
                        </h2>
                        <p className="text-xs text-gray-400">{ledger.notes}</p>
                    </div>
                    <div className="text-right">
                        <p className="text-2xl font-bold text-green-600 dark:text-green-400">
                            ₱{' '}
                            {ledger.credit === 0
                                ? '- ' + ledger.debit
                                : ledger.credit}
                        </p>
                    </div>
                </div>
                <Accordion type="single" collapsible>
                    <AccordionItem value="1">
                        <AccordionTrigger>
                            <div className="flex items-center justify-between">
                                <p className="text-sm text-gray-500">
                                    {' '}
                                    Transaction Details
                                </p>
                            </div>
                        </AccordionTrigger>
                        <AccordionContent>
                            <div className="space-y-4">
                                <div className="flex justify-between">
                                    <span className="text-sm text-gray-500">
                                        Transaction Number
                                    </span>
                                    <span className="font-medium">
                                        {ledger.or_number}
                                    </span>
                                </div>
                                <div className="flex items-center justify-between">
                                    <span className="text-sm text-gray-500">
                                        Status
                                    </span>
                                    <div className="flex items-center gap-2 text-green-600 dark:text-green-400">
                                        <CheckCircle size={16} /> Success
                                    </div>
                                </div>
                                <div className="flex justify-between">
                                    <span className="text-sm text-gray-500">
                                        Transaction Date
                                    </span>
                                    <span>
                                        {ledger.transaction_date.toLocaleString()}
                                    </span>
                                </div>
                                <div className="flex justify-between">
                                    <span className="text-sm text-gray-500">
                                        Recipient
                                    </span>
                                    <span className="font-medium">
                                        {ledger.member_profile_id}
                                    </span>
                                </div>
                                <div className="flex justify-between">
                                    <span className="text-sm text-gray-500">
                                        Amount Sent
                                    </span>
                                    <span>₱ {ledger.debit.toFixed(2)}</span>
                                </div>
                                <div className="flex justify-between">
                                    <span className="text-sm text-gray-500">
                                        Completed On
                                    </span>
                                    <span>
                                        {ledger.entry_date.toLocaleString()}
                                    </span>
                                </div>
                            </div>
                        </AccordionContent>
                    </AccordionItem>
                </Accordion>
            </CardContent>
        </Card>
    )
}

export default LedgerCard
