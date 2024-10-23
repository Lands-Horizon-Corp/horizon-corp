import { createContext, useContext } from 'react'

export const ExpandContext = createContext<boolean | null>(true)

export const useSidebarExpandContext = () => {
    const sidebarExpandContext = useContext(ExpandContext)

    if (sidebarExpandContext === null)
        throw new Error('This hook must only be used within the sidebar')

    return sidebarExpandContext
}
