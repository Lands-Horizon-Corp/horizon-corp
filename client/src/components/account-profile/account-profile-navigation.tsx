import { cn } from '@/lib'
import { useEffect } from 'react'
import { IconType } from 'react-icons/lib'
import { useRouter, useRouterState } from '@tanstack/react-router'

import {
    AccountSettingIcon,
    PaymentsIcon,
    ShieldIcon,
    TransactionListIcon,
} from '@/components/icons'

const PATHS: { hash: string; name: string; Icon: IconType }[] = [
    {
        hash: 'account-settings',
        name: 'Account Settings',
        Icon: AccountSettingIcon,
    },
    { hash: 'security', name: 'Security', Icon: ShieldIcon },
    { hash: 'mode-of-payments', name: 'Mode of Payments', Icon: PaymentsIcon },
    { hash: 'transactions', name: 'Transactions', Icon: TransactionListIcon },
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
        <div className="mx-auto flex w-fit max-w-full justify-center gap-x-1 overflow-x-scroll rounded-xl bg-secondary p-1 text-sm sm:mx-0 sm:w-full sm:justify-start">
            {PATHS.map((hashPath) => (
                <button
                    key={hashPath.hash}
                    className={cn(
                        'flex w-fit cursor-pointer items-center gap-x-1 whitespace-nowrap rounded-lg p-2 text-secondary-foreground/70 duration-300 ease-in-out hover:text-secondary-foreground',
                        routeHash === hashPath.hash && 'bg-background'
                    )}
                    onClick={() => router.navigate({ hash: hashPath.hash })}
                >
                    <hashPath.Icon className="size-6 sm:size-4" />
                    <span className="hidden sm:block">{hashPath.name}</span>
                </button>
            ))}
        </div>
    )
}

export default AccountProfileNavigation
