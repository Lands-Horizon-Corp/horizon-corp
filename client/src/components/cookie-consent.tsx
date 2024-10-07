import { addYears } from 'date-fns'
import { useCookies } from 'react-cookie'

import { Button } from '@/components/ui/button'
import { CookieIcon } from '@/components/icons'
import { Alert, AlertDescription } from '@/components/ui/alert'

const CookieConsent = () => {
    const [cookies, setCookie] = useCookies(['ecoop-cookie-consent'])

    const onAccept = () => {
        const expirationDate = addYears(new Date(), 1)
        setCookie('ecoop-cookie-consent', true, {
            path: '/',
            secure: true,
            sameSite: 'lax',
            expires: expirationDate,
        })
    }

    if (cookies['ecoop-cookie-consent']) return null

    return (
        <Alert className="fixed bottom-0 left-1/2 mb-3 flex w-full max-w-sm -translate-x-1/2 flex-col items-center gap-x-4 gap-y-4 rounded-xl bg-popover/95 text-center shadow-center-md backdrop-blur md:mb-6 lg:w-fit lg:max-w-none lg:flex-row lg:text-left">
            <span className="absolute -top-6 block size-fit rounded-full bg-secondary p-1 lg:static">
                <CookieIcon className="size-9 text-[#f3cc83] lg:size-6" />
            </span>
            <AlertDescription className="pt-4 lg:pt-0">
                ECoop uses cookies to enhance the user experience.
            </AlertDescription>
            <Button
                onClick={onAccept}
                variant="secondary"
                className="w-full rounded-xl lg:w-fit"
            >
                I understand
            </Button>
        </Alert>
    )
}

export default CookieConsent
