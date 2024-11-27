import { Input } from '@/components/ui/input'

interface Props {
    value: { from?: number; to?: number }
    onChange: (newValue: { from?: number; to?: number }) => void
}

const NumberRange = ({ value, onChange }: Props) => {
    return (
        <div className="flex items-center gap-x-1">
            <Input
                type="number"
                className="min-w-40"
                value={value.from ?? ''}
                onChange={(inpt) =>
                    onChange({
                        ...value,
                        from: inpt.target.value as unknown as number,
                    })
                }
                placeholder="Min"
            />
            <Input
                type="number"
                className="min-w-40"
                value={value.to ?? ''}
                onChange={(inpt) =>
                    onChange({
                        ...value,
                        to: inpt.target.value as unknown as number,
                    })
                }
                placeholder="Max"
            />
        </div>
    )
}

export default NumberRange
