import { useRouterState } from '@tanstack/react-router'

import { Separator } from '@/components/ui/separator'

import UsernameOption from './username-option'

import { useUserAuthStore } from '@/store/user-auth-store'
import EmailOption from './email-option'
import ContactOption from './contact-option'

const AccountSecurity = () => {
    const { currentUser, setCurrentUser } = useUserAuthStore()

    const hash = useRouterState({
        select: ({ location }) => location.hash,
    })

    if (hash !== 'security' || !currentUser) return null

    return (
        <div>
            <div className="flex items-center justify-between">
                <p className="text-lg font-medium">Security</p>
            </div>
            <p className="mt-1 text-xs text-foreground/60">
                Manage or Update your account security
            </p>
            <Separator className="my-2 sm:my-4" />
            <div className="space-y-4">
                <UsernameOption
                    username={currentUser.username}
                    onSave={(newUserData) => setCurrentUser(newUserData)}
                />
                <EmailOption
                    email={currentUser.email}
                    onSave={(newUserData) => setCurrentUser(newUserData)}
                />
                <ContactOption 
                    contact={currentUser.contactNumber}
                    onSave={(newUserData) => setCurrentUser(newUserData)}
                />
            </div>
        </div>
    )
}

export default AccountSecurity
