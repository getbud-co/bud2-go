import { ComponentType, ReactNode } from 'react';
interface IconProps {
    size?: number | string;
    weight?: "regular";
    className?: string;
}
export interface CommandItem {
    id: string;
    label: string;
    icon?: ComponentType<IconProps>;
    /** Palavras-chave para busca (não exibidas) */
    keywords?: string[];
    /** Texto auxiliar à direita (ex: atalho de teclado) */
    hint?: string;
}
export interface CommandGroup {
    label: string;
    items: CommandItem[];
}
interface CommandPaletteProps {
    /** Controla abertura da paleta */
    open: boolean;
    /** Callback para fechar */
    onClose: () => void;
    /** Callback ao selecionar um item */
    onSelect: (itemId: string) => void;
    /** Itens agrupados */
    groups: CommandGroup[];
    /** Placeholder do campo de busca */
    placeholder?: string;
    /** Texto exibido quando não há resultados */
    emptyMessage?: string;
    /** Conteúdo extra no footer (ex: atalhos) */
    footer?: ReactNode;
}
export declare function CommandPalette({ open, onClose, onSelect, groups, placeholder, emptyMessage, footer, }: CommandPaletteProps): import('react').ReactPortal | null;
export {};
