import { KeyboardEvent as ReactKeyboardEvent, RefObject } from 'react';
export declare const FOCUSABLE_SELECTOR = "button:not([disabled]), [href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex=\"-1\"])";
type FocusScope = ParentNode | null;
type AnyKeyboardEvent = globalThis.KeyboardEvent | ReactKeyboardEvent<HTMLElement>;
export declare function getFocusableElements(container: FocusScope): HTMLElement[];
export declare function trapFocusWithin(container: HTMLElement | null, event: AnyKeyboardEvent): void;
export declare function useHasOpened(open: boolean): boolean;
export declare function useDocumentEscape(active: boolean, onEscape: () => void): void;
export declare function useBodyScrollLock(active: boolean): void;
interface OpenFocusOptions<TContainer extends HTMLElement = HTMLElement, TInitial extends HTMLElement = HTMLElement> {
    active: boolean;
    containerRef: RefObject<TContainer | null>;
    initialFocusRef?: RefObject<TInitial | null>;
    restoreFocus?: boolean;
}
interface ClickOutsideOptions {
    active: boolean;
    refs: Array<RefObject<HTMLElement | null>>;
    onOutside: () => void;
    shouldIgnoreTarget?: (target: Node) => boolean;
    relatedAnchorRef?: RefObject<HTMLElement | null>;
    relatedPortalSelectors?: string[];
}
interface RelatedOverlayIgnoreOptions {
    anchorRef?: RefObject<HTMLElement | null>;
    portalSelectors?: string[];
}
interface ClampToViewportOptions {
    value: number;
    size: number;
    viewportSize: number;
    margin: number;
}
interface VerticalPositionOptions {
    anchorTop: number;
    anchorBottom: number;
    overlayHeight: number;
    viewportHeight: number;
    gap: number;
    margin: number;
    preferred?: "bottom" | "top";
}
interface SidePositionOptions {
    anchorLeft: number;
    anchorRight: number;
    overlayWidth: number;
    viewportWidth: number;
    gap: number;
    margin: number;
    preferred?: "right" | "left";
}
interface AnchoredOverlayPositionOptions {
    anchorTop: number;
    anchorBottom: number;
    anchorLeft: number;
    anchorRight: number;
    overlayWidth: number;
    overlayHeight: number;
    viewportWidth: number;
    viewportHeight: number;
    gap: number;
    margin: number;
    horizontalAlign?: "start" | "end";
    preferredVertical?: "bottom" | "top";
}
interface SideStartOverlayPositionOptions {
    anchorTop: number;
    anchorLeft: number;
    anchorRight: number;
    overlayWidth: number;
    overlayHeight: number;
    viewportWidth: number;
    viewportHeight: number;
    gap: number;
    margin: number;
    preferredSide?: "right" | "left";
}
export declare function useOpenFocus<TContainer extends HTMLElement = HTMLElement, TInitial extends HTMLElement = HTMLElement>({ active, containerRef, initialFocusRef, restoreFocus, }: OpenFocusOptions<TContainer, TInitial>): () => void;
export declare function useDocumentClickOutside({ active, refs, onOutside, shouldIgnoreTarget, relatedAnchorRef, relatedPortalSelectors, }: ClickOutsideOptions): void;
export declare function shouldIgnoreTargetFromRelatedOverlays(target: Node, { anchorRef, portalSelectors }: RelatedOverlayIgnoreOptions): boolean;
export declare function useInitialReposition(active: boolean, reposition: () => void): void;
export declare function useViewportReposition(active: boolean, reposition: () => void): void;
export declare function clampToViewport({ value, size, viewportSize, margin, }: ClampToViewportOptions): number;
export declare function resolveVerticalPosition({ anchorTop, anchorBottom, overlayHeight, viewportHeight, gap, margin, preferred, }: VerticalPositionOptions): {
    top: number;
    placement: "bottom" | "top";
};
export declare function resolveSidePosition({ anchorLeft, anchorRight, overlayWidth, viewportWidth, gap, margin, preferred, }: SidePositionOptions): {
    left: number;
    side: "right" | "left";
};
export declare function resolveAnchoredOverlayPosition({ anchorTop, anchorBottom, anchorLeft, anchorRight, overlayWidth, overlayHeight, viewportWidth, viewportHeight, gap, margin, horizontalAlign, preferredVertical, }: AnchoredOverlayPositionOptions): {
    top: number;
    left: number;
    verticalPlacement: "bottom" | "top";
};
export declare function resolveSideStartOverlayPosition({ anchorTop, anchorLeft, anchorRight, overlayWidth, overlayHeight, viewportWidth, viewportHeight, gap, margin, preferredSide, }: SideStartOverlayPositionOptions): {
    top: number;
    left: number;
    side: "right" | "left";
};
export {};
