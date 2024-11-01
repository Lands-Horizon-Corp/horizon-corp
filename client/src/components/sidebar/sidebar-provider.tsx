import { useEffect, useState } from 'react'
import SidebarContext from './sidebar-context'

import { IBaseCompChildOnly } from '@/types'
import { useIsMobile } from '@/hooks/use-mobile'

export const SidebarProvider = ({ children }: IBaseCompChildOnly) => {
    const isMobile = useIsMobile()
    const [isExpanded, toggleExpanded] = useState(true)
    const [isMobileSidebarVisible, toggleMobileSidebar] = useState(false)

    useEffect(() => {
        if (isMobile) toggleExpanded(true)
    }, [isMobile])

    return (
        <SidebarContext.Provider
            value={{
                isMobile,
                isExpanded,
                isMobileSidebarVisible,
                toggleExpanded,
                toggleMobileSidebar,
            }}
        >
            {children}
        </SidebarContext.Provider>
    )
}
