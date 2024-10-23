import { create } from 'zustand'
import { UserData } from '@/horizon-corp/types'

type TAuthStoreStatus = 'loading' | 'authorized' | 'unauthorized'

interface UserAuthStore {
    currentUser: UserData | null
    authStatus: TAuthStoreStatus
    setCurrentUser: (newUserData: UserData | null) => void
    setAuthStatus: (status: TAuthStoreStatus) => void
}

export const useUserAuthStore = create<UserAuthStore>((set) => ({
    currentUser : null,
    authStatus: 'loading',
    setCurrentUser : (newUserData : UserData | null) =>
        set({ currentUser : newUserData, authStatus: newUserData ? 'authorized' : 'unauthorized' }),
    setAuthStatus: (authStatus: TAuthStoreStatus) => set({ authStatus }),
}))
