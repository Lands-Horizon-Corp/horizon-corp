import TextEditor from '@/components/text-editor'

const PublicPracticeLayout = () => {
    return (
        <>
            <div className="m-auto mt-10 h-[100vh] w-[90%]">
                <TextEditor
                    className=""
                    spellCheck={true}
                    content={`<p>start typing...</p>`}
                />
                <li>hello</li>
            </div>
        </>
    )
}

export default PublicPracticeLayout
