import { ComponentType, ReactNode } from 'react';
interface IconProps {
    size?: number | string;
    weight?: "regular";
    className?: string;
}
export interface DropdownItem {
    id: string;
    label: string;
    icon?: ComponentType<IconProps>;
    description?: string;
}
type DropdownVariant = "primary" | "secondary" | "tertiary";
type DropdownSize = "sm" | "md" | "lg";
interface DropdownButtonProps {
    /** Lista de itens do menu */
    items: DropdownItem[];
    /** Callback ao selecionar um item */
    onSelect: (item: DropdownItem) => void;
    /** Ícone à esquerda do botão */
    leftIcon?: ComponentType<IconProps>;
    /** Variante visual do botão */
    variant?: DropdownVariant;
    /** Tamanho do botão */
    size?: DropdownSize;
    /** Habilita campo de busca no menu */
    searchable?: boolean;
    /** Placeholder do campo de busca */
    searchPlaceholder?: string;
    /** Conteúdo do botão (label) */
    children: ReactNode;
    /** Desabilitado */
    disabled?: boolean;
    /** Classe CSS adicional */
    className?: string;
}
export declare function DropdownButton({ items, onSelect, leftIcon: LeftIcon, variant, size, searchable, searchPlaceholder, children, disabled, className, }: DropdownButtonProps): import("react/jsx-runtime").JSX.Element;
export {};
