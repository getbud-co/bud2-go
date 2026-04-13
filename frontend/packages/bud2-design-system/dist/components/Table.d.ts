import { ReactNode, HTMLAttributes, ThHTMLAttributes, TdHTMLAttributes } from 'react';
interface TableProps extends HTMLAttributes<HTMLDivElement> {
    variant?: "divider" | "striped";
    /** Adiciona box-shadow ao container (default: true) */
    elevated?: boolean;
    /** Exibe borda e border-radius no container (default: true). Use false quando a tabela está dentro de um Card. */
    bordered?: boolean;
    selectable?: boolean;
    selectedRows?: Set<string>;
    rowIds?: string[];
    onSelectRow?: (id: string, checked: boolean) => void;
    onSelectAll?: (checked: boolean) => void;
}
export declare function Table({ variant, elevated, bordered, selectable, selectedRows, rowIds, onSelectRow, onSelectAll, children, className, ...rest }: TableProps): import("react/jsx-runtime").JSX.Element;
interface TableCardHeaderProps {
    title: string;
    badge?: ReactNode;
    actions?: ReactNode;
}
export declare function TableCardHeader({ title, badge, actions }: TableCardHeaderProps): import("react/jsx-runtime").JSX.Element;
interface TableBulkActionsProps {
    /** Quantidade de itens selecionados */
    count: number;
    /** Callback ao clicar em desmarcar tudo */
    onClear: () => void;
    /** Botões de ação (ex: Excluir, Exportar) */
    children: ReactNode;
}
export declare function TableBulkActions({ count, onClear, children }: TableBulkActionsProps): import('react').ReactPortal | null;
export declare function TableContent({ children, className, ...rest }: HTMLAttributes<HTMLDivElement>): import("react/jsx-runtime").JSX.Element;
export declare function TableHead({ children, ...rest }: HTMLAttributes<HTMLTableSectionElement>): import("react/jsx-runtime").JSX.Element;
export declare function TableBody({ children, ...rest }: HTMLAttributes<HTMLTableSectionElement>): import("react/jsx-runtime").JSX.Element;
interface TableRowProps extends HTMLAttributes<HTMLTableRowElement> {
    rowId?: string;
}
export declare function TableRow({ rowId, children, className, ...rest }: TableRowProps): import("react/jsx-runtime").JSX.Element;
type SortDirection = "asc" | "desc" | undefined;
interface TableHeaderCellProps extends ThHTMLAttributes<HTMLTableCellElement> {
    sortable?: boolean;
    sortDirection?: SortDirection;
    onSort?: () => void;
    isCheckbox?: boolean;
    selectAllLabel?: string;
    sortLabel?: string;
}
export declare function TableHeaderCell({ sortable, sortDirection, onSort, isCheckbox, selectAllLabel, sortLabel, children, className, ...rest }: TableHeaderCellProps): import("react/jsx-runtime").JSX.Element;
interface TableCellProps extends TdHTMLAttributes<HTMLTableCellElement> {
    isCheckbox?: boolean;
    rowId?: string;
    selectionLabel?: string;
}
export declare function TableCell({ isCheckbox, rowId, selectionLabel, children, className, ...rest }: TableCellProps): import("react/jsx-runtime").JSX.Element;
interface TablePaginationProps {
    currentPage: number;
    totalPages: number;
    onPageChange: (page: number) => void;
}
export declare function TablePagination({ currentPage, totalPages, onPageChange, }: TablePaginationProps): import("react/jsx-runtime").JSX.Element;
export {};
