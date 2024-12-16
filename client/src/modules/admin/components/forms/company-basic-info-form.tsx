import z from 'zod'
import { toast } from 'sonner'
import { useForm } from 'react-hook-form'
import { useMutation } from '@tanstack/react-query'
import { zodResolver } from '@hookform/resolvers/zod'

import { Form } from '@/components/ui/form'

import { cn } from '@/lib'
import { withCatchAsync } from '@/utils'
import { IBaseCompNoChild } from '@/types'
import { IForm } from '@/types/component/form'
import { serverRequestErrExtractor } from '@/helpers'
import { CompanyResource } from '@/horizon-corp/types'
import CompanyService from '@/horizon-corp/server/admin/CompanyService'

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

const CompanyBasicInfoForm = ({
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
        <form className={cn('', className)}>
            <Form {...form}>
                <fieldset
                    disabled={readOnly || isPending}
                    className="grid gap-x-4 gap-y-6 sm:grid-cols-2"
                >

                </fieldset>
            </Form>
        </form>
    )
}

export default CompanyBasicInfoForm
