import { useState } from 'react'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'

import useConfirmModalStore from '@/store/confirm-modal-store'
import { ITransactionTypeTableActionComponentProp } from '../columns'
import { useDeleteTransactionType } from '@/hooks/api-hooks/transactions/use-transaction-type'
import { TransactionTypeCreateUpdateFormModal } from '@/components/forms/transactions/transaction-create-update-form'

interface ITransactionTypeTableOwnerActionProps
    extends ITransactionTypeTableActionComponentProp {
    onTransactionTypeUpdate?: () => void
    onDeleteSuccess?: () => void
}

const TransactionTypeTableOwnerAction = ({
    row,
    onDeleteSuccess,
}: ITransactionTypeTableOwnerActionProps) => {
    const [updateModalForm, setUpdateModalForm] = useState(false)
    const transactionType = row.original

    const { onOpen } = useConfirmModalStore()

    const {
        isPending: isDeletingTransactionType,
        mutate: deleteTransactionType,
    } = useDeleteTransactionType({
        onSuccess: onDeleteSuccess,
    })

    return (
        <>
            <div onClick={(e) => e.stopPropagation()}>
                <TransactionTypeCreateUpdateFormModal
                    formProps={{
                        transactionTypeId: transactionType.id,
                        defaultValues: {
                            ...transactionType,
                        },
                    }}
                    title="Update Transaction Type"
                    description="Modify/Update transaction type details."
                    open={updateModalForm}
                    onOpenChange={setUpdateModalForm}
                />
            </div>
            <RowActionsGroup
                onDelete={{
                    text: 'Delete',
                    isAllowed: !isDeletingTransactionType,
                    onClick: () => {
                        onOpen({
                            title: 'Delete Transaction Type',
                            description:
                                'Are you sure you want to delete this Transaction Type?',
                            onConfirm: () =>
                                deleteTransactionType(transactionType.id),
                        })
                    },
                }}
                onEdit={{
                    text: 'Edit',
                    isAllowed: true,
                    onClick: () => setUpdateModalForm(true),
                }}
                otherActions={<>{/* Additional actions can be added here */}</>}
            />
        </>
    )
}

export default TransactionTypeTableOwnerAction
