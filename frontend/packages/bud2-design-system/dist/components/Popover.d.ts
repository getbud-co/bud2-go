import { ComponentType, ReactNode } from 'react';
interface IconProps {
    size?: number | string;
    weight?: "regular";
    className?: string;
}
export interface PopoverItem {
    id: string;
    label: string;
    icon?: ComponentType<IconProps>;
    /** URL de imagem exibida no lugar do ícone (ex: logo de empresa) */
    image?: string;
    onClick?: () => void;
    submenu?: ReactNode;
    /** Número exibido como badge (ex: contagem de seleções no submenu) */
    badge?: number;
    /** Estilo destrutivo (vermelho) para ações como logout/excluir */
    danger?: boolean;
    /** Renderiza um separador (linha horizontal) antes deste item */
    divider?: boolean;
}
interface PopoverProps {
    items: PopoverItem[];
    open: boolean;
    onClose: () => void;
    /** Ref do elemento trigger para posicionamento e click-outside */
    anchorRef: React.RefObject<HTMLElement | null>;
    /** Rótulo acessível para o menu */
    ariaLabel?: string;
}
export declare function Popover({ items, open, onClose, anchorRef, ariaLabel }: PopoverProps): import('react').ReactPortal | null;
export {};
