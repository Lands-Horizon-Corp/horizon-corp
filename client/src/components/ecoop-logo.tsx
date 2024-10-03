import { cn } from '@/lib/utils'
import { IBaseCompNoChild } from '@/types/component/base'
import { ResolvedTheme } from './providers/theme-provider'

type TEcoopThemeMode = 'dynamic' | ResolvedTheme

interface Props extends IBaseCompNoChild {
    url?: string
    blurDisabled?: boolean
    themeMode?: TEcoopThemeMode
}

const EcoopLogo = ({
    className,
    blurDisabled = false,
    url = '/e-coop-logo-1.webp',
}: Props) => {
    return (
        <div className={cn('relative size-8', className)}>
            <img src={url} alt="logo" className="h-full w-full" />
            {!blurDisabled && (
                <img
                    src={url}
                    alt="logo-blur"
                    className="pointer-events-none absolute inset-0 left-1/2 top-1/2 z-0 hidden size-full -translate-x-1/2 -translate-y-[45%] blur-xl selection:bg-none dark:block"
                />
            )}
        </div>
    )
}

export default EcoopLogo
