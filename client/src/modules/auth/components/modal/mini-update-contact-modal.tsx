import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'
import MiniUpdateContactsForm, {
    IMiniUpdateContactsProps,
} from '../forms/mini-update-contacts-form'

interface Props extends IMiniUpdateContactsProps {
    open: boolean
    onOpenChange: (newState: boolean) => void
}

const MiniUpdateContactsModal = ({
    open,
    onOpenChange,
    ...miniUpdateContactsFormProps
}: Props) => {
    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent
                overlayClassName="backdrop-blur"
                className="!rounded-2xl sm:w-fit"
            >
                <DialogHeader className="">
                    <DialogTitle>Update Contact Info</DialogTitle>
                    <DialogDescription>
                        Change your email or contact number
                    </DialogDescription>
                </DialogHeader>
                <MiniUpdateContactsForm {...miniUpdateContactsFormProps} />
            </DialogContent>
        </Dialog>
    )
}

export default MiniUpdateContactsModal
