import { create } from 'zustand'
import { IUserData } from '@/server'

type TAuthStoreStatus = 'loading' | 'authorized' | 'unauthorized' | 'error'

interface UserAuthStore {
    currentUser: IUserData | null
    authStatus: TAuthStoreStatus
    setCurrentUser: (newUserData: IUserData | null) => void
    setAuthStatus: (status: TAuthStoreStatus) => void
}

export const useUserAuthStore = create<UserAuthStore>((set) => ({
    currentUser: null,
    authStatus: 'loading',
    setCurrentUser: (newUserData: IUserData | null) =>
        set({
            currentUser: newUserData,
            authStatus: newUserData ? 'authorized' : 'unauthorized',
        }),
    setAuthStatus: (authStatus: TAuthStoreStatus) => set({ authStatus }),
}))
