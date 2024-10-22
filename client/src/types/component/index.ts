import { ReactNode } from 'react'

export * from './sidebar'

export interface IBaseComp {
    className?: string
    children?: ReactNode
}

export interface IBaseCompNoChild extends Omit<IBaseComp, 'children'> {}
