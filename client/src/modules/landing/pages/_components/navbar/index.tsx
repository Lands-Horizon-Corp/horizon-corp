type Props = {}
import { Link } from '@tanstack/react-router'
import ecoop_logo from '@/assets/images/ecoop_logo.png'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'

const NavBar = ({}: Props) => {
    return (
        <>
            <nav className="flex h-[68px] justify-between gap-x-2 px-4 lg:px-16">
                <div className="flex w-fit items-center justify-center">
                    <img
                        src={ecoop_logo}
                        className="h-auto w-[46px]"
                        alt="Logo"
                    />
                </div>
                <div className="flex h-full items-center justify-center text-[15px]">
                    <div className="flex space-x-[14px] font-medium">
                        <Link className="" to="/">
                            <p className="scale-effects nav-links">Home</p>
                        </Link>
                        <Link to="/about">
                            <p className="scale-effects nav-links">About</p>
                        </Link>
                        <Link to="/contact">
                            <p className="scale-effects nav-links">Contact</p>
                        </Link>
                    </div>
                </div>
                <div className="flex h-full items-center justify-center">
                    <div className="space-x-4">
                        <Button
                            className={cn(
                                'scale-effects h-[30px] rounded-full bg-green-500 text-white hover:bg-green-500'
                            )}
                        >
                            <Link to="/auth/sign-in">Sign-In</Link>
                        </Button>
                        <Button
                            className={cn(
                                'scale-effects h-[30px] rounded-full border border-black bg-transparent text-black hover:bg-transparent'
                            )}
                        >
                            <Link to="/auth/sign-up">Sign-Up</Link>
                        </Button>
                    </div>
                </div>
            </nav>
        </>
    )
}

export default NavBar
