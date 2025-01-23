import FootstepTable from '@/components/tables/footsteps-table'

const OwnerFootstepTrackingPage = () => {
    return (
        <div className="flex w-full max-w-full flex-col items-center px-4 pb-6 sm:px-8">
            <FootstepTable mode="team" className="min-h-[90vh] w-full" />
        </div>
    )
}

export default OwnerFootstepTrackingPage
