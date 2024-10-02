import ActionTooltip, { IActionTooltipProps } from '@/components/action-tooltip'

interface Props extends IActionTooltipProps {
    enableTooltip?: boolean
}

const SidebarItemWithTooltip = ({
    children,
    enableTooltip = true,
    side = 'right',
    align = 'center',
    ...otherProps
}: Props) => {
    if (!enableTooltip) return children
    return (
        <ActionTooltip
            delayDuration={300}
            side={side}
            align={align}
            {...otherProps}
        >
            {children}
        </ActionTooltip>
    )
}

export default SidebarItemWithTooltip
