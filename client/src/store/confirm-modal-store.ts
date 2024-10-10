import { create } from 'zustand'
import { ReactNode } from 'react'

interface ISharedConfirmModalProperty {
    showCancel: boolean
    showConfirm: boolean

    onClose: () => void
    onConfirm: () => void
    onCancel: () => void
}

interface IConfirmModalOnOpenData extends Partial<ISharedConfirmModalProperty> {
    title: string
    description?: string
    content?: ReactNode

    confirmString?: string
    cancelString?: string
    onClose?: () => void
    onCancel?: () => void
    onConfirm?: () => void
}

interface IConfirmModalStore extends ISharedConfirmModalProperty {
    isOpen: boolean
    showCancel: boolean
    showConfirm: boolean
    modalData?: IConfirmModalOnOpenData

    onOpen: (newModalData: IConfirmModalOnOpenData) => void
}

const useConfirmModalStore = create<IConfirmModalStore>((set) => ({
    isOpen: false,
    showCancel: false,
    showConfirm: false,
    onOpen: ({
        confirmString,
        cancelString,
        showConfirm,
        showCancel,
        ...newModalData
    }) =>
        set({
            isOpen: true,
            showConfirm: showConfirm ?? false,
            showCancel: showCancel ?? false,
            modalData: {
                ...newModalData,
                confirmString: confirmString ?? 'confirm',
                cancelString: cancelString ?? 'cancel',
            },
        }),

    onConfirm: () => {
        set((state) => {
            state.modalData?.onConfirm?.()
            return { isOpen: false }
        })
        setTimeout(() => set(() => ({ modalData: undefined })), 100)
    },
    onClose: () => {
        set((state) => {
            state.modalData?.onClose?.()
            return { isOpen: false }
        })
        setTimeout(() => set(() => ({ modalData: undefined })), 100)
    },
    onCancel: () => {
        set((state) => {
            state.modalData?.onCancel?.()
            return { isOpen: false }
        })
        setTimeout(() => set(() => ({ modalData: undefined })), 100)
    },
}))

export default useConfirmModalStore
