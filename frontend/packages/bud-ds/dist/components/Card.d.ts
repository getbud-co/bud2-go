import { ReactNode, HTMLAttributes } from 'react';
interface CardProps extends HTMLAttributes<HTMLDivElement> {
    padding?: "none" | "sm" | "md" | "lg";
    shadow?: boolean;
    children: ReactNode;
}
export declare function Card({ padding, shadow, children, className, ...rest }: CardProps): import("react/jsx-runtime").JSX.Element;
interface CardHeaderProps extends HTMLAttributes<HTMLDivElement> {
    title: string;
    description?: string;
    action?: ReactNode;
}
export declare function CardHeader({ title, description, action, className, ...rest }: CardHeaderProps): import("react/jsx-runtime").JSX.Element;
interface CardBodyProps extends HTMLAttributes<HTMLDivElement> {
    children: ReactNode;
}
export declare function CardBody({ children, className, ...rest }: CardBodyProps): import("react/jsx-runtime").JSX.Element;
interface CardFooterProps extends HTMLAttributes<HTMLDivElement> {
    children: ReactNode;
}
export declare function CardFooter({ children, className, ...rest }: CardFooterProps): import("react/jsx-runtime").JSX.Element;
export declare function CardDivider(): import("react/jsx-runtime").JSX.Element;
export {};
