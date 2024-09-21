import { Link, Outlet } from '@tanstack/react-router'

interface Props {}

const PublicLayout = (_props: Props) => {
    return (
        <div>
            <nav className="flex gap-x-2">
                <Link to="/">Home</Link>
                <Link to="/about">About</Link>
                <Link to="/contact">Contact</Link>
                <Link to="/auth/sign-in">Sign-In</Link>
                <Link to="/auth/sign-up">Sign-Up</Link>
            </nav>
            <Outlet />
        </div>
    )
}

export default PublicLayout
