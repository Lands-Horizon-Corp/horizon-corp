import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'

import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'

import useConfirmModalStore from '@/store/confirm-modal-store'

const ConfirmModal = () => {
    const { isOpen, modalData, onConfirm, onClose, onCancel } =
        useConfirmModalStore()

    return (
        <Dialog open={isOpen} onOpenChange={onClose}>
            <DialogContent
                overlayClassName="backdrop-blur"
                className="shadow-2 !rounded-2xl border font-inter"
            >
                <DialogHeader>
                    <DialogTitle className="font-medium">
                        {modalData?.title}
                    </DialogTitle>
                </DialogHeader>
                <DialogDescription className="mb-4">
                    {modalData?.description}
                </DialogDescription>
                {modalData?.content}
                <Separator className="bg-muted/70" />
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
