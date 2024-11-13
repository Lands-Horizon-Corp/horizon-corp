import { Outlet } from '@tanstack/react-router'
import {
    ImagePreview,
    ImagePreviewContent,
} from '@/components/ui/image-preview'
import { Button } from '@/components/ui/button'
import { sampleMediaResourceList } from './testSampleData'
import { useImagePreview } from '@/store/image-preview-store'
import Signature from '@/components/signature'
const TestLayout = () => {
    const { isOpen, setIsOpen } = useImagePreview()

    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <main>
                <Outlet />
            </main>
            <Signature/>
            <div className="mx-auto h-[100vh] w-[80%]">
                <Button onClick={() => setIsOpen(true)}>view</Button>
                <ImagePreview
                    open={isOpen}
                    onOpenChange={() => setIsOpen(false)}
                >
                    <ImagePreviewContent Images={sampleMediaResourceList} />
                </ImagePreview>
            </div>
        </div>
    )
}

export default TestLayout
