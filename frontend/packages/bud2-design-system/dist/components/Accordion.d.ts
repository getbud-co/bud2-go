import { ComponentType, ReactNode, HTMLAttributes } from 'react';
interface IconProps {
    size?: number | string;
    weight?: "regular";
    className?: string;
}
interface AccordionItemProps extends Omit<HTMLAttributes<HTMLDivElement>, "title" | "onToggle"> {
    /** Ícone à esquerda do título */
    icon?: ComponentType<IconProps>;
    /** Título exibido no trigger */
    title: ReactNode;
    /** Descrição opcional abaixo do título */
    description?: string;
    /** Conteúdo extra à direita do trigger (badge, ícone, etc.) */
    action?: ReactNode;
    /** Inicia aberto */
    defaultOpen?: boolean;
    /** Controlado: estado aberto */
    open?: boolean;
    /** Callback ao alternar */
    onToggle?: (open: boolean) => void;
    /** Desabilita interação */
    disabled?: boolean;
    children: ReactNode;
}
export declare function AccordionItem({ icon: Icon, title, description, action, defaultOpen, open: openProp, onToggle, disabled, children, className, ...rest }: AccordionItemProps): import("react/jsx-runtime").JSX.Element;
interface AccordionProps extends HTMLAttributes<HTMLDivElement> {
    /** Aplica fundo caramel-50 nos triggers */
    header?: boolean;
    children: ReactNode;
}
export declare function Accordion({ header, children, className, ...rest }: AccordionProps): import("react/jsx-runtime").JSX.Element;
export {};
