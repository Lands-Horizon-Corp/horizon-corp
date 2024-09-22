import { Outlet } from '@tanstack/react-router'

interface Props {}

const LandingPage = (_props: Props) => {
    return (
        <div>
            Public
            <Outlet />
        </div>
    )
}

export default LandingPage
