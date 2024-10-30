import { MediaResource } from "@/horizon-corp/types"



const Album = () => {

}

const Actions = () => {
  return <>Actions</>
}

interface ImagePreviewProps {
  className: string,
  isOpen: boolean,
  Images: MediaResource[]

}

const ImagePreview = ({className, isOpen, Images}:ImagePreviewProps) => {
  

  return (
    <div className={`${className ?? ''}`}>
        ImagePreview
      
      <Actions/>
    </div>
  )
}

export default ImagePreview