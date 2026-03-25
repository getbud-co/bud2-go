import { ReactNode } from 'react';
type Placement = "top" | "bottom" | "left" | "right";
interface TooltipProps {
    /** Conteúdo do tooltip (texto ou ReactNode). */
    content: ReactNode;
    /** Posição preferida. O tooltip reposiciona se não couber. */
    placement?: Placement;
    /** Delay em ms antes de exibir (padrão: 200). */
    delay?: number;
    /** Desabilita o tooltip sem removê-lo da árvore. */
    disabled?: boolean;
    /** Elemento filho que ativa o tooltip. */
    children: ReactNode;
}
export declare function Tooltip({ content, placement, delay, disabled, children, }: TooltipProps): import("react/jsx-runtime").JSX.Element;
export {};
