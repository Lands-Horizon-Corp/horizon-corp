import { createContext, useContext } from 'react'

interface ISidebarContext {
    isMobileSidebarVisible: boolean
    toggleMobileSidebar: (newState: boolean) => void
    isExpanded: boolean
    toggleExpanded: (newState: boolean) => void
}

export const SidebarContext = createContext<ISidebarContext | null>(null)

export const useSidebarContext = () => {
    const sidebarContext = useContext(SidebarContext)

    if (!sidebarContext)
        throw new Error('This hook is only usable within SidebarProvider ')

    return sidebarContext
}

export default SidebarContext
