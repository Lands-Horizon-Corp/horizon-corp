import { useState } from 'react'
import { IMemberClassificationTableActionComponentProp } from '../columns'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'

import useConfirmModalStore from '@/store/confirm-modal-store'
import { useDeleteMemberClassification } from '@/hooks/api-hooks/member/use-member-classification'
import { MemberClassificationCreateUpdateFormModal } from '@/components/forms/member-forms/member-classification-create-update-form'

interface IMemberClassificationTableOwnerActionProps
    extends IMemberClassificationTableActionComponentProp {
    onMemberClassificationUpdate?: () => void
    onDeleteSuccess?: () => void
}

const MemberClassificationTableOwnerAction = ({
    row,
    onDeleteSuccess,
}: IMemberClassificationTableOwnerActionProps) => {
    const [updateModalForm, setUpdateModalForm] = useState(false)
    const memberClassification = row.original

    const { onOpen } = useConfirmModalStore()

    const { isPending: isDeleting, mutate: deleteMemberClassification } =
        useDeleteMemberClassification({
            onSuccess: onDeleteSuccess,
        })

    return (
        <>
            <div onClick={(e) => e.stopPropagation()}>
                <MemberClassificationCreateUpdateFormModal
                    formProps={{
                        memberClassificationId: memberClassification.id,
                        defaultValues: {
                            ...memberClassification,
                        },
                    }}
                    title="Update Member Classification"
                    description="Modify/Update the member classification..."
                    open={updateModalForm}
                    onOpenChange={setUpdateModalForm}
                />
            </div>
            <RowActionsGroup
                onDelete={{
                    text: 'Delete',
                    isAllowed: !isDeleting,
                    onClick: () => {
                        onOpen({
                            title: 'Delete Member Classification',
                            description:
                                'Are you sure you want to delete this Member Classification?',
                            onConfirm: () =>
                                deleteMemberClassification(
                                    memberClassification.id
                                ),
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

export default MemberClassificationTableOwnerAction
