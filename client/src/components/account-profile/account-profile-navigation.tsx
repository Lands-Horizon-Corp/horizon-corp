import { cn } from '@/lib'
import { useEffect } from 'react'
import { useRouter, useRouterState } from '@tanstack/react-router'

const PATHS: { hash: string; name: string }[] = [
    { hash: 'account-settings', name: 'Account Settings' },
    { hash: 'security', name: 'Security' },
    { hash: 'mode-of-payments', name: 'Mode of Payments' },
    { hash: 'transactions', name: 'Transactions' },
]

const AccountProfileNavigation = () => {
    const router = useRouter()

    const routeHash = useRouterState({
        select: (state) => state.location.hash,
    })

    useEffect(() => {
        if (!router.history.location.hash)
            router.navigate({
                hash: PATHS[0].hash,
            })
    }, [router])

    return (
        <div className="flex items-center gap-x-1 rounded-xl bg-secondary p-1 text-sm">
            {PATHS.map((hashPath) => (
                <span
                    key={hashPath.hash}
                    className={cn(
                        'cursor-pointer rounded-xl p-2 text-secondary-foreground/70 duration-300 ease-in-out hover:text-secondary-foreground',
                        routeHash === hashPath.hash && 'bg-background'
                    )}
                    onClick={() => router.navigate({ hash: hashPath.hash })}
                >
                    {hashPath.name}
                </span>
            ))}
        </div>
    )
}

export default AccountProfileNavigation
