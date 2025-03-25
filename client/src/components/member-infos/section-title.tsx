import { cn } from '@/lib'
import { IconType } from 'react-icons/lib'

interface Props {
    Icon?: IconType
    title: string
    subTitle?: string
}

const SectionTitle = ({ Icon, title, subTitle }: Props) => {
    return (
        <div
            className={cn(
                'flex items-start gap-x-2',
                !subTitle && 'items-center'
            )}
        >
            {Icon && <Icon className="inline size-8 text-foreground/70" />}
            <div className="space-y-2">
                <p>{title}</p>
                {subTitle && (
                    <p className="!mt-0 text-sm text-muted-foreground/60">
                        {subTitle}
                    </p>
                )}
            </div>
        </div>
    )
}

export default SectionTitle
