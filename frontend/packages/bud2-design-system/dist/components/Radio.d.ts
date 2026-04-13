import { InputHTMLAttributes } from 'react';
type RadioSize = "sm" | "md";
interface RadioProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "size" | "type"> {
    size?: RadioSize;
    label?: string;
    description?: string;
}
export declare const Radio: import('react').ForwardRefExoticComponent<RadioProps & import('react').RefAttributes<HTMLInputElement>>;
export {};
