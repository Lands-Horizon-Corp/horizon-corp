import { EditorContent, useEditor } from '@tiptap/react';
import { Pagination } from '@/components/document-builder/pagination'; // Ensure the path is correct
import StarterKit from '@tiptap/starter-kit'; // Include basic editor functionality

const TestPaging = () => {

  const editor = useEditor({
    extensions: [
      StarterKit, // Base extensions for a working editor
      Pagination.configure({
        pageHeight: 1020, // Default height of the page
        pageWidth: 816,   // Default width of the page
        pageMargin: 96,   // Default margin of the page
      }),
    ],
    content: '<p>Hello, World!</p>', // Initial content for the editor
  });

  if (!editor) {
    return null;
  }

  console.log(editor)
  return (
    <div>
      <EditorContent className='w-full' editor={editor} />
    </div>
  );
};

export default TestPaging;
