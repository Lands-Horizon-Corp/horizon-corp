import { cn } from '@/lib'
import { IBaseComp } from '@/types'
import { IMemberProfileResource, TEntityId } from '@/server'
// import { useMemberProfile } from '@/hooks/api-hooks/member/use-member-profile'

interface Props extends IBaseComp {
    profileId: TEntityId
    defaultData?: IMemberProfileResource
}

const MemberGeneralMembershipInfo = ({
    // profileId,
    className,
    // defaultData,
}: Props) => {
    // const { data, isPending } = useMemberProfile({
    //     profileId,
    //     initialData: defaultData,
    // })

    return (
        <div
            className={cn(
                'flex flex-1 flex-col gap-y-4 rounded-xl bg-background p-4',
                className
            )}
        >
            <div>
                <p>Membership Information</p>
                <div className=""></div>
            </div>

            <div>
                <p>Member Classification & Type</p>
                <div className=""></div>
            </div>
        </div>
    )
}

export default MemberGeneralMembershipInfo
