import { ComponentType, InputHTMLAttributes, ReactNode } from 'react';
type MessageType = "error" | "attention" | "success";
interface IconProps {
    size?: number | string;
    weight?: "regular";
    "aria-hidden"?: boolean | "true" | "false";
}
type InputSize = "sm" | "md" | "lg";
interface InputProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "size"> {
    /** Tamanho do input — combina com Button sizes (default: "md") */
    size?: InputSize;
    label?: ReactNode;
    leftIcon?: ComponentType<IconProps>;
    rightIcon?: ComponentType<IconProps>;
    message?: string;
    messageType?: MessageType;
}
export declare const Input: import('react').ForwardRefExoticComponent<InputProps & import('react').RefAttributes<HTMLInputElement>>;
export {};
