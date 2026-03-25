import { ComponentType, ReactNode } from 'react';
interface IconProps {
    size?: number | string;
    weight?: "regular";
    className?: string;
}
interface FilterChipProps {
    /** Texto exibido no chip */
    label: string;
    /** Ícone Phosphor à esquerda */
    icon?: ComponentType<IconProps>;
    /** Callback ao clicar no chip (ex: abrir popover de edição) */
    onClick?: () => void;
    /** Callback ao remover o chip */
    onRemove?: () => void;
    /** Indica que o chip está sendo editado (focus ring) */
    active?: boolean;
    className?: string;
}
export declare function FilterChip({ label, icon: Icon, onClick, onRemove, active, className, }: FilterChipProps): import("react/jsx-runtime").JSX.Element;
interface FilterDropdownProps {
    /** Controla se o dropdown está aberto */
    open: boolean;
    /** Callback para fechar */
    onClose: () => void;
    /** Ref do elemento âncora para posicionamento */
    anchorRef: React.RefObject<HTMLElement | null>;
    /** Conteúdo do dropdown */
    children: ReactNode;
    className?: string;
    /** Posicionamento relativo ao âncora. Default: "bottom-start" */
    placement?: "bottom-start" | "right-start";
    /** Refs de elementos que NÃO devem disparar click-outside */
    ignoreRefs?: React.RefObject<HTMLElement | null>[];
    /** Se true, não renderiza overlay de fundo (útil para sub-menus) */
    noOverlay?: boolean;
}
export declare function FilterDropdown({ open, onClose, anchorRef, children, className, placement, ignoreRefs, noOverlay, }: FilterDropdownProps): import('react').ReactPortal | null;
export interface FilterOption {
    id: string;
    label: string;
    icon?: ComponentType<IconProps>;
}
interface FilterBarProps {
    /** Filtros disponíveis exibidos no popover "Adicionar filtro" */
    filters: FilterOption[];
    /** Callback quando um filtro é selecionado no popover */
    onAddFilter: (filterId: string) => void;
    /** Limpa todos os filtros — exibido somente quando há children */
    onClearAll?: () => void;
    /** Salvar visualização — exibido somente quando há children */
    onSaveView?: () => void;
    /** Label do botão salvar (default: "Salvar visualização") */
    saveViewLabel?: string;
    /** Ação primária adicional à direita */
    primaryAction?: ReactNode;
    /** Placeholder do campo de busca no popover */
    searchPlaceholder?: string;
    /** Abre o popover de adicionar filtro ao montar */
    defaultOpen?: boolean;
    /** FilterChips renderizados como children */
    children?: ReactNode;
    className?: string;
}
export declare function FilterBar({ filters, onAddFilter, onClearAll, onSaveView, saveViewLabel, primaryAction, searchPlaceholder, defaultOpen, children, className, }: FilterBarProps): import("react/jsx-runtime").JSX.Element;
export {};
