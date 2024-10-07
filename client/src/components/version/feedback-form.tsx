import z from 'zod'
import { cn } from '@/lib/utils'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { Button } from '@/components/ui/button'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'
import { Textarea } from '../ui/textarea'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import { Input } from '../ui/input'

import { FeedbackFormSchema } from './validations'
import FormErrorMessage from '@/modules/auth/components/form-error-message'
import { UpdateStatus } from '@/types/constants'

type TFeedBack = z.infer<typeof FeedbackFormSchema>

const FeedbackForm = () => {
    const defaultValues = {
        description: '',
        email: '',
        message: '',
    }

    const feedbackForm = useForm<TFeedBack>({
        resolver: zodResolver(FeedbackFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })

    const handleFeedBackSubmit = (data: TFeedBack) => {
        const parsedData = FeedbackFormSchema.parse(data)
        console.log(parsedData)
        // TODO: Logic
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
                            <FormItem className="">
                                <div className="flex flex-col justify-start">
                                    <FormLabel className="h-[24px] w-full text-[14px]">
                                        Feedback type
                                    </FormLabel>
                                    <FormControl>
                                        <Select
                                            onValueChange={field.onChange}
                                            defaultValue={field.value}
                                        >
                                            <SelectTrigger
                                                className={cn(
                                                    'rounded-[10px] bg-transparent placeholder:text-[#838383]'
                                                )}
                                            >
                                                <SelectValue placeholder="Choose feedback type" />
                                            </SelectTrigger>
                                            <SelectContent>
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
                                                    value={
                                                        UpdateStatus.IMPROVEMENT
                                                    }
                                                >
                                                    {UpdateStatus.IMPROVEMENT}
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
                            <FormItem>
                                <FormLabel className="h-[24px] w-full text-[14px]">
                                    Description
                                </FormLabel>
                                <FormControl>
                                    <Textarea
                                        {...field}
                                        id="width"
                                        placeholder="Leave us some feedback..."
                                        className={cn(
                                            'max-h-40 rounded-[10px] bg-transparent placeholder:text-[#838383]'
                                        )}
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        name="email"
                        control={feedbackForm.control}
                        render={({ field }) => (
                            <FormItem className="">
                                <div className="flex flex-col justify-start">
                                    <FormLabel className="h-[24px] w-full text-[14px]">
                                        Email
                                    </FormLabel>
                                    <FormControl>
                                        <div className="flex-1 space-y-2">
                                            <Input
                                                className={cn(
                                                    'rounded-[10px] bg-transparent placeholder:text-[#838383]'
                                                )}
                                                placeholder="ecoop@email.com"
                                                {...field}
                                            />
                                        </div>
                                    </FormControl>
                                </div>
                            </FormItem>
                        )}
                    />
                    <FormErrorMessage
                        className="text-[12px]"
                        errorMessage={showFieldError}
                    />
                    <Button className={cn('w-full')}>send</Button>
                </form>
            </Form>
        </div>
    )
}

export default FeedbackForm
