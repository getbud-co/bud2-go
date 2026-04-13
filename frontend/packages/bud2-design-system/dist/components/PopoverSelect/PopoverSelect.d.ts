import { ReactNode, RefObject } from 'react';
import { Icon } from '@phosphor-icons/react';
export interface PopoverSelectOption {
    id: string;
    label: string;
    /** Avatar initials (renders Avatar when provided) */
    initials?: string;
    /** Avatar image src */
    avatarSrc?: string;
    /** Icon component rendered before the label */
    icon?: Icon;
}
interface BaseProps {
    /** Whether the popover is open */
    open: boolean;
    /** Called when the popover should close */
    onClose: () => void;
    /** Anchor element ref for positioning */
    anchorRef: RefObject<HTMLElement | null>;
    /** FilterDropdown placement */
    placement?: "bottom-start" | "right-start";
    /** Skip overlay (useful for nested/sub-panel popovers) */
    noOverlay?: boolean;
    /** Refs to ignore for click-outside (passed to FilterDropdown) */
    ignoreRefs?: RefObject<HTMLElement | null>[];
    /** List of selectable options */
    options: PopoverSelectOption[];
    /** Render custom content after the indicator (Checkbox/Radio) and before the label */
    renderOptionPrefix?: (option: PopoverSelectOption) => ReactNode;
    /** Enable search filtering */
    searchable?: boolean;
    /** Search placeholder text */
    searchPlaceholder?: string;
    /** Content rendered in a footer section (below the list, with top border) */
    footer?: ReactNode;
    /** Text shown when filtered list is empty */
    emptyText?: string;
}
interface SingleSelectProps extends BaseProps {
    /** Single selection mode (Radio indicators) */
    mode: "single";
    value: string | null;
    onChange: (id: string) => void;
    /** Whether to close popover after selecting */
    closeOnSelect?: boolean;
}
interface MultiSelectProps extends BaseProps {
    /** Multiple selection mode (Checkbox indicators) */
    mode: "multiple";
    value: string[];
    onChange: (value: string[]) => void;
    /** Enable inline creation of new options */
    creatable?: boolean;
    /** Placeholder for the create input */
    createPlaceholder?: string;
    /** Called when a new option is created; should return the new option to add */
    onCreateOption?: (label: string) => PopoverSelectOption;
    /** Icon rendered next to each creatable option */
    creatableIcon?: Icon;
}
export type PopoverSelectProps = SingleSelectProps | MultiSelectProps;
/**
 * Formats a multi-select value into a display label.
 * Returns first selected label + count of remaining, or the fallback.
 */
export declare function formatMultiLabel(ids: string[], options: {
    id: string;
    label: string;
}[], fallback: string): string;
export declare function PopoverSelect(props: PopoverSelectProps): import("react/jsx-runtime").JSX.Element;
export default PopoverSelect;
