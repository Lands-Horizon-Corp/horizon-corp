import { Toggle } from '@/components/ui/toggle'
import { Button } from '@/components/ui/button'

//icons
import {
    ToolbarItalicIcon,
    LuListOrderedIcon,
    LuUndo2Icon,
    LuRedo2Icon,
    PiTextStrikethroughLightIcon,
    PiListBulletsBoldIcon,
    PiCodeBlockIcon,
    FaBoldIcon,
    FaRulerHorizontalIcon,
    IoIosCodeIcon,
    TbClearFormattingIcon,
    BsTextParagraphIcon,
    TbBlockquoteIcon,
} from '@/components/icons'

import { THeadingLevel } from '.'

type Props = {
    editor: any
    toggleHeading: (level: THeadingLevel) => void
    activeHeading: THeadingLevel | null
}

const Toolbar = ({ editor, toggleHeading, activeHeading }: Props) => {
    if (!editor) {
        return null
    }
    return (
        <div className="flex w-fit min-w-fit flex-wrap space-x-2 border">
            <Toggle
                pressed={editor.isActive('italic')}
                onClick={() => editor.chain().focus().toggleItalic().run()}
                size="sm"
                aria-label="Toggle italic"
            >
                <ToolbarItalicIcon className="size-4" />
            </Toggle>
            <Toggle
                pressed={editor.isActive('bold')}
                onClick={() => editor.chain().focus().toggleBold().run()}
                size="sm"
                aria-label="Toggle Bold"
            >
                <FaBoldIcon className="size-4" />
            </Toggle>
            <Toggle
                pressed={editor.isActive('strike')}
                onClick={() => editor.chain().focus().toggleStrike().run()}
            >
                <PiTextStrikethroughLightIcon className="size-4" />
            </Toggle>
            <Toggle
                pressed={editor.isActive('code')}
                onClick={() => editor.chain().focus().toggleCode().run()}
            >
                <IoIosCodeIcon className="size-4" />
            </Toggle>
            <Toggle
                onClick={() => editor.chain().focus().unsetAllMarks().run()}
            >
                <TbClearFormattingIcon className="size-4" />
            </Toggle>
            <Toggle
                onClick={() => editor.chain().focus().setParagraph().run()}
                className={editor.isActive('paragraph') ? 'is-active' : ''}
            >
                <BsTextParagraphIcon className="size-4" />
            </Toggle>
            {Array.from({ length: 4 }, (_, i) => {
                const level = i + 1
                return (
                    <Toggle
                        key={level}
                        onClick={() => toggleHeading(level as THeadingLevel)}
                        pressed={activeHeading === level}
                    >
                        {`H${level}`}
                    </Toggle>
                )
            })}
            <Toggle
                pressed={editor.isActive('bulletList')}
                onPressedChange={() =>
                    editor.chain().focus().toggleBulletList().run()
                }
            >
                <PiListBulletsBoldIcon className="size-4" />
            </Toggle>
            <Toggle
                onClick={() => editor.chain().focus().toggleOrderedList().run()}
                className={editor.isActive('orderedList') ? 'is-active' : ''}
            >
                <LuListOrderedIcon className="size-4" />
            </Toggle>
            <Button
                variant={'ghost'}
                onClick={() => editor.chain().focus().toggleCodeBlock().run()}
                className={editor.isActive('codeBlock') ? 'is-active' : ''}
            >
                <PiCodeBlockIcon className="size-4" />
            </Button>
            <Button
                variant={'ghost'}
                onClick={() => editor.chain().focus().toggleBlockquote().run()}
                className={editor.isActive('blockquote') ? 'is-active' : ''}
            >
                <TbBlockquoteIcon className="size-4" />
            </Button>
            <Button
                variant={'ghost'}
                onClick={() => editor.chain().focus().setHorizontalRule().run()}
            >
                <FaRulerHorizontalIcon className="size-4" />
            </Button>
            <Button
                variant={'ghost'}
                onClick={() => editor.chain().focus().undo().run()}
            >
                <LuUndo2Icon className="size-4" />
            </Button>
            <Button
                variant={'ghost'}
                onClick={() => editor.chain().focus().redo().run()}
            >
                <LuRedo2Icon className="size-4" />
            </Button>
        </div>
    )
}

export default Toolbar
