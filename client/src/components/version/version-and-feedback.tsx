import { useState } from 'react'
import packageJson from '../../../package.json'
import z from 'zod'
import { cn } from '@/lib/utils'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { Button } from '@/components/ui/button'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { Badge } from '@/components/ui/badge'
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

import { MdOutlineExpandLess } from 'react-icons/md'
import { MdExpandMore } from 'react-icons/md'

import { FeedbackFormSchema } from './validations'
import FormErrorMessage from '@/modules/auth/components/form-error-message'

type TFeedBack = z.infer<typeof FeedbackFormSchema>

export const VersionAndFeedBack = () => {
    const [isOpen, setIsOpen] = useState(false)

    const FeedbackType = [
        {
            name: 'Bug',
        },
        {
            name: 'improvement',
        },
        {
            name: 'Feature',
        },
        {
            name: 'Gratitude',
        },
    ]

    const defaultValues = {
        feedbackType: '',
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

    const togglePopover = () => {
        setIsOpen(!isOpen)
    }

    return (
        <div className="fixed bottom-3 right-3 z-50">
            <Popover open={isOpen} onOpenChange={setIsOpen}>
                <PopoverTrigger asChild>
                    <Badge
                        onClick={() => {
                            togglePopover()
                        }}
                        variant="outline"
                        className={cn(
                            'cursor-pointer border-primary/50 dark:bg-background/80 dark:text-white'
                        )}
                    >
                        <span className="font-bold">Ecoop</span>&nbsp;Beta
                        &nbsp;v{packageJson.version}
                        {isOpen ? (
                            <MdOutlineExpandLess className="size-5" />
                        ) : (
                            <MdExpandMore className="size-5" />
                        )}
                    </Badge>
                </PopoverTrigger>
                <PopoverContent className="w-80 bg-background">
                    <div className="grid gap-4">
                        <div className="space-y-2">
                            <h4 className="font-medium leading-none">
                                Current Version
                            </h4>
                            <p className="text-sm text-muted-foreground">
                                v{packageJson.version}
                            </p>
                        </div>
                        <div className="space-y-2 px-2">
                            <Form {...feedbackForm}>
                                <form
                                    onSubmit={feedbackForm.handleSubmit(
                                        handleFeedBackSubmit
                                    )}
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
                                                            onValueChange={
                                                                field.onChange
                                                            }
                                                            defaultValue={
                                                                field.value
                                                            }
                                                        >
                                                            <SelectTrigger
                                                                className={cn(
                                                                    'rounded-[10px] bg-transparent placeholder:text-[#838383]'
                                                                )}
                                                            >
                                                                <SelectValue placeholder="Choose feedback type" />
                                                            </SelectTrigger>
                                                            <SelectContent>
                                                                {FeedbackType.map(
                                                                    (
                                                                        feedback,
                                                                        index
                                                                    ) => {
                                                                        return (
                                                                            <SelectItem
                                                                                key={
                                                                                    index
                                                                                }
                                                                                value={
                                                                                    feedback.name
                                                                                }
                                                                            >
                                                                                {
                                                                                    feedback.name
                                                                                }
                                                                            </SelectItem>
                                                                        )
                                                                    }
                                                                )}
                                                            </SelectContent>
                                                        </Select>
                                                    </FormControl>
                                                </div>
                                            </FormItem>
                                        )}
                                    />
                                    <FormField
                                        name="message"
                                        control={feedbackForm.control}
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel className="h-[24px] w-full text-[14px]">
                                                    Message
                                                </FormLabel>
                                                <FormControl>
                                                    <Textarea
                                                        {...field}
                                                        id="width"
                                                        placeholder="Leave us some feedback..."
                                                        className={cn(
                                                            'rounded-[10px] bg-transparent placeholder:text-[#838383]'
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
                                    <p className="text-[12px]">
                                        Your feedback helps us identify areas
                                        for improvement.
                                    </p>
                                    <FormErrorMessage
                                        className="text-[12px]"
                                        errorMessage={showFieldError}
                                    />
                                    <Button className={cn('w-full')}>
                                        send
                                    </Button>
                                </form>
                            </Form>
                        </div>
                    </div>
                </PopoverContent>
            </Popover>
        </div>
    )
}
