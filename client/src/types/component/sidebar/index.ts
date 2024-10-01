import { IBaseCompNoChild } from '../base'
import { IconType } from 'react-icons/lib'

export interface ISidebarItem extends IBaseCompNoChild {
    Icon?: IconType
    text: string
    url?: string
    baseUrl?: string
    isSub?: boolean
    collapsed?: boolean
    subItems?: ISidebarItem[]
}
