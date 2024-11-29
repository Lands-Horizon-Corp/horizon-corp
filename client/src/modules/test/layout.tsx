// import MainMapContainer from '@/components/map'
import { Outlet } from '@tanstack/react-router'
// import {
//     ImagePreview,
//     ImagePreviewContent,
// } from '@/components/ui/image-preview'
// import { Button } from '@/components/ui/button'
// import { sampleMediaResourceList } from './testSampleData'
// import { useImagePreview } from '@/store/image-preview-store'
// import Signature from '@/components/signature'
// import { Separator } from '@radix-ui/react-dropdown-menu'
// import DocumentBuilder from '@/components/document-builder'
const TestLayout = () => {1
    // const { isOpen, setIsOpen } = useImagePreview()
    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <main>
                <Outlet />
            </main>
            {/* <div className="hidden mx-auto h-[100vh] w-[80%] border">
                <Signature/>
                <h1>Signature testing</h1>
                <Separator  className='my-5 h-[1px] dark:bg-white w-full'/>
                <Button onClick={() => setIsOpen(true)}>view</Button>
                <ImagePreview
                    open={isOpen}
                    onOpenChange={() => setIsOpen(false)}
                >
                    <ImagePreviewContent Images={sampleMediaResourceList} />
                </ImagePreview>
                <h1>Image preview testing</h1>
                <Separator  className='my-5 h-[1px] dark:bg-white w-full'/>
            </div> */}
        </div>

    )
}

export default TestLayout
