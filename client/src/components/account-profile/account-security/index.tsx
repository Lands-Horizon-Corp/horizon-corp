import { useRouterState } from '@tanstack/react-router'

import { Separator } from '@/components/ui/separator'

import UsernameOption from './username-option'

import { useUserAuthStore } from '@/store/user-auth-store'

const AccountSecurity = () => {
    const {} = useUserAuthStore()

    const hash = useRouterState({
        select: ({ location }) => location.hash,
    })

    if (hash !== 'security') return null

    return (
        <div>
            <div className="flex items-center justify-between">
                <p className="text-lg font-medium">Security</p>
            </div>
            <p className="mt-1 text-xs text-foreground/60">
                Manage or Update your account security
            </p>
            <Separator className="my-2 sm:my-4" />
            <UsernameOption />
        </div>
    )
}

export default AccountSecurity
