export interface SortableItem {
    /** Unique identifier */
    id: string;
    /** Display label */
    label: string;
}
export interface SortableListProps {
    /** Ordered list of items */
    items: SortableItem[];
    /** Called with the reordered items after a drop */
    onChange?: (items: SortableItem[]) => void;
    /** Disables drag-and-drop */
    disabled?: boolean;
    /** Visual size */
    size?: "sm" | "md";
    className?: string;
}
export declare function SortableList({ items, onChange, disabled, size, className, }: SortableListProps): import("react/jsx-runtime").JSX.Element;
