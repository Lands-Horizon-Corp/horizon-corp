import { create } from 'zustand'
import { IMediaResource } from '@/server/types'

interface ImagePreviewState {
    isOpen: boolean
    Images?: IMediaResource[]
    focusIndex: number
}

interface ImagePreviewActions {
    setIsOpen: (isOpen: boolean) => void
    setFocusIndex: (index: number) => void
    onClose: () => void
}

interface ImagePreviewProps extends ImagePreviewState, ImagePreviewActions {}

export const useImagePreview = create<ImagePreviewProps>()((set) => ({
    focusIndex: 0,
    isOpen: false,
    setFocusIndex: (index) =>
        set((state) => ({
            ...state,
            focusIndex: index,
        })),
    onClose: () =>
        set((state) => ({
            ...state,
            isOpen: false,
        })),
    setIsOpen: (isOpen) =>
        set((state) => ({
            ...state,
            isOpen: isOpen,
        })),
}))
