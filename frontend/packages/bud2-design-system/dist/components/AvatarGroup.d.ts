export type AvatarGroupSize = "xs" | "sm" | "md";
export interface AvatarGroupItem {
    /** URL da foto. */
    src?: string;
    /** Iniciais (fallback sem foto). */
    initials?: string;
    /** Texto alternativo. */
    alt?: string;
}
interface AvatarGroupProps {
    /** Tamanho do grupo (xs | sm | md). */
    size?: AvatarGroupSize;
    /** Lista de avatares a exibir. */
    avatars: AvatarGroupItem[];
    /** Quantidade máxima de avatares visíveis antes do "+N". */
    maxVisible?: number;
    /** Exibe botão de adicionar usuário. */
    showAddButton?: boolean;
    /** Callback do botão de adicionar. */
    onAddClick?: () => void;
    className?: string;
}
export declare function AvatarGroup({ size, avatars, maxVisible, showAddButton, onAddClick, className, }: AvatarGroupProps): import("react/jsx-runtime").JSX.Element;
export {};
