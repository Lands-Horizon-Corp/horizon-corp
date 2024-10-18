import StarterKit from '@tiptap/starter-kit'
import { useEditor, EditorContent } from '@tiptap/react'
import { useState } from 'react'
import Toolbar from './toolbar'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {
    content?: string
    spellCheck?: boolean
    showToolbar?: boolean
}

export type THeadingLevel = 1 | 2 | 3 | 4

const TextEditor = ({
    content = '',
    className = '',
    spellCheck = true,
    showToolbar = true,
}: Props) => {
    const [activeHeading, setActiveHeading] = useState<THeadingLevel | null>(
        null
    )

    const editor = useEditor({
        extensions: [
            StarterKit.configure({
                bulletList: {
                    keepMarks: false,
                    keepAttributes: false,
                },
            }),
        ],
        content: content,
        editorProps: {
            attributes: {
                spellcheck: spellCheck ? 'true' : 'false',
                class: `toolbar-custom ${className ?? ''} `,
            },
        },
    })

    const toggleHeading = (level: THeadingLevel) => {
        editor?.chain().focus().toggleHeading({ level }).run()
        setActiveHeading(level)
    }

    return (
        <div className="w-fit">
            {showToolbar && editor && (
                <Toolbar
                    activeHeading={activeHeading}
                    toggleHeading={toggleHeading}
                    editor={editor}
                />
            )}
            <EditorContent editor={editor} />
        </div>
    )
}

export default TextEditor
