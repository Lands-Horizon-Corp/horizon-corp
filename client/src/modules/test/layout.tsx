// import MainMapContainer from '@/components/map'
import UploadSignature from '@/components/signature'
import { Outlet } from '@tanstack/react-router'

const TestLayout = () => {
    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <main>
                <Outlet />
            </main>
            {/* <UploadSignature/> */}
        </div>
    )
}

export default TestLayout
