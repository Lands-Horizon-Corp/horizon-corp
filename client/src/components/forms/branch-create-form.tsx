import z from 'zod'
import { useEffect } from 'react'
import { LatLngLiteral } from 'leaflet'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import {
    Form,
    FormItem,
    FormField,
    FormLabel,
    FormControl,
} from '@/components/ui/form'

import { Input } from '@/components/ui/input'
import MainMapContainer from '@/components/map'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { LoadingSpinnerIcon } from '@/components/icons'
import Modal, { IModalProps } from '@/components/modals/modal'
import CompanyPicker from '@/components/pickers/company-picker'
import FormErrorMessage from '@/components/ui/form-error-message'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { contactNumberSchema } from '@/validations/common'
import { useCreateBranch } from '@/hooks/api-hooks/use-branch'
import { BranchResource, CompanyRequest } from '@/horizon-corp/types'

type TBranchBasicInfo = Partial<CompanyRequest>

interface BranchCreateFormProps
    extends IBaseCompNoChild,
        IForm<TBranchBasicInfo, BranchResource, string> {}

const BranchBasicInfoFormSchema = z.object({
    name: z.string().min(1, 'Branch name is required'),
    companyId: z.coerce.number({ required_error: 'Company is required', invalid_type_error : 'Company ID is invalid' }),
    address: z.string().min(1, 'Branch address is required').optional(),
    longitude: z.coerce.number().optional(),
    latitude: z.coerce.number().optional(),
    email: z
        .string()
        .min(1, 'Email is required')
        .email('Invalid email address'),
    contactNumber: contactNumberSchema,
    isAdminVerified: z.boolean().default(false),
})

const BranchCreateForm = ({
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
    onLoading,
}: BranchCreateFormProps) => {
    const form = useForm<z.infer<typeof BranchBasicInfoFormSchema>>({
        defaultValues: {
            name: '',
            email: '',
            address: '',
            contactNumber: '',
            latitude: 14.5842,
            longitude: 120.9876,
            ...defaultValues,
        },
        reValidateMode: 'onChange',
        resolver: zodResolver(BranchBasicInfoFormSchema),
    })

    const firstError = Object.values(form.formState.errors)[0]?.message

    const defaultCenter = {
        lat: form.getValues('latitude') ?? 14.5842,
        lng: form.getValues('longitude') ?? 120.9876,
    }

    const {
        error,
        isPending,
        mutate: create,
        reset,
    } = useCreateBranch({
        onSuccess: (data) => {
            onSuccess?.(data)
            form.reset(data)
        },
        onError,
    })

    useEffect(() => {
        onLoading?.(isPending)
    }, [isPending, onLoading])

    return (
        <form
            className={cn('flex flex-col gap-y-2', className)}
            onSubmit={form.handleSubmit((data) => create(data))}
        >
            <Form {...form}>
                <fieldset
                    disabled={readOnly || isPending}
                    className="grid gap-x-4 gap-y-4 sm:grid-cols-2"
                >
                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Branch Name
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        placeholder="Branch Name"
                                        autoComplete="off"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="address"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Branch Address
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        autoComplete="off"
                                        placeholder="Address"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Email
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        type="email"
                                        autoComplete="off"
                                        placeholder="Email"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="contactNumber"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Contact Number
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        autoComplete="off"
                                        placeholder="Contact Number"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="companyId"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Company
                                </FormLabel>
                                <FormControl>
                                    <CompanyPicker
                                        value={field.value}
                                        onSelect={(company) =>
                                            field.onChange(company.id)
                                        }
                                        placeholder='Select company'
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <h4 className="col-span-2 text-sm font-normal text-foreground/60">
                        Branch Map Location
                    </h4>
                    <FormField
                        control={form.control}
                        name="latitude"
                        render={({ field }) => (
                            <FormItem className="col-span-1 space-y-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Latitude
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        type="number"
                                        autoComplete="off"
                                        placeholder="Latitude"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="longitude"
                        render={({ field }) => (
                            <FormItem className="col-span-1 space-y-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Longitude
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        type="number"
                                        autoComplete="off"
                                        placeholder="Longitude"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                </fieldset>
                <div className="">
                    <h4 className="text-sm font-normal text-foreground/60">
                        Branch Map Location
                    </h4>
                    <p className="mt-2 text-sm text-foreground/60">
                        Use the map to set the coordinates of the branch.
                    </p>
                    <MainMapContainer
                        zoom={13}
                        center={defaultCenter}
                        defaultMarkerPins={[defaultCenter]}
                        onCoordinateClick={(coord) => {
                            const { lat, lng } = coord as LatLngLiteral
                            if (!lat || !lng || isPending || readOnly) return

                            form.setValue('latitude', lat, {
                                shouldDirty: true,
                            })
                            form.setValue('longitude', lng, {
                                shouldDirty: true,
                            })
                        }}
                        className="min-h-[500px] p-0 py-2"
                    />
                </div>
                <FormErrorMessage errorMessage={firstError || error} />
                <Separator className="my-2 sm:my-4" />
                <div className="flex items-center justify-end gap-x-2">
                    <Button
                        type="button"
                        variant="ghost"
                        onClick={() => {
                            form.reset()
                            reset()
                        }}
                        className="w-full self-end px-8 sm:w-fit"
                    >
                        Reset
                    </Button>
                    <Button
                        type="submit"
                        disabled={isPending}
                        className="w-full self-end px-8 sm:w-fit"
                    >
                        {isPending ? <LoadingSpinnerIcon /> : 'Create'}
                    </Button>
                </div>
            </Form>
        </form>
    )
}

export const BranchCreateFormModal = ({
    title = 'Create Branch',
    description = 'Fill out the form to create a new branch.',
    formProps,
    className,
    ...props
}: IModalProps & { formProps?: Omit<BranchCreateFormProps, 'className'> }) => {
    return (
        <Modal title={title} className={cn('sm:max-w-5xl', className)} description={description} {...props}>
            <BranchCreateForm {...formProps} />
        </Modal>
    )
}

export default BranchCreateForm
