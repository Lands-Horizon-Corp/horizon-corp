import { ReactNode } from '@tanstack/react-router'

import { IBaseCompNoChild } from '..'
import { IconType } from 'react-icons/lib'

export interface ISidebarItemBase extends IBaseCompNoChild {
    url?: string
    text: string
    isSub?: boolean
    Icon?: IconType
}

export interface IClickableSidebarItem extends ISidebarItemBase {
    url: string
    baseUrl?: never
    subItems?: never
    component?: never
    isCollapseEnabled?: never
}

export interface IExpandableSidebarItem extends ISidebarItemBase {
    url?: never
    baseUrl?: string
    component?: never
    subItems: TSidebarItem[]
    isCollapseEnabled?: boolean
}

export interface IComponentSidebarItem extends ISidebarItemBase {
    url?: never
    Icon?: never
    isSub?: never
    subItems?: never
    component?: ReactNode
    isCollapseEnabled?: never
}

export type TSidebarItem =
    | IClickableSidebarItem
    | IExpandableSidebarItem
    | IComponentSidebarItem
