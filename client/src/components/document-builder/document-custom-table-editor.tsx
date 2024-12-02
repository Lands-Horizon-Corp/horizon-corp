import Table from '@tiptap/extension-table'
import TableCell from '@tiptap/extension-table-cell'
import TableRow from '@tiptap/extension-table-row'
// import TableHeader from '@tiptap/extension-table-header'

export const CustomTable = Table.extend({
    addAttributes() {
        return {
            class: {
                renderHTML: (attributes) => {
                    return {
                        class: attributes.class,
                    }
                },
            },
        }
    },
})

export const CustomTableRow = TableRow.extend({
    addAttributes() {
        return {
            ...this.parent?.(),
            rowheight: {
                default: 60, // Default height for rows
                renderHTML: (attributes) => {
                    console.log('rowheight attribute:', attributes.rowheight)
                    if (!attributes.rowheight) return {}
                    return {
                        style: `height: ${attributes.rowheight}px;`,
                    }
                },
                parseHTML: (element) => {
                    const height = element.style.height?.replace('px', '')
                    return height ? { rowheight: parseInt(height, 10) } : {}
                },
            },
        }
    },
})

// export const CustomTableHeader = TableHeader.extend({
//     addAttributes() {
//         return {
//             ...this.parent?.(),
//             headerRowheight: {
//                 default: 10,
//                 renderHTML: (attributes) =>{
//                     console.log('height attrib', attributes)
//                     if(!attributes.headerRowheight) return
//                     return {
//                         style: `height: ${attributes.headerRowheight}px;`,
//                     }
//                 }
//             }
//         }
//     },
// })

export const CustomTableCell = TableCell.extend({
    addAttributes() {
        return {
            ...this.parent?.(),
            backgroundColor: {
                default: null,
                parseHTML: (element) => element.style.backgroundColor || null,
                renderHTML: (attributes) => {
                    if (!attributes.backgroundColor) {
                        return {}
                    }
                    return {
                        style: `background-color: ${attributes.backgroundColor};`,
                    }
                },
            },
            borderLeft: {
                default: null,
                parseHTML: (element) => element.style.borderLeft || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderLeft) {
                        return {}
                    }
                    return {
                        style: `border-left: ${attributes.borderLeft};`,
                    }
                },
            },
            borderRight: {
                default: null,
                parseHTML: (element) => element.style.borderRight || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderRight) {
                        return {}
                    }
                    return {
                        style: `border-right: ${attributes.borderRight};`,
                    }
                },
            },
            borderTop: {
                default: null,
                parseHTML: (element) => element.style.borderTop || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderTop) {
                        return {}
                    }
                    return {
                        style: `border-top: ${attributes.borderTop};`,
                    }
                },
            },
            borderBottom: {
                default: null,
                parseHTML: (element) => element.style.borderBottom || null,
                renderHTML: (attributes) => {
                    if (!attributes.borderBottom) {
                        return {}
                    }
                    return {
                        style: `border-bottom: ${attributes.borderBottom};`,
                    }
                },
            },
        }
    },
})
