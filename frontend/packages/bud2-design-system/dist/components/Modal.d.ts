import { ReactNode } from 'react';
type ModalSize = "sm" | "md" | "lg";
export interface ModalProps {
    open: boolean;
    onClose: () => void;
    size?: ModalSize;
    children: ReactNode;
    sidePanel?: ReactNode;
    width?: string;
    className?: string;
    "aria-label"?: string;
}
export declare function Modal({ open, onClose, size, children, sidePanel, width, className, "aria-label": ariaLabel, }: ModalProps): import('react').ReactPortal | null;
interface ModalHeaderProps {
    title: ReactNode;
    description?: ReactNode;
    onClose?: () => void;
    children?: ReactNode;
    afterTitle?: ReactNode;
}
export declare function ModalHeader({ title, description, onClose, children, afterTitle, }: ModalHeaderProps): import("react/jsx-runtime").JSX.Element;
interface ModalBodyProps {
    children: ReactNode;
}
export declare function ModalBody({ children }: ModalBodyProps): import("react/jsx-runtime").JSX.Element;
interface ModalFooterProps {
    children: ReactNode;
    align?: "end" | "between";
}
export declare function ModalFooter({ children, align }: ModalFooterProps): import("react/jsx-runtime").JSX.Element;
export {};
