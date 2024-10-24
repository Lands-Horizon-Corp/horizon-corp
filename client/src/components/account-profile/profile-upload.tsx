import { DragEvent, useState } from 'react'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { ImageIcon } from '../icons'
import { toast } from 'sonner'
import { cn } from '@/lib'

const ProfileUpload = () => {
    const [isDragging, setIsDragging] = useState(false)
    const [files, setFiles] = useState<FileList | null>(null)

    const handleDrop = (event: DragEvent<HTMLDivElement>) => {
        event.preventDefault()

        const droppedFiles = event.dataTransfer.files
        const droppedFilesArray = Array.from(droppedFiles)

        if (droppedFilesArray.length === 1) {
            const file = droppedFilesArray[0]

            if (file.type.startsWith('image/')) {
                const dataTransfer = new DataTransfer()
                dataTransfer.items.add(file)
                setFiles(dataTransfer.files)
            } else {
                toast.error('Only image files are allowed!')
            }
        } else if (droppedFilesArray.length > 1) {
            toast.error('Please drop only one image file!')
        } else {
            toast.error('No file was dropped!')
        }

        setIsDragging(false)
    }
    return (
        <div>
            <div
                className={cn(
                    'inset pointer-events-none absolute left-0 top-0 flex size-full items-center justify-center rounded-2xl bg-background/70 opacity-0 backdrop-blur duration-500 ease-in-out',
                    isDragging && 'opacity-100'
                )}
            >
                <p className="text-foreground/80">Drop Picture</p>
            </div>
            {files === null && (
                <div>
                    <Label htmlFor="picture" className="group cursor-pointer">
                        <div
                            onDrop={handleDrop}
                            onDragOver={(e) => {
                                e.preventDefault()
                                setIsDragging(true)
                            }}
                            onDragLeave={(e) => {
                                e.preventDefault()
                                setIsDragging(false)
                            }}
                            className="flex flex-col items-center justify-center rounded-xl border border-dashed p-8"
                        >
                            <ImageIcon className="size-16 text-secondary" />
                            <span className="duration-300 ease-in-out hover:text-foreground/80">
                                Drop your image here
                            </span>
                        </div>
                    </Label>
                    <Input id="picture" type="file" className="hidden" onChange={(e)=> {
                        if(e.target.files){
                            setFiles(e.target.files)
                        }
                    }} />
                </div>
            )}
            {files !== null && files.item(0) !== null && (
                <img src={URL.createObjectURL(files.item(0)!)} alt="Preview" />
            )}
        </div>
    )
}

export default ProfileUpload
