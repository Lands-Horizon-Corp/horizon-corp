import { ReactNode } from 'react'

import {
    TrashIcon,
    EyeNoneIcon,
    DotsVerticalIcon,
    PencilOutlineIcon,
} from '@/components/icons'
import {
    DropdownMenu,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuContent,
    DropdownMenuTrigger,
    DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'

export interface IRowActionOption {
    text: string
    isAllowed?: boolean
    onClick: () => void
}

interface Props {
    onDelete?: IRowActionOption
    onView?: IRowActionOption
    onEdit?: IRowActionOption
    canEdit?: boolean
    canView?: boolean
    canDelete?: boolean
    otherActions?: ReactNode
}

const RowActionsGroup = ({ onDelete, onView, onEdit, otherActions }: Props) => {
    if (!onDelete && !onView && !onEdit) return null

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button className="size-fit p-1" variant="ghost">
                    <DotsVerticalIcon />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent onClick={(e) => e.stopPropagation()}>
                <DropdownMenuLabel>Action</DropdownMenuLabel>
                <DropdownMenuSeparator />
                {onView && (
                    <DropdownMenuItem
                        disabled={!onView.isAllowed}
                        onClick={onView.onClick}
                    >
                        <EyeNoneIcon className="mr-2" />
                        {onView.text}
                    </DropdownMenuItem>
                )}
                {onEdit && (
                    <DropdownMenuItem
                        disabled={!onEdit.isAllowed}
                        onClick={onEdit.onClick}
                    >
                        <PencilOutlineIcon className="mr-2" />
                        {onEdit.text}
                    </DropdownMenuItem>
                )}
                {onDelete && (
                    <DropdownMenuItem
                        disabled={!onDelete.isAllowed}
                        className="text-rose-400 focus:bg-destructive"
                        onClick={onDelete.onClick}
                    >
                        <TrashIcon className="mr-2" /> {onDelete.text}
                    </DropdownMenuItem>
                )}
                {otherActions}
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default RowActionsGroup
