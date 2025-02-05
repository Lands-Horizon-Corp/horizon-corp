import { UseFormReturn, FieldValues, Path } from 'react-hook-form'

import {
    FormItem,
    FormLabel,
    FormField,
    FormControl,
    FormMessage,
    FormDescription,
    FormHidableItem,
} from '@/components/ui/form'
import { Input, InputProps } from '@/components/ui/input'

import { cn } from '@/lib/utils'

type FormLabeledInputFieldProps<T extends FieldValues> = {
    name: Path<T>
    form: UseFormReturn<T>
    hiddenFields?: Array<Path<T>>

    isDisabled?: boolean
    showMessage?: boolean

    label?: string
    placeholder?: string
    description?: string
    inputProps?: Omit<InputProps, 'id' | 'disabled' | 'placeholder'>

    className?: string
    labelClassName?: string
    messageClassName?: string
    descriptionClassName?: string
}

export const FormLabeledInputField = <T extends FieldValues>({
    name,
    form,
    label,
    className,
    inputProps,
    isDisabled,
    hiddenFields,
    placeholder,
    description,
    labelClassName,
    messageClassName,
    showMessage = true,
    descriptionClassName,
}: FormLabeledInputFieldProps<T>) => {


    return (
        <FormField<T, Path<T>>
            name={name}
            control={form.control}
            render={({ field }) => (
                <FormHidableItem<T>
                    field={field.name}
                    hiddenFields={hiddenFields}
                >
                    <FormItem
                        className={cn('col-span-1 w-full space-y-1', className)}
                    >
                        {label && (
                            <FormLabel
                                htmlFor={field.name}
                                className={cn(labelClassName)}
                            >
                                {label}
                            </FormLabel>
                        )}
                        {description && (
                            <FormDescription
                                className={cn(descriptionClassName)}
                            >
                                {description}
                            </FormDescription>
                        )}
                        <FormControl>
                            <Input
                                {...field}
                                {...inputProps}
                                id={field.name}
                                disabled={isDisabled}
                                placeholder={placeholder}
                            />
                        </FormControl>
                        {showMessage && (
                            <FormMessage
                                className={cn('text-xs', messageClassName)}
                            />
                        )}
                    </FormItem>
                </FormHidableItem>
            )}
        />
    )
}
