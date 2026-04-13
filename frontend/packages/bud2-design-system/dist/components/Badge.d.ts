import { ComponentType, HTMLAttributes } from 'react';
type BadgeColor = "neutral" | "orange" | "wine" | "caramel" | "error" | "warning" | "success";
type BadgeSize = "sm" | "md" | "lg";
interface IconProps {
    size?: number | string;
    weight?: "regular";
}
interface BadgeProps extends Omit<HTMLAttributes<HTMLSpanElement>, "color"> {
    color?: BadgeColor;
    size?: BadgeSize;
    leftIcon?: ComponentType<IconProps>;
    rightIcon?: ComponentType<IconProps>;
}
export declare function Badge({ color, size, leftIcon: LeftIcon, rightIcon: RightIcon, children, className, ...rest }: BadgeProps): import("react/jsx-runtime").JSX.Element;
export {};
