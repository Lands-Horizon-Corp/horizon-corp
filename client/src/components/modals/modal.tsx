import * as DialogPrimitive from '@radix-ui/react-dialog'

import {
    Dialog,
    DialogTitle,
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
    titleClassName?: string
    descriptionClassName?: string
}

const Modal = ({
    title,
    footer,
    children,
    className,
    description,
    titleClassName,
    overlayClassName,
    closeButtonClassName,
    descriptionClassName,
    ...other
}: IModalProps) => {
    return (
        <Dialog {...other}>
            <DialogContent
                closeButtonClassName={closeButtonClassName}
                overlayClassName={cn('backdrop-blur', overlayClassName)}
                className={cn(
                    'shadow-2 ecoop-scroll max-h-[95vh] max-w-xl overflow-y-auto !rounded-2xl border font-inter',
                    className
                )}
            >
                <DialogTitle className={cn('font-medium', titleClassName)}>
                    {title}
                </DialogTitle>
                <DialogDescription className={cn('mb-4', descriptionClassName)}>
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
