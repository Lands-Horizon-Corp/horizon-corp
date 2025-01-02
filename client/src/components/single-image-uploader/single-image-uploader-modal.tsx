import {
    Dialog,
    DialogTitle,
    DialogHeader,
    DialogContent,
    DialogDescription,
} from '../ui/dialog'
import SingleImageUpload, {
    ISingleImageUploadProps,
} from './single-image-uploader'

import { IModalBase } from '@/types/component/modal'

interface Props extends IModalBase, ISingleImageUploadProps {}

const SingleImageUploaderModal = ({
    modalState,
    modalTitle,
    modalDescription,
    setModalState,
    ...singleImageUploadProps
}: Props) => {
    return (
        <Dialog open={modalState} onOpenChange={setModalState}>
            <DialogContent
                closeButtonClassName="sm:hidden"
                className="!rounded-2xl"
            >
                <DialogHeader>
                    <DialogTitle>{modalTitle}</DialogTitle>
                    <DialogDescription>{modalDescription}</DialogDescription>
                </DialogHeader>
                <SingleImageUpload {...singleImageUploadProps} />
            </DialogContent>
        </Dialog>
    )
}

export default SingleImageUploaderModal
