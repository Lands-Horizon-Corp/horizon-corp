// Dependencies
import { TbExternalLink } from 'react-icons/tb'
import { ReactNode, useEffect, useState } from 'react'
import { useRouter } from '@tanstack/react-router'
import { Link } from '@tanstack/react-router'

// Components
import RootNav from '@/components/navbars/root-nav'
import { ThemeToggleMenu } from '@/components/theme-toggle'
import NavAuthContents from '@/components/navbars/nav-auth-contents'
import { cn } from '@/lib/utils'

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
    const router = useRouter()
    const currentPath = router.state.location.pathname

    const [currentTab, setCurrentTab] = useState(currentPath)

    useEffect(() => {
        setCurrentTab(currentPath)
    }, [currentPath])

    return (
        <RootNav
            midContents={navLinks.map((link, index) => {
                const isCurrentTab = currentTab === link.path
                const isExternalLink =
                    link.path.startsWith('http://') ||
                    link.path.startsWith('https://')

                return (
                    <div key={index} className="relative flex space-x-1">
                        {isExternalLink ? (
                            <a
                                href={link.path}
                                target="_blank"
                                rel="noopener noreferrer"
                                className={cn(
                                    'scale-effects nav-links flex items-center gap-x-2 font-normal',
                                    isCurrentTab && 'font-bold'
                                )}
                            >
                                {link.name}
                                <div className="self-center">{link.icon}</div>
                            </a>
                        ) : (
                            <Link
                                className={cn(
                                    'scale-effects nav-links flex items-center gap-x-2 font-normal',
                                    isCurrentTab && 'font-bold'
                                )}
                                onClick={() => setCurrentTab(link.path)}
                                to={link.path}
                            >
                                {link.name}
                                <div className="self-center">{link.icon}</div>
                            </Link>
                        )}

                        {isCurrentTab && (
                            <div className="absolute -bottom-2 h-[5px] w-[20px] rounded-full bg-green-500"></div>
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
