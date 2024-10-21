import { create } from 'zustand'
import { UserData } from '@/horizon-corp/types'

type AuthState =
    | 'Idle'
    | 'Loading'
    | 'Authenticated'
    | 'UnAuthenticated'
    | 'Error'

interface IAuthStore {
    authState: AuthState
    currentUser: UserData | null
    setCurrentUser: (user: UserData | null) => void
    setAuthState: (newAuthState: AuthState) => void
}

const useAuthStore = create<IAuthStore>((set) => ({
    authState: 'Idle',
    currentUser: null,
    setCurrentUser: (user) => set({ currentUser: user }),
    setAuthState: (newAuthState) => set({ authState: newAuthState }),
}))

export default useAuthStore
