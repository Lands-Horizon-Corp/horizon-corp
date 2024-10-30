import { useState } from 'react'

import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import PasswordInput from '@/components/ui/password-input'

type Props<T> = {
    state: boolean
    payloadData: T
    onClose: (newState: boolean) => void
    onSubmit: (updatedData: T & { password: string }) => void
}

const PasswordInputModal = <T extends object>({
    state,
    onClose,
    onSubmit,
    payloadData,
}: Props<T>) => {
    const [password, setPassword] = useState('')

    return (
        <Dialog open={state} onOpenChange={onClose}>
            <DialogContent
                overlayClassName="backdrop-blur"
                className="shadow-2 !rounded-2xl border font-inter"
            >
                <DialogHeader>
                    <DialogTitle className="font-medium">Password Confirm</DialogTitle>
                </DialogHeader>
                <DialogDescription className="mb-4">
                    Please enter password to proceed
                </DialogDescription>
                <PasswordInput
                    hidden
                    value={password}
                    autoComplete="new-password"
                    placeholder="Enter Password"
                    onChange={(e) => setPassword(e.target.value)}
                />
                <Separator className="bg-muted/70" />
                <div className="flex justify-end gap-x-2">
                    <Button
                        onClick={() => onClose(false)}
                        variant={'ghost'}
                        className="bg-muted/60 hover:bg-muted"
                    >
                        Back
                    </Button>
                    <Button
                        onClick={() => {
                            onSubmit({ ...payloadData, password })
                            setPassword('')
                        }}
                    >
                        Okay
                    </Button>
                </div>
            </DialogContent>
        </Dialog>
    )
}

export default PasswordInputModal
