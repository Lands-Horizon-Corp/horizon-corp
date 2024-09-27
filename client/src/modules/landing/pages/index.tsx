import home_image_1 from '@/assets/images/home_image_1.png'
import bg_element_4 from '@/assets/images/bg_element_4.png'
import bg_element_5 from '@/assets/images/bg_element_5.png'
import home_image_2 from '@/assets/images/home_image_2.png'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'

interface Props {}

const LandingPage = (_props: Props) => {
    return (
        <div className="flex h-[100vh] justify-center px-6 font-inter sm:px-8 lg:px-[60px] xl:px-[124px]">
            <div className="h-fit w-full max-w-[1240px]">
                <h1 className="w-[80%] pt-2 text-[min(64px,5.5vw)] font-black capitalize md:pt-20 lg:leading-[4.8rem]">
                    Empowering Communities Through Cooperative Ownership
                </h1>
                <div className="flex w-full">
                    <div className="grow"></div>
                    <p className="w-full flex-none text-justify text-[min(25px,3.0vw)] font-semibold text-[#5A5A5A] md:w-[468px]">
                        Cooperatives embody the power of community, where shared
                        ownership and mutual aid transform economic challenges
                        into opportunities for progress and empowerment.
                    </p>
                </div>
                <img
                    src={home_image_1}
                    className="h-auto w-[350px] pt-5 md:w-[500px] xl:w-[651px] xl:-translate-y-20"
                    alt="iamge-1"
                ></img>
                <div className="relative space-y-5">
                    <img
                        src={bg_element_4}
                        className="absolute -top-12 -z-40 h-auto w-[985px] 2xl:left-0"
                        alt="background"
                    />
                    <div className="">
                        <p className="text-[min(25px,3.5vw)] font-bold">
                            Mission
                        </p>
                        <p className="mt-3 text-[min(25px,3.0vw)] font-normal lg:w-[80%] lg:leading-10">
                            We offer top-tier financial products and related
                            services to all sectors of society, supporting their
                            sustainable growth.
                        </p>
                    </div>
                    <div>
                        <p className="text-[min(25px,3.5vw)] font-bold">
                            Vision
                        </p>
                        <p className="mt-3 text-[min(25px,3.0vw)] font-normal lg:w-[80%] lg:leading-10">
                            A premier multi-purpose cooperative in Luzon
                            dedicated to serving both its members and the
                            community.
                        </p>
                    </div>
                </div>
                <div className="relative mt-20 w-full">
                    <img
                        src={bg_element_5}
                        className="absolute -top-20 right-0 -z-40 h-auto w-[178px]"
                        alt="background"
                    />
                    <h3 className="text-[min(25px,3.5vw)] font-bold">
                        A Glimpse of Visions and Successes Realized
                    </h3>
                    <div className="h-fit w-full space-y-5 self-center">
                        <img
                            src={home_image_2}
                            className="-top-20 right-0 -z-40 m-auto mt-5 h-auto w-[75%] text-center 2xl:left-0"
                            alt="background"
                        />
                        <div>
                            <p className="mt-10 text-justify indent-8 text-[min(23px,2.5vw)] xl:leading-[41px]">
                                Cooperatives originated in the early 19th
                                century, with the first established by the
                                Rochdale Pioneers in 1844 to help working-class
                                communities combat economic challenges. Over
                                time, cooperatives expanded across sectors like
                                agriculture, housing, and credit, providing
                                marginalized individuals with resources and
                                opportunities. Today, cooperatives empower
                                millions worldwide, promoting shared ownership
                                and mutual support.
                            </p>
                        </div>
                        <div className=" h-[130px] w-full flex justify-center items-center">
                            <Button
                                className={cn(
                                    'text-[min(18px,2.5vw)] h-10 xl:h-14 xl:px-5 rounded-full bg-green-500 hover:bg-green-500'
                                )}
                            >
                                Read more about us
                            </Button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default LandingPage
