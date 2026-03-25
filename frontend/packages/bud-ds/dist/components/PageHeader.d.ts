import { ReactNode, ButtonHTMLAttributes } from 'react';
interface SearchButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    className?: string;
}
export declare const SearchButton: import('react').ForwardRefExoticComponent<SearchButtonProps & import('react').RefAttributes<HTMLButtonElement>>;
interface NotificationButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    /** Exibe indicador de notificações não lidas */
    hasUnread?: boolean;
    className?: string;
}
export declare const NotificationButton: import('react').ForwardRefExoticComponent<NotificationButtonProps & import('react').RefAttributes<HTMLButtonElement>>;
interface AssistantButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    /** Label exibido ao lado do ícone (default: "Assistente") */
    label?: string;
    /** Indica que o painel do assistente está aberto — troca ícone para X e aplica estado ativo */
    active?: boolean;
    className?: string;
}
export declare const AssistantButton: import('react').ForwardRefExoticComponent<AssistantButtonProps & import('react').RefAttributes<HTMLButtonElement>>;
interface PageHeaderProps {
    /** Título da página */
    title: string;
    /** Ações à direita (SearchButton, NotificationButton, AssistantButton, etc.) */
    children?: ReactNode;
    className?: string;
}
export declare function PageHeader({ title, children, className, }: PageHeaderProps): import("react/jsx-runtime").JSX.Element;
export {};
