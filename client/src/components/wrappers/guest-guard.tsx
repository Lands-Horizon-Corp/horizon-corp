import { useRouter } from "@tanstack/react-router"

import { IBaseComp } from "@/types"
import useCurrentUser from "@/hooks/use-current-user"
import { getUsersAccountTypeRedirectPage } from "@/helpers"
import { toast } from "sonner"

interface IGuestGuardProps extends Omit<IBaseComp, "className"> {
    allowAuthenticatedUser? : false
}

const GuestGuard = ({ allowAuthenticatedUser = false, children } : IGuestGuardProps ) => {
    const router = useRouter();
    const { data : currentUser } = useCurrentUser()

    if(!allowAuthenticatedUser && currentUser){ 
        const redirectUrl = getUsersAccountTypeRedirectPage(currentUser)
        toast.message("You are signed in, system redirected you.")
        router.navigate({ to : redirectUrl })
        return <></>
    }

    return <>{children}</> 
}

export default GuestGuard
