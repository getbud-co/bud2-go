import { ReactNode } from 'react';
/** Preset de alturas para casos comuns */
export declare const SKELETON_HEIGHTS: {
    readonly text: 14;
    readonly heading: 24;
    readonly subheading: 18;
    readonly button: 40;
    readonly input: 40;
    readonly avatar: 40;
    readonly avatarLg: 48;
};
export interface SkeletonProps {
    /** Forma do skeleton */
    variant?: "text" | "circular" | "rectangular" | "rounded";
    /** Largura (px ou string CSS) */
    width?: number | string;
    /** Altura (px ou string CSS) */
    height?: number | string;
    /** Desabilitar animação */
    animation?: boolean;
    /** Classe CSS adicional */
    className?: string;
}
export interface SkeletonContainerProps {
    /** Conteúdo (skeletons) */
    children: ReactNode;
    /** Texto para screen readers (default: "Carregando...") */
    loadingText?: string;
    /** Classe CSS adicional */
    className?: string;
}
/**
 * Wrapper acessível para composições de Skeleton.
 *
 * Adiciona `role="status"` e texto oculto para screen readers,
 * notificando tecnologias assistivas de que o conteúdo está sendo carregado.
 *
 * @example
 * ```tsx
 * <SkeletonContainer>
 *   <Skeleton variant="text" width={200} height={14} />
 *   <Skeleton variant="circular" width={40} height={40} />
 * </SkeletonContainer>
 * ```
 */
export declare function SkeletonContainer({ children, loadingText, className, }: SkeletonContainerProps): import("react/jsx-runtime").JSX.Element;
export declare function Skeleton({ variant, width, height, animation, className, }: SkeletonProps): import("react/jsx-runtime").JSX.Element;
