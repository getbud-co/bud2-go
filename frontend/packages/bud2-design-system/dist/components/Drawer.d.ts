import { ReactNode } from 'react';
type DrawerSide = "right" | "left";
type DrawerSize = "sm" | "md" | "lg";
export interface DrawerProps {
    /** Whether the drawer is open */
    open: boolean;
    /** Callback when the drawer should close */
    onClose: () => void;
    /** Which side the drawer slides from */
    side?: DrawerSide;
    /** Predefined width: sm (380px), md (480px), lg (640px). Ignored when `width` is set. */
    size?: DrawerSize;
    /** Custom CSS width (e.g. "33.333%", "520px"). Overrides `size`. */
    width?: string;
    /** Extra CSS class applied to the panel element */
    className?: string;
    /** Accessible label when not using DrawerHeader (alternative to aria-labelledby) */
    "aria-label"?: string;
    children: ReactNode;
}
export declare function Drawer({ open, onClose, side, size, width, className, "aria-label": ariaLabel, children, }: DrawerProps): import('react').ReactPortal | null;
interface DrawerHeaderProps {
    title: ReactNode;
    description?: ReactNode;
    onClose?: () => void;
    /** Extra action elements rendered before the close button (in the top row) */
    children?: ReactNode;
    /** Content rendered below the title/description area (e.g. breadcrumb links, chips) */
    afterTitle?: ReactNode;
}
export declare function DrawerHeader({ title, description, onClose, children, afterTitle, }: DrawerHeaderProps): import("react/jsx-runtime").JSX.Element;
interface DrawerBodyProps {
    children: ReactNode;
}
export declare function DrawerBody({ children }: DrawerBodyProps): import("react/jsx-runtime").JSX.Element;
interface DrawerFooterProps {
    children: ReactNode;
    align?: "end" | "between";
}
export declare function DrawerFooter({ children, align }: DrawerFooterProps): import("react/jsx-runtime").JSX.Element;
export {};
