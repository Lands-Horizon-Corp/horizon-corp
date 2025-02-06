import { IMemberTypeTableActionComponentProp } from '../columns'
import RowActionsGroup from '@/components/data-table/data-table-row-actions'

import useConfirmModalStore from '@/store/confirm-modal-store'
import { useDeleteMemberType } from '@/hooks/api-hooks/member/use-member-type'

interface IMemberTypeTableOwnerActionProps
    extends IMemberTypeTableActionComponentProp {
    onMemberTypeUpdate?: () => void
    onDeleteSuccess?: () => void
}

const MemberTypeTableOwnerAction = ({
    row,
    onDeleteSuccess,
}: IMemberTypeTableOwnerActionProps) => {
    const memberType = row.original

    const { onOpen } = useConfirmModalStore()

    const { isPending: isDeletingMemberType, mutate: deleteMemberType } =
        useDeleteMemberType({
            onSuccess: onDeleteSuccess,
        })

    return (
        <RowActionsGroup
            onDelete={{
                text: 'Delete',
                isAllowed: !isDeletingMemberType,
                onClick: () => {
                    onOpen({
                        title: 'Delete Member Type',
                        description: 'Are you sure to delete this Member Type?',
                        onConfirm: () => deleteMemberType(memberType.id),
                    })
                },
            }}
            onEdit={{
                text: 'Edit',
                isAllowed: true,
                onClick: () => {
                    // router.navigate({
                    //     to: '/admin/member-types-management/$memberTypeId/edit',
                    //     params: { memberTypeId: memberType.id },
                    // })
                },
            }}
            onView={{
                text: 'View',
                isAllowed: true,
                onClick: () => {
                    // router.navigate({
                    //     to: '/admin/member-types-management/$memberTypeId/view',
                    //     params: { memberTypeId: memberType.id },
                    // })
                },
            }}
            otherActions={<>{/* Additional actions can be added here */}</>}
        />
    )
}

export default MemberTypeTableOwnerAction
