import z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useMutation } from '@tanstack/react-query'
import { serverRequestErrExtractor } from '@/helpers'
import { cn, withCatchAsync } from '@/lib/utils'

import { Button } from '@/components/ui/button'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import { Input } from '../ui/input'

import { FeedbackFormSchema } from './validations'
import FormErrorMessage from '@/components/ui/form-error-message'
import { UpdateStatus } from '@/types/constants'
import TextEditor from '../text-editor'
import FeedbackService from '@/horizon-corp/server/common/FeedbackService'
import { toast } from 'sonner'

import { LoadingCircleIcon } from '../icons'

type TFeedBack = z.infer<typeof FeedbackFormSchema>

const FeedbackForm = () => {
    const defaultValues = {
        feedbackType: '',
        email: '',
        description: '',
    }

    const feedbackForm = useForm<TFeedBack>({
        resolver: zodResolver(FeedbackFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })

    const { mutate: sendFeedbackMessage, isPending } = useMutation<
        void,
        string,
        TFeedBack
    >({
        mutationKey: ['send-feedback-message'],
        mutationFn: async (feedbackRequest: TFeedBack) => {
            if (!feedbackRequest) return

            const [error] = await withCatchAsync(
                FeedbackService.create(feedbackRequest)
            )

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                return
            }

            feedbackForm.reset()
            toast.success(`Feedback Message was Already Sent!`)
        },
    })

    const handleFeedBackSubmit = (data: TFeedBack) => {
        const parsedData = FeedbackFormSchema.parse(data)
        sendFeedbackMessage(parsedData)
    }

    const showFieldError = Object.values(feedbackForm.formState.errors)[0]
        ?.message
    return (
        <div className="space-y-2 px-2">
            <Form {...feedbackForm}>
                <form
                    onSubmit={feedbackForm.handleSubmit(handleFeedBackSubmit)}
                    className="space-y-3"
                >
                    <FormField
                        name="feedbackType"
                        control={feedbackForm.control}
                        render={({ field }) => (
                            <FormItem>
                                <div className="flex flex-col justify-start">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="h-[24px] w-full text-[14px]"
                                    >
                                        Feedback type
                                    </FormLabel>
                                    <FormControl>
                                        <Select
                                            onValueChange={field.onChange}
                                            defaultValue={field.value}
                                            name={field.name}
                                        >
                                            <SelectTrigger
                                                className={cn(
                                                    'max-w-80 rounded-[10px] bg-transparent placeholder:text-[#838383]'
                                                )}
                                                id={field.name}
                                            >
                                                <SelectValue placeholder="Choose feedback type" />
                                            </SelectTrigger>
                                            <SelectContent id="helo">
                                                <SelectItem
                                                    value={UpdateStatus.FEATURE}
                                                >
                                                    {UpdateStatus.FEATURE}
                                                </SelectItem>
                                                <SelectItem
                                                    value={UpdateStatus.BUG}
                                                >
                                                    {UpdateStatus.BUG}
                                                </SelectItem>
                                                <SelectItem
                                                    value={UpdateStatus.GENERAL}
                                                >
                                                    {UpdateStatus.GENERAL}
                                                </SelectItem>
                                            </SelectContent>
                                        </Select>
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormField
                        name="description"
                        control={feedbackForm.control}
                        render={({ field }) => (
                            <FormItem className="">
                                <FormLabel
                                    htmlFor={field.name}
                                    className="h-[24px] w-full text-[14px]"
                                >
                                    Description
                                </FormLabel>
                                <FormControl>
                                    <TextEditor
                                        content={field.value}
                                        onChange={field.onChange}
                                    ></TextEditor>
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        name="email"
                        control={feedbackForm.control}
                        render={({ field }) => (
                            <FormItem className="max-w-80">
                                <div className="flex flex-col justify-start">
                                    <FormLabel
                                        htmlFor={field.name}
                                        className="h-[24px] w-full text-[14px]"
                                    >
                                        Email
                                    </FormLabel>
                                    <FormControl>
                                        <div className="flex-1 space-y-2">
                                            <Input
                                                id={field.name}
                                                className={cn(
                                                    'rounded-[10px] bg-transparent placeholder:text-[#838383]'
                                                )}
                                                placeholder="ecoop@email.com"
                                                {...field}
                                                autoComplete="email"
                                            />
                                        </div>
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormErrorMessage
                        className="w-fit text-[12px]"
                        errorMessage={showFieldError}
                    />
                    <Button className={cn('w-full')}>
                        {isPending ? (
                            <LoadingCircleIcon className="animate-spin" />
                        ) : (
                            'Send Feedback'
                        )}
                    </Button>
                </form>
            </Form>
        </div>
    )
}

export default FeedbackForm
