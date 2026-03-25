import { ComponentType, ButtonHTMLAttributes } from 'react';
type ButtonVariant = "primary" | "secondary" | "tertiary" | "danger";
type ButtonSize = "sm" | "md" | "lg";
interface IconProps {
    size?: number | string;
    weight?: "regular";
}
interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    variant?: ButtonVariant;
    size?: ButtonSize;
    leftIcon?: ComponentType<IconProps>;
    rightIcon?: ComponentType<IconProps>;
    loading?: boolean;
}
export declare const Button: import('react').ForwardRefExoticComponent<ButtonProps & import('react').RefAttributes<HTMLButtonElement>>;
export {};
