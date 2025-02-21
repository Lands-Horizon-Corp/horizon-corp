import { Card, CardContent } from '@/components/ui/card'
import { ArrowUpRight } from 'lucide-react'
import { Link } from '@tanstack/react-router'
import { ImagePreview, ImagePreviewContent } from '../ui/image-preview'
import { useImagePreview } from '@/store/image-preview-store'
import { IMemberCardResource } from '../forms/transactions/payments-entry-form'

const PaymentsEntryProfile = ({
    id,
    passbookNumber,
    fullName,
    permanentAddress,
    media,
}: IMemberCardResource) => {
    const { isOpen, setIsOpen } = useImagePreview()

    return (
        <Card className="flex items-center gap-6 rounded-2xl p-6 shadow-md">
            <div className="relative h-32 w-32 overflow-hidden rounded-2xl">
                <img
                    onClick={() => setIsOpen(true)}
                    className="h-full w-full cursor-pointer object-cover"
                    src={
                        media?.url
                            ? media.url
                            : 'https://via.placeholder.com/150'
                    }
                ></img>
                <ImagePreview
                    open={isOpen}
                    onOpenChange={() => setIsOpen(false)}
                >
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
                        Passbook No: {passbookNumber}
                    </p>
                    <div className="flex items-center gap-2 text-sm text-gray-600">
                        <Link
                            href={'/members/' + id}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="flex items-center gap-1 text-primary-foreground hover:underline"
                        >
                            profile <ArrowUpRight size={14} />
                        </Link>
                        <span>• {permanentAddress}</span>
                    </div>
                </div>
                {/* <div className="flex flex-wrap gap-3">
                    {badges.map((badge, index) => (
                        <Badge
                            key={index}
                            className="rounded-full bg-green-100 px-3 py-1 text-xs text-green-700"
                        >
                            {badge.label} {badge.years}
                        </Badge>
                    ))}
                </div> */}
            </CardContent>
        </Card>
    )
}

export default PaymentsEntryProfile
