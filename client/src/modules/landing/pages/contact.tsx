import z from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Link } from '@tanstack/react-router'
import { useMutation } from '@tanstack/react-query'
import { serverRequestErrExtractor } from '@/helpers'
import UseCooldown from '@/hooks/use-cooldown'
import { contactFormSchema } from '@/modules/landing/validations/contact-form'
import { cn, withCatchAsync } from '@/lib/utils'

import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import ContactService from '@/horizon-corp/server/common/ContactService'
import { toast } from 'sonner'

import {
    LoadingCircleIcon,
    AiOutlineMessageIcon,
    MdOutlineEmailIcon,
    IoCallOutlineIcon,
    CgFacebookIcon,
} from '@/components/icons'
import { PhoneInput } from '@/components/contact-input/contact-input'

type TContact = z.infer<typeof contactFormSchema>

const ContactPage = () => {
    const defaultValues = {
        firstName: '',
        lastName: '',
        email: '',
        contactNumber: '',
        description: '',
    }

    const form = useForm<TContact>({
        resolver: zodResolver(contactFormSchema),
        reValidateMode: 'onChange',
        mode: 'onChange',
        defaultValues,
    })
    const { startCooldown } = UseCooldown({
        cooldownDuration: 12,
        counterInterval: 1000,
    })

    const { mutate: sendContactMessage, isPending } = useMutation<
        void,
        string,
        TContact
    >({
        mutationKey: ['send-contact-message'],
        mutationFn: async (contact: TContact) => {
            if (!contact) return

            const [error] = await withCatchAsync(ContactService.create(contact))

            if (error) {
                const errorMessage = serverRequestErrExtractor({ error })
                toast.error(errorMessage)
                return
            }

            startCooldown()
            form.reset()
            toast.success(`Message was Already Sent!`)
        },
    })

    const onSubmitContactForm = async (data: TContact) => {
        const parsedData = contactFormSchema.parse(data)
        sendContactMessage(parsedData)
    }

    const showFieldError = Object.values(form.formState.errors)[0]?.message

    return (
        <div className="flex justify-center px-6 py-5 font-inter sm:px-8 lg:px-[60px] lg:py-10 xl:px-[124px]">
            <div className="mt-3 flex max-w-[1300px] flex-col items-center justify-center space-y-4 md:mt-5 lg:mt-16 lg:space-y-7 xl:space-y-10">
                <h1 className="max-w-[1100px] text-center text-[min(64px,5.5vw)] font-bold">
                    Contact our Team
                </h1>
                <h2 className="max-w-[1100px] text-center text-[min(24px,3.5vw)] font-medium">
                    Got any questions about the product or scaling on our
                    platform? We’re here to help. Chat our friendly team 24/7
                    and get onboard in less than 5 minutes.
                </h2>
                <div className="flex w-full flex-col justify-evenly md:flex-row">
                    <Form {...form}>
                        <form
                            onSubmit={form.handleSubmit(onSubmitContactForm)}
                            className="max-w-[456px] space-y-2"
                        >
                            <div className="flex flex-col gap-[26px] md:flex-row">
                                <FormField
                                    name="firstName"
                                    control={form.control}
                                    render={({ field }) => (
                                        <FormItem className="">
                                            <div className="flex flex-col justify-start">
                                                <FormLabel
                                                    htmlFor={field.name}
                                                    className="h-[38px] w-full max-w-[90px] text-[15px] font-semibold"
                                                >
                                                    First Name
                                                </FormLabel>
                                                <FormControl>
                                                    <div className="flex-1 space-y-2">
                                                        <Input
                                                            id={field.name}
                                                            className={cn(
                                                                'rounded-[10px] border border-[#4D4C4C] bg-transparent placeholder:text-[#838383]'
                                                            )}
                                                            placeholder="First Name"
                                                            {...field}
                                                        />
                                                    </div>
                                                </FormControl>
                                            </div>
                                        </FormItem>
                                    )}
                                />
                                <FormField
                                    name="lastName"
                                    control={form.control}
                                    render={({ field }) => (
                                        <FormItem className="">
                                            <div className="flex flex-col justify-start">
                                                <FormLabel
                                                    htmlFor={field.name}
                                                    className="h-[38px] w-full max-w-[90px] text-[15px] font-semibold"
                                                >
                                                    Last Name
                                                </FormLabel>
                                                <FormControl>
                                                    <div className="flex-1 space-y-2">
                                                        <Input
                                                            id={field.name}
                                                            className={cn(
                                                                'rounded-[10px] border border-[#4D4C4C] bg-transparent placeholder:text-[#838383]'
                                                            )}
                                                            placeholder="Last Name"
                                                            {...field}
                                                        />
                                                    </div>
                                                </FormControl>
                                            </div>
                                        </FormItem>
                                    )}
                                />
                            </div>
                            <FormField
                                name="email"
                                control={form.control}
                                render={({ field }) => (
                                    <FormItem className="">
                                        <div className="flex flex-col justify-start">
                                            <FormLabel
                                                htmlFor={field.name}
                                                className="h-[38px] w-full max-w-[90px] text-[15px] font-semibold"
                                            >
                                                Email
                                            </FormLabel>
                                            <FormControl>
                                                <div className="flex-1 space-y-2">
                                                    <Input
                                                        id={field.name}
                                                        className={cn(
                                                            'rounded-[10px] border border-[#4D4C4C] bg-transparent placeholder:text-[#838383]'
                                                        )}
                                                        placeholder="you@company.com"
                                                        {...field}
                                                        autoComplete="email"
                                                    />
                                                </div>
                                            </FormControl>
                                        </div>
                                    </FormItem>
                                )}
                            />
                            <FormField
                                name="contactNumber"
                                control={form.control}
                                render={({ field }) => (
                                    <FormItem className="">
                                        <div className="flex flex-col justify-start">
                                            <FormLabel
                                                htmlFor={field.name}
                                                className="h-[38px] w-full max-w-[90px] text-[15px] font-semibold"
                                            >
                                                Phone
                                            </FormLabel>
                                            <FormControl>
                                                <div className="flex-1 space-y-2">
                                                    <PhoneInput
                                                        placeholder="Enter a phone number"
                                                        className={cn('rounded-[10px] border border-[#4D4C4C] !bg-transparent placeholder:text-[#838383]')}
                                                        {...field}
                                                    />
                                                </div>
                                            </FormControl>
                                        </div>
                                    </FormItem>
                                )}
                            />

                            <FormField
                                name="description"
                                control={form.control}
                                render={({ field }) => (
                                    <FormItem className="">
                                        <div className="flex flex-col justify-start">
                                            <FormLabel
                                                htmlFor={field.name}
                                                className="h-[38px] w-full max-w-[90px] text-[15px] font-semibold"
                                            >
                                                Message
                                            </FormLabel>
                                            <FormControl>
                                                <div className="flex-1 space-y-2">
                                                    <Textarea
                                                        id={field.name}
                                                        className={cn(
                                                            'max-h-[122px] min-h-[122px] rounded-[10px] border border-[#4D4C4C] bg-transparent placeholder:text-[#838383]'
                                                        )}
                                                        placeholder="Leave us message..."
                                                        {...field}
                                                    />
                                                </div>
                                            </FormControl>
                                        </div>
                                    </FormItem>
                                )}
                            />
                            {showFieldError && (
                                <span className="mt-2 rounded-md bg-destructive/10 py-2 text-center text-sm text-destructive">
                                    {showFieldError}
                                </span>
                            )}
                            <div className="bg- flex flex-col space-y-2">
                                <Button
                                    disabled={isPending}
                                    type="submit"
                                    className="mt-6 bg-[#34C759] hover:bg-[#38b558]"
                                >
                                    {isPending ? (
                                        <LoadingCircleIcon className="animate-spin" />
                                    ) : (
                                        'Send Message'
                                    )}
                                </Button>
                            </div>
                        </form>
                    </Form>

                    <div className="flex h-full flex-col justify-evenly">
                        <div className="space-y-2">
                            <div>
                                <h1 className="text-[18px] font-bold">
                                    Chat with us
                                </h1>
                                <p className="text-sm font-semibold text-[#686868]">
                                    Speak to our friendly team via live chat.
                                </p>
                            </div>
                            <div className="flex space-x-2">
                                <AiOutlineMessageIcon className="self-center" />
                                <Link className="text-sm font-semibold" to="/">
                                    start a live chat
                                </Link>
                            </div>
                            <div className="flex space-x-2">
                                <MdOutlineEmailIcon className="self-center" />
                                <Link className="text-sm font-semibold" to="/">
                                    shoot us an email
                                </Link>
                            </div>
                            <div className="flex space-x-2">
                                <CgFacebookIcon className="self-center" />
                                <Link className="text-sm font-semibold" to="/">
                                    Message us on Facebook
                                </Link>
                            </div>
                        </div>
                        <div className="space-y-2">
                            <div>
                                <h1 className="text-[18px] font-bold">
                                    Call us
                                </h1>
                                <p className="text-sm font-semibold text-[#686868]">
                                    Call our team on Mon-Fri 8am to 5pm.
                                </p>
                            </div>
                            <div className="flex space-x-2">
                                <IoCallOutlineIcon className="self-center" />
                                <Link
                                    className="text-sm font-semibold underline"
                                    to="/"
                                >
                                    +639321784911
                                </Link>
                            </div>
                        </div>
                        <div className="space-y-2">
                            <div>
                                <h1 className="text-[18px] font-bold">
                                    Visit Us
                                </h1>
                                <p className="text-sm font-semibold text-[#686868]">
                                    Reach our team on Mon-Fri 8am to 5pm.
                                </p>
                            </div>
                            <div className="flex space-x-2">
                                <p className="text-sm font-semibold underline">
                                    Bagong Street, Norway VC 3066
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default ContactPage
