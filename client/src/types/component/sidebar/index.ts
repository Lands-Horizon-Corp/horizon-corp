import { ReactNode } from '@tanstack/react-router'
import { IBaseCompNoChild } from '../base'
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
    collapsed?: never
    component?: never
}

export interface IExpandableSidebarItem extends ISidebarItemBase {
    baseUrl: string
    subItems: TSidebarItem[]

    url?: never
    collapsed?: boolean
    component?: never
}

export interface IComponentSidebarItem extends ISidebarItemBase {
    component?: ReactNode
    url?: never
    isSub?: never
    Icon?: never
    subItems?: never
    collapsed?: never
}

export type TSidebarItem =
    | IClickableSidebarItem
    | IExpandableSidebarItem
    | IComponentSidebarItem
