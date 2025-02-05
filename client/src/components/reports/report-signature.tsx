import { useForm, SubmitHandler } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button } from '@/components/ui/button'
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'

export enum SignedBy {
    APPROVED_BY = 'Approved by',
    PREPARED_BY = 'Prepared by',
    CERTIFIED_BY = 'Certified by',
    VERIFIED_BY = 'Verified by',
    CHECK_BY = 'Check by',
    ACKNOWLEDGE_BY = 'Acknowledge by',
    NOTED_BY = 'Noted by',
    POSTED_BY = 'Posted by',
    PAID_BY = 'Paid by',
}

export enum Position {
    LEFT = 'Left',
    RIGHT = 'Right',
    CENTER = 'Center',
}

const signatureSchema = z.object({
    name: z.string().optional(),
    position: z.nativeEnum(Position).optional(),
})

type SignatureFormData = Record<
    keyof typeof SignedBy,
    { name?: string; position?: Position }
>

const SignatureForm: React.FC<{
    open: boolean
    onSubmit: (data: SignatureFormData) => void
    onClose: () => void
}> = ({ open, onSubmit, onClose }) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        setValue,
        watch,
    } = useForm<SignatureFormData>({
        resolver: zodResolver(z.record(signatureSchema)),
    })

    const formValues = watch()

    const handleSelectChange = (
        key: keyof SignatureFormData,
        value: Position
    ) => {
        setValue(`${key}.position`, value)
    }

    const onSubmitHandler: SubmitHandler<SignatureFormData> = (data) => {
        onSubmit(data)
    }

    return (
        <Dialog open={open} onOpenChange={onClose}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Signature Name</DialogTitle>
                </DialogHeader>
                <form
                    onSubmit={handleSubmit(onSubmitHandler)}
                    className="space-y-4"
                >
                    <div className="grid grid-cols-3 items-center gap-4 font-medium">
                        <span className="col-span-1" />
                        <span className="col-span-1">Name</span>
                        <span className="col-span-1">Position</span>
                    </div>
                    {Object.entries(SignedBy).map(([key, label]) => (
                        <div
                            key={key}
                            className="grid grid-cols-3 items-center gap-4"
                        >
                            <label className="font-medium">{label}:</label>
                            <Input
                                type="text"
                                {...register(
                                    `${key}.name` as keyof SignatureFormData
                                )}
                                className="w-full"
                                placeholder="Enter Name"
                            />
                            <Select
                                value={
                                    formValues[key as keyof SignatureFormData]
                                        ?.position || ''
                                }
                                onValueChange={(value) =>
                                    handleSelectChange(
                                        key as keyof SignatureFormData,
                                        value as Position
                                    )
                                }
                            >
                                <SelectTrigger className="w-full">
                                    <SelectValue placeholder="Select Position" />
                                </SelectTrigger>
                                <SelectContent>
                                    {Object.values(Position).map((pos) => (
                                        <SelectItem key={pos} value={pos}>
                                            {pos}
                                        </SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                            {errors[key as keyof SignatureFormData]?.name && (
                                <p className="col-span-3 text-sm text-red-500">
                                    {
                                        errors[key as keyof SignatureFormData]
                                            ?.name?.message
                                    }
                                </p>
                            )}
                            {errors[key as keyof SignatureFormData]
                                ?.position && (
                                <p className="col-span-3 text-sm text-red-500">
                                    {
                                        errors[key as keyof SignatureFormData]
                                            ?.position?.message
                                    }
                                </p>
                            )}
                        </div>
                    ))}
                    <div className="mt-4 flex justify-end gap-2">
                        <Button
                            type="button"
                            variant="outline"
                            onClick={onClose}
                        >
                            Cancel
                        </Button>
                        <Button type="submit">Save</Button>
                    </div>
                </form>
            </DialogContent>
        </Dialog>
    )
}

export default SignatureForm
