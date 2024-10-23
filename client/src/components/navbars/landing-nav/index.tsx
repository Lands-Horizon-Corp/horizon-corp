// Dependencies
import { ReactNode } from 'react'
import { TbExternalLink } from 'react-icons/tb'
import { useLocation, Link } from '@tanstack/react-router'

// Components
import RootNav from '@/components/navbars/root-nav'
import { ThemeToggleMenu } from '@/components/theme-toggle'
import NavAuthContents from '@/components/navbars/auth-nav/nav-auth-contents'

import { cn } from '@/lib/utils'
import { EnvironmentManager } from '@/manager'

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
        path: import.meta.env.VITE_CLIENT_DOCUMENT_URL,
        icon: <TbExternalLink />,
    },
]

const NavBar = () => {
    const pathName = useLocation({
        select: (location) => location.pathname,
    })

    return (
        <RootNav
            midContents={navLinks.map((link, index) => {
                const isCurrentTab = pathName === link.path
                const isExternalLink = link.path?.charAt(0) !== '/'

                return (
                    <div key={index} className="relative flex space-x-1">
                        {isExternalLink ? (
                            <a
                                href={link.path}
                                target="_blank"
                                rel="noopener noreferrer"
                                className={cn(
                                    'scale-effects nav-links hidden items-center gap-x-2 font-normal sm:flex',
                                    isCurrentTab && 'font-bold'
                                )}
                            >
                                {link.name}
                                <div className="self-center">{link.icon}</div>
                            </a>
                        ) : (
                            <Link
                                className={cn(
                                    'scale-effects nav-links hidden items-center gap-x-2 font-normal sm:flex',
                                    isCurrentTab && 'font-bold'
                                )}
                                to={link.path}
                            >
                                {link.name}
                                <div className="self-center">{link.icon}</div>
                            </Link>
                        )}
                        {isCurrentTab && (
                            <div className="absolute -bottom-2 hidden h-[5px] w-[20px] rounded-full bg-green-500 sm:block"></div>
                        )}
                    </div>
                )
            })}
            rightContents={
                <>
                    <NavAuthContents />
                    <ThemeToggleMenu />
                </>
            }
        />
    )
}

export default NavBar
