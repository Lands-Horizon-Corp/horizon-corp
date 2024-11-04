import { useState } from 'react'
import SidebarContext from './sidebar-context'

import { IBaseCompChildOnly } from '@/types'

export const SidebarProvider = ({ children }: IBaseCompChildOnly) => {
    const [isExpanded, toggleExpanded] = useState(true)
    const [isMobileSidebarVisible, toggleMobileSidebar] = useState(false)

    return (
        <SidebarContext.Provider
            value={{
                isMobileSidebarVisible,
                toggleMobileSidebar,
                isExpanded,
                toggleExpanded,
            }}
        >
            {children}
        </SidebarContext.Provider>
    )
}
