import { useState } from 'react'
import { Link } from '@tanstack/react-router'

import MainMapContainer from '@/components/map'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { PencilOutlineIcon } from '@/components/icons'
import EnsureOwnerCompany from '../../components/ensure-company'
import PageContainer from '@/components/containers/page-container'
import CompanyBanner from '@/components/company-profile/company-banner'
import CompanyDescription from '@/components/company-profile/company-description'
import CompanyBasicDetailsList from '@/components/company-profile/company-basic-details-list'
import { CompanyEditFormModal } from '@/components/forms/company-forms/company-edit-basic-info-form'

import { ICompanyResource } from '@/server'

const OwnerCompanyProfilePage = () => {
    const [editModal, setEditModal] = useState(false)
    const [company, setCompany] = useState<ICompanyResource | undefined>(
        undefined
    )

    return (
        <PageContainer className="gap-y-4">
            <EnsureOwnerCompany onSuccess={setCompany}>
                {company && (
                    <>
                        <CompanyEditFormModal
                            formProps={{
                                companyId: company.id,
                                defaultValues: company,
                            }}
                            open={editModal}
                            onOpenChange={setEditModal}
                        />
                        <CompanyBanner company={company} />
                        <div className="my-0 flex w-full items-center justify-between">
                            <div className="flex items-center gap-x-4 text-foreground/60 dark:text-foreground/70">
                                <p className="text-xs font-medium text-foreground">
                                    Quick Links:
                                </p>
                                <Link
                                    className="border-b border-transparent hover:border-primary hover:text-foreground"
                                    to="/owner/company/branches"
                                >
                                    Branches
                                </Link>
                                <Link
                                    className="border-b border-transparent hover:border-primary hover:text-foreground"
                                    to="/owner/users/members"
                                >
                                    Members
                                </Link>
                                <Link
                                    className="border-b border-transparent hover:border-primary hover:text-foreground"
                                    to="/owner/users/employees"
                                >
                                    Employees
                                </Link>
                            </div>
                            <div className="flex items-center gap-x-2">
                                <Button
                                    size="sm"
                                    variant="default"
                                    onClick={() => setEditModal((val) => !val)}
                                >
                                    <PencilOutlineIcon className="mr-2" />
                                    Edit Company
                                </Button>
                            </div>
                        </div>
                        <div className="w-full rounded-xl border bg-popover dark:bg-secondary">
                            <p className="p-4 font-medium">Company Details</p>
                            <Separator className="dark:bg-background" />
                            <div className="grid grid-cols-5 gap-4 p-4">
                                {company.latitude && company.longitude ? (
                                    <MainMapContainer
                                        viewOnly
                                        zoom={13}
                                        hideControls
                                        className="pointer-events-nonee pointer-events-none z-10 col-span-3 min-h-40 overflow-clip rounded-xl p-0"
                                        mapContainerClassName="sm:rounded-none"
                                        center={{
                                            lng: company.longitude,
                                            lat: company.latitude,
                                        }}
                                        defaultMarkerPins={[
                                            {
                                                lng: company.longitude,
                                                lat: company.latitude,
                                            },
                                        ]}
                                    />
                                ) : (
                                    <div className="col-span-3 flex min-h-40 items-center justify-center rounded-lg border bg-secondary p-4 text-foreground/50 dark:bg-popover">
                                        <p>
                                            Unable to show company location map
                                        </p>
                                    </div>
                                )}
                                <CompanyBasicDetailsList
                                    className="col-span-2"
                                    company={company}
                                />
                            </div>
                        </div>
                        <div className="w-full rounded-xl border bg-popover dark:bg-secondary">
                            <p className="p-4 font-medium">Company About</p>
                            <Separator className="dark:bg-background" />
                            <CompanyDescription
                                className="!bg-transparent p-4 !text-base"
                                description={company.description}
                            />
                        </div>
                    </>
                )}
            </EnsureOwnerCompany>
        </PageContainer>
    )
}

export default OwnerCompanyProfilePage
