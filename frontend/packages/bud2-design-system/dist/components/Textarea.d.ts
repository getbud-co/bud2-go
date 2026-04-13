import { TextareaHTMLAttributes, ReactNode } from 'react';
type MessageType = "error" | "attention" | "success";
interface TextareaProps extends Omit<TextareaHTMLAttributes<HTMLTextAreaElement>, "size"> {
    label?: ReactNode;
    message?: string;
    messageType?: MessageType;
}
export declare const Textarea: import('react').ForwardRefExoticComponent<TextareaProps & import('react').RefAttributes<HTMLTextAreaElement>>;
export {};
