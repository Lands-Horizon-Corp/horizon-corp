import { useState } from 'react'

import { Button } from '@/components/ui/button'
import { CameraFillIcon } from '@/components/icons'
import ImageDisplay from '@/components/image-display'
import ActionTooltip from '@/components/action-tooltip'
import SingleImageUploaderModal from '@/components/single-image-uploader/single-image-uploader-modal'

import { IBaseCompNoChild } from '@/types'
import { CompanyResource } from '@/horizon-corp/types'
import { useUpdateCompanyProfilePicture } from '@/hooks/api-hooks/use-company'
import LoadingSpinner from '@/components/spinners/loading-spinner'

interface Props extends IBaseCompNoChild {
    company: CompanyResource
}

const CompanyLogo = ({ company }: Props) => {
    const [modal, toggleModal] = useState(false)

    const {
        mutate: updateCompanyProfile,
        isPending: isUpdatingCompanyProfilePicture,
    } = useUpdateCompanyProfilePicture({})

    return (
        <div className="absolute -top-28 left-4 z-20 size-fit">
            <SingleImageUploaderModal
                modalState={modal}
                setModalState={toggleModal}
                modalTitle="Update Company Image"
                defaultFileName={`company-${company.id}`}
                onUploadComplete={(newMediaResource) => {
                    updateCompanyProfile({
                        companyId: company.id,
                        mediaResource: newMediaResource,
                    })
                    toggleModal(false)
                }}
            />
            <ImageDisplay
                className="size-36 border-4 border-popover shadow-sm"
                src={company.media?.downloadURL}
                fallback={company.name.charAt(0) ?? '-'}
            />
            <ActionTooltip tooltipContent="Change" align="center" side="right">
                <Button
                    variant="secondary"
                    disabled={isUpdatingCompanyProfilePicture}
                    onClick={() => toggleModal((prev) => !prev)}
                    className="absolute bottom-2 right-2 size-fit rounded-full border border-transparent p-1 hover:border-foreground/20"
                >
                    {isUpdatingCompanyProfilePicture ? (
                        <LoadingSpinner />
                    ) : (
                        <CameraFillIcon className="size-4 opacity-50 duration-300 ease-in-out group-hover:opacity-80" />
                    )}
                </Button>
            </ActionTooltip>
        </div>
    )
}

export default CompanyLogo
