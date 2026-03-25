import { HTMLAttributes } from 'react';
interface PaginationProps extends Omit<HTMLAttributes<HTMLElement>, "onChange"> {
    currentPage: number;
    totalPages: number;
    onPageChange: (page: number) => void;
}
export declare function Pagination({ currentPage, totalPages, onPageChange, className, ...rest }: PaginationProps): import("react/jsx-runtime").JSX.Element;
export {};
