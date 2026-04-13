export interface BreadcrumbItem {
    label: string;
    href?: string;
    onClick?: () => void;
}
interface BreadcrumbProps {
    items: BreadcrumbItem[];
    /** Índice do passo atual (0-based). Itens antes deste são links clicáveis; o atual e posteriores não são. */
    current?: number;
}
export declare function Breadcrumb({ items, current }: BreadcrumbProps): import("react/jsx-runtime").JSX.Element;
export {};
