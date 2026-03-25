import { InputHTMLAttributes } from 'react';
type CheckboxSize = "sm" | "md";
interface CheckboxProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "size" | "type"> {
    size?: CheckboxSize;
    indeterminate?: boolean;
    label?: string;
    description?: string;
}
export declare const Checkbox: import('react').ForwardRefExoticComponent<CheckboxProps & import('react').RefAttributes<HTMLInputElement>>;
export {};
