import { Button } from '../ui/button'
import { SidebarLeftIcon } from '../icons'

const AppSidebarToggle = () => {
    return (
        <Button
            variant="secondary"
            className='size-fit p-2'
        >
            <SidebarLeftIcon />
        </Button>
    )
}

export default AppSidebarToggle
