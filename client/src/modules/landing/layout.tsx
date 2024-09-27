import { Outlet } from '@tanstack/react-router'
import { ReactNode } from '@tanstack/react-router'
import NavBar from './pages/_components/navbar'
import bg_element_1 from '@/assets/images/bg_element_1.png'
import bg_element_2 from '@/assets/images/bg_element_2.png'
import bg_element_3 from '@/assets/images/bg_element_3.png'

interface Props {
    children?: ReactNode
}

const PublicLayout = (_props: Props) => {
    return (
        <>
            <div className="relative h-full">
                <img
                    src={bg_element_1}
                    className="absolute left-[50%] -translate-x-[50%] -z-30 h-auto w-[1681px]"
                    alt="background"
                />
                <img
                    src={bg_element_2}
                    className="absolute 2xl:left-[60%] left-[90%] -translate-x-[50%] -z-30 h-auto w-[1681px]"
                    alt="background"
                />
                <img
                    src={bg_element_3}
                    className="absolute 2xl:left-[50%] left-[50%] -translate-x-[50%] -z-40 h-auto w-[1440px]"
                    alt="background"
                />
                <main className="h-full w-full">
                    <NavBar></NavBar>
                    <Outlet />
                </main>
            </div>
        </>
    )
}

export default PublicLayout
