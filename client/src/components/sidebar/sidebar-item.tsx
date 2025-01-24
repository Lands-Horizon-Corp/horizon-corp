import { FC, useState } from 'react'
import { useLocation, useRouter } from '@tanstack/react-router'

import type { TSidebarItem } from '@/components/sidebar/sidebar-types'
import SidebarItemContent from '@/components/sidebar/sidebar-item-content'
import SidebarItemWithTooltip from '@/components/sidebar/sidebar-with-tooltip-wrapper'

import { concatParentUrl, sidebarRouteMatcher } from './sidebar-utils'
import { cn } from '@/lib/utils'
import { useSidebarContext } from './sidebar-context'

const SidebarItem: FC<TSidebarItem> = (props) => {
    const router = useRouter()
    const pathname = useLocation({
        select: (location) => location.pathname,
    })

    const { isExpanded } = useSidebarContext()

    const [isCollapsed, setIsCollapsed] = useState(props.isCollapseEnabled)

    if (props.component) {
        return props.component
    } else if (props.subItems) {
        const { text, Icon, subItems, baseUrl, isSub } = props

        const isRouteMatched = sidebarRouteMatcher(baseUrl, pathname)

        return (
            <>
                <SidebarItemWithTooltip
                    enableTooltip={!isExpanded}
                    tooltipContent={text}
                >
                    <SidebarItemContent
                        Icon={Icon}
                        text={text}
                        isSub={isSub}
                        isExpanded={isExpanded}
                        isCollapsed={isCollapsed}
                        isActive={isRouteMatched}
                        onCollapse={() => setIsCollapsed((val) => !val)}
                    />
                </SidebarItemWithTooltip>
                {isCollapsed && (
                    <div
                        className={cn(
                            'ml-3.5 flex flex-col gap-y-2 border-l pl-1',
                            isExpanded && 'gap-y-0 pl-2'
                        )}
                    >
                        {subItems.map((subItem) => {
                            const builtUrl = concatParentUrl({
                                baseUrl: props.baseUrl,
                                url: subItem.url,
                            })
                            if (subItem.component) {
                                return subItem.component
                            } else if (subItem.subItems) {
                                return (
                                    <SidebarItem
                                        isSub={true}
                                        text={subItem.text}
                                        Icon={subItem.Icon}
                                        baseUrl={subItem.baseUrl}
                                        subItems={subItem.subItems}
                                        key={`${subItem.text}-${builtUrl}`}
                                    />
                                )
                            } else {
                                return (
                                    <SidebarItem
                                        isSub={true}
                                        url={builtUrl}
                                        text={subItem.text}
                                        Icon={subItem.Icon}
                                        key={`${subItem.text}-${builtUrl}`}
                                    />
                                )
                            }
                        })}
                    </div>
                )}
            </>
        )
    } else {
        const { text, Icon, url, isSub } = props

        const isRouteMatched = sidebarRouteMatcher(url, pathname)

        return (
            <SidebarItemWithTooltip
                tooltipContent={text}
                enableTooltip={!isExpanded}
            >
                <SidebarItemContent
                    Icon={Icon}
                    text={text}
                    isSub={isSub}
                    isExpanded={isExpanded}
                    isActive={isRouteMatched}
                    onClick={() => router.navigate({ to: url })}
                />
            </SidebarItemWithTooltip>
        )
    }
}

export default SidebarItem
