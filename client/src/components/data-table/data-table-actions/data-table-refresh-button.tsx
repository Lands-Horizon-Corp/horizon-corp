import { ReloadIcon } from '@/components/icons'
import { Button } from '@/components/ui/button'
import LoadingSpinner from '@/components/spinners/loading-spinner'

interface Props {
    isLoading?: boolean
    onClick: () => void
}

const DataTableRefreshButton = ({ isLoading, onClick }: Props) => {
    return (
        <Button
            disabled={isLoading}
            onClick={onClick}
            className=""
            variant="secondary"
            size="icon"
        >
            {isLoading ? <LoadingSpinner /> : <ReloadIcon />}
        </Button>
    )
}

export default DataTableRefreshButton
