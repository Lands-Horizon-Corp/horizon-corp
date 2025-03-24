import RawDescription from './raw-description'
import CopyTextButton from '@/components/copy-text-button'

import { toReadableDate } from '@/utils'
import { abbreviateUUID } from '@/utils/formatting-utils'
import { IMemberDescriptionResource } from '@/server'

interface Props {
    descriptions?: IMemberDescriptionResource[]
}

const MemberDescriptionDisplays = ({ descriptions }: Props) => {
    return (
        <div className="space-y-4">
            <p>Descriptions</p>
            <p className="!mt-0 text-sm text-muted-foreground/60">
                Other descriptions about the member are shown here...
            </p>
            {(!descriptions || descriptions.length === 0) && (
                <p className="w-full text-center text-xs text-muted-foreground/70">
                    no other descriptions
                </p>
            )}
            {descriptions?.map((description) => {
                return (
                    <div className="space-y-4 rounded-xl border p-4">
                        <div className="flex items-center justify-between">
                            <div>
                                <p>{description.name}</p>
                                <p className="text-xs text-muted-foreground/70">
                                    ID: {abbreviateUUID(description.id)}...
                                    <CopyTextButton
                                        className="ml-1"
                                        textContent={description.id}
                                    />
                                </p>
                            </div>
                            <p className="text-xs">
                                {toReadableDate(
                                    description.date,
                                    "MMMM d yyyy 'at' h:mm a"
                                )}
                            </p>
                        </div>
                        <RawDescription
                            className="rounded-xl bg-popover/80 p-4"
                            content={description.description}
                        />
                    </div>
                )
            })}
        </div>
    )
}

export default MemberDescriptionDisplays
