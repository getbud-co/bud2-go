import { ComponentType, ReactNode } from 'react';
interface IconProps {
    size?: number | string;
    weight?: "regular";
    className?: string;
}
export interface NotificationItem {
    id: string;
    /** Ícone ou avatar à esquerda */
    icon?: ComponentType<IconProps>;
    /** Avatar image URL (prioridade sobre icon) */
    avatarUrl?: string;
    /** Título da notificação */
    title: ReactNode;
    /** Descrição complementar */
    description?: string;
    /** Timestamp relativo (ex: "há 5 min") */
    time: string;
    /** Indica se a notificação não foi lida */
    unread?: boolean;
}
interface NotificationPanelProps {
    /** Controla se o painel está aberto */
    open: boolean;
    /** Callback para fechar */
    onClose: () => void;
    /** Ref do elemento âncora para posicionamento */
    anchorRef: React.RefObject<HTMLElement | null>;
    /** Lista de notificações */
    notifications: NotificationItem[];
    /** Callback ao clicar numa notificação */
    onClickItem?: (id: string) => void;
    /** Callback para marcar todas como lidas */
    onMarkAllRead?: () => void;
    /** Callback para "Ver todas" */
    onViewAll?: () => void;
    /** Título do painel (default: "Notificações") */
    title?: string;
    /** Label de "Ver todas" (default: "Ver todas as notificações") */
    viewAllLabel?: string;
    /** Texto do estado vazio */
    emptyMessage?: string;
    className?: string;
}
export declare function NotificationPanel({ open, onClose, anchorRef, notifications, onClickItem, onMarkAllRead, onViewAll, title, viewAllLabel, emptyMessage, className, }: NotificationPanelProps): import('react').ReactPortal | null;
export {};
