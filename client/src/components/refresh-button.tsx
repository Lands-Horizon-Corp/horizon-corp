import { ReloadIcon } from '@/components/icons'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

export interface IRefreshButtonProps {
    isLoading?: boolean
    onClick: () => void
    className?: string
}

const RefreshButton = ({
    isLoading,
    onClick,
    className,
}: IRefreshButtonProps) => {
    return (
        <Button
            size="icon"
            onClick={onClick}
            variant="secondary"
            disabled={isLoading}
            className={className}
        >
            {isLoading ? <LoadingSpinner /> : <ReloadIcon />}
        </Button>
    )
}

export default RefreshButton
