import { InputHTMLAttributes } from 'react';
interface ToggleProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "size" | "type"> {
    label?: string;
    description?: string;
}
export declare const Toggle: import('react').ForwardRefExoticComponent<ToggleProps & import('react').RefAttributes<HTMLInputElement>>;
export {};
