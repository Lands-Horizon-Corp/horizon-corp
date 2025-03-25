import DOMPurify from 'dompurify'

import {
    Accordion,
    AccordionItem,
    AccordionContent,
    AccordionTrigger,
} from '@/components/ui/accordion'
import ImageDisplay from '@/components/image-display'
import { IMemberJointAccountsResource } from '@/server'
import SectionTitle from '../section-title'
import { HandShakeHeartIcon } from '@/components/icons'

interface Props {
    jointAccounts?: IMemberJointAccountsResource[]
}

const JointAccountsDisplay = ({ jointAccounts }: Props) => {
    return (
        <div className="space-y-4">
            <SectionTitle
                title="Joint Accounts"
                subTitle="Co-owners of this account that have the access and share
                financial responsibility of this account "
                Icon={HandShakeHeartIcon}
            />
            {(!jointAccounts || jointAccounts.length === 0) && (
                <p className="w-full text-center text-xs text-muted-foreground/70">
                    no joint accounts
                </p>
            )}
            {jointAccounts?.map((jointAcc) => (
                <div
                    key={jointAcc.id}
                    className="space-y-2 rounded-xl bg-secondary/20 p-4"
                >
                    <div className="flex gap-x-4">
                        <div className="grid w-80 grid-cols-2 gap-x-4">
                            <div className="">
                                <ImageDisplay
                                    src={jointAcc.media?.downloadURL}
                                    className="h-24 w-full rounded-xl"
                                />
                                <p className="mt-1 text-xs text-muted-foreground/70">
                                    Photo
                                </p>
                            </div>

                            <div className="">
                                <ImageDisplay
                                    src={jointAcc.signatureMedia?.downloadURL}
                                    className="h-24 w-full rounded-xl"
                                />
                                <p className="mt-1 text-xs text-muted-foreground/70">
                                    Signature
                                </p>
                            </div>
                        </div>

                        <div className="grid flex-1 gap-2 md:grid-cols-4">
                            <div className="space-y-2">
                                <p>{jointAcc.firstName ?? '-'}</p>
                                <p className="text-xs text-muted-foreground/70">
                                    First Name
                                </p>
                            </div>
                            <div className="space-y-2">
                                <p>{jointAcc.middleName ?? '-'}</p>
                                <p className="text-xs text-muted-foreground/70">
                                    Middle Name
                                </p>
                            </div>
                            <div className="space-y-2">
                                <p>{jointAcc.lastName ?? '-'}</p>
                                <p className="text-xs text-muted-foreground/70">
                                    Last Name
                                </p>
                            </div>
                            <div className="space-y-2">
                                <p>{jointAcc.suffix ?? '-'}</p>
                                <p className="text-xs text-muted-foreground/70">
                                    Suffix
                                </p>
                            </div>
                            <div className="space-y-2">
                                <p>{jointAcc.familyRelationship ?? '-'}</p>
                                <p className="text-xs text-muted-foreground/70">
                                    Relationship
                                </p>
                            </div>
                        </div>
                    </div>
                    <Accordion collapsible type="single" className="w-full">
                        <AccordionItem value="item-1">
                            <AccordionTrigger className="font-normal">
                                Description
                            </AccordionTrigger>
                            <AccordionContent className="prose-h1: prose w-full !max-w-full rounded-xl p-4 text-sm text-foreground/70 dark:prose-invert prose-p:text-foreground/80 prose-strong:text-foreground sm:text-sm">
                                <div
                                    dangerouslySetInnerHTML={{
                                        __html: DOMPurify.sanitize(
                                            jointAcc.description &&
                                                jointAcc.description.length > 0
                                                ? jointAcc.description
                                                : '<i>No Description</i>'
                                        ),
                                    }}
                                />
                            </AccordionContent>
                        </AccordionItem>
                    </Accordion>
                </div>
            ))}
        </div>
    )
}

export default JointAccountsDisplay
