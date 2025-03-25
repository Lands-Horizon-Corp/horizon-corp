import { forwardRef } from 'react'

import { cn } from '@/lib'
import { IBaseComp } from '@/types'

interface Props extends IBaseComp {
    lon: number
    lat: number
}

const OpenExternalMap = forwardRef<HTMLAnchorElement, Props>(
    ({ lon, lat, className, ...props }, ref) => {
        return (
            <a
                ref={ref}
                href={`https://www.google.com/maps?q=${lat},${lon}`}
                target="_blank"
                className={cn('', className)}
                {...props}
            ></a>
        )
    }
)

OpenExternalMap.displayName = 'OpenExternalMap'

export default OpenExternalMap
