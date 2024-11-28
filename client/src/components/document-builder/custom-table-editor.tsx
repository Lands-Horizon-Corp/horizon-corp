import Table from "@tiptap/extension-table"
import TableCell from "@tiptap/extension-table-cell"

export const CustomTableCell = TableCell.extend({
    addAttributes() {
        return {
            ...this.parent?.(),
            backgroundColor: {
                default: null,
                parseHTML: (element) => element.style.backgroundColor || null,
                renderHTML: (attributes) => {
                    console.log('attributes', attributes)
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