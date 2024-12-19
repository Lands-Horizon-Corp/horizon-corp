import { ReactNode } from 'react'

export * from '../../components/sidebar'

export interface IBaseComp {
    className?: string
    children?: ReactNode
}

export interface IBaseCompChildOnly extends Omit<IBaseComp, 'className'> {}

export interface IBaseCompNoChild extends Omit<IBaseComp, 'children'> {}
