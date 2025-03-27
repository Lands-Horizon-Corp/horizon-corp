import { IconType } from 'react-icons/lib'

export type TNavItemType = 'dropdown' | 'item'

export interface INavItemBase {
    url?: string
    title: string
    depth?: number
    isSub?: boolean
    icon?: IconType
}

export interface INavItemSingle extends INavItemBase {
    type: 'item'
}

type ExceededDepthError = never

type TDepth = {
    3: 2
    2: 1
    1: 0
    0: ExceededDepthError
}

export interface INavItemDropdown<D extends keyof TDepth = 3>
    extends INavItemBase {
    type: 'dropdown'
    items: D extends 1
        ? INavItemSingle[]
        : [INavItem<TDepth[D]>, ...INavItem<TDepth[D]>[]]
}

export type INavItem<D extends keyof TDepth = 2> = D extends 0
    ? never
    : INavItemDropdown<D> | INavItemSingle

export interface INavGroupItem {
    title: string
    navItems: INavItem[]
}
