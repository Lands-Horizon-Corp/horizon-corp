import {
    DotsVerticalIcon,
    EyeNoneIcon,
    PencilOutlineIcon,
    TrashIcon,
} from '@/components/icons'
import { Button } from '@/components/ui/button'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { ReactNode } from '@tanstack/react-router'

export interface IRowActionOption {
    text: string
    onClick: () => void
}

interface Props {
    onDelete?: IRowActionOption
    onView?: IRowActionOption
    onEdit?: IRowActionOption
    otherActions? : ReactNode
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
            <DropdownMenuContent>
                <DropdownMenuLabel>Action</DropdownMenuLabel>
                <DropdownMenuSeparator />
                {onView && (
                    <DropdownMenuItem onClick={onView.onClick}>
                        <EyeNoneIcon className="mr-2" />
                        {onView.text}
                    </DropdownMenuItem>
                )}
                {onEdit && (
                    <DropdownMenuItem onClick={onEdit.onClick}>
                        <PencilOutlineIcon className="mr-2" />
                        {onEdit.text}
                    </DropdownMenuItem>
                )}
                {onDelete && (
                    <DropdownMenuItem className="text-rose-400 focus:bg-destructive" onClick={onDelete.onClick}>
                        <TrashIcon className="mr-2" /> {onDelete.text}
                    </DropdownMenuItem>
                )}
                {
                    otherActions
                }
            </DropdownMenuContent>
        </DropdownMenu>
    )
}

export default RowActionsGroup
