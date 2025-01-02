import * as DialogPrimitive from '@radix-ui/react-dialog'

import {
    Dialog,
    DialogTitle,
    DialogHeader,
    DialogContent,
    DialogDescription,
    DialogExtraProps,
} from '@/components/ui/dialog'
import { Separator } from '@/components/ui/separator'

import { cn } from '@/lib'
import { IBaseComp } from '@/types'

export interface IModalProps
    extends IBaseComp,
        DialogPrimitive.DialogProps,
        DialogExtraProps {
    title?: string
    description?: string
    footer?: React.ReactNode
}

const Modal = ({
    title,
    footer,
    children,
    className,
    description,
    overlayClassName,
    closeButtonClassName,
    ...other
}: IModalProps) => {
    return (
        <Dialog {...other}>
            <DialogContent
                closeButtonClassName={closeButtonClassName}
                overlayClassName={cn('backdrop-blur', overlayClassName)}
                className={cn(
                    'shadow-2 !rounded-2xl border font-inter max-w-5xl max-h-[95vh] overflow-y-auto ecoop-scroll',
                    className
                )}
            >
                <DialogHeader>
                    <DialogTitle className="font-medium">{title}</DialogTitle>
                </DialogHeader>
                <DialogDescription className="mb-4">
                    {description}
                </DialogDescription>
                {children}
                {footer && <Separator className="bg-muted/70" />}
                {footer}
            </DialogContent>
        </Dialog>
    )
}

export default Modal
