import StarterKit from '@tiptap/starter-kit'
import { useEditor, EditorContent } from '@tiptap/react'
import { useState } from 'react'
import Toolbar from './toolbar'
import { IBaseComp } from '@/types/component/base'

interface Props extends IBaseComp {
    content: string
    spellCheck: boolean
}

export type THeadingLevel = 1 | 2 | 3 | 4

const TextEditor = ({ content, className, spellCheck }: Props) => {
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
                class: `toolbar-custom prose-sm sm:prose lg:prose-lg rounded-lg p-2 px-5 focus:outline-none ${className ?? ''} `,
            },
        },
    })

    const toggleHeading = (level: THeadingLevel) => {
        editor?.chain().focus().toggleHeading({ level }).run()
        setActiveHeading(level)
    }

    return (
        <div className="w-fit">
            <Toolbar
                activeHeading={activeHeading}
                toggleHeading={toggleHeading}
                editor={editor}
            />
            <EditorContent editor={editor} />
        </div>
    )
}

export default TextEditor
