import TextEditor from '@/components/text-editor'

const PublicPracticeLayout = () => {
    return (
        <>
            <div className="m-auto mt-10 h-[100vh] w-[90%]">
                <TextEditor
                    spellCheck={true}
                    content={`<p>start typing...</p>`}
                />
            </div>
        </>
    )
}

export default PublicPracticeLayout
