'use client'

import { Fragment, useMemo } from 'react'
import { Link, useRouterState } from '@tanstack/react-router'

import { ChevronRightIcon, HomeFillIcon } from '../icons'
import {
    Breadcrumb,
    BreadcrumbItem,
    BreadcrumbLink,
    BreadcrumbList,
    BreadcrumbPage,
    BreadcrumbSeparator,
} from '@/components/ui/breadcrumb'

import { cn } from '@/lib/utils'
import PageNavigator from './page-navigator'

type Props = {
    homeUrl?: string
    className?: string
}

const PageBreadCrumb = ({ className, homeUrl }: Props) => {
    const router = useRouterState()
    const pathName = router.location.pathname

    const matches = router.matches

    const paths = useMemo(() => {
        const paths: { name: string; urlPath: string; isClickable: boolean }[] =
            []
        const components = pathName.split('/').filter((path) => path)

        let currentPath = ''

        components.forEach((component) => {
            if ('/' + component === homeUrl) return

            currentPath += '/' + component
            paths.push({
                name: component,
                urlPath: currentPath,
                isClickable: matches.some(
                    (match) => match.fullPath === currentPath
                ),
            })
        })

        return paths
    }, [pathName, homeUrl, matches])

    return (
        <Breadcrumb className={cn('capitalize', className)}>
            <BreadcrumbList>
                <PageNavigator />
                {homeUrl && (
                    <>
                        <BreadcrumbItem>
                            <BreadcrumbLink asChild>
                                <Link to={homeUrl}>
                                    <HomeFillIcon className="size-4" />
                                </Link>
                            </BreadcrumbLink>
                        </BreadcrumbItem>

                        <BreadcrumbSeparator>
                            <ChevronRightIcon className="size-4" />
                        </BreadcrumbSeparator>
                    </>
                )}
                {paths.map((path, i) => {
                    return (
                        <Fragment key={path.urlPath}>
                            <BreadcrumbItem className="text-foreground/40">
                                {i !== paths.length - 1 ? (
                                    <BreadcrumbLink
                                        className={cn(
                                            'text-inherit',
                                            !path.isClickable &&
                                                'pointer-events-none'
                                        )}
                                        asChild
                                    >
                                        <Link to={path.urlPath}>
                                            {path.name}
                                        </Link>
                                    </BreadcrumbLink>
                                ) : (
                                    <BreadcrumbPage className="cursor-default">
                                        {path.name}
                                    </BreadcrumbPage>
                                )}
                            </BreadcrumbItem>

                            {i < paths.length - 1 && (
                                <BreadcrumbSeparator>
                                    <ChevronRightIcon className="size-4" />
                                </BreadcrumbSeparator>
                            )}
                        </Fragment>
                    )
                })}
            </BreadcrumbList>
        </Breadcrumb>
    )
}

export default PageBreadCrumb