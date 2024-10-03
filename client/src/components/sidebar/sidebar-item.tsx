import { FC, useState } from 'react'
import { useLocation, useRouter } from '@tanstack/react-router'

import { useSidebarContext } from '@/components/sidebar'
import type { TSidebarItem } from '@/types/component/sidebar'
import SidebarItemContent from '@/components/sidebar/sidebar-item-content'
import SidebarItemWithTooltip from '@/components/sidebar/sidebar-with-tooltip-wrapper'

import { cn } from '@/lib/utils'
import {
    concatParentUrl,
    sidebarRouteMatcher,
} from '@/components/sidebar/sidebar-utils'

const SidebarItem: FC<TSidebarItem> = (props) => {
    const router = useRouter()
    const pathname = useLocation({
        select: (location) => location.pathname,
    })

    const expand = useSidebarContext()

    const [collapse, setcollapse] = useState(props.collapsed)

    if (props.component) {
        return props.component
    } else if (props.subItems) {
        const { text, Icon, subItems, baseUrl, isSub } = props

        const routeMatched = sidebarRouteMatcher(baseUrl, pathname)

        return (
            <>
                <SidebarItemWithTooltip
                    enableTooltip={!expand}
                    tooltipContent={text}
                >
                    <SidebarItemContent
                        Icon={Icon}
                        text={text}
                        isSub={isSub}
                        expand={expand}
                        collapse={collapse}
                        active={routeMatched}
                        onCollapse={() => setcollapse((val) => !val)}
                    />
                </SidebarItemWithTooltip>
                {collapse && (
                    <div
                        className={cn(
                            'ml-3.5 flex flex-col gap-y-1 border-l pl-2',
                            !expand && 'gap-y-2 pl-1'
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

        const routeMatched = sidebarRouteMatcher(url, pathname)

        return (
            <SidebarItemWithTooltip
                enableTooltip={!expand}
                tooltipContent={text}
            >
                <SidebarItemContent
                    Icon={Icon}
                    text={text}
                    isSub={isSub}
                    expand={expand}
                    active={routeMatched}
                    onClick={() => router.navigate({ to: url })}
                />
            </SidebarItemWithTooltip>
        )
    }
}

export default SidebarItem
