import { ReactNode } from 'react';
interface BaseGroupProps {
    label?: string;
    name?: string;
    disabled?: boolean;
    className?: string;
    children: ReactNode;
}
interface SingleGroupProps extends BaseGroupProps {
    multiple?: false;
    value?: string;
    defaultValue?: string;
    onChange?: (value: string | undefined) => void;
}
interface MultipleGroupProps extends BaseGroupProps {
    multiple: true;
    value?: string[];
    defaultValue?: string[];
    onChange?: (value: string[]) => void;
}
type ChoiceBoxGroupProps = SingleGroupProps | MultipleGroupProps;
export declare function ChoiceBoxGroup(props: ChoiceBoxGroupProps): import("react/jsx-runtime").JSX.Element;
interface ChoiceBoxLink {
    text: string;
    href?: string;
    onClick?: () => void;
}
interface ChoiceBoxProps {
    value: string;
    title: string;
    description?: string;
    link?: ChoiceBoxLink;
    disabled?: boolean;
    className?: string;
}
export declare function ChoiceBox({ value, title, description, link, disabled: itemDisabled, className, }: ChoiceBoxProps): import("react/jsx-runtime").JSX.Element;
export {};
