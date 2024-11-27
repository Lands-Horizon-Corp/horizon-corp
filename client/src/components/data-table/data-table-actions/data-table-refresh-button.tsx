import { ReloadIcon } from '@/components/icons'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

interface Props {
    isLoading?: boolean
    onClick: () => void
    className? : string
}

const DataTableRefreshButton = ({ isLoading, onClick, className }: Props) => {
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

export default DataTableRefreshButton
