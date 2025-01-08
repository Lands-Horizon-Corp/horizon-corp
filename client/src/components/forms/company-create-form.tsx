import z from 'zod'
import { useEffect } from 'react'
import { LatLngLiteral } from 'leaflet'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'

import { Input } from '@/components/ui/input'
import MainMapContainer from '@/components/map'
import { Button } from '@/components/ui/button'
import TextEditor from '@/components/text-editor'
import { Separator } from '@/components/ui/separator'
import Modal, { IModalProps } from '@/components/modals/modal'
import FormErrorMessage from '@/components/ui/form-error-message'
import { PhoneInput } from '@/components/contact-input/contact-input'
import { LoadingSpinnerIcon, VerifiedPatchIcon } from '@/components/icons'

import { cn } from '@/lib'
import { useCreateCompany } from '@/hooks/api-hooks/use-company'

import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { CompanyRequest, CompanyResource } from '@/horizon-corp/types'
import { contactNumberSchema } from '@/validations/common'
import { Textarea } from '../ui/textarea'

interface CompanyCreateFormProps
    extends IBaseCompNoChild,
        IForm<CompanyRequest, CompanyResource> {}

const CompanyBasicInfoFormSchema = z.object({
    name: z.string().min(1, 'Company name is required'),
    description: z
        .string()
        .min(1, 'Company description is required')
        .optional(),
    address: z.string().min(1, 'Company address is required').optional(),
    longitude: z.coerce.number().optional(),
    latitude: z.coerce.number().optional(),
    contactNumber: contactNumberSchema,
})

const CompanyCreateForm = ({
    readOnly,
    className,
    defaultValues,
    onError,
    onSuccess,
    onLoading,
}: CompanyCreateFormProps) => {
    const form = useForm<z.infer<typeof CompanyBasicInfoFormSchema>>({
        defaultValues: {
            address: '',
            contactNumber: '',
            description: '',
            latitude: 14.58423341171918,
            longitude: 120.987654321,
            name: '',
            ...defaultValues,
        },
        reValidateMode: 'onChange',
        resolver: zodResolver(CompanyBasicInfoFormSchema),
    })

    const firstError = Object.values(form.formState.errors)[0]?.message

    const defaultCenter = {
        lat: form.getValues('latitude') ?? 14.58423341171918,
        lng: form.getValues('longitude') ?? 120.987654321,
    }

    const {
        isPending,
        mutate: create,
        error,
        reset,
    } = useCreateCompany({
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
                                    Company Name
                                </FormLabel>
                                <FormControl>
                                    <Input
                                        {...field}
                                        id={field.name}
                                        placeholder="Company Name"
                                        autoComplete="off"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="contactNumber"
                        render={({ field, fieldState: { invalid, error } }) => (
                            <FormItem className="col-span-2 space-y-1 sm:col-span-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Contact Number
                                </FormLabel>
                                <FormControl>
                                    <div className="relative flex w-full flex-1 items-center gap-x-2">
                                        <PhoneInput
                                            {...field}
                                            defaultCountry="PH"
                                            className="w-full"
                                        />
                                        <VerifiedPatchIcon
                                            className={cn(
                                                'absolute right-2 top-1/2 size-4 -translate-y-1/2 text-primary delay-300 duration-300 ease-in-out',
                                                (invalid || error) &&
                                                    'text-destructive'
                                            )}
                                        />
                                    </div>
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="description"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Description
                                </FormLabel>
                                <FormControl>
                                    <TextEditor
                                        content={field.value}
                                        onChange={field.onChange}
                                        className="!max-w-none"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="address"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Company Address
                                </FormLabel>
                                <FormControl>
                                    <Textarea
                                        {...field}
                                        id={field.name}
                                        autoComplete="off"
                                        placeholder="Address"
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <h4 className="col-span-2 text-sm font-normal text-foreground/60">
                        Company Map Location
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
                        Company Map Location
                    </h4>
                    <p className="mt-2 text-sm text-foreground/60">
                        Use the map to set the coordinates of the company.
                        Coordinates will display your company&apos;s location on
                        the interactive map.
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
                        className="min-h-[500px] w-full p-0 py-2"
                    />
                </div>
                <FormErrorMessage errorMessage={firstError || error} />
                <div>
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
                </div>
            </Form>
        </form>
    )
}

export const CompanyCreateFormModal = ({
    title = 'Company Create',
    description = 'Setup a new company',
    formProps,
    className,
    ...props
}: IModalProps & { formProps?: Omit<CompanyCreateFormProps, 'className'> }) => {
    return (
        <Modal
            title={title}
            className={cn('sm:max-w-5xl', className)}
            description={description}
            {...props}
        >
            <CompanyCreateForm {...formProps} />
        </Modal>
    )
}

export default CompanyCreateForm
