// import MainMapContainer from '@/components/map'
import { Outlet } from '@tanstack/react-router'

const TestLayout = () => {1
    return (
        <div className="grid min-h-[100dvh] grid-cols-[auto_1fr]">
            <main>
                <Outlet />
            </main>
        </div>
    )
}

export default TestLayout
