import { Link } from '@tanstack/react-router'

interface Props {}

const AuthFooter = (_props: Props) => {
    return (
        <footer className="flex flex-col items-center gap-y-2 py-8 text-sm">
            <p className="text-foreground/40">
                e-Coop @ All Rights Reserved 2024
            </p>
            <div className="flex w-full justify-center gap-x-3 font-inter text-green-400">
                <Link to="/contact">Contact Us</Link>
                <span>・</span>
                <p>Terms of Use</p>
                <span>・</span>
                <p>Privacy Policy</p>
            </div>
        </footer>
    )
}

export default AuthFooter
