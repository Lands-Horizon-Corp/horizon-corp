import { Attribute, Extension } from '@tiptap/core';
import { Plugin, PluginKey } from 'prosemirror-state';
import { Decoration, DecorationSet } from 'prosemirror-view';
import { Node } from 'prosemirror-model';

export interface PaginationOptions {
  pageHeight: number;
  pageWidth: number;
  pageMargin: number;
}

declare module '@tiptap/core' {
  interface Commands<ReturnType> {
    pagination: {
      setPaginationOptions: (options: Partial<PaginationOptions>) => ReturnType;
    };
  }
}

const createPageBreak = (pos: number) => {
  return Decoration.widget(pos, () => {
    const pageBreak = document.createElement('div');
    pageBreak.className = 'page-break';
    pageBreak.style.height = '20px';
    pageBreak.style.width = '100%';
    pageBreak.style.borderTop = '2px dashed #ccc';
    pageBreak.style.marginTop = '20px';
    pageBreak.style.marginBottom = '20px';
    pageBreak.style.pageBreakAfter = 'always';
    return pageBreak;
  });
};

export const Pagination = Extension.create<PaginationOptions>({
  name: 'pagination',

  addOptions() {
    return {
      pageHeight: 1056,
      pageWidth: 816,
      pageMargin: 96,
    };
  },

  addCommands() {
    return {
      setPaginationOptions:
        (options: Partial<PaginationOptions>) =>
        ({ tr, dispatch }) => {
          if (dispatch) {
            tr.setMeta('paginationOptions', options);
          }
          return true;
        },
    };
  },

  addProseMirrorPlugins() {
    const pluginKey = new PluginKey('pagination');

    return [
      new Plugin({
        key: pluginKey,
        state: {
          init: () => ({ ...this.options }),
          apply: (tr, value) => {
            const newOptions = tr.getMeta('paginationOptions');
            if (newOptions) {
              return { ...value, ...newOptions };
            }
            return value;
          },
        },
        props: {
          decorations: (state) => {
            const { doc } = state;
            const decorations: Decoration[] = [];
            let currentHeight = 0;
          
            const options = pluginKey.getState(state);
            const { pageHeight, pageMargin } = options;

            doc.descendants((node: Node, pos: number) => {
              if (!this.editor || !this.editor.view || this.editor.view.isDestroyed) {
                return;
              }
          
              if (pos < 0 || pos > doc.content.size) {
                return;
              }
          
              const dom = this.editor.view.nodeDOM(pos);
              if (!dom || !(dom instanceof HTMLElement)) {
                return;
              }

              if (['tableRow', 'tableCell', 'tableHeader'].includes(node.type.name)) {
                currentHeight = 0
                return;
              }
              
              const nodeHeight = node.isBlock ? dom.offsetHeight || 0 : 0;
              console.log('dom offset', dom.offsetHeight)
              if (currentHeight + nodeHeight > pageHeight - 2 * pageMargin) {
                console.log('page break', currentHeight)
                decorations.push(createPageBreak(pos));
                currentHeight = 0;
              }
              currentHeight += nodeHeight;
            });
             return DecorationSet.create(doc, decorations);
          }
        },
      }),
    ];
  },

  addGlobalAttributes() {
    return [
      {
        types: ['textStyle'],
        attributes: {
          class: {
            default: null,
            parseHTML: (element: HTMLElement) => element.getAttribute('class'),
            renderHTML: (attributes: Record<string, any>) => {
              if (!attributes.class) {
                return {};
              }
              return { class: attributes.class };
            },
          } as Attribute,
        },
      },
    ];
  },
});