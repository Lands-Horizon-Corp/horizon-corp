import React, { useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'

import { Input } from '@/components/ui/input'
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group'
import { Checkbox } from '@/components/ui/checkbox'
import { Button } from '@/components/ui/button'
import {
    Select,
    SelectTrigger,
    SelectContent,
    SelectItem,
    SelectValue,
} from '@/components/ui/select'
import {
    Command,
    CommandInput,
    CommandGroup,
    CommandItem,
    CommandList,
    CommandEmpty,
} from '@/components/ui/command'

import {
    Popover,
    PopoverTrigger,
    PopoverContent,
} from '@/components/ui/popover'
import { cn } from '@/lib'
import { ChevronsUpDown, Check } from 'lucide-react'
import { ReportProps } from '@/components/reports/config'
import SignatureForm, { Position, SignedBy } from '@/components/reports/report-signature'

// Define Zod schema
type SignatureFormData = Record<keyof typeof SignedBy, { name?: string; position?: Position }>

const dailyCollDetailSchema = z.object({
    title: z.string().min(1, 'Title is required'),
    startDate: z.string().min(1, 'Start Date is required'),
    endDate: z.string().min(1, 'End Date is required'),
    tellerFrom: z.string().optional(),
    tellerTo: z.string().optional(),
    account: z.string().min(1, 'Account is required'),
    batchNo: z.string().min(1, 'Batch No is required'),
    grouping: z.enum(['by-teller', 'no-grouping']),
    optionType: z.enum(['option-1', 'option-2']),
    format: z.enum([
        'format-1',
        'format-2',
        'format-3',
        'atm-format',
        'short-over',
        'single-col',
    ]),
    type: z.enum(['standard', 'micro-finance', 'all']),
    printSummary: z.boolean(),
    sundriesSeparate: z.boolean(),
    genericTextPrinting: z.boolean(),
    signature: z.record(z.object({
        name: z.string().optional(),
        position: z.nativeEnum(Position).optional()
    })).optional()
})

const accounts = [
    { code: '160.00', description: 'ACCOUNT RECEIVABLES' },
    { code: '305.00', description: 'ACCOUNTS PAYABLE' },
    { code: '310.00', description: 'ACCRUED EXPENSES' },
    { code: '195.15', description: "ACCUM. DEPR'N-FUR.&FIX." },
    { code: '192.50', description: "ACCUM.DEPR'N-BUILDING IMP" },
    { code: '196.50', description: "ACCUM.DEPR'N-OFFICE EQUIPMENT" },
    { code: '198.00', description: 'ACCUMULATED DEP. - LODGING' },
    { code: '192.01', description: "ACCUMULATED DEPR'N - BUILDING" },
    { code: '252.10', description: 'ACCUMULATED DEPRECIATION-COM.' },
    { code: '530.00', description: 'ADMINISTRATIVE COST' },
    { code: '161.10', description: 'ADVANCES TO OFF AND EMP' },
    { code: '565.00', description: 'AFFILIATION FEES' },
    { code: '189.05', description: 'ALILEM CONS.MKTG.COOPERATIVE' },
    { code: '373.00', description: 'ALILEM MUNICIPAL ASSISTANCE' },
    { code: '154.00', description: 'ALLOW FOR PROBABLE LOAN LOSS' },
    { code: '532.00', description: 'ANNUAL DUES' },
    { code: '150.90', description: 'APPLIANCE LOAN' },
    { code: '253.00', description: 'ASSET ACQUIRED' },
    { code: '534.00', description: 'BANK CHARGES' },
]

const tellers = [
    { code: '00', name: 'ECOOP' },
    { code: '01', name: 'PRECILA BALTAZAR' },
    { code: '02', name: 'VENUS JOY PILAYAN' },
    { code: '03', name: 'EMMANUEL RUBANG' },
    { code: '04', name: 'CLAUDETTE LIYO' },
    { code: '06', name: 'MARIO ARZABAL' },
    { code: '07', name: 'BENILDA A. KILIP' },
    { code: '10', name: 'IMELDA LUBRIN' },
    { code: '11', name: 'ROJAN' },
    { code: '12', name: 'JERBEE' },
]

export const formatOptions = [
    { value: 'format-1', label: 'Format #1' },
    { value: 'format-2', label: 'Format #2' },
    { value: 'format-3', label: 'Format #3' },
    { value: 'atm-format', label: 'ATM Format' },
    { value: 'short-over', label: 'Short/Over' },
    { value: 'single-col', label: 'Single Col.' },
]

// Define TypeScript type from Zod schema
type DailyCollDetailSchema = z.infer<typeof dailyCollDetailSchema>

const DailyCollDetail: React.FC<ReportProps<DailyCollDetailSchema>> = ({
    onSubmit,
    onClose,
}) => {
    const {
        register,
        handleSubmit,
        setValue,
        watch,
        formState: { errors },
    } = useForm<DailyCollDetailSchema>({
        resolver: zodResolver(dailyCollDetailSchema),
        defaultValues: {
            title: 'COLLECTION REPORT',
            startDate: '',
            endDate: '',
            tellerFrom: '',
            tellerTo: '',
            account: 'ALL',
            batchNo: '1',
            grouping: 'by-teller',
            optionType: 'option-1',
            format: 'format-1',
            type: 'all',
            printSummary: false,
            sundriesSeparate: false,
            genericTextPrinting: true,
            signature: {} as SignatureFormData,
        },
    })

    const [open, setOpen] = useState(false)

    const selectedAccount = accounts.find(
        (acc) => acc.code === watch('account')
    )

    const [isSignatureFormOpen, setIsSignatureFormOpen] = useState(false)
    const [_, setSignatureData] = useState({})

    const handleOpen = () => {
        setIsSignatureFormOpen(true)
    }

    const handleClose = () => {
        setIsSignatureFormOpen(false)
    }

    const handleSubmitSignature = (data: any) => {
        console.log("Submitted Signature Data:", data);
        setSignatureData(data)
        setValue('signature', data)
        handleClose()
    }

    return (
        <form
            onSubmit={handleSubmit(onSubmit)}
            className="w-full space-y-4 text-xs"
        >
            <SignatureForm
                open={isSignatureFormOpen}
                onClose={handleClose}
                onSubmit={handleSubmitSignature}
            />

            <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
                <div className="col-span-2">
                    <label className="block text-xs font-medium text-gray-700">
                        TITLE
                    </label>
                    <Input {...register('title')} />
                    {errors.title && (
                        <p className="text-xs text-red-500">
                            {errors.title.message}
                        </p>
                    )}
                </div>
                <Button
                    className="col-span-2"
                    variant={'outline'}
                    onClick={handleOpen}
                >
                    Sign
                </Button>
                <div>
                    <label className="block text-xs font-medium text-gray-700">
                        Start Date
                    </label>
                    <Input type="date" {...register('startDate')} />
                    {errors.startDate && (
                        <p className="text-xs text-red-500">
                            {errors.startDate.message}
                        </p>
                    )}
                </div>

                <div>
                    <label className="block text-xs font-medium text-gray-700">
                        End Date
                    </label>
                    <Input type="date" {...register('endDate')} />
                    {errors.endDate && (
                        <p className="text-xs text-red-500">
                            {errors.endDate.message}
                        </p>
                    )}
                </div>

                <div>
                    <label className="block text-xs font-medium text-gray-700">
                        Teller From
                    </label>
                    <Select
                        onValueChange={(value) => setValue('tellerFrom', value)}
                    >
                        <SelectTrigger>
                            <SelectValue placeholder="Select Teller" />
                        </SelectTrigger>
                        <SelectContent>
                            {tellers.map((teller) => (
                                <SelectItem
                                    key={teller.code}
                                    value={teller.code}
                                >
                                    {teller.name}
                                </SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                    {errors.tellerFrom && (
                        <p className="text-xs text-red-500">
                            {errors.tellerFrom.message}
                        </p>
                    )}
                </div>

                {/* TELLER TO (Dropdown) */}
                <div>
                    <label className="block text-xs font-medium text-gray-700">
                        Teller To
                    </label>
                    <Select
                        onValueChange={(value) => setValue('tellerTo', value)}
                    >
                        <SelectTrigger>
                            <SelectValue placeholder="Select Teller" />
                        </SelectTrigger>
                        <SelectContent>
                            {tellers.map((teller) => (
                                <SelectItem
                                    key={teller.code}
                                    value={teller.code}
                                >
                                    {teller.name}
                                </SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                    {errors.tellerTo && (
                        <p className="text-xs text-red-500">
                            {errors.tellerTo.message}
                        </p>
                    )}
                </div>

                <div className="col-span-2">
                    <label className="block text-xs font-medium text-gray-700">
                        Account
                    </label>
                    <Popover open={open} onOpenChange={setOpen}>
                        <PopoverTrigger asChild>
                            <Button
                                variant="outline"
                                role="combobox"
                                aria-expanded={open}
                                className="w-full justify-between"
                            >
                                {selectedAccount
                                    ? `${selectedAccount.code} - ${selectedAccount.description}`
                                    : 'Select Account'}
                                <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                            </Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-full p-0">
                            <Command>
                                <CommandInput placeholder="Search account..." />
                                <CommandList>
                                    <CommandEmpty>
                                        No account found.
                                    </CommandEmpty>
                                    <CommandGroup>
                                        {accounts.map((account) => (
                                            <CommandItem
                                                key={account.code}
                                                value={account.code}
                                                onSelect={(currentValue) => {
                                                    setValue(
                                                        'account',
                                                        currentValue
                                                    )
                                                    setOpen(false)
                                                }}
                                            >
                                                <Check
                                                    className={cn(
                                                        'mr-2 h-4 w-4',
                                                        watch('account') ===
                                                            account.code
                                                            ? 'opacity-100'
                                                            : 'opacity-0'
                                                    )}
                                                />
                                                {account.code} -{' '}
                                                {account.description}
                                            </CommandItem>
                                        ))}
                                    </CommandGroup>
                                </CommandList>
                            </Command>
                        </PopoverContent>
                    </Popover>
                    {errors.account && (
                        <p className="text-xs text-red-500">
                            {errors.account.message}
                        </p>
                    )}
                </div>

                <div>
                    <label className="block text-xs font-medium text-gray-700">
                        Batch No
                    </label>
                    <Input {...register('batchNo')} />
                    {errors.batchNo && (
                        <p className="text-xs text-red-500">
                            {errors.batchNo.message}
                        </p>
                    )}
                </div>
                <div>
                    <label className="block text-xs font-medium text-gray-700">
                        Format
                    </label>
                    <Select
                        onValueChange={(value) =>
                            setValue(
                                'format',
                                value as DailyCollDetailSchema['format']
                            )
                        }
                    >
                        <SelectTrigger>
                            <SelectValue placeholder="Select format" />
                        </SelectTrigger>
                        <SelectContent>
                            {formatOptions.map((format) => (
                                <SelectItem
                                    key={format.value}
                                    value={format.value}
                                >
                                    {format.label}
                                </SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                    {errors.format && (
                        <p className="text-xs text-red-500">
                            {errors.format.message}
                        </p>
                    )}
                </div>
            </div>

            <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
                <div className="space-y-2">
                    <label className="font-medium">Groupings</label>
                    <RadioGroup name="grouping" defaultValue="by-teller">
                        <label className="flex items-center gap-2">
                            <RadioGroupItem
                                value="by-teller"
                                {...register('grouping')}
                            />
                            By Teller
                        </label>
                        <label className="flex items-center gap-2">
                            <RadioGroupItem
                                value="no-grouping"
                                {...register('grouping')}
                            />
                            No Grouping
                        </label>
                    </RadioGroup>
                </div>
                <div className="space-y-2">
                    <label className="font-medium">Option Type</label>
                    <RadioGroup name="optionType" defaultValue="option-1">
                        <label className="flex items-center gap-2">
                            <RadioGroupItem
                                value="option-1"
                                {...register('optionType')}
                            />
                            Option #1
                        </label>
                        <label className="flex items-center gap-2">
                            <RadioGroupItem
                                value="option-2"
                                {...register('optionType')}
                            />
                            Option #2
                        </label>
                    </RadioGroup>
                </div>

                <div className="space-y-2">
                    <label className="font-medium">Type</label>
                    <RadioGroup name="type" defaultValue="all">
                        <label className="flex items-center gap-2">
                            <RadioGroupItem
                                value="standard"
                                {...register('type')}
                            />
                            Standard
                        </label>
                        <label className="flex items-center gap-2">
                            <RadioGroupItem
                                value="micro-finance"
                                {...register('type')}
                            />
                            Micro Finance
                        </label>
                        <label className="flex items-center gap-2">
                            <RadioGroupItem value="all" {...register('type')} />
                            ALL
                        </label>
                    </RadioGroup>
                </div>
            </div>
            <div className="block space-y-2">
                <div className="flex items-center gap-2">
                    <Checkbox
                        id="printSummary"
                        checked={watch('printSummary')}
                        onCheckedChange={(checked) =>
                            setValue('printSummary', checked as boolean)
                        }
                    />
                    <label htmlFor="printSummary" className="text-xs">
                        Print summary (cash/check)
                    </label>
                </div>

                <div className="flex items-center gap-2">
                    <Checkbox
                        id="sundriesSeparate"
                        checked={watch('sundriesSeparate')}
                        onCheckedChange={(checked) =>
                            setValue('sundriesSeparate', checked as boolean)
                        }
                    />
                    <label htmlFor="sundriesSeparate" className="text-xs">
                        Sundries print separate page?
                    </label>
                </div>

                <div className="flex items-center gap-2">
                    <Checkbox
                        id="genericTextPrinting"
                        checked={watch('genericTextPrinting')}
                        onCheckedChange={(checked) =>
                            setValue('genericTextPrinting', checked as boolean)
                        }
                    />
                    <label htmlFor="genericTextPrinting" className="text-xs">
                        Generic Text Printing
                    </label>
                </div>
            </div>
            <div className="flex justify-end">
                <Button variant={'outline'} onClick={onClose}>
                    Cancel
                </Button>
                <Button type="submit" className="">
                    Submit
                </Button>
            </div>
        </form>
    )
}

export default DailyCollDetail
