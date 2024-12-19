import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { useMutation } from '@tanstack/react-query'
import { zodResolver } from '@hookform/resolvers/zod'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'

import { cn } from '@/lib'
import { withCatchAsync } from '@/utils'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { serverRequestErrExtractor } from '@/helpers'
import { CompanyResource } from '@/horizon-corp/types'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'
import { Input } from '@/components/ui/input'
import TextEditor from '@/components/text-editor'
import MainMapContainer from '@/components/map'
import logger from '@/helpers/loggers/logger'

type TCompanyBasicInfo = Omit<CompanyResource, 'owner' | 'media' | 'branches'>

interface Props
    extends IBaseCompNoChild,
        IForm<Partial<TCompanyBasicInfo>, CompanyResource, string> {
    companyId: number
}

const CompanyBasicInfoFormSchema = z.object({
    name: z.string(),
    description: z.string(),
    address: z.string(),
    longitude: z.number().optional(),
    latitude: z.number().optional(),
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
}: Props) => {
    const form = useForm({
        defaultValues,
        reValidateMode: 'onChange',
        resolver: zodResolver(CompanyBasicInfoFormSchema),
    })

    const { isPending } = useMutation<
        CompanyResource,
        string,
        z.infer<typeof CompanyBasicInfoFormSchema>
    >({
        mutationKey: ['admin', 'company', companyId],
        mutationFn: async (data) => {
            onLoading?.(true)

            const [error, response] = await withCatchAsync(
                CompanyService.update(companyId, data)
            )

            onLoading?.(false)

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                if (onError) onError?.(errorMessage)
                else toast.error(errorMessage)
                throw errorMessage
            }

            onSuccess?.(response)
            return response
        },
    })

    return (
        <form
            className={cn('', className)}
            onSubmit={form.handleSubmit((data) => {})}
        >
            <Form {...form}>
                <fieldset
                    disabled={readOnly || isPending}
                    className="grid gap-x-4 gap-y-6 sm:grid-cols-2"
                >
                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem className="space-y-0">
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
                            <FormItem className="space-y-0">
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
                    <div className="col-span-2">
                        <h4 className="text-sm font-normal text-foreground/60">
                            Company Map Location
                        </h4>
                        <p className="text-sm mt-2 text-foreground/60">Use the map to set the coordinates of the company. Coordinates use to display your company&apos;s location to interactive map</p>
                        <MainMapContainer
                            zoom={13}
                            center={{
                                lat:
                                    form.getValues('latitude') ??
                                    14.58423341171918,
                                lng:
                                    form.getValues('longitude') ??
                                    -239.01863962431653,
                            }}
                        />
                    </div>
                </fieldset>
            </Form>
        </form>
    )
}

export default CompanyEditBasicInfoForm
