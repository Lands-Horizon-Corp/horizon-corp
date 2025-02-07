import Modal from '@/components/modals/modal'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'

import useConfirmModalStore from '@/store/confirm-modal-store'

const ConfirmModal = () => {
    const { isOpen, modalData, onConfirm, onClose, onCancel } =
        useConfirmModalStore()

    return (
        <Modal
            open={isOpen}
            onOpenChange={onClose}
            title={modalData?.title}
            description={modalData?.description}
        >
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
                <Button onClick={onConfirm}>{modalData?.confirmString}</Button>
            </div>
        </Modal>
    )
}

export default ConfirmModal
