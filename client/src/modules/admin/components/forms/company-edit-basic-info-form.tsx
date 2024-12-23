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
import { LoadingSpinnerIcon } from '@/components/icons'
import FormErrorMessage from '@/components/ui/form-error-message'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { CompanyResource } from '@/horizon-corp/types'
import { useUpdateCompany } from '@/hooks/api-hooks/use-company'

type TCompanyBasicInfo = Omit<CompanyResource, 'owner' | 'media' | 'branches'>

interface CompanyEditBasicInfoFormProps
    extends IBaseCompNoChild,
        IForm<TCompanyBasicInfo, CompanyResource, string> {
    companyId: number
}

const CompanyBasicInfoFormSchema = z.object({
    name: z.string().min(1, 'Company name is required'),
    description: z
        .string()
        .min(1, 'Company description is required')
        .optional(),
    address: z.string().min(1, 'Company address is required').optional(),
    longitude: z.coerce.number().optional(),
    latitude: z.coerce.number().optional(),
    contactNumber: z.string(),
    isAdminVerified: z.boolean(),
})

const CompanyEditBasicInfoForm = ({
    readOnly,
    companyId,
    className,
    defaultValues,
    onError,
    onSuccess,
    onLoading,
}: CompanyEditBasicInfoFormProps) => {
    const form = useForm<z.infer<typeof CompanyBasicInfoFormSchema>>({
        defaultValues,
        reValidateMode: 'onChange',
        resolver: zodResolver(CompanyBasicInfoFormSchema),
    })

    const hasChanges = form.formState.isDirty
    const firstError = Object.values(form.formState.errors)[0]?.message

    const defaultCenter = {
        lat: form.getValues('latitude') ?? 14.58423341171918,
        lng: form.getValues('longitude') ?? -239.01863962431653,
    }

    const {
        isPending,
        mutate: save,
        error,
        reset,
    } = useUpdateCompany((data) => {
        onSuccess?.(data)
        form.reset(data)
    }, onError)

    useEffect(() => {
        onLoading?.(isPending)
    }, [isPending, onLoading])

    return (
        <form
            className={cn('flex flex-col gap-y-2', className)}
            onSubmit={form.handleSubmit((data) =>
                save({ id: companyId, data })
            )}
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
                            <FormItem className="col-span-2 space-y-0 sm:col-span-1">
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
                        name="address"
                        render={({ field }) => (
                            <FormItem className="col-span-2 space-y-0 sm:col-span-1">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="text-right text-sm font-normal text-foreground/60"
                                >
                                    Company Address
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
                        Coordinates use to display your company&apos;s location
                        to interactive map
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
                {hasChanges && (
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
                                {isPending ? <LoadingSpinnerIcon /> : 'Save'}
                            </Button>
                        </div>
                    </div>
                )}
            </Form>
        </form>
    )
}

export default CompanyEditBasicInfoForm
