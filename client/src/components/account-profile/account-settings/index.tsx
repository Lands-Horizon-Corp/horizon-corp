import React, { useState } from 'react'
import { useRouterState } from '@tanstack/react-router'
import { useUserAuthStore } from '@/store/user-auth-store'

const AccountSettings = ({}: Props) => {
    const [mode, setMode] = useState(false)
    const { currentUser } = useUserAuthStore()

    const hash = useRouterState({
        select: ({ location }) => location.hash,
    })

    if (hash !== 'account-settings') return null

    return (
        <div className="rounded-xl border bg-secondary p-4">
            <p className="text-lg font-medium">Account Settings</p>
        </div>
    )
}

export default AccountSettings
