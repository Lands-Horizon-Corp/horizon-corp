import z from "zod"
import { useMutation } from "@tanstack/react-query"

import { withCatchAsync } from "@/lib"
import { ChangeUsernameRequest, UserData } from "@/horizon-corp/types"
import ProfileService from "@/horizon-corp/server/auth/ProfileService"

interface Props {
    onSave : (newUserData : UserData) => void
}

const userNameOptionSchema = z.object({
    
})

const UsernameOption = ({ onSave } : Props) => {

    const req = useMutation<any, string, ChangeUsernameRequest>({
        mutationKey : ['account-security-username'],
        mutationFn : async (data) => {
            const [error, response] = await withCatchAsync(ProfileService.ChangeUsername(data))
        }
    })

    return (
        <div>
            <p className="text-foreground/70">Username *</p> 
        </div>
    )
}

export default UsernameOption
