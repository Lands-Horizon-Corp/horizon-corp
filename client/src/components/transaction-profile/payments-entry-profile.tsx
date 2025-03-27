import { Card, CardContent } from '@/components/ui/card'
import { ArrowUpRight, UserX } from 'lucide-react'
import { Link } from '@tanstack/react-router'
import { ImagePreview, ImagePreviewContent } from '../ui/image-preview'
import { useImagePreview } from '@/store/image-preview-store'
import { useCallback } from 'react'
import { Image2Icon } from '../icons'
import { IMemberResource } from '@/server'
import CopyTextButton from '@/components/copy-text-button'

type PaymentsEntryProfileProps = {
    profile: IMemberResource | null
}

const NoMemberSelected = () => (
    <Card className="flex flex-col items-center justify-center gap-4 rounded-2xl p-5 shadow-md">
        <UserX size={48} className="text-gray-400" />
        <h2 className="text-2xl font-semibold text-gray-700">
            No Member Selected
        </h2>
        <p className="text-center text-sm text-gray-500">
            There’s no member selected. Please add a member first to continue.
        </p>
    </Card>
)

const PaymentsEntryProfile = ({ profile }: PaymentsEntryProfileProps) => {
    const { isOpen, setIsOpen } = useImagePreview()

    const handleImageClick = useCallback(() => setIsOpen(true), [setIsOpen])
    const handlePreviewClose = useCallback(() => setIsOpen(false), [setIsOpen])

    if (!profile) return <NoMemberSelected />

    const { id, memberProfile, fullName, permanentAddress, media } = profile

    return (
        <Card className="flex items-center gap-6 rounded-2xl p-5 shadow-md">
            <div className="relative h-32 w-32 overflow-hidden rounded-2xl">
                {media?.url ? (
                    <img
                        onClick={handleImageClick}
                        className="h-full w-full cursor-pointer object-cover"
                        src={media?.url ?? 'https://via.placeholder.com/150'}
                        alt={fullName ?? 'Member Image'}
                    />
                ) : (
                    <div className="flex h-full w-full items-center justify-center bg-secondary dark:bg-popover">
                        <Image2Icon
                            className="text-muted-foreground"
                            size={100}
                        />
                    </div>
                )}
                <ImagePreview open={isOpen} onOpenChange={handlePreviewClose}>
                    <ImagePreviewContent Images={media ? [media] : []} />
                </ImagePreview>
            </div>

            <CardContent className="flex flex-col justify-between gap-4 p-0">
                <div>
                    <p className="text-lg font-semibold text-primary">#{id}</p>
                    <h2 className="text-2xl font-bold text-accent-foreground">
                        {fullName}
                    </h2>
                    <p className="text-sm text-gray-500">
                        <span className="text-xs font-light">
                            {' '}
                            Passbook No:
                        </span>{' '}
                        {memberProfile?.passbookNumber}
                        <CopyTextButton
                            className="ml-2"
                            textContent={memberProfile?.passbookNumber ?? ''}
                        />
                    </p>
                    <div className="flex items-center gap-2 text-sm text-gray-600">
                        <Link
                            href={`/members/${id}`}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="flex items-center gap-1 text-black hover:underline dark:text-primary-foreground"
                        >
                            profile <ArrowUpRight size={14} />
                        </Link>
                        <span>• {permanentAddress}</span>
                    </div>
                </div>
            </CardContent>
        </Card>
    )
}

export default PaymentsEntryProfile
