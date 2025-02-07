export interface IModalBase {
    modalState? : boolean
    setModalState? : (newState : boolean) => void
    modalTitle? : string
    modalDescription? : string
}