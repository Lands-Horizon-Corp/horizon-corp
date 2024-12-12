import { useRouter } from '@tanstack/react-router'

import { Button } from '@/components/ui/button'
import { ArrowLeftIcon } from '@/components/icons'

const CompanyErrorPage = () => {
    const router = useRouter()

    return (
        <div className="flex h-full w-full flex-col items-center justify-center gap-y-4 p-4">
            <h1 className="text-4xl font-medium">Oops!</h1>
            <p className="text-foreground/70">
                Sorry, an unexpected error has occurred.
            </p>
            <Button
                variant="secondary"
                className="gap-x-2 rounded-full"
                onClick={() => router.history.back()}
            >
                <ArrowLeftIcon />
                Go Back
            </Button>
        </div>
    )
}

export default CompanyErrorPage
