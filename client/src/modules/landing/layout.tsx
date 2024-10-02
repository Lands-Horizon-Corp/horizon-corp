import { Outlet } from '@tanstack/react-router'
import { ReactNode } from '@tanstack/react-router'
import NavBar from './pages/_components/navbar'
import bg_element_1 from '@/assets/images/bg_element_1.png'
import bg_element_2 from '@/assets/images/bg_element_2.png'
import bg_element_3 from '@/assets/images/bg_element_3.png'
import Footer from './pages/_components/footer'

interface Props {
    children?: ReactNode
}

const PublicLayout = (_props: Props) => {
    return (
        <>
            <div className="relative overflow-hidden">
                <img
                    src={bg_element_1}
                    className="absolute left-[50%] -z-30 h-auto w-[1681px] -translate-x-[50%]"
                    alt="background"
                />
                <img
                    src={bg_element_2}
                    className="absolute left-[90%] -z-30 h-auto w-[1681px] -translate-x-[50%] 2xl:left-[60%]"
                    alt="background"
                />
                <img
                    src={bg_element_3}
                    className="absolute left-[50%] -z-40 h-auto w-[1440px] -translate-x-[50%] opacity-30 2xl:left-[50%]"
                    alt="background"
                />
                <main className="w-full">
                    <NavBar />
                    <Outlet />
                    <Footer />
                </main>
            </div>
        </>
    )
}

export default PublicLayout
