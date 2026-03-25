import { ReactNode, HTMLAttributes } from 'react';
type AlertVariant = "info" | "success" | "warning" | "error";
interface AlertProps extends HTMLAttributes<HTMLDivElement> {
    variant?: AlertVariant;
    title: string;
    children?: ReactNode;
    onDismiss?: () => void;
    action?: {
        label: string;
        onClick: () => void;
    };
}
export declare function Alert({ variant, title, children, onDismiss, action, className, ...rest }: AlertProps): import("react/jsx-runtime").JSX.Element;
export {};
