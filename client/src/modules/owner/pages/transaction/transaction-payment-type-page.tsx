import PageContainer from '@/components/containers/page-container'
import { TransactionPaymentTypesCreateUpdateFormModal } from '@/components/forms/transactions/transaction-payment-types-create-update-form'
import TransactionPaymentTypesTable from '@/components/tables/transactions-table/transaction-payment-types'
import TransactionPaymentTypesTableOwnerAction from '@/components/tables/transactions-table/transaction-payment-types/row-actions/transaction-type-actions'
import { useState } from 'react'

const ownerTransactionPaymentTypePage = () => {
    const [createModal, setCreateModal] = useState(false)

    return (
        <PageContainer>
            <TransactionPaymentTypesCreateUpdateFormModal
                open={createModal}
                onOpenChange={setCreateModal}
                className="w-[30rem]"
            />

            <TransactionPaymentTypesTable
                toolbarProps={{
                    createActionProps: {
                        onClick: () => setCreateModal(true),
                    },
                }}
                className="min-h-[90vh] w-full"
                actionComponent={(props) => (
                    <TransactionPaymentTypesTableOwnerAction {...props} />
                )}
            />
        </PageContainer>
    )
}

export default ownerTransactionPaymentTypePage
