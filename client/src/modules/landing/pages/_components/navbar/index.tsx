type Props = {}
import { Link } from '@tanstack/react-router'
import ecoop_logo from '@/assets/images/ecoop_logo.png'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'
import { TbExternalLink } from 'react-icons/tb'
import { ReactNode, useEffect, useState } from 'react'
import { useRouter } from '@tanstack/react-router'

type NavLink = {
    name: string
    path: string
    icon?: ReactNode
}

const navLinks: NavLink[] = [
    {
        name: 'Home',
        path: '/',
    },
    {
        name: 'About',
        path: '/about',
    },
    {
        name: 'Contact',
        path: '/contact',
    },
    {
        name: 'Developers',
        path: '/dev',
        icon: <TbExternalLink />,
    },
]

const NavBar = ({}: Props) => {
    const router = useRouter()
    const currentPath = router.state.location.pathname

    const [currentTab, setCurrentTab] = useState<string>(currentPath)

    useEffect(() => {
        setCurrentTab(currentPath)
    }, [])

    return (
        <>
            <nav className="flex h-[68px] justify-between gap-x-2 px-4 lg:px-16">
                <div className="flex w-fit items-center justify-center">
                    <Link to="/">
                        <img
                            src={ecoop_logo}
                            className="h-auto w-[46px] cursor-pointer"
                            alt="Logo"
                        />
                    </Link>
                </div>
                <div className="flex h-full items-center justify-center text-[15px]">
                    <div className="flex space-x-[14px] font-medium">
                        {navLinks.map((link, index) => (
                            <div
                                key={index}
                                className="relative flex space-x-1"
                            >
                                <Link
                                    className="scale-effects nav-links"
                                    onClick={() => setCurrentTab(link.path)}
                                    to={link.path}
                                >
                                    {link.name}
                                </Link>
                                <div className="self-center">{link.icon}</div>
                                {currentTab === link.path && (
                                    <div className="absolute -bottom-2 h-[5px] w-[20px] rounded-full bg-black"></div>
                                )}
                            </div>
                        ))}
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
