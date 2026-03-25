import { ComponentType, ReactNode } from 'react';
type MessageType = "error" | "attention" | "success";
interface IconProps {
    size?: number | string;
    weight?: "regular";
}
export interface SelectOption {
    value: string;
    label: string;
}
type SelectSize = "sm" | "md" | "lg";
interface SelectBaseProps {
    /** Tamanho do trigger — combina com Button sizes (default: "md") */
    size?: SelectSize;
    label?: ReactNode;
    leftIcon?: ComponentType<IconProps>;
    placeholder?: string;
    options: SelectOption[];
    searchable?: boolean;
    searchPlaceholder?: string;
    message?: string;
    messageType?: MessageType;
    disabled?: boolean;
    className?: string;
}
interface SingleSelectProps extends SelectBaseProps {
    multiple?: false;
    value?: string;
    defaultValue?: string;
    onChange?: (value: string) => void;
}
interface MultiSelectProps extends SelectBaseProps {
    multiple: true;
    value?: string[];
    defaultValue?: string[];
    onChange?: (value: string[]) => void;
}
type SelectProps = SingleSelectProps | MultiSelectProps;
export declare function Select(props: SelectProps): import("react/jsx-runtime").JSX.Element;
export {};
