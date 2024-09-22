import { Link } from '@tanstack/react-router'

interface Props {}

const AuthFooter = (_props: Props) => {
    return (
        <footer className="flex flex-col text-sm items-center gap-y-2 py-8">
            <p className="text-foreground/40">
                e-Coop @ All Rights Reserved 2024
            </p>
            <div className="flex w-full font-inter justify-center gap-x-3 text-green-400">
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
