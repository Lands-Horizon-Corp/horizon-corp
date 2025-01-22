import { useEffect, useState } from 'react'
import StarterKit from '@tiptap/starter-kit'
import { useEditor, EditorContent } from '@tiptap/react'

import Toolbar from './toolbar'

import { cn } from '@/lib'
import { IBaseComp } from '@/types'

interface Props extends IBaseComp {
    content?: string
    spellCheck?: boolean
    placeholder?: string
    showToolbar?: boolean
    isHeadingDisabled?: boolean
    onChange: (content: string) => void
}

export type THeadingLevel = 1 | 2 | 3 | 4

const TextEditor = ({
    className,
    content = '',
    spellCheck = true,
    showToolbar = true,
    isHeadingDisabled = true,
    onChange,
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
                class: `toolbar-custom ${cn('', className)} `,
            },
        },
        onUpdate({ editor }) {
            onChange(editor.getHTML())
        },
    })

    useEffect(() => {
        if (editor && content !== editor.getHTML()) {
            editor.commands.setContent(content || '')
        }
    }, [content, editor])

    const toggleHeading = (level: THeadingLevel) => {
        editor?.chain().focus().toggleHeading({ level }).run()
        setActiveHeading(level)
    }
    return (
        <div className="w-full space-y-2">
            {showToolbar && editor && (
                <Toolbar
                    editor={editor}
                    activeHeading={activeHeading}
                    toggleHeading={toggleHeading}
                    isHeadingDisabled={isHeadingDisabled}
                />
            )}
            <EditorContent editor={editor} />
        </div>
    )
}

export default TextEditor
