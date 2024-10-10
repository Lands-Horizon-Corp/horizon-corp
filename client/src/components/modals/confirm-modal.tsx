import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'

import useConfirmModalStore from '@/store/confirm-store'
import { Separator } from '../ui/separator'
import { Button } from '../ui/button'

interface Props {}

const ConfirmModal = ({}: Props) => {
    const { isOpen, modalData, onConfirm, onClose, onCancel } =
        useConfirmModalStore()

    return (
        <Dialog open={isOpen} onOpenChange={onClose}>
            <DialogContent className="shadow-2 border-none font-inter sm:rounded-2xl">
                <DialogHeader>
                    <DialogTitle className="font-medium">
                        {modalData?.title}
                    </DialogTitle>
                </DialogHeader>
                <DialogDescription className="my-4">
                    {modalData?.description}
                </DialogDescription>
                <Separator className="bg-muted/70" />
                {modalData?.content}
                <div className="flex justify-end gap-x-2">
                    <Button
                        onClick={onCancel}
                        variant={'ghost'}
                        className="bg-muted/60 hover:bg-muted"
                    >
                        {modalData?.cancelString}
                    </Button>
                    <Button onClick={onConfirm}>
                        {modalData?.confirmString}
                    </Button>
                </div>
            </DialogContent>
        </Dialog>
    )
}

export default ConfirmModal
