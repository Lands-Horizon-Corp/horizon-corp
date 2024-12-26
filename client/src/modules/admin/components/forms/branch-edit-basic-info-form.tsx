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
import FormErrorMessage from '@/components/ui/form-error-message'

import { cn } from '@/lib'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { BranchResource } from '@/horizon-corp/types'
import { useUpdateBranch } from '@/hooks/api-hooks/use-branch'

type TBranchBasicInfo = Omit<BranchResource, 'id' | 'media' | 'company' | 'employees' | 'members' | 'createdAt' | 'updatedAt'>

interface BranchEditBasicInfoFormProps
    extends IBaseCompNoChild,
        IForm<TBranchBasicInfo, BranchResource, string> {
    branchId: number
}

const BranchBasicInfoFormSchema = z.object({
    name: z.string().min(1, 'Branch name is required'),
    address: z.string().min(1, 'Branch address is required').optional(),
    longitude: z.coerce.number().optional(),
    latitude: z.coerce.number().optional(),
    email: z.string().email('Invalid email address'),
    contactNumber: z.string(),
    isAdminVerified: z.boolean(),
})

const BranchEditBasicInfoForm = ({
    readOnly,
    branchId,
    className,
    defaultValues,
    onError,
    onSuccess,
    onLoading,
}: BranchEditBasicInfoFormProps) => {
    const form = useForm<z.infer<typeof BranchBasicInfoFormSchema>>({
        defaultValues,
        reValidateMode: 'onChange',
        resolver: zodResolver(BranchBasicInfoFormSchema),
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
    } = useUpdateBranch({
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
            onSubmit={form.handleSubmit((data) =>
                save({ id: branchId, data })
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
                            <FormItem className="col-span-2 space-y-0 sm:col-span-1">
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
                            <FormItem className="col-span-2 space-y-0 sm:col-span-1">
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
                            <FormItem className="col-span-2 space-y-0 sm:col-span-1">
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

export default BranchEditBasicInfoForm
