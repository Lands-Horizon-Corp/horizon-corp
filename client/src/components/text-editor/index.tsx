import { useState } from 'react'
import StarterKit from '@tiptap/starter-kit'
import { useEditor, EditorContent } from '@tiptap/react'

import Toolbar from './toolbar'
import { IBaseComp } from '@/types'

interface Props extends IBaseComp {
    content?: string
    spellCheck?: boolean
    showToolbar?: boolean
    onChange: (content: string) => void
    isHeadingDisabled?: boolean
}

export type THeadingLevel = 1 | 2 | 3 | 4

const TextEditor = ({
    content = '',
    className = '',
    spellCheck = true,
    showToolbar = true,
    onChange,
    isHeadingDisabled = true,
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
        onUpdate({ editor }) {
            onChange(editor.getHTML())
        },
    })

    const toggleHeading = (level: THeadingLevel) => {
        editor?.chain().focus().toggleHeading({ level }).run()
        setActiveHeading(level)
    }
    return (
        <div className="w-full space-y-2">
            {showToolbar && editor && (
                <Toolbar
                    isHeadingDisabled={isHeadingDisabled}
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
