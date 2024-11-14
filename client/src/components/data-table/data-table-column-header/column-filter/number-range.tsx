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
                className="w-full"
                value={value.from ?? ''}
                onChange={(inpt) =>
                    onChange({ ...value, from: inpt.target.value as any })
                }
                placeholder="Min"
            />
            <Input
                type="number"
                className="w-full"
                value={value.to ?? ''}
                onChange={(inpt) =>
                    onChange({ ...value, to: inpt.target.value as any })
                }
                placeholder="Max"
            />
        </div>
    )
}

export default NumberRange
