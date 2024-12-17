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
    pageBreak.style.borderTop = '1px dashed #ccc';
    pageBreak.style.marginTop = '10px';
    pageBreak.style.marginBottom = '10px';
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

            console.log('currentHeight', currentHeight)
          
            doc.descendants((node: Node, pos: number) => {
              // Check if the editor and view are initialized
              if (!this.editor || !this.editor.view || this.editor.view.isDestroyed) {
                // console.log('Editor or editor view is not initialized or has been destroyed.');
                return;
              }
          
              // Ensure the position is within bounds
              if (pos < 0 || pos > doc.content.size) {
                // console.log(Invalid position ${pos});
                return;
              }
          
              // Delay execution to ensure DOM is stable
              const dom = this.editor.view.nodeDOM(pos);
              if (!dom || !(dom instanceof HTMLElement)) {
                // console.log(nodeDOM is null for position ${pos}. Skipping node processing.);
                return;
              }
        
            //   // Handle tables specifically
            //   // if (node.type.name === 'table') {
            //   //   const tableRows = Array.from(dom.querySelectorAll('tr'));
            //   //   let tableHeight = 0;
            //   //   let remainingHeight = pageHeight - currentHeight - 2 * pageMargin;
        
            //   //   let fragment = document.createElement('table');
            //   //   fragment.innerHTML = '<tbody></tbody>';
            //   //   let tbody = fragment.querySelector('tbody')!;
                
            //   //   for (let row of tableRows) {
            //   //     const rowHeight = row.offsetHeight || 0;
        
            //   //     if (tableHeight + rowHeight > remainingHeight) {
            //   //       if (tbody.children.length > 0) {
            //   //         decorations.push(Decoration.widget(pos, () => fragment));
            //   //         decorations.push(createPageBreak(pos));
            //   //       }
        
            //   //       currentHeight = 0;
            //   //       remainingHeight = pageHeight - 2 * pageMargin;
            //   //       tableHeight = 0;
        
            //   //       fragment = document.createElement('table');
            //   //       fragment.innerHTML = '<tbody></tbody>';
            //   //       tbody = fragment.querySelector('tbody')!;
            //   //     }
        
            //   //     tbody.appendChild(row.cloneNode(true));
            //   //     tableHeight += rowHeight;
            //   //   }
        
            //   //   if (tbody.children.length > 0) {
            //   //     decorations.push(Decoration.widget(pos, () => fragment));
            //   //   }
        
            //   //   currentHeight += tableHeight;
            //   //   return;
            //   // }
        
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

            // console.log('decorations', decorations)
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