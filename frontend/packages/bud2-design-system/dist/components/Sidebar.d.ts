import { ReactNode, ComponentType, HTMLAttributes } from 'react';
interface SidebarProps extends HTMLAttributes<HTMLElement> {
    children: ReactNode;
    /** Estado colapsado do sidebar */
    collapsed?: boolean;
    /** Callback ao clicar no botão de colapsar/expandir. Se omitido, o botão não aparece. */
    onCollapse?: () => void;
    /** Drawer aberto em mobile (≤768px) */
    mobileOpen?: boolean;
    /** Callback ao fechar o drawer mobile */
    onMobileClose?: () => void;
}
export declare function Sidebar({ children, className, collapsed, onCollapse, mobileOpen, onMobileClose, ...rest }: SidebarProps): import("react/jsx-runtime").JSX.Element;
interface SidebarHeaderProps {
    children: ReactNode;
    /** Conteúdo exibido no estado colapsado (ex: símbolo da marca) */
    collapsedContent?: ReactNode;
}
export declare function SidebarHeader({ children, collapsedContent }: SidebarHeaderProps): import("react/jsx-runtime").JSX.Element;
interface SidebarOrgSwitcherProps {
    icon?: ComponentType<{
        size?: number;
    }>;
    /** URL de imagem/logo da organização (substitui o ícone) */
    image?: string;
    label: string;
    onClick?: () => void;
}
export declare const SidebarOrgSwitcher: import('react').ForwardRefExoticComponent<SidebarOrgSwitcherProps & import('react').RefAttributes<HTMLButtonElement>>;
export declare function SidebarDivider(): import("react/jsx-runtime").JSX.Element;
interface SidebarNavProps {
    children: ReactNode;
    "aria-label"?: string;
}
export declare function SidebarNav({ children, "aria-label": ariaLabel, }: SidebarNavProps): import("react/jsx-runtime").JSX.Element;
interface SidebarGroupProps {
    label: string;
    children: ReactNode;
}
export declare function SidebarGroup({ label, children }: SidebarGroupProps): import("react/jsx-runtime").JSX.Element;
interface SidebarItemProps {
    icon: ComponentType<{
        size?: number;
    }>;
    label: string;
    href?: string;
    active?: boolean;
    defaultExpanded?: boolean;
    onClick?: () => void;
    children?: ReactNode;
}
export declare function SidebarItem({ icon: Icon, label, href, active, defaultExpanded, onClick, children, }: SidebarItemProps): import("react/jsx-runtime").JSX.Element;
interface SidebarSubItemProps {
    href?: string;
    active?: boolean;
    onClick?: () => void;
    children: ReactNode;
}
export declare function SidebarSubItem({ href, active, onClick, children, }: SidebarSubItemProps): import("react/jsx-runtime").JSX.Element;
interface SidebarFooterProps {
    children: ReactNode;
}
export declare function SidebarFooter({ children }: SidebarFooterProps): import("react/jsx-runtime").JSX.Element;
interface SidebarUserProps {
    name: string;
    role?: string;
    avatar: ReactNode;
    onClick?: () => void;
}
export declare const SidebarUser: import('react').ForwardRefExoticComponent<SidebarUserProps & import('react').RefAttributes<HTMLButtonElement>>;
export {};
