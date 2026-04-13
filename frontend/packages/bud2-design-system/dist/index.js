import { jsx, jsxs, Fragment } from "react/jsx-runtime";
import { useState, useId, useCallback, forwardRef, useEffect, useRef, useLayoutEffect, useContext, createContext, useMemo, useSyncExternalStore } from "react";
import { CaretDown, CircleNotch, Minus, Check, CaretRight, Lightning, SidebarSimple, File, Link, ArrowUp, X, Plus, FileArrowUp, MagnifyingGlass, WarningCircle, Warning, CheckCircle, Info, User, CalendarBlank, CaretLeft, Broom, FloppyDisk, BellSlash, Bell, DotsThreeVertical, DotsSixVertical, ArrowDown, CaretLineRight, CaretLineLeft, DotsThree } from "@phosphor-icons/react";
import { createPortal } from "react-dom";
const root$8 = "_root_bgflm_3";
const item$9 = "_item_bgflm_13";
const rootHeader = "_rootHeader_bgflm_23";
const trigger$3 = "_trigger_bgflm_23";
const icon$2 = "_icon_bgflm_62";
const itemDisabled = "_itemDisabled_bgflm_69";
const triggerContent = "_triggerContent_bgflm_75";
const title$b = "_title_bgflm_83";
const description$9 = "_description_bgflm_91";
const action$1 = "_action_bgflm_100";
const caret$3 = "_caret_bgflm_108";
const itemOpen = "_itemOpen_bgflm_116";
const panel$2 = "_panel_bgflm_122";
const panelInner = "_panelInner_bgflm_132";
const s$F = {
  root: root$8,
  item: item$9,
  rootHeader,
  trigger: trigger$3,
  icon: icon$2,
  itemDisabled,
  triggerContent,
  title: title$b,
  description: description$9,
  action: action$1,
  caret: caret$3,
  itemOpen,
  panel: panel$2,
  panelInner
};
function AccordionItem({
  icon: Icon,
  title: title2,
  description: description2,
  action: action2,
  defaultOpen = false,
  open: openProp,
  onToggle,
  disabled: disabled2 = false,
  children,
  className,
  ...rest
}) {
  const [internalOpen, setInternalOpen] = useState(defaultOpen);
  const isControlled = openProp !== void 0;
  const isOpen = isControlled ? openProp : internalOpen;
  const uid2 = useId();
  const triggerId = `acc-trigger-${uid2}`;
  const panelId = `acc-panel-${uid2}`;
  const toggle = useCallback(() => {
    if (disabled2) return;
    const next2 = !isOpen;
    if (!isControlled) setInternalOpen(next2);
    onToggle?.(next2);
  }, [disabled2, isOpen, isControlled, onToggle]);
  const classes = [
    s$F.item,
    isOpen ? s$F.itemOpen : "",
    disabled2 ? s$F.itemDisabled : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("div", { className: classes, ...rest, children: [
    /* @__PURE__ */ jsxs(
      "button",
      {
        type: "button",
        id: triggerId,
        className: s$F.trigger,
        "aria-expanded": isOpen,
        "aria-controls": panelId,
        disabled: disabled2,
        onClick: toggle,
        children: [
          Icon && /* @__PURE__ */ jsx(Icon, { size: 16, className: s$F.icon }),
          /* @__PURE__ */ jsxs("div", { className: s$F.triggerContent, children: [
            /* @__PURE__ */ jsx("span", { className: s$F.title, children: title2 }),
            description2 && /* @__PURE__ */ jsx("span", { className: s$F.description, children: description2 })
          ] }),
          action2 && /* @__PURE__ */ jsx("div", { className: s$F.action, children: action2 }),
          /* @__PURE__ */ jsx(CaretDown, { size: 16, className: s$F.caret })
        ]
      }
    ),
    /* @__PURE__ */ jsx(
      "div",
      {
        id: panelId,
        role: "region",
        "aria-labelledby": triggerId,
        className: s$F.panel,
        inert: !isOpen || void 0,
        children: /* @__PURE__ */ jsx("div", { className: s$F.panelInner, children })
      }
    )
  ] });
}
function Accordion({ header: header2 = false, children, className, ...rest }) {
  const classes = [s$F.root, header2 ? s$F.rootHeader : "", className ?? ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx("div", { className: classes, ...rest, children });
}
const button = "_button_1g5v3_2";
const sm$e = "_sm_1g5v3_23";
const md$c = "_md_1g5v3_30";
const lg$9 = "_lg_1g5v3_37";
const iconOnly$1 = "_iconOnly_1g5v3_45";
const hasLeftIcon$1 = "_hasLeftIcon_1g5v3_61";
const hasRightIcon = "_hasRightIcon_1g5v3_62";
const primary$1 = "_primary_1g5v3_71";
const hovered$6 = "_hovered_1g5v3_78";
const active$1 = "_active_1g5v3_79";
const secondary$1 = "_secondary_1g5v3_90";
const tertiary$1 = "_tertiary_1g5v3_109";
const danger = "_danger_1g5v3_126";
const loading = "_loading_1g5v3_163";
const spinner = "_spinner_1g5v3_167";
const spin = "_spin_1g5v3_167";
const srOnly = "_srOnly_1g5v3_177";
const btnStyles = {
  button,
  sm: sm$e,
  md: md$c,
  lg: lg$9,
  iconOnly: iconOnly$1,
  hasLeftIcon: hasLeftIcon$1,
  hasRightIcon,
  primary: primary$1,
  hovered: hovered$6,
  active: active$1,
  secondary: secondary$1,
  tertiary: tertiary$1,
  danger,
  loading,
  spinner,
  spin,
  srOnly
};
const iconSize$4 = {
  sm: 14,
  md: 16,
  lg: 20
};
const Button = forwardRef(function Button2({
  variant = "primary",
  size = "md",
  leftIcon: LeftIcon,
  rightIcon: RightIcon,
  loading: loading2 = false,
  disabled: disabled2 = false,
  children,
  className,
  ...rest
}, ref) {
  const hasLeft = !loading2 && !!LeftIcon;
  const hasRight = !loading2 && !!RightIcon;
  const isIconOnly = !children && (hasLeft || hasRight);
  const classes = [
    btnStyles.button,
    btnStyles[variant],
    btnStyles[size],
    loading2 ? btnStyles.loading : "",
    isIconOnly ? btnStyles.iconOnly : "",
    !isIconOnly && hasLeft ? btnStyles.hasLeftIcon : "",
    !isIconOnly && hasRight ? btnStyles.hasRightIcon : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  const iSize = iconSize$4[size];
  return /* @__PURE__ */ jsx(
    "button",
    {
      ref,
      type: "button",
      className: classes,
      disabled: disabled2 || loading2,
      "aria-busy": loading2 || void 0,
      ...rest,
      children: loading2 ? /* @__PURE__ */ jsxs(Fragment, { children: [
        /* @__PURE__ */ jsx(CircleNotch, { size: iSize, className: btnStyles.spinner }),
        /* @__PURE__ */ jsx("span", { className: btnStyles.srOnly, children: "Carregando..." })
      ] }) : /* @__PURE__ */ jsxs(Fragment, { children: [
        LeftIcon && /* @__PURE__ */ jsx(LeftIcon, { size: iSize }),
        children && /* @__PURE__ */ jsx("span", { children }),
        RightIcon && /* @__PURE__ */ jsx(RightIcon, { size: iSize })
      ] })
    }
  );
});
const wrapper$9 = "_wrapper_1v4ah_2";
const hasLabel$2 = "_hasLabel_1v4ah_10";
const disabled$a = "_disabled_1v4ah_14";
const input$6 = "_input_1v4ah_19";
const box$1 = "_box_1v4ah_31";
const icon$1 = "_icon_1v4ah_42";
const sm$d = "_sm_1v4ah_49";
const md$b = "_md_1v4ah_53";
const hovered$5 = "_hovered_1v4ah_123";
const focused$6 = "_focused_1v4ah_133";
const text$5 = "_text_1v4ah_146";
const text_sm$1 = "_text_sm_1v4ah_155";
const title$a = "_title_1v4ah_159";
const text_md$1 = "_text_md_1v4ah_165";
const description$8 = "_description_1v4ah_179";
const s$E = {
  wrapper: wrapper$9,
  hasLabel: hasLabel$2,
  disabled: disabled$a,
  input: input$6,
  box: box$1,
  icon: icon$1,
  sm: sm$d,
  md: md$b,
  hovered: hovered$5,
  focused: focused$6,
  text: text$5,
  text_sm: text_sm$1,
  title: title$a,
  text_md: text_md$1,
  description: description$8
};
const boxSize$1 = { sm: 16, md: 20 };
const iconSize$3 = { sm: 12, md: 16 };
const Checkbox = forwardRef(
  ({
    size = "md",
    indeterminate = false,
    label: label2,
    description: description2,
    disabled: disabled2 = false,
    className,
    ...rest
  }, ref) => {
    const wrapperClasses = [
      s$E.wrapper,
      label2 ? s$E.hasLabel : "",
      disabled2 ? s$E.disabled : "",
      className ?? ""
    ].filter(Boolean).join(" ");
    const iSize = iconSize$3[size];
    return /* @__PURE__ */ jsxs("label", { className: wrapperClasses, children: [
      /* @__PURE__ */ jsx(
        "input",
        {
          ref: (node) => {
            if (node) node.indeterminate = indeterminate;
            if (typeof ref === "function") ref(node);
            else if (ref) ref.current = node;
          },
          type: "checkbox",
          className: s$E.input,
          disabled: disabled2,
          ...rest
        }
      ),
      /* @__PURE__ */ jsx(
        "span",
        {
          className: `${s$E.box} ${s$E[size]}`,
          style: {
            width: boxSize$1[size],
            height: boxSize$1[size]
          },
          children: indeterminate ? /* @__PURE__ */ jsx(Minus, { size: iSize, className: s$E.icon }) : /* @__PURE__ */ jsx(Check, { size: iSize, className: s$E.icon })
        }
      ),
      label2 && /* @__PURE__ */ jsxs("span", { className: `${s$E.text} ${s$E[`text_${size}`]}`, children: [
        /* @__PURE__ */ jsx("span", { className: s$E.title, children: label2 }),
        description2 && /* @__PURE__ */ jsx("span", { className: s$E.description, children: description2 })
      ] })
    ] });
  }
);
Checkbox.displayName = "Checkbox";
const FOCUSABLE_SELECTOR = 'button:not([disabled]), [href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex="-1"])';
function getFocusableElements(container2) {
  if (!container2) return [];
  return Array.from(container2.querySelectorAll(FOCUSABLE_SELECTOR));
}
function trapFocusWithin(container2, event) {
  if (event.key !== "Tab") return;
  const focusable = getFocusableElements(container2);
  if (focusable.length === 0) return;
  const first = focusable[0];
  const last = focusable[focusable.length - 1];
  if (event.shiftKey && document.activeElement === first) {
    event.preventDefault();
    last.focus();
  } else if (!event.shiftKey && document.activeElement === last) {
    event.preventDefault();
    first.focus();
  }
}
function useHasOpened(open2) {
  const [hasOpened, setHasOpened] = useState(open2);
  useEffect(() => {
    if (open2) setHasOpened(true);
  }, [open2]);
  return hasOpened;
}
function useDocumentEscape(active2, onEscape) {
  useEffect(() => {
    if (!active2) return;
    function handleKeyDown(event) {
      if (event.key === "Escape") onEscape();
    }
    document.addEventListener("keydown", handleKeyDown);
    return () => document.removeEventListener("keydown", handleKeyDown);
  }, [active2, onEscape]);
}
function useBodyScrollLock(active2) {
  useEffect(() => {
    if (!active2) return;
    const previousOverflow = document.body.style.overflow;
    document.body.style.overflow = "hidden";
    return () => {
      document.body.style.overflow = previousOverflow;
    };
  }, [active2]);
}
function useOpenFocus({
  active: active2,
  containerRef,
  initialFocusRef,
  restoreFocus = true
}) {
  const previousFocusRef = useRef(null);
  const restorePreviousFocus = useCallback(() => {
    if (!restoreFocus) return;
    previousFocusRef.current?.focus();
    previousFocusRef.current = null;
  }, [restoreFocus]);
  useEffect(() => {
    if (!active2) return;
    previousFocusRef.current = document.activeElement;
    const raf = requestAnimationFrame(() => {
      const explicitTarget = initialFocusRef?.current;
      if (explicitTarget) {
        explicitTarget.focus();
        return;
      }
      const firstFocusable = getFocusableElements(containerRef.current)[0];
      firstFocusable?.focus();
    });
    return () => {
      cancelAnimationFrame(raf);
      restorePreviousFocus();
    };
  }, [active2, containerRef, initialFocusRef, restorePreviousFocus]);
  return restorePreviousFocus;
}
function useDocumentClickOutside({
  active: active2,
  refs,
  onOutside,
  shouldIgnoreTarget,
  relatedAnchorRef,
  relatedPortalSelectors
}) {
  useEffect(() => {
    if (!active2) return;
    function handleMouseDown(event) {
      const target = event.target;
      if (shouldIgnoreTarget?.(target)) return;
      if (shouldIgnoreTargetFromRelatedOverlays(target, {
        anchorRef: relatedAnchorRef,
        portalSelectors: relatedPortalSelectors
      })) {
        return;
      }
      const isInsideAnyRef = refs.some((ref) => ref.current?.contains(target));
      if (!isInsideAnyRef) {
        onOutside();
      }
    }
    document.addEventListener("mousedown", handleMouseDown);
    return () => document.removeEventListener("mousedown", handleMouseDown);
  }, [
    active2,
    refs,
    onOutside,
    shouldIgnoreTarget,
    relatedAnchorRef,
    relatedPortalSelectors
  ]);
}
function shouldIgnoreTargetFromRelatedOverlays(target, { anchorRef, portalSelectors = [] }) {
  if (!(target instanceof Element)) return false;
  if (portalSelectors.some((selector) => target.closest(selector))) {
    return true;
  }
  const dialog = target.closest('[role="dialog"]');
  if (dialog && anchorRef?.current && !dialog.contains(anchorRef.current)) {
    return true;
  }
  return false;
}
function useInitialReposition(active2, reposition) {
  useLayoutEffect(() => {
    if (!active2) return;
    reposition();
  }, [active2, reposition]);
}
function useViewportReposition(active2, reposition) {
  useEffect(() => {
    if (!active2) return;
    window.addEventListener("scroll", reposition, true);
    window.addEventListener("resize", reposition);
    return () => {
      window.removeEventListener("scroll", reposition, true);
      window.removeEventListener("resize", reposition);
    };
  }, [active2, reposition]);
}
function clampToViewport({
  value: value2,
  size,
  viewportSize,
  margin
}) {
  return Math.max(margin, Math.min(value2, viewportSize - size - margin));
}
function resolveVerticalPosition({
  anchorTop,
  anchorBottom,
  overlayHeight,
  viewportHeight,
  gap,
  margin,
  preferred = "bottom"
}) {
  const bottomTop = anchorBottom + gap;
  const topTop = anchorTop - overlayHeight - gap;
  const fitsBottom = bottomTop + overlayHeight <= viewportHeight - margin;
  const fitsTop = topTop >= margin;
  if (preferred === "top") {
    if (fitsTop) return { top: topTop, placement: "top" };
    if (fitsBottom) return { top: bottomTop, placement: "bottom" };
  } else {
    if (fitsBottom) return { top: bottomTop, placement: "bottom" };
    if (fitsTop) return { top: topTop, placement: "top" };
  }
  const clampedTop = clampToViewport({
    value: preferred === "top" ? topTop : bottomTop,
    size: overlayHeight,
    viewportSize: viewportHeight,
    margin
  });
  return { top: clampedTop, placement: preferred };
}
function resolveSidePosition({
  anchorLeft,
  anchorRight,
  overlayWidth,
  viewportWidth,
  gap,
  margin,
  preferred = "right"
}) {
  const rightLeft = anchorRight + gap;
  const leftLeft = anchorLeft - overlayWidth - gap;
  const fitsRight = rightLeft + overlayWidth <= viewportWidth - margin;
  const fitsLeft = leftLeft >= margin;
  if (preferred === "left") {
    if (fitsLeft) return { left: leftLeft, side: "left" };
    if (fitsRight) return { left: rightLeft, side: "right" };
  } else {
    if (fitsRight) return { left: rightLeft, side: "right" };
    if (fitsLeft) return { left: leftLeft, side: "left" };
  }
  const clampedLeft = clampToViewport({
    value: preferred === "left" ? leftLeft : rightLeft,
    size: overlayWidth,
    viewportSize: viewportWidth,
    margin
  });
  return { left: clampedLeft, side: preferred };
}
function resolveAnchoredOverlayPosition({
  anchorTop,
  anchorBottom,
  anchorLeft,
  anchorRight,
  overlayWidth,
  overlayHeight,
  viewportWidth,
  viewportHeight,
  gap,
  margin,
  horizontalAlign = "start",
  preferredVertical = "bottom"
}) {
  const { top: top2, placement } = resolveVerticalPosition({
    anchorTop,
    anchorBottom,
    overlayHeight,
    viewportHeight,
    gap,
    margin,
    preferred: preferredVertical
  });
  const rawLeft = horizontalAlign === "end" ? anchorRight - overlayWidth : anchorLeft;
  const left2 = clampToViewport({
    value: rawLeft,
    size: overlayWidth,
    viewportSize: viewportWidth,
    margin
  });
  return { top: top2, left: left2, verticalPlacement: placement };
}
function resolveSideStartOverlayPosition({
  anchorTop,
  anchorLeft,
  anchorRight,
  overlayWidth,
  overlayHeight,
  viewportWidth,
  viewportHeight,
  gap,
  margin,
  preferredSide = "right"
}) {
  const { left: left2, side } = resolveSidePosition({
    anchorLeft,
    anchorRight,
    overlayWidth,
    viewportWidth,
    gap,
    margin,
    preferred: preferredSide
  });
  const top2 = clampToViewport({
    value: anchorTop,
    size: overlayHeight,
    viewportSize: viewportHeight,
    margin
  });
  return { top: top2, left: left2, side };
}
const popover$1 = "_popover_1yz8f_2";
const divider$2 = "_divider_1yz8f_25";
const item$8 = "_item_1yz8f_32";
const itemDanger = "_itemDanger_1yz8f_55";
const itemIcon$2 = "_itemIcon_1yz8f_59";
const itemImage = "_itemImage_1yz8f_72";
const itemLabel$2 = "_itemLabel_1yz8f_81";
const itemBadge = "_itemBadge_1yz8f_85";
const itemCaret$1 = "_itemCaret_1yz8f_102";
const submenuWrapper = "_submenuWrapper_1yz8f_109";
const submenu = "_submenu_1yz8f_109";
const submenuWrapperOpen = "_submenuWrapperOpen_1yz8f_117";
const submenuFlip = "_submenuFlip_1yz8f_135";
const s$D = {
  popover: popover$1,
  divider: divider$2,
  item: item$8,
  itemDanger,
  itemIcon: itemIcon$2,
  itemImage,
  itemLabel: itemLabel$2,
  itemBadge,
  itemCaret: itemCaret$1,
  submenuWrapper,
  submenu,
  submenuWrapperOpen,
  submenuFlip
};
function adjustSubmenuOverflow(wrapperEl) {
  const submenu2 = wrapperEl.querySelector('[class*="submenu"]');
  if (!submenu2) return;
  submenu2.style.top = "";
  submenu2.style.left = "";
  submenu2.style.right = "";
  submenu2.style.maxHeight = "";
  void submenu2.offsetHeight;
  const rect = submenu2.getBoundingClientRect();
  if (rect.width === 0 && rect.height === 0) return;
  const gap = 8;
  let topAdjust = 0;
  if (rect.bottom > window.innerHeight - gap) {
    topAdjust = -(rect.bottom - (window.innerHeight - gap));
  }
  if (rect.top + topAdjust < gap) {
    topAdjust = gap - rect.top;
  }
  if (topAdjust !== 0) {
    submenu2.style.top = `${topAdjust}px`;
  }
  const availableHeight = window.innerHeight - gap * 2;
  if (rect.height > availableHeight) {
    submenu2.style.maxHeight = `${availableHeight}px`;
    submenu2.style.overflowY = "auto";
    submenu2.style.top = `${gap - rect.top + topAdjust}px`;
  }
  const updatedRect = submenu2.getBoundingClientRect();
  if (updatedRect.left < gap) {
    const shift = gap - updatedRect.left;
    submenu2.style.right = `calc(100% + var(--sp-3xs) - ${shift}px)`;
  }
  if (updatedRect.right > window.innerWidth - gap) {
    const shift = updatedRect.right - (window.innerWidth - gap);
    submenu2.style.left = `calc(100% + var(--sp-3xs) - ${shift}px)`;
  }
}
function Popover({ items: items2, open: open2, onClose, anchorRef, ariaLabel }) {
  const popoverRef = useRef(null);
  const [submenuFlip2, setSubmenuFlip] = useState(false);
  const [focusedIndex, setFocusedIndex] = useState(-1);
  const [openSubmenuId, setOpenSubmenuId] = useState(null);
  const previousFocusRef = useRef(null);
  const getMenuItems = useCallback(() => {
    if (!popoverRef.current) return [];
    return Array.from(popoverRef.current.querySelectorAll('[role="menuitem"]'));
  }, []);
  const applyPosition = useCallback(() => {
    const anchor2 = anchorRef.current;
    const el = popoverRef.current;
    if (!anchor2 || !el) return;
    const ar = anchor2.getBoundingClientRect();
    const gap = 4;
    el.style.position = "fixed";
    el.style.top = "auto";
    el.style.right = "auto";
    el.style.bottom = `${window.innerHeight - ar.top + gap}px`;
    el.style.left = `${ar.left}px`;
    const pr = el.getBoundingClientRect();
    if (pr.top < gap) {
      el.style.bottom = "auto";
      el.style.top = `${ar.bottom + gap}px`;
    }
    if (pr.right > window.innerWidth - gap) {
      el.style.left = `${Math.max(gap, window.innerWidth - pr.width - gap)}px`;
    }
    const updatedPr = el.getBoundingClientRect();
    setSubmenuFlip(updatedPr.right + 220 > window.innerWidth);
  }, [anchorRef]);
  useInitialReposition(open2, applyPosition);
  const getSubmenuFocusable = useCallback((wrapperEl) => {
    const submenu2 = wrapperEl.querySelector(`.${s$D.submenu}`);
    if (!submenu2) return [];
    return Array.from(
      submenu2.querySelectorAll(
        '[role="menuitem"], button:not([disabled]), a[href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex="-1"])'
      )
    );
  }, []);
  useEffect(() => {
    if (!open2) return;
    previousFocusRef.current = document.activeElement;
    const raf = requestAnimationFrame(() => {
      const menuItems = getMenuItems();
      if (menuItems.length > 0) {
        setOpenSubmenuId(null);
        setFocusedIndex(0);
        menuItems[0].focus();
      }
    });
    return () => cancelAnimationFrame(raf);
  }, [open2, getMenuItems]);
  useEffect(() => {
    const menuItems = getMenuItems();
    menuItems.forEach((item2, i) => {
      item2.setAttribute("tabindex", i === focusedIndex ? "0" : "-1");
    });
    if (focusedIndex >= 0 && menuItems[focusedIndex]) {
      menuItems[focusedIndex].focus();
    }
  }, [focusedIndex, getMenuItems]);
  const restoreFocus = useCallback(() => {
    const el = previousFocusRef.current;
    if (el && el instanceof HTMLElement) {
      el.focus();
    }
    previousFocusRef.current = null;
  }, []);
  const handleClose = useCallback(() => {
    onClose();
    restoreFocus();
  }, [onClose, restoreFocus]);
  useViewportReposition(open2, applyPosition);
  useDocumentClickOutside({
    active: open2,
    refs: [anchorRef, popoverRef],
    onOutside: handleClose
  });
  useDocumentEscape(open2, handleClose);
  const handlePopoverKeyDown = useCallback(
    (e) => {
      const activeWrapper = document.activeElement instanceof HTMLElement ? document.activeElement.closest(`.${s$D.submenuWrapper}`) : null;
      if (e.key === "ArrowRight" && activeWrapper instanceof HTMLElement) {
        const submenuItems = getSubmenuFocusable(activeWrapper);
        if (submenuItems.length > 0) {
          e.preventDefault();
          const triggerId = activeWrapper.dataset.submenuTrigger;
          if (triggerId) setOpenSubmenuId(triggerId);
          requestAnimationFrame(() => {
            adjustSubmenuOverflow(activeWrapper);
            submenuItems[0]?.focus();
          });
        }
        return;
      }
      if (e.key === "ArrowLeft") {
        const activeSubmenu = document.activeElement instanceof HTMLElement ? document.activeElement.closest(`.${s$D.submenu}`) : null;
        if (activeSubmenu instanceof HTMLElement) {
          e.preventDefault();
          const wrapper2 = activeSubmenu.closest(`.${s$D.submenuWrapper}`);
          if (wrapper2 instanceof HTMLElement) {
            const triggerId = wrapper2.dataset.submenuTrigger;
            if (triggerId) {
              setOpenSubmenuId(null);
              const trigger2 = wrapper2.querySelector(`[data-submenu-trigger="${triggerId}"]`);
              trigger2?.focus();
            }
          }
        }
        return;
      }
      if (e.key === "Escape") {
        const activeSubmenu = document.activeElement instanceof HTMLElement ? document.activeElement.closest(`.${s$D.submenu}`) : null;
        if (activeSubmenu instanceof HTMLElement) {
          e.preventDefault();
          const wrapper2 = activeSubmenu.closest(`.${s$D.submenuWrapper}`);
          if (wrapper2 instanceof HTMLElement) {
            const triggerId = wrapper2.dataset.submenuTrigger;
            setOpenSubmenuId(null);
            if (triggerId) {
              const trigger2 = wrapper2.querySelector(`[data-submenu-trigger="${triggerId}"]`);
              trigger2?.focus();
            }
          }
          return;
        }
      }
      const menuItems = getMenuItems();
      const count = menuItems.length;
      if (count === 0) return;
      let nextIndex = focusedIndex;
      switch (e.key) {
        case "ArrowDown":
          e.preventDefault();
          nextIndex = focusedIndex < count - 1 ? focusedIndex + 1 : 0;
          break;
        case "ArrowUp":
          e.preventDefault();
          nextIndex = focusedIndex > 0 ? focusedIndex - 1 : count - 1;
          break;
        case "Home":
          e.preventDefault();
          nextIndex = 0;
          break;
        case "End":
          e.preventDefault();
          nextIndex = count - 1;
          break;
        default:
          return;
      }
      setFocusedIndex(nextIndex);
    },
    [focusedIndex, getMenuItems, getSubmenuFocusable]
  );
  if (!open2) return null;
  return createPortal(
    /* @__PURE__ */ jsx(
      "div",
      {
        ref: popoverRef,
        className: s$D.popover,
        role: "menu",
        "aria-label": ariaLabel ?? "Popover",
        onKeyDown: handlePopoverKeyDown,
        children: items2.map((item2, index) => {
          const dividerEl = item2.divider ? /* @__PURE__ */ jsx("hr", { className: s$D.divider }) : null;
          if (item2.submenu) {
            return /* @__PURE__ */ jsxs("div", { children: [
              dividerEl,
              /* @__PURE__ */ jsxs(
                "div",
                {
                  className: `${s$D.submenuWrapper}${submenuFlip2 ? ` ${s$D.submenuFlip}` : ""}${openSubmenuId === item2.id ? ` ${s$D.submenuWrapperOpen}` : ""}`,
                  "data-submenu-trigger": item2.id,
                  onBlur: (e) => {
                    const currentTarget = e.currentTarget;
                    requestAnimationFrame(() => {
                      if (!currentTarget.contains(document.activeElement)) {
                        setOpenSubmenuId((current) => current === item2.id ? null : current);
                      }
                    });
                  },
                  children: [
                    /* @__PURE__ */ jsxs(
                      "button",
                      {
                        type: "button",
                        className: s$D.item,
                        role: "menuitem",
                        tabIndex: index === focusedIndex ? 0 : -1,
                        "aria-haspopup": "menu",
                        "aria-expanded": openSubmenuId === item2.id,
                        "aria-controls": `${item2.id}-submenu`,
                        "data-submenu-trigger": item2.id,
                        onClick: (e) => {
                          e.preventDefault();
                          const wrapper2 = e.currentTarget.closest(`.${s$D.submenuWrapper}`);
                          setOpenSubmenuId((current) => {
                            const nextOpen = current === item2.id ? null : item2.id;
                            if (nextOpen === item2.id && wrapper2 instanceof HTMLElement) {
                              requestAnimationFrame(() => {
                                adjustSubmenuOverflow(wrapper2);
                                const submenuItems = getSubmenuFocusable(wrapper2);
                                submenuItems[0]?.focus();
                              });
                            }
                            return nextOpen;
                          });
                        },
                        children: [
                          item2.image ? /* @__PURE__ */ jsx("img", { src: item2.image, alt: "", className: s$D.itemImage }) : item2.icon ? /* @__PURE__ */ jsx(item2.icon, { size: 16, className: s$D.itemIcon }) : null,
                          /* @__PURE__ */ jsx("span", { className: s$D.itemLabel, children: item2.label }),
                          item2.badge != null && item2.badge > 0 && /* @__PURE__ */ jsx("span", { className: s$D.itemBadge, children: item2.badge }),
                          /* @__PURE__ */ jsx(CaretRight, { size: 12, className: s$D.itemCaret })
                        ]
                      }
                    ),
                    /* @__PURE__ */ jsx(
                      "div",
                      {
                        id: `${item2.id}-submenu`,
                        className: s$D.submenu,
                        role: "menu",
                        "aria-label": `${item2.label} submenu`,
                        children: item2.submenu
                      }
                    )
                  ]
                }
              )
            ] }, item2.id);
          }
          return /* @__PURE__ */ jsxs("div", { children: [
            dividerEl,
            /* @__PURE__ */ jsxs(
              "button",
              {
                type: "button",
                className: `${s$D.item}${item2.danger ? ` ${s$D.itemDanger}` : ""}`,
                role: "menuitem",
                tabIndex: index === focusedIndex ? 0 : -1,
                onClick: () => {
                  item2.onClick?.();
                  handleClose();
                },
                children: [
                  item2.image ? /* @__PURE__ */ jsx("img", { src: item2.image, alt: "", className: s$D.itemImage }) : item2.icon ? /* @__PURE__ */ jsx(item2.icon, { size: 16, className: s$D.itemIcon }) : null,
                  /* @__PURE__ */ jsx("span", { children: item2.label })
                ]
              }
            )
          ] }, item2.id);
        })
      }
    ),
    document.body
  );
}
const container$2 = "_container_k1loz_2";
const header$7 = "_header_k1loz_13";
const headerIcon$1 = "_headerIcon_k1loz_20";
const title$9 = "_title_k1loz_25";
const body$5 = "_body_k1loz_37";
const symbol = "_symbol_k1loz_47";
const budSymbol = "_budSymbol_k1loz_51";
const heading = "_heading_k1loz_57";
const suggestions = "_suggestions_k1loz_69";
const suggestion = "_suggestion_k1loz_69";
const messageList = "_messageList_k1loz_108";
const userMessage = "_userMessage_k1loz_118";
const userContext = "_userContext_k1loz_137";
const userContextBadge = "_userContextBadge_k1loz_144";
const userText$1 = "_userText_k1loz_159";
const aiMessage = "_aiMessage_k1loz_174";
const aiIcon = "_aiIcon_k1loz_181";
const aiText = "_aiText_k1loz_187";
const aiContent = "_aiContent_k1loz_197";
const useAsBaseBtn = "_useAsBaseBtn_k1loz_204";
const typingDots = "_typingDots_k1loz_226";
const dot$3 = "_dot_k1loz_233";
const cursor = "_cursor_k1loz_259";
const inputArea = "_inputArea_k1loz_279";
const inputGlow = "_inputGlow_k1loz_284";
const inputBox$1 = "_inputBox_k1loz_294";
const inputRow$1 = "_inputRow_k1loz_306";
const input$5 = "_input_k1loz_279";
const attachButton = "_attachButton_k1loz_342";
const fileInput = "_fileInput_k1loz_373";
const contextStrip = "_contextStrip_k1loz_386";
const contextBadge = "_contextBadge_k1loz_393";
const contextBadgeIcon = "_contextBadgeIcon_k1loz_415";
const contextBadgeLabel = "_contextBadgeLabel_k1loz_420";
const contextBadgeRemove = "_contextBadgeRemove_k1loz_432";
const missionSearchBox = "_missionSearchBox_k1loz_454";
const missionSearchIcon = "_missionSearchIcon_k1loz_462";
const missionSearchInput = "_missionSearchInput_k1loz_467";
const missionList = "_missionList_k1loz_486";
const missionListItem = "_missionListItem_k1loz_491";
const missionLabel = "_missionLabel_k1loz_504";
const submitButton = "_submitButton_k1loz_517";
const s$C = {
  container: container$2,
  header: header$7,
  headerIcon: headerIcon$1,
  title: title$9,
  body: body$5,
  symbol,
  budSymbol,
  heading,
  suggestions,
  suggestion,
  messageList,
  userMessage,
  userContext,
  userContextBadge,
  userText: userText$1,
  aiMessage,
  aiIcon,
  aiText,
  aiContent,
  useAsBaseBtn,
  typingDots,
  dot: dot$3,
  cursor,
  inputArea,
  inputGlow,
  inputBox: inputBox$1,
  inputRow: inputRow$1,
  input: input$5,
  attachButton,
  fileInput,
  contextStrip,
  contextBadge,
  contextBadgeIcon,
  contextBadgeLabel,
  contextBadgeRemove,
  missionSearchBox,
  missionSearchIcon,
  missionSearchInput,
  missionList,
  missionListItem,
  missionLabel,
  submitButton
};
let nextId = 0;
function uid() {
  return `msg-${++nextId}-${Date.now()}`;
}
function AiAssistant({
  title: title2 = "Assistente de criação",
  heading: heading2 = "Tenha ajuda na criação das suas missões",
  suggestions: suggestions2 = [
    "Faça uma revisão do meu Q1",
    "Dê sugestões de OKR",
    "Sugestões de PDI com base na minha performance"
  ],
  placeholder: placeholder2 = "Pergunte ao bud...",
  onClose,
  onMessage,
  onUseAsBase,
  allowUpload,
  missions,
  selectedMissions,
  onMissionsChange
}) {
  const [value2, setValue] = useState("");
  const [messages, setMessages] = useState([]);
  const [status, setStatus] = useState("idle");
  const [streamingText, setStreamingText] = useState("");
  const [attachOpen, setAttachOpen] = useState(false);
  const [missionSearch, setMissionSearch] = useState("");
  const [attachedFiles, setAttachedFiles] = useState([]);
  const [pendingUpload, setPendingUpload] = useState(false);
  const messageListRef = useRef(null);
  const attachButtonRef = useRef(null);
  const fileInputRef = useRef(null);
  const busy = status !== "idle";
  const hasAttach = !!(allowUpload || missions);
  const scrollToBottom = useCallback(() => {
    const el = messageListRef.current;
    if (el) {
      el.scrollTo({ top: el.scrollHeight, behavior: "smooth" });
    }
  }, []);
  useEffect(() => {
    if (messages.length > 0) {
      scrollToBottom();
    }
  }, [messages, scrollToBottom]);
  useEffect(() => {
    if (status === "streaming") {
      scrollToBottom();
    }
  }, [streamingText, status, scrollToBottom]);
  useLayoutEffect(() => {
    if (pendingUpload) {
      setPendingUpload(false);
      fileInputRef.current?.click();
    }
  }, [pendingUpload]);
  async function sendMessage(text2) {
    const fileCount = attachedFiles.length;
    const missionCount = selectedMissionCount;
    const context = fileCount > 0 || missionCount > 0 ? { fileCount, missionCount } : void 0;
    const userMsg = { id: uid(), role: "user", content: text2, context };
    setMessages((prev2) => [...prev2, userMsg]);
    setValue("");
    setAttachedFiles([]);
    if (missionCount > 0) onMissionsChange?.([]);
    if (!onMessage) return;
    setStatus("thinking");
    try {
      const response = await onMessage(text2);
      setStatus("streaming");
      setStreamingText("");
      await new Promise((resolve) => {
        let i = 0;
        const interval = setInterval(() => {
          i++;
          setStreamingText(response.slice(0, i));
          if (i >= response.length) {
            clearInterval(interval);
            resolve();
          }
        }, 30);
      });
      const aiMsg = { id: uid(), role: "assistant", content: response };
      setMessages((prev2) => [...prev2, aiMsg]);
    } finally {
      setStatus("idle");
      setStreamingText("");
    }
  }
  function handleSubmit(e) {
    e.preventDefault();
    const trimmed = value2.trim();
    if (trimmed && !busy) {
      sendMessage(trimmed);
    }
  }
  function handleSuggestion(text2) {
    if (!busy) {
      sendMessage(text2);
    }
  }
  function handleToggleAttach() {
    if (busy) return;
    setAttachOpen((prev2) => {
      if (prev2) setMissionSearch("");
      return !prev2;
    });
  }
  function handleFileChange(e) {
    const files = e.target.files;
    if (files && files.length > 0) {
      setAttachedFiles((prev2) => [...prev2, ...Array.from(files)]);
    }
    e.target.value = "";
  }
  function handleMissionToggle(id) {
    if (!onMissionsChange || !selectedMissions) return;
    const next2 = selectedMissions.includes(id) ? selectedMissions.filter((m) => m !== id) : [...selectedMissions, id];
    onMissionsChange(next2);
  }
  const filteredMissions = missions?.filter(
    (m) => m.label.toLowerCase().includes(missionSearch.toLowerCase())
  );
  const selectedMissionCount = selectedMissions?.filter(
    (id) => missions?.some((m) => m.id === id)
  ).length ?? 0;
  const hasContext = attachedFiles.length > 0 || selectedMissionCount > 0;
  const hasMessages = messages.length > 0 || busy;
  return /* @__PURE__ */ jsxs("div", { className: s$C.container, role: "region", "aria-label": title2, children: [
    /* @__PURE__ */ jsxs("div", { className: s$C.header, children: [
      /* @__PURE__ */ jsx(Lightning, { size: 16, className: s$C.headerIcon }),
      /* @__PURE__ */ jsx("h2", { className: s$C.title, children: title2 }),
      onClose && /* @__PURE__ */ jsx(
        Button,
        {
          variant: "tertiary",
          size: "md",
          leftIcon: SidebarSimple,
          onClick: onClose,
          "aria-label": "Recolher assistente"
        }
      )
    ] }),
    !hasMessages ? /* @__PURE__ */ jsxs("div", { className: s$C.body, children: [
      /* @__PURE__ */ jsx("div", { className: s$C.symbol, "aria-hidden": "true", children: /* @__PURE__ */ jsx(BudSymbol, {}) }),
      /* @__PURE__ */ jsx("p", { className: s$C.heading, children: heading2 }),
      suggestions2.length > 0 && /* @__PURE__ */ jsx("div", { className: s$C.suggestions, children: suggestions2.map((text2) => /* @__PURE__ */ jsxs(
        "button",
        {
          type: "button",
          className: s$C.suggestion,
          onClick: () => handleSuggestion(text2),
          children: [
            /* @__PURE__ */ jsx(Lightning, { size: 14, "aria-hidden": "true", style: { flexShrink: 0, marginTop: 2 } }),
            /* @__PURE__ */ jsx("span", { children: text2 })
          ]
        },
        text2
      )) })
    ] }) : /* @__PURE__ */ jsxs(
      "div",
      {
        className: s$C.messageList,
        ref: messageListRef,
        role: "log",
        "aria-live": "polite",
        children: [
          messages.map(
            (msg) => msg.role === "user" ? /* @__PURE__ */ jsxs("div", { className: s$C.userMessage, children: [
              msg.context && /* @__PURE__ */ jsxs("div", { className: s$C.userContext, children: [
                msg.context.fileCount > 0 && /* @__PURE__ */ jsxs("span", { className: s$C.userContextBadge, children: [
                  /* @__PURE__ */ jsx(File, { size: 10, "aria-hidden": "true" }),
                  msg.context.fileCount,
                  " ",
                  msg.context.fileCount === 1 ? "arquivo" : "arquivos"
                ] }),
                msg.context.missionCount > 0 && /* @__PURE__ */ jsxs("span", { className: s$C.userContextBadge, children: [
                  /* @__PURE__ */ jsx(Link, { size: 10, "aria-hidden": "true" }),
                  msg.context.missionCount,
                  " ",
                  msg.context.missionCount === 1 ? "missão / pesquisa" : "missões / pesquisas"
                ] })
              ] }),
              /* @__PURE__ */ jsx("p", { className: s$C.userText, children: msg.content })
            ] }, msg.id) : /* @__PURE__ */ jsxs("div", { className: s$C.aiMessage, children: [
              /* @__PURE__ */ jsx(
                Lightning,
                {
                  size: 16,
                  className: s$C.aiIcon,
                  "aria-hidden": "true"
                }
              ),
              /* @__PURE__ */ jsxs("div", { className: s$C.aiContent, children: [
                /* @__PURE__ */ jsx("p", { className: s$C.aiText, children: msg.content }),
                onUseAsBase && /* @__PURE__ */ jsxs(
                  "button",
                  {
                    type: "button",
                    className: s$C.useAsBaseBtn,
                    onClick: () => onUseAsBase(msg.content),
                    children: [
                      /* @__PURE__ */ jsx(ArrowUp, { size: 12, "aria-hidden": "true" }),
                      "Usar como base"
                    ]
                  }
                )
              ] })
            ] }, msg.id)
          ),
          status === "thinking" && /* @__PURE__ */ jsxs(
            "div",
            {
              className: s$C.aiMessage,
              "aria-label": "Assistente está digitando",
              children: [
                /* @__PURE__ */ jsx(
                  Lightning,
                  {
                    size: 16,
                    className: s$C.aiIcon,
                    "aria-hidden": "true"
                  }
                ),
                /* @__PURE__ */ jsxs("div", { className: s$C.typingDots, children: [
                  /* @__PURE__ */ jsx("span", { className: s$C.dot }),
                  /* @__PURE__ */ jsx("span", { className: s$C.dot }),
                  /* @__PURE__ */ jsx("span", { className: s$C.dot })
                ] })
              ]
            }
          ),
          status === "streaming" && /* @__PURE__ */ jsxs("div", { className: s$C.aiMessage, children: [
            /* @__PURE__ */ jsx(
              Lightning,
              {
                size: 16,
                className: s$C.aiIcon,
                "aria-hidden": "true"
              }
            ),
            /* @__PURE__ */ jsxs("p", { className: s$C.aiText, children: [
              streamingText,
              /* @__PURE__ */ jsx("span", { className: s$C.cursor, "aria-hidden": "true" })
            ] })
          ] })
        ]
      }
    ),
    /* @__PURE__ */ jsxs("form", { className: s$C.inputArea, onSubmit: handleSubmit, children: [
      /* @__PURE__ */ jsx("div", { className: s$C.inputGlow, "aria-hidden": "true" }),
      allowUpload && /* @__PURE__ */ jsx(
        "input",
        {
          ref: fileInputRef,
          type: "file",
          multiple: true,
          className: s$C.fileInput,
          onChange: handleFileChange,
          tabIndex: -1
        }
      ),
      /* @__PURE__ */ jsxs("div", { className: s$C.inputBox, children: [
        hasContext && /* @__PURE__ */ jsxs("div", { className: s$C.contextStrip, children: [
          attachedFiles.length > 0 && /* @__PURE__ */ jsxs("span", { className: s$C.contextBadge, children: [
            /* @__PURE__ */ jsx(File, { size: 12, className: s$C.contextBadgeIcon }),
            /* @__PURE__ */ jsxs("span", { className: s$C.contextBadgeLabel, children: [
              attachedFiles.length,
              " ",
              attachedFiles.length === 1 ? "arquivo" : "arquivos"
            ] }),
            /* @__PURE__ */ jsx(
              "button",
              {
                type: "button",
                className: s$C.contextBadgeRemove,
                onClick: () => setAttachedFiles([]),
                "aria-label": "Remover todos os arquivos",
                children: /* @__PURE__ */ jsx(X, { size: 10 })
              }
            )
          ] }),
          selectedMissionCount > 0 && /* @__PURE__ */ jsxs("span", { className: s$C.contextBadge, children: [
            /* @__PURE__ */ jsx(Link, { size: 12, className: s$C.contextBadgeIcon }),
            /* @__PURE__ */ jsxs("span", { className: s$C.contextBadgeLabel, children: [
              selectedMissionCount,
              " ",
              selectedMissionCount === 1 ? "missão / pesquisa" : "missões / pesquisas"
            ] }),
            /* @__PURE__ */ jsx(
              "button",
              {
                type: "button",
                className: s$C.contextBadgeRemove,
                onClick: () => onMissionsChange?.([]),
                "aria-label": "Remover todas as missões e pesquisas",
                children: /* @__PURE__ */ jsx(X, { size: 10 })
              }
            )
          ] })
        ] }),
        /* @__PURE__ */ jsxs("div", { className: s$C.inputRow, children: [
          /* @__PURE__ */ jsx(
            "input",
            {
              className: s$C.input,
              value: value2,
              onChange: (e) => setValue(e.target.value),
              placeholder: placeholder2,
              "aria-label": placeholder2,
              disabled: busy
            }
          ),
          hasAttach && /* @__PURE__ */ jsxs(Fragment, { children: [
            /* @__PURE__ */ jsx(
              "button",
              {
                type: "button",
                className: s$C.attachButton,
                ref: attachButtonRef,
                onClick: handleToggleAttach,
                "aria-label": "Opções de anexo",
                disabled: busy,
                children: /* @__PURE__ */ jsx(Plus, { size: 14 })
              }
            ),
            /* @__PURE__ */ jsx(
              Popover,
              {
                items: (() => {
                  const menuItems = [];
                  if (allowUpload) {
                    menuItems.push({
                      id: "upload",
                      label: "Enviar arquivo",
                      icon: FileArrowUp,
                      onClick: () => setPendingUpload(true)
                    });
                  }
                  if (missions && missions.length > 0) {
                    menuItems.push({
                      id: "missions",
                      label: "Missão / Pesquisa",
                      icon: Link,
                      badge: selectedMissionCount,
                      submenu: /* @__PURE__ */ jsxs(Fragment, { children: [
                        /* @__PURE__ */ jsxs("div", { className: s$C.missionSearchBox, children: [
                          /* @__PURE__ */ jsx(
                            MagnifyingGlass,
                            {
                              size: 14,
                              className: s$C.missionSearchIcon
                            }
                          ),
                          /* @__PURE__ */ jsx(
                            "input",
                            {
                              className: s$C.missionSearchInput,
                              value: missionSearch,
                              onChange: (e) => setMissionSearch(e.target.value),
                              placeholder: "Buscar missão...",
                              "aria-label": "Buscar missão"
                            }
                          )
                        ] }),
                        /* @__PURE__ */ jsx("div", { className: s$C.missionList, children: filteredMissions?.map((m) => /* @__PURE__ */ jsxs("label", { className: s$C.missionListItem, children: [
                          /* @__PURE__ */ jsx(
                            Checkbox,
                            {
                              size: "sm",
                              checked: selectedMissions?.includes(m.id) ?? false,
                              onChange: () => handleMissionToggle(m.id)
                            }
                          ),
                          /* @__PURE__ */ jsx("span", { className: s$C.missionLabel, children: m.label })
                        ] }, m.id)) })
                      ] })
                    });
                  }
                  return menuItems;
                })(),
                open: attachOpen,
                onClose: () => {
                  setAttachOpen(false);
                  setMissionSearch("");
                },
                anchorRef: attachButtonRef
              }
            )
          ] }),
          /* @__PURE__ */ jsx(
            "button",
            {
              type: "submit",
              className: s$C.submitButton,
              "aria-label": "Enviar mensagem",
              disabled: !value2.trim() || busy,
              children: /* @__PURE__ */ jsx(ArrowUp, { size: 14 })
            }
          )
        ] })
      ] })
    ] })
  ] });
}
function BudSymbol() {
  return /* @__PURE__ */ jsx("svg", { viewBox: "0 0 240 92", fill: "none", className: s$C.budSymbol, children: /* @__PURE__ */ jsx(
    "path",
    {
      d: "M239.986 44.8463C239.406 21.0338 220.559 1.4407 196.862 0.0776031C189.718 -0.333 182.901 0.900098 176.753 3.43709V3.43451L155.54 12.2361C153.849 12.9376 151.913 12.6686 150.476 11.532L148.329 9.83555C140.532 3.67521 130.693 0.000373357 119.999 0.000373357C113.82 0.000373357 107.926 1.22832 102.545 3.45511V3.45253L6.85696 43.3789L6.85439 43.3802L6.50423 43.5269L6.47217 43.5398V43.5411C5.14977 44.1216 3.9659 44.9634 2.98469 45.9983C1.13641 47.947 0 50.5831 0 53.4869C0 56.3882 1.13385 59.0243 2.9834 60.9756C4.41226 62.4816 6.26695 63.5757 8.35507 64.0648L8.37816 64.0712L107.653 90.31C111.58 91.4105 115.722 91.9987 119.999 91.9987C130.693 91.9987 140.53 88.3239 148.329 82.1635C148.969 81.6577 149.595 81.1364 150.207 80.5984C151.668 79.3138 153.718 78.9521 155.513 79.7012L176.325 88.3844L176.755 88.5646C182.124 90.7785 188.003 92 194.166 92C214.255 92 231.327 79.0306 237.515 60.9782C239.243 55.9364 240.124 50.5046 239.986 44.8463ZM148.33 46.2827V46.3342C148.33 46.3754 148.33 46.4166 148.329 46.4565C148.245 51.7789 146.704 56.7447 144.087 60.9743C141.591 65.0083 138.118 68.3678 133.998 70.7207C129.87 73.0788 125.094 74.4264 120.004 74.4264C115.121 74.4264 110.528 73.1869 106.52 71.0052C102.178 68.6419 98.5204 65.1731 95.9217 60.9743C94.498 58.6729 93.3924 56.1539 92.6664 53.4754C92.02 51.0915 91.6762 48.5854 91.6762 45.997C91.6762 43.3287 92.0418 40.7441 92.7293 38.2946C94.5083 31.9335 98.4383 26.4734 103.683 22.7586C108.297 19.4879 113.927 17.5675 120.006 17.5675C125.053 17.5675 129.792 18.892 133.897 21.2153C142.39 26.019 148.167 35.0947 148.33 45.5387C148.331 45.5799 148.331 45.6198 148.331 45.6597V45.7112C148.334 45.8065 148.334 45.9017 148.334 45.997C148.333 46.0922 148.333 46.1875 148.33 46.2827ZM218.252 60.9743C213.256 69.05 204.338 74.4264 194.169 74.4264C187.602 74.4264 181.558 72.1829 176.755 68.4205C174.119 66.3572 171.857 63.8344 170.087 60.9743C167.396 56.625 165.842 51.4931 165.842 45.9983C165.842 36.8878 170.113 28.7774 176.755 23.576C181.558 19.8123 187.602 17.5701 194.169 17.5701C209.814 17.5701 222.498 30.2975 222.498 45.9995C222.497 51.4931 220.943 56.6237 218.252 60.9743Z",
      fill: "currentColor"
    }
  ) });
}
const alert = "_alert_ix5n1_3";
const icon = "_icon_ix5n1_15";
const content = "_content_ix5n1_23";
const title$8 = "_title_ix5n1_31";
const description$7 = "_description_ix5n1_39";
const action = "_action_ix5n1_52";
const dismiss = "_dismiss_ix5n1_74";
const info = "_info_ix5n1_101";
const success$6 = "_success_ix5n1_124";
const warning$2 = "_warning_ix5n1_151";
const error$6 = "_error_ix5n1_178";
const s$B = {
  alert,
  icon,
  content,
  title: title$8,
  description: description$7,
  action,
  dismiss,
  info,
  success: success$6,
  warning: warning$2,
  error: error$6
};
const variantIcon$1 = {
  info: Info,
  success: CheckCircle,
  warning: Warning,
  error: WarningCircle
};
function Alert({
  variant = "info",
  title: title2,
  children,
  onDismiss,
  action: action2,
  className,
  ...rest
}) {
  const Icon = variantIcon$1[variant];
  const classes = [s$B.alert, s$B[variant], className ?? ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("div", { className: classes, role: "alert", ...rest, children: [
    /* @__PURE__ */ jsx("span", { className: s$B.icon, children: /* @__PURE__ */ jsx(Icon, { size: 20, "aria-hidden": true }) }),
    /* @__PURE__ */ jsxs("div", { className: s$B.content, children: [
      /* @__PURE__ */ jsx("p", { className: s$B.title, children: title2 }),
      children && /* @__PURE__ */ jsx("div", { className: s$B.description, children }),
      action2 && /* @__PURE__ */ jsx(
        "button",
        {
          type: "button",
          className: s$B.action,
          onClick: action2.onClick,
          children: action2.label
        }
      )
    ] }),
    onDismiss && /* @__PURE__ */ jsx(
      "button",
      {
        type: "button",
        className: s$B.dismiss,
        onClick: onDismiss,
        "aria-label": "Fechar alerta",
        children: /* @__PURE__ */ jsx(X, { size: 16, "aria-hidden": true })
      }
    )
  ] });
}
const avatar = "_avatar_xixil_2";
const hasInitials = "_hasInitials_xixil_18";
const placeholder$2 = "_placeholder_xixil_24";
const placeholderIcon = "_placeholderIcon_xixil_28";
const image = "_image_xixil_34";
const initials = "_initials_xixil_45";
const xs$1 = "_xs_xixil_55";
const sm$c = "_sm_xixil_56";
const md$a = "_md_xixil_57";
const lg$8 = "_lg_xixil_58";
const xl$1 = "_xl_xixil_59";
const xxl = "_xxl_xixil_60";
const online = "_online_xixil_70";
const company = "_company_xixil_87";
const s$A = {
  avatar,
  hasInitials,
  placeholder: placeholder$2,
  placeholderIcon,
  image,
  initials,
  xs: xs$1,
  sm: sm$c,
  md: md$a,
  lg: lg$8,
  xl: xl$1,
  xxl,
  online,
  company
};
const sizeClass$2 = {
  xs: s$A.xs,
  sm: s$A.sm,
  md: s$A.md,
  lg: s$A.lg,
  xl: s$A.xl,
  "2xl": s$A.xxl
};
const iconSize$2 = {
  xs: 14,
  sm: 16,
  md: 24,
  lg: 28,
  xl: 32,
  "2xl": 36
};
function Avatar({
  size = "md",
  src,
  initials: initials2,
  alt = "",
  online: online2,
  companyLogo,
  className
}) {
  const isImage = !!src;
  const isInitials = !src && !!initials2;
  const classes = [
    s$A.avatar,
    sizeClass$2[size],
    isInitials ? s$A.hasInitials : "",
    !isImage && !isInitials ? s$A.placeholder : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("span", { className: classes, children: [
    isImage && /* @__PURE__ */ jsx("img", { className: s$A.image, src, alt }),
    isInitials && /* @__PURE__ */ jsx("span", { className: s$A.initials, children: initials2 }),
    !isImage && !isInitials && /* @__PURE__ */ jsx(User, { size: iconSize$2[size], className: s$A.placeholderIcon }),
    online2 && !companyLogo && /* @__PURE__ */ jsx("span", { className: s$A.online }),
    companyLogo && /* @__PURE__ */ jsx("span", { className: s$A.company, children: /* @__PURE__ */ jsx("img", { src: companyLogo, alt: "" }) })
  ] });
}
const group$3 = "_group_a73zn_2";
const stack = "_stack_a73zn_7";
const stackItem = "_stackItem_a73zn_13";
const more = "_more_a73zn_19";
const moreText = "_moreText_a73zn_30";
const addButton = "_addButton_a73zn_39";
const addIcon = "_addIcon_a73zn_56";
const xs = "_xs_a73zn_61";
const gapAdd = "_gapAdd_a73zn_66";
const gapMore = "_gapMore_a73zn_67";
const sm$b = "_sm_a73zn_70";
const md$9 = "_md_a73zn_79";
const s$z = {
  group: group$3,
  stack,
  stackItem,
  more,
  moreText,
  addButton,
  addIcon,
  xs,
  gapAdd,
  gapMore,
  sm: sm$b,
  md: md$9
};
const sizeClass$1 = {
  xs: s$z.xs,
  sm: s$z.sm,
  md: s$z.md
};
function AvatarGroup({
  size = "sm",
  avatars,
  maxVisible = 5,
  showAddButton,
  onAddClick,
  className
}) {
  const visible = avatars.slice(0, maxVisible);
  const remaining = avatars.length - maxVisible;
  const hasMore = remaining > 0;
  const classes = [
    s$z.group,
    sizeClass$1[size],
    showAddButton ? hasMore ? s$z.gapMore : s$z.gapAdd : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("div", { className: classes, children: [
    /* @__PURE__ */ jsxs("div", { className: s$z.stack, children: [
      visible.map((avatar2, i) => /* @__PURE__ */ jsx(
        Avatar,
        {
          size,
          src: avatar2.src,
          initials: avatar2.initials,
          alt: avatar2.alt,
          className: s$z.stackItem
        },
        i
      )),
      hasMore && /* @__PURE__ */ jsx("span", { className: s$z.more, children: /* @__PURE__ */ jsxs("span", { className: s$z.moreText, children: [
        "+",
        remaining
      ] }) })
    ] }),
    showAddButton && /* @__PURE__ */ jsx(
      "button",
      {
        type: "button",
        className: s$z.addButton,
        onClick: onAddClick,
        "aria-label": "Adicionar usuário",
        children: /* @__PURE__ */ jsx(Plus, { size: 16, className: s$z.addIcon })
      }
    )
  ] });
}
const group$2 = "_group_12ql5_2";
const text$4 = "_text_12ql5_8";
const name$1 = "_name_12ql5_14";
const supporting = "_supporting_12ql5_22";
const sm$a = "_sm_12ql5_29";
const md$8 = "_md_12ql5_30";
const lg$7 = "_lg_12ql5_31";
const xl = "_xl_12ql5_32";
const s$y = {
  group: group$2,
  text: text$4,
  name: name$1,
  supporting,
  sm: sm$a,
  md: md$8,
  lg: lg$7,
  xl
};
const sizeClass = {
  sm: s$y.sm,
  md: s$y.md,
  lg: s$y.lg,
  xl: s$y.xl
};
function AvatarLabelGroup({
  size = "md",
  name: name2,
  supportingText,
  src,
  initials: initials2,
  alt,
  online: online2,
  companyLogo,
  className
}) {
  const classes = [s$y.group, sizeClass[size], className ?? ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("div", { className: classes, children: [
    /* @__PURE__ */ jsx(
      Avatar,
      {
        size,
        src,
        initials: initials2,
        alt,
        online: online2,
        companyLogo
      }
    ),
    /* @__PURE__ */ jsxs("div", { className: s$y.text, children: [
      /* @__PURE__ */ jsx("span", { className: s$y.name, children: name2 }),
      supportingText && /* @__PURE__ */ jsx("span", { className: s$y.supporting, children: supportingText })
    ] })
  ] });
}
const badge$1 = "_badge_1mkcp_2";
const sm$9 = "_sm_1mkcp_16";
const iconOnly = "_iconOnly_1mkcp_22";
const md$7 = "_md_1mkcp_27";
const lg$6 = "_lg_1mkcp_38";
const label$c = "_label_1mkcp_49";
const neutral$2 = "_neutral_1mkcp_55";
const orange$1 = "_orange_1mkcp_63";
const wine = "_wine_1mkcp_69";
const caramel = "_caramel_1mkcp_75";
const error$5 = "_error_1mkcp_83";
const warning$1 = "_warning_1mkcp_91";
const success$5 = "_success_1mkcp_97";
const s$x = {
  badge: badge$1,
  sm: sm$9,
  iconOnly,
  md: md$7,
  lg: lg$6,
  label: label$c,
  neutral: neutral$2,
  orange: orange$1,
  wine,
  caramel,
  error: error$5,
  warning: warning$1,
  success: success$5
};
const iconSize$1 = { sm: 12, md: 12, lg: 16 };
function Badge({
  color = "neutral",
  size = "sm",
  leftIcon: LeftIcon,
  rightIcon: RightIcon,
  children,
  className,
  ...rest
}) {
  const isIconOnly = !children && (LeftIcon || RightIcon);
  const Icon = isIconOnly ? LeftIcon || RightIcon : null;
  const classes = [
    s$x.badge,
    s$x[color],
    s$x[size],
    isIconOnly ? s$x.iconOnly : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx(
    "span",
    {
      className: classes,
      role: isIconOnly ? "img" : void 0,
      "aria-hidden": isIconOnly && !rest["aria-label"] ? true : void 0,
      ...rest,
      children: isIconOnly ? Icon && /* @__PURE__ */ jsx(Icon, { size: iconSize$1[size], "aria-hidden": true }) : /* @__PURE__ */ jsxs(Fragment, { children: [
        LeftIcon && /* @__PURE__ */ jsx(LeftIcon, { size: iconSize$1[size], "aria-hidden": true }),
        children && /* @__PURE__ */ jsx("span", { className: s$x.label, children }),
        RightIcon && /* @__PURE__ */ jsx(RightIcon, { size: iconSize$1[size], "aria-hidden": true })
      ] })
    }
  );
}
const breadcrumb = "_breadcrumb_17owf_2";
const list$4 = "_list_17owf_6";
const step$1 = "_step_17owf_15";
const item$7 = "_item_17owf_22";
const itemCompleted = "_itemCompleted_17owf_31";
const clickable = "_clickable_17owf_38";
const separator$1 = "_separator_17owf_64";
const s$w = {
  breadcrumb,
  list: list$4,
  step: step$1,
  item: item$7,
  itemCompleted,
  clickable,
  separator: separator$1
};
function Breadcrumb({ items: items2, current = 0 }) {
  return /* @__PURE__ */ jsx("nav", { className: s$w.breadcrumb, "aria-label": "Breadcrumb", children: /* @__PURE__ */ jsx("ol", { className: s$w.list, children: items2.map((item2, i) => {
    const isPast = i < current;
    const isCurrent = i === current;
    const isLast = i === items2.length - 1;
    const itemClass = [
      s$w.item,
      isPast || isCurrent ? s$w.itemCompleted : ""
    ].filter(Boolean).join(" ");
    return /* @__PURE__ */ jsxs("li", { className: s$w.step, children: [
      isPast && item2.href ? /* @__PURE__ */ jsx("a", { href: item2.href, className: itemClass, children: item2.label }) : isPast && item2.onClick ? /* @__PURE__ */ jsx(
        "button",
        {
          type: "button",
          className: `${itemClass} ${s$w.clickable}`,
          onClick: item2.onClick,
          children: item2.label
        }
      ) : /* @__PURE__ */ jsx(
        "span",
        {
          className: itemClass,
          "aria-current": isCurrent ? "step" : void 0,
          children: item2.label
        }
      ),
      !isLast && /* @__PURE__ */ jsx(
        CaretRight,
        {
          size: 16,
          className: s$w.separator,
          "aria-hidden": true
        }
      )
    ] }, i);
  }) }) });
}
const card$1 = "_card_sow7f_3";
const shadow = "_shadow_sow7f_12";
const header$6 = "_header_sow7f_32";
const body$4 = "_body_sow7f_40";
const description$6 = "_description_sow7f_45";
const headerText$2 = "_headerText_sow7f_49";
const title$7 = "_title_sow7f_57";
const headerAction = "_headerAction_sow7f_74";
const footer$7 = "_footer_sow7f_100";
const divider$1 = "_divider_sow7f_123";
const s$v = {
  card: card$1,
  shadow,
  "padding-sm": "_padding-sm_sow7f_18",
  "padding-md": "_padding-md_sow7f_22",
  "padding-lg": "_padding-lg_sow7f_26",
  header: header$6,
  body: body$4,
  description: description$6,
  headerText: headerText$2,
  title: title$7,
  headerAction,
  footer: footer$7,
  divider: divider$1
};
function Card({
  padding = "md",
  shadow: shadow2 = false,
  children,
  className,
  ...rest
}) {
  const classes = [
    s$v.card,
    padding !== "none" ? s$v[`padding-${padding}`] : "",
    shadow2 ? s$v.shadow : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx("div", { className: classes, ...rest, children });
}
function CardHeader({
  title: title2,
  description: description2,
  action: action2,
  className,
  ...rest
}) {
  const classes = [s$v.header, className ?? ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("div", { className: classes, ...rest, children: [
    /* @__PURE__ */ jsxs("div", { className: s$v.headerText, children: [
      /* @__PURE__ */ jsx("h3", { className: s$v.title, children: title2 }),
      description2 && /* @__PURE__ */ jsx("p", { className: s$v.description, children: description2 })
    ] }),
    action2 && /* @__PURE__ */ jsx("div", { className: s$v.headerAction, children: action2 })
  ] });
}
function CardBody({ children, className, ...rest }) {
  const classes = [s$v.body, className ?? ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx("div", { className: classes, ...rest, children });
}
function CardFooter({ children, className, ...rest }) {
  const classes = [s$v.footer, className ?? ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx("div", { className: classes, ...rest, children });
}
function CardDivider() {
  return /* @__PURE__ */ jsx("hr", { className: s$v.divider });
}
const chart = "_chart_y4gra_5";
const ring = "_ring_y4gra_16";
const track$2 = "_track_y4gra_23";
const progress = "_progress_y4gra_29";
const textGroup = "_textGroup_y4gra_40";
const value$3 = "_value_y4gra_47";
const percent = "_percent_y4gra_56";
const halfChart = "_halfChart_y4gra_69";
const halfProgress = "_halfProgress_y4gra_76";
const halfText = "_halfText_y4gra_85";
const halfValue = "_halfValue_y4gra_89";
const halfPercent = "_halfPercent_y4gra_96";
const s$u = {
  chart,
  ring,
  track: track$2,
  progress,
  textGroup,
  value: value$3,
  percent,
  halfChart,
  halfProgress,
  halfText,
  halfValue,
  halfPercent
};
const RADIUS = 18;
const FULL_CIRCUMFERENCE = 2 * Math.PI * RADIUS;
const HALF_CIRCUMFERENCE = Math.PI * RADIUS;
function clamp$3(v, lo, hi) {
  return Math.min(hi, Math.max(lo, v));
}
function Chart({ value: value2, variant = "full", size = 40, className }) {
  const clamped = clamp$3(Math.round(value2), 0, 100);
  const vars = { "--chart-size": `${size}px` };
  if (variant === "half") {
    const offset2 = HALF_CIRCUMFERENCE - clamped / 100 * HALF_CIRCUMFERENCE;
    return /* @__PURE__ */ jsxs(
      "svg",
      {
        className: `${s$u.halfChart} ${className ?? ""}`,
        style: vars,
        viewBox: "0 0 40 22",
        role: "img",
        "aria-label": `${clamped}%`,
        children: [
          /* @__PURE__ */ jsx(
            "path",
            {
              d: "M 2,20 A 18,18 0 0,1 38,20",
              className: s$u.track
            }
          ),
          clamped > 0 && /* @__PURE__ */ jsx(
            "path",
            {
              d: "M 2,20 A 18,18 0 0,1 38,20",
              className: s$u.halfProgress,
              strokeDasharray: HALF_CIRCUMFERENCE,
              strokeDashoffset: offset2
            }
          ),
          /* @__PURE__ */ jsxs("text", { x: "20", y: "20", textAnchor: "middle", className: s$u.halfText, children: [
            /* @__PURE__ */ jsx("tspan", { className: s$u.halfValue, children: clamped }),
            /* @__PURE__ */ jsx("tspan", { className: s$u.halfPercent, children: "%" })
          ] })
        ]
      }
    );
  }
  const offset = FULL_CIRCUMFERENCE - clamped / 100 * FULL_CIRCUMFERENCE;
  return /* @__PURE__ */ jsxs(
    "div",
    {
      className: `${s$u.chart} ${className ?? ""}`,
      style: vars,
      role: "img",
      "aria-label": `${clamped}%`,
      children: [
        /* @__PURE__ */ jsxs("svg", { className: s$u.ring, viewBox: "0 0 40 40", "aria-hidden": "true", children: [
          /* @__PURE__ */ jsx("circle", { cx: "20", cy: "20", r: RADIUS, className: s$u.track }),
          clamped > 0 && /* @__PURE__ */ jsx(
            "circle",
            {
              cx: "20",
              cy: "20",
              r: RADIUS,
              className: s$u.progress,
              strokeDasharray: FULL_CIRCUMFERENCE,
              strokeDashoffset: offset
            }
          )
        ] }),
        /* @__PURE__ */ jsxs("span", { className: s$u.textGroup, children: [
          /* @__PURE__ */ jsx("span", { className: s$u.value, children: clamped }),
          /* @__PURE__ */ jsx("span", { className: s$u.percent, children: "%" })
        ] })
      ]
    }
  );
}
const tooltip$2 = "_tooltip_10xca_1";
const label$b = "_label_10xca_16";
const items$1 = "_items_10xca_25";
const item$6 = "_item_10xca_25";
const dot$2 = "_dot_10xca_37";
const name = "_name_10xca_44";
const value$2 = "_value_10xca_53";
const s$t = {
  tooltip: tooltip$2,
  label: label$b,
  items: items$1,
  item: item$6,
  dot: dot$2,
  name,
  value: value$2
};
function ChartTooltipContent({
  active: active2,
  payload,
  label: label2,
  labelFormatter,
  valueFormatter
}) {
  if (!active2 || !payload?.length) return null;
  const formattedLabel = labelFormatter ? labelFormatter(String(label2)) : label2;
  return /* @__PURE__ */ jsxs("div", { className: s$t.tooltip, children: [
    formattedLabel && /* @__PURE__ */ jsx("p", { className: s$t.label, children: formattedLabel }),
    /* @__PURE__ */ jsx("div", { className: s$t.items, children: payload.map((entry, i) => /* @__PURE__ */ jsxs("div", { className: s$t.item, children: [
      /* @__PURE__ */ jsx(
        "span",
        {
          className: s$t.dot,
          style: { backgroundColor: entry.color }
        }
      ),
      /* @__PURE__ */ jsx("span", { className: s$t.name, children: entry.name }),
      /* @__PURE__ */ jsx("span", { className: s$t.value, children: valueFormatter ? valueFormatter(entry.value ?? 0) : entry.value?.toLocaleString("pt-BR") })
    ] }, i)) })
  ] });
}
const group$1 = "_group_19zz7_2";
const groupLabel$2 = "_groupLabel_19zz7_8";
const groupGrid = "_groupGrid_19zz7_16";
const card = "_card_19zz7_29";
const input$4 = "_input_19zz7_44";
const cardContent = "_cardContent_19zz7_56";
const radio = "_radio_19zz7_63";
const radioDot = "_radioDot_19zz7_76";
const checkbox = "_checkbox_19zz7_86";
const checkIcon = "_checkIcon_19zz7_99";
const text$3 = "_text_19zz7_106";
const title$6 = "_title_19zz7_113";
const description$5 = "_description_19zz7_124";
const linkRow = "_linkRow_19zz7_133";
const link = "_link_19zz7_133";
const disabled$9 = "_disabled_19zz7_152";
const groupDisabled = "_groupDisabled_19zz7_257";
const s$s = {
  group: group$1,
  groupLabel: groupLabel$2,
  groupGrid,
  card,
  input: input$4,
  cardContent,
  radio,
  radioDot,
  checkbox,
  checkIcon,
  text: text$3,
  title: title$6,
  description: description$5,
  linkRow,
  link,
  disabled: disabled$9,
  groupDisabled
};
const ChoiceBoxContext = createContext(null);
function ChoiceBoxGroup(props) {
  const {
    label: label2,
    name: name2,
    disabled: disabled2 = false,
    className,
    children,
    multiple = false
  } = props;
  const autoId = useId();
  const groupName = name2 ?? autoId;
  const [singleInternal, setSingleInternal] = useState(
    !multiple ? props.defaultValue : void 0
  );
  const [multiInternal, setMultiInternal] = useState(
    multiple ? props.defaultValue ?? [] : []
  );
  const isControlled = props.value !== void 0;
  const isSelected = useCallback(
    (val) => {
      if (multiple) {
        const current2 = isControlled ? props.value : multiInternal;
        return current2.includes(val);
      }
      const current = isControlled ? props.value : singleInternal;
      return current === val;
    },
    [multiple, isControlled, props.value, singleInternal, multiInternal]
  );
  const toggle = useCallback(
    (val) => {
      if (multiple) {
        const current = isControlled ? props.value : multiInternal;
        const next2 = current.includes(val) ? current.filter((v) => v !== val) : [...current, val];
        if (!isControlled) setMultiInternal(next2);
        props.onChange?.(next2);
      } else {
        const current = isControlled ? props.value : singleInternal;
        const next2 = current === val ? void 0 : val;
        if (!isControlled) setSingleInternal(next2);
        props.onChange?.(next2);
      }
    },
    [multiple, isControlled, props, singleInternal, multiInternal]
  );
  const groupClasses = [
    s$s.group,
    disabled2 ? s$s.groupDisabled : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx(
    ChoiceBoxContext.Provider,
    {
      value: { name: groupName, multiple, isSelected, toggle, disabled: disabled2 },
      children: /* @__PURE__ */ jsxs(
        "div",
        {
          className: groupClasses,
          role: multiple ? "group" : "radiogroup",
          "aria-label": label2,
          children: [
            label2 && /* @__PURE__ */ jsx("span", { className: s$s.groupLabel, children: label2 }),
            /* @__PURE__ */ jsx("div", { className: s$s.groupGrid, children })
          ]
        }
      )
    }
  );
}
function ChoiceBox({
  value: value2,
  title: title2,
  description: description2,
  link: link2,
  disabled: itemDisabled2 = false,
  className
}) {
  const ctx = useContext(ChoiceBoxContext);
  if (!ctx) throw new Error("ChoiceBox must be used inside ChoiceBoxGroup");
  const isDisabled2 = ctx.disabled || itemDisabled2;
  const isChecked = ctx.isSelected(value2);
  const cardClasses = [s$s.card, isDisabled2 ? s$s.disabled : "", className ?? ""].filter(Boolean).join(" ");
  const handleLinkClick = (e) => {
    e.stopPropagation();
    e.preventDefault();
    link2?.onClick?.();
    if (link2?.href) window.open(link2.href, "_blank", "noopener");
  };
  return /* @__PURE__ */ jsxs("label", { className: cardClasses, children: [
    /* @__PURE__ */ jsx(
      "input",
      {
        type: ctx.multiple ? "checkbox" : "radio",
        className: s$s.input,
        name: ctx.name,
        value: value2,
        checked: isChecked,
        disabled: isDisabled2,
        onChange: () => {
        },
        onClick: () => ctx.toggle(value2)
      }
    ),
    /* @__PURE__ */ jsxs("div", { className: s$s.cardContent, children: [
      ctx.multiple ? /* @__PURE__ */ jsx("span", { className: s$s.checkbox, children: /* @__PURE__ */ jsx(Check, { size: 12, className: s$s.checkIcon }) }) : /* @__PURE__ */ jsx("span", { className: s$s.radio, children: /* @__PURE__ */ jsx("span", { className: s$s.radioDot }) }),
      /* @__PURE__ */ jsxs("div", { className: s$s.text, children: [
        /* @__PURE__ */ jsx("span", { className: s$s.title, children: title2 }),
        description2 && /* @__PURE__ */ jsx("span", { className: s$s.description, children: description2 })
      ] })
    ] }),
    link2 && /* @__PURE__ */ jsx("div", { className: s$s.linkRow, children: /* @__PURE__ */ jsx(
      "a",
      {
        className: s$s.link,
        href: link2.href,
        onClick: handleLinkClick,
        tabIndex: isDisabled2 ? -1 : 0,
        children: link2.text
      }
    ) })
  ] });
}
const WEEKDAY_LABELS = ["Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"];
const MONTH_LABELS = [
  "Janeiro",
  "Fevereiro",
  "Março",
  "Abril",
  "Maio",
  "Junho",
  "Julho",
  "Agosto",
  "Setembro",
  "Outubro",
  "Novembro",
  "Dezembro"
];
function daysInMonth(year, month) {
  return new Date(year, month, 0).getDate();
}
function firstDayOfWeek(year, month) {
  return new Date(year, month - 1, 1).getDay();
}
function prevMonth(d) {
  return d.month === 1 ? { year: d.year - 1, month: 12, day: 1 } : { year: d.year, month: d.month - 1, day: 1 };
}
function nextMonth(d) {
  return d.month === 12 ? { year: d.year + 1, month: 1, day: 1 } : { year: d.year, month: d.month + 1, day: 1 };
}
function isSameDay(a, b) {
  return a.year === b.year && a.month === b.month && a.day === b.day;
}
function today() {
  const now = /* @__PURE__ */ new Date();
  return { year: now.getFullYear(), month: now.getMonth() + 1, day: now.getDate() };
}
function isToday(d) {
  return isSameDay(d, today());
}
function compareDates(a, b) {
  if (a.year !== b.year) return a.year < b.year ? -1 : 1;
  if (a.month !== b.month) return a.month < b.month ? -1 : 1;
  if (a.day !== b.day) return a.day < b.day ? -1 : 1;
  return 0;
}
function isInRange(d, start, end) {
  return compareDates(d, start) >= 0 && compareDates(d, end) <= 0;
}
function isDisabled(d, minDate, maxDate) {
  if (minDate && compareDates(d, minDate) < 0) return true;
  if (maxDate && compareDates(d, maxDate) > 0) return true;
  return false;
}
function pad2(n) {
  return n < 10 ? `0${n}` : `${n}`;
}
function formatDate(d) {
  return `${pad2(d.day)}/${pad2(d.month)}/${d.year}`;
}
function parseDate(str) {
  const match = str.match(/^(\d{2})\/(\d{2})\/(\d{4})$/);
  if (!match) return null;
  const day2 = parseInt(match[1], 10);
  const month = parseInt(match[2], 10);
  const year = parseInt(match[3], 10);
  if (!isValidDate(year, month, day2)) return null;
  return { year, month, day: day2 };
}
function isValidDate(year, month, day2) {
  if (month < 1 || month > 12) return false;
  if (day2 < 1 || day2 > daysInMonth(year, month)) return false;
  if (year < 1900 || year > 2100) return false;
  return true;
}
const wrapper$8 = "_wrapper_tvj5n_2";
const label$a = "_label_tvj5n_10";
const anchor$1 = "_anchor_tvj5n_19";
const trigger$2 = "_trigger_tvj5n_24";
const sm$8 = "_sm_tvj5n_40";
const md$6 = "_md_tvj5n_45";
const lg$5 = "_lg_tvj5n_50";
const rangeTrigger = "_rangeTrigger_tvj5n_56";
const disabled$8 = "_disabled_tvj5n_68";
const hovered$4 = "_hovered_tvj5n_69";
const open$2 = "_open_tvj5n_74";
const focused$5 = "_focused_tvj5n_75";
const error$4 = "_error_tvj5n_80";
const value$1 = "_value_tvj5n_104";
const placeholder$1 = "_placeholder_tvj5n_117";
const rangeTriggerInput = "_rangeTriggerInput_tvj5n_193";
const rangeSeparator = "_rangeSeparator_tvj5n_219";
const caret$2 = "_caret_tvj5n_246";
const caretOpen$2 = "_caretOpen_tvj5n_251";
const popover = "_popover_tvj5n_256";
const popoverHeader = "_popoverHeader_tvj5n_269";
const headerInput = "_headerInput_tvj5n_277";
const headerSeparator = "_headerSeparator_tvj5n_302";
const calendarNav = "_calendarNav_tvj5n_311";
const monthYear = "_monthYear_tvj5n_318";
const weekdays = "_weekdays_tvj5n_326";
const weekday = "_weekday_tvj5n_326";
const grid$1 = "_grid_tvj5n_344";
const day = "_day_tvj5n_352";
const dayDisabled = "_dayDisabled_tvj5n_370";
const daySelected = "_daySelected_tvj5n_370";
const dayRangeStart = "_dayRangeStart_tvj5n_370";
const dayRangeEnd = "_dayRangeEnd_tvj5n_370";
const dayToday = "_dayToday_tvj5n_380";
const dayInRange = "_dayInRange_tvj5n_408";
const dayOtherMonth = "_dayOtherMonth_tvj5n_433";
const popoverFooter = "_popoverFooter_tvj5n_444";
const todayBtn = "_todayBtn_tvj5n_453";
const message$3 = "_message_tvj5n_463";
const attention$4 = "_attention_tvj5n_477";
const success$4 = "_success_tvj5n_481";
const s$r = {
  wrapper: wrapper$8,
  label: label$a,
  anchor: anchor$1,
  trigger: trigger$2,
  sm: sm$8,
  md: md$6,
  lg: lg$5,
  rangeTrigger,
  disabled: disabled$8,
  hovered: hovered$4,
  open: open$2,
  focused: focused$5,
  error: error$4,
  value: value$1,
  placeholder: placeholder$1,
  rangeTriggerInput,
  rangeSeparator,
  caret: caret$2,
  caretOpen: caretOpen$2,
  popover,
  popoverHeader,
  headerInput,
  headerSeparator,
  calendarNav,
  monthYear,
  weekdays,
  weekday,
  grid: grid$1,
  day,
  dayDisabled,
  daySelected,
  dayRangeStart,
  dayRangeEnd,
  dayToday,
  dayInRange,
  dayOtherMonth,
  popoverFooter,
  todayBtn,
  message: message$3,
  attention: attention$4,
  success: success$4
};
const ICON_SIZES$2 = { sm: 14, md: 16, lg: 20 };
const messageIconMap$3 = {
  error: WarningCircle,
  attention: WarningCircle,
  success: CheckCircle
};
function buildGrid(year, month) {
  const cells = [];
  const first = firstDayOfWeek(year, month);
  const total = daysInMonth(year, month);
  const prev2 = prevMonth({ year, month });
  const prevDays = daysInMonth(prev2.year, prev2.month);
  for (let i = first - 1; i >= 0; i--) {
    cells.push({
      date: { year: prev2.year, month: prev2.month, day: prevDays - i },
      isCurrentMonth: false
    });
  }
  for (let d = 1; d <= total; d++) {
    cells.push({ date: { year, month, day: d }, isCurrentMonth: true });
  }
  const next2 = nextMonth({ year, month });
  const remaining = 42 - cells.length;
  for (let d = 1; d <= remaining; d++) {
    cells.push({
      date: { year: next2.year, month: next2.month, day: d },
      isCurrentMonth: false
    });
  }
  return cells;
}
function applyDateMask(raw) {
  const digits = raw.replace(/\D/g, "").slice(0, 8);
  if (digits.length <= 2) return digits;
  if (digits.length <= 4) return `${digits.slice(0, 2)}/${digits.slice(2)}`;
  return `${digits.slice(0, 2)}/${digits.slice(2, 4)}/${digits.slice(4)}`;
}
function findNextEnabledDate(start, step2, minDate, maxDate) {
  let current = start;
  for (let i = 0; i < 366; i += 1) {
    current = addDays(current, step2);
    if (!isDisabled(current, minDate, maxDate)) {
      return current;
    }
  }
  return null;
}
function getInitialFocusableDate(preferred, minDate, maxDate) {
  if (!isDisabled(preferred, minDate, maxDate)) {
    return preferred;
  }
  return findNextEnabledDate(preferred, 1, minDate, maxDate) ?? findNextEnabledDate(preferred, -1, minDate, maxDate);
}
function DatePicker(props) {
  const {
    size = "md",
    label: label2,
    placeholder: placeholder2,
    message: message2,
    messageType,
    disabled: disabled2 = false,
    minDate,
    maxDate,
    className
  } = props;
  const iconSize2 = ICON_SIZES$2[size];
  const mode = props.mode ?? "single";
  const isRange = mode === "range";
  const isSingleControlled = !isRange && props.value !== void 0;
  const [singleInternal, setSingleInternal] = useState(
    (!isRange ? props.defaultValue : null) ?? null
  );
  const singleValue = isSingleControlled ? props.value : singleInternal;
  const isRangeControlled = isRange && props.value !== void 0;
  const [rangeInternal, setRangeInternal] = useState(
    (isRange ? props.defaultValue : [null, null]) ?? [
      null,
      null
    ]
  );
  const rangeValue = isRangeControlled ? props.value : rangeInternal;
  const [open2, setOpen] = useState(false);
  const [viewMonth, setViewMonth] = useState(() => {
    const t = today();
    return { year: t.year, month: t.month, day: 1 };
  });
  const [rangeStep, setRangeStep] = useState("start");
  const [hoverDate, setHoverDate] = useState(null);
  const [focusedDay, setFocusedDay] = useState(null);
  const [popoverStyle, setPopoverStyle] = useState({});
  const [startText, setStartText] = useState("");
  const [endText, setEndText] = useState("");
  const wrapperRef = useRef(null);
  const popoverRef = useRef(null);
  const triggerRef = useRef(null);
  const rangeTriggerRef = useRef(null);
  const startInputRef = useRef(null);
  const endInputRef = useRef(null);
  const autoId = useId();
  const triggerId = `${autoId}-trigger`;
  const messageId = `${autoId}-message`;
  const startInputId = `${autoId}-start`;
  const hasMessage = !!message2 && !!messageType;
  const isError = messageType === "error";
  const MsgIcon = messageType ? messageIconMap$3[messageType] : null;
  const grid2 = useMemo(
    () => buildGrid(viewMonth.year, viewMonth.month),
    [viewMonth.year, viewMonth.month]
  );
  const triggerText = useMemo(() => {
    if (isRange) {
      const [s2, e] = rangeValue;
      if (s2 && e) return `${formatDate(s2)} → ${formatDate(e)}`;
      if (s2) return `${formatDate(s2)} → ...`;
      return null;
    }
    return singleValue ? formatDate(singleValue) : null;
  }, [isRange, rangeValue, singleValue]);
  const defaultPlaceholder = isRange ? "DD/MM/AAAA → DD/MM/AAAA" : "DD/MM/AAAA";
  const wrapperClasses = [s$r.wrapper, className ?? ""].filter(Boolean).join(" ");
  const triggerClasses = [
    s$r.trigger,
    s$r[size],
    isError ? s$r.error : "",
    disabled2 ? s$r.disabled : "",
    open2 ? s$r.open : ""
  ].filter(Boolean).join(" ");
  const rangeTriggerClasses = [
    s$r.rangeTrigger,
    s$r[size],
    isError ? s$r.error : "",
    disabled2 ? s$r.disabled : "",
    open2 ? s$r.open : ""
  ].filter(Boolean).join(" ");
  const updatePosition = useCallback(() => {
    const trigger2 = isRange ? rangeTriggerRef.current : triggerRef.current;
    if (!trigger2) return;
    const rect = trigger2.getBoundingClientRect();
    const popoverWidth = 320;
    let left2 = rect.left;
    if (left2 + popoverWidth > window.innerWidth - 8) {
      left2 = window.innerWidth - popoverWidth - 8;
    }
    if (left2 < 8) left2 = 8;
    setPopoverStyle({
      position: "fixed",
      top: rect.bottom + 4,
      left: left2
    });
  }, [isRange]);
  const selectSingle = useCallback(
    (d) => {
      if (!isSingleControlled) setSingleInternal(d);
      props.onChange?.(d);
      setOpen(false);
    },
    [isSingleControlled, props]
  );
  const selectRange = useCallback(
    (d) => {
      if (rangeStep === "start") {
        const next2 = [d, null];
        if (!isRangeControlled) setRangeInternal(next2);
        props.onChange?.(next2);
        setRangeStep("end");
        setStartText(formatDate(d));
        setEndText("");
        requestAnimationFrame(() => endInputRef.current?.focus());
      } else {
        let start = rangeValue[0];
        let end = d;
        if (start && compareDates(end, start) < 0) {
          [start, end] = [end, start];
        }
        const next2 = [start, end];
        if (!isRangeControlled) setRangeInternal(next2);
        props.onChange?.(next2);
        setRangeStep("start");
        setStartText(start ? formatDate(start) : "");
        setEndText(formatDate(end));
        setOpen(false);
        requestAnimationFrame(() => {
          startInputRef.current?.blur();
          endInputRef.current?.blur();
        });
      }
    },
    [rangeStep, rangeValue, isRangeControlled, props]
  );
  const handleDayClick = useCallback(
    (d) => {
      if (isDisabled(d, minDate, maxDate)) return;
      if (isRange) {
        selectRange(d);
      } else {
        selectSingle(d);
      }
    },
    [isRange, selectRange, selectSingle, minDate, maxDate]
  );
  const openPopover = useCallback(
    (focusField) => {
      if (disabled2) return;
      updatePosition();
      setOpen(true);
      setHoverDate(null);
      if (isRange) {
        setRangeStep(focusField === "end" ? "end" : "start");
        const nav2 = rangeValue[0] ?? today();
        setViewMonth({ year: nav2.year, month: nav2.month, day: 1 });
      } else {
        setStartText(singleValue ? formatDate(singleValue) : "");
        const nav2 = singleValue ?? today();
        setViewMonth({ year: nav2.year, month: nav2.month, day: 1 });
        setFocusedDay(getInitialFocusableDate(nav2, minDate, maxDate));
        requestAnimationFrame(() => startInputRef.current?.focus());
      }
    },
    [disabled2, updatePosition, isRange, rangeValue, singleValue, minDate, maxDate]
  );
  const closePopover = useCallback(() => {
    setOpen(false);
    setFocusedDay(null);
    if (isRange) {
      startInputRef.current?.focus();
    } else {
      triggerRef.current?.focus();
    }
  }, [isRange]);
  const goPrevMonth = useCallback(() => {
    setViewMonth((v) => prevMonth(v));
  }, []);
  const goNextMonth = useCallback(() => {
    setViewMonth((v) => nextMonth(v));
  }, []);
  const handleToday = useCallback(() => {
    const t = today();
    setViewMonth({ year: t.year, month: t.month, day: 1 });
    if (!isDisabled(t, minDate, maxDate)) {
      handleDayClick(t);
    }
  }, [handleDayClick, minDate, maxDate]);
  const handleStartTextChange = (e) => {
    const masked = applyDateMask(e.target.value);
    setStartText(masked);
    if (masked.length === 10) {
      const parsed = parseDate(masked);
      if (parsed && !isDisabled(parsed, minDate, maxDate)) {
        setViewMonth({ year: parsed.year, month: parsed.month, day: 1 });
        if (!isRange) {
          if (!isSingleControlled) setSingleInternal(parsed);
          props.onChange?.(parsed);
        } else {
          const next2 = [
            parsed,
            rangeValue[1]
          ];
          if (!isRangeControlled) setRangeInternal(next2);
          props.onChange?.(next2);
        }
      }
    }
  };
  const handleEndTextChange = (e) => {
    const masked = applyDateMask(e.target.value);
    setEndText(masked);
    if (masked.length === 10) {
      const parsed = parseDate(masked);
      if (parsed && !isDisabled(parsed, minDate, maxDate)) {
        setViewMonth({ year: parsed.year, month: parsed.month, day: 1 });
        let start = rangeValue[0];
        let end = parsed;
        if (start && compareDates(end, start) < 0) {
          [start, end] = [end, start];
        }
        const next2 = [start, end];
        if (!isRangeControlled) setRangeInternal(next2);
        props.onChange?.(next2);
      }
    }
  };
  const handleInputKeyDown = (e) => {
    if (e.key === "Enter") {
      e.preventDefault();
      closePopover();
    } else if (e.key === "Escape") {
      e.preventDefault();
      closePopover();
    }
  };
  const handleRangeStartFocus = () => {
    if (!open2) openPopover("start");
    else setRangeStep("start");
  };
  const handleRangeEndFocus = () => {
    if (!open2) openPopover("end");
    else setRangeStep("end");
  };
  const handleStartBlur = () => {
    if (isRange) {
      setStartText(rangeValue[0] ? formatDate(rangeValue[0]) : "");
    } else {
      setStartText(singleValue ? formatDate(singleValue) : "");
    }
  };
  const handleEndBlur = () => {
    setEndText(rangeValue[1] ? formatDate(rangeValue[1]) : "");
  };
  const handleGridKeyDown = (e) => {
    if (!focusedDay) return;
    let next2 = null;
    switch (e.key) {
      case "ArrowRight":
        e.preventDefault();
        next2 = addDays(focusedDay, 1);
        break;
      case "ArrowLeft":
        e.preventDefault();
        next2 = addDays(focusedDay, -1);
        break;
      case "ArrowDown":
        e.preventDefault();
        next2 = addDays(focusedDay, 7);
        break;
      case "ArrowUp":
        e.preventDefault();
        next2 = addDays(focusedDay, -7);
        break;
      case "PageDown":
        e.preventDefault();
        next2 = {
          ...nextMonth(focusedDay),
          day: Math.min(
            focusedDay.day,
            daysInMonth(
              nextMonth(focusedDay).year,
              nextMonth(focusedDay).month
            )
          )
        };
        break;
      case "PageUp":
        e.preventDefault();
        next2 = {
          ...prevMonth(focusedDay),
          day: Math.min(
            focusedDay.day,
            daysInMonth(
              prevMonth(focusedDay).year,
              prevMonth(focusedDay).month
            )
          )
        };
        break;
      case "Enter":
      case " ":
        e.preventDefault();
        handleDayClick(focusedDay);
        return;
      case "Escape":
        e.preventDefault();
        closePopover();
        return;
    }
    if (next2) {
      const target = isDisabled(next2, minDate, maxDate) ? findNextEnabledDate(next2, e.key === "ArrowLeft" || e.key === "ArrowUp" || e.key === "PageUp" ? -1 : 1, minDate, maxDate) : next2;
      if (!target) return;
      setFocusedDay(target);
      if (target.year !== viewMonth.year || target.month !== viewMonth.month) {
        setViewMonth({ year: target.year, month: target.month, day: 1 });
      }
    }
  };
  const handleTriggerKeyDown = (e) => {
    if (disabled2) return;
    if (e.key === "Enter" || e.key === " " || e.key === "ArrowDown") {
      e.preventDefault();
      if (!open2) openPopover();
    } else if (e.key === "Escape" && open2) {
      e.preventDefault();
      closePopover();
    }
  };
  useEffect(() => {
    if (!isRange) return;
    if (document.activeElement !== startInputRef.current) {
      setStartText(rangeValue[0] ? formatDate(rangeValue[0]) : "");
    }
    if (document.activeElement !== endInputRef.current) {
      setEndText(rangeValue[1] ? formatDate(rangeValue[1]) : "");
    }
  }, [isRange, rangeValue]);
  useDocumentClickOutside({
    active: open2,
    refs: [wrapperRef, popoverRef],
    onOutside: () => {
      setOpen(false);
      setFocusedDay(null);
    }
  });
  useDocumentEscape(open2, closePopover);
  useViewportReposition(open2, updatePosition);
  useEffect(() => {
    if (!open2 || !focusedDay) return;
    const key = `${focusedDay.year}-${focusedDay.month}-${focusedDay.day}`;
    const el = popoverRef.current?.querySelector(
      `[data-date="${key}"]`
    );
    el?.focus();
  }, [open2, focusedDay]);
  const getDayClasses = (cell2) => {
    const d = cell2.date;
    const classes = [s$r.day];
    if (!cell2.isCurrentMonth) classes.push(s$r.dayOtherMonth);
    if (isDisabled(d, minDate, maxDate)) classes.push(s$r.dayDisabled);
    if (isToday(d)) classes.push(s$r.dayToday);
    if (isRange) {
      const [rs, re] = rangeValue;
      const previewEnd = rangeStep === "end" && rs && hoverDate ? hoverDate : null;
      const effectiveStart = rs;
      const effectiveEnd = re ?? previewEnd;
      if (effectiveStart && isSameDay(d, effectiveStart)) {
        if (effectiveEnd && !isSameDay(effectiveStart, effectiveEnd)) {
          classes.push(s$r.dayRangeStart);
        } else if (effectiveEnd && isSameDay(effectiveStart, effectiveEnd)) {
          classes.push(s$r.daySelected);
        } else {
          classes.push(s$r.daySelected);
        }
      } else if (effectiveEnd && isSameDay(d, effectiveEnd)) {
        classes.push(s$r.dayRangeEnd);
      } else if (effectiveStart && effectiveEnd) {
        const lo = compareDates(effectiveStart, effectiveEnd) <= 0 ? effectiveStart : effectiveEnd;
        const hi = compareDates(effectiveStart, effectiveEnd) <= 0 ? effectiveEnd : effectiveStart;
        if (isInRange(d, lo, hi)) {
          classes.push(s$r.dayInRange);
        }
      }
    } else {
      if (singleValue && isSameDay(d, singleValue)) {
        classes.push(s$r.daySelected);
      }
    }
    return classes.filter(Boolean).join(" ");
  };
  return /* @__PURE__ */ jsxs("div", { className: wrapperClasses, ref: wrapperRef, children: [
    label2 && /* @__PURE__ */ jsx("label", { className: s$r.label, htmlFor: isRange ? startInputId : triggerId, children: label2 }),
    /* @__PURE__ */ jsx("div", { className: s$r.anchor, children: isRange ? (
      /* ——— Range trigger: inline inputs ——— */
      /* @__PURE__ */ jsxs(
        "div",
        {
          ref: rangeTriggerRef,
          className: rangeTriggerClasses,
          onClick: () => {
            if (disabled2) return;
            if (!open2) openPopover("start");
          },
          role: "group",
          "aria-expanded": open2,
          "aria-haspopup": "dialog",
          "aria-describedby": hasMessage ? messageId : void 0,
          "aria-invalid": isError || void 0,
          children: [
            /* @__PURE__ */ jsx(CalendarBlank, { size: iconSize2 }),
            /* @__PURE__ */ jsx(
              "input",
              {
                ref: startInputRef,
                id: startInputId,
                type: "text",
                className: s$r.rangeTriggerInput,
                placeholder: "DD/MM/AAAA",
                value: startText,
                onChange: handleStartTextChange,
                onKeyDown: handleInputKeyDown,
                onFocus: handleRangeStartFocus,
                onBlur: handleStartBlur,
                disabled: disabled2,
                "aria-label": "Data inicial"
              }
            ),
            /* @__PURE__ */ jsx("span", { className: s$r.rangeSeparator, children: "→" }),
            /* @__PURE__ */ jsx(
              "input",
              {
                ref: endInputRef,
                type: "text",
                className: s$r.rangeTriggerInput,
                placeholder: "DD/MM/AAAA",
                value: endText,
                onChange: handleEndTextChange,
                onKeyDown: handleInputKeyDown,
                onFocus: handleRangeEndFocus,
                onBlur: handleEndBlur,
                disabled: disabled2,
                "aria-label": "Data final"
              }
            ),
            /* @__PURE__ */ jsx(
              CaretDown,
              {
                size: iconSize2,
                className: `${s$r.caret} ${open2 ? s$r.caretOpen : ""}`
              }
            )
          ]
        }
      )
    ) : (
      /* ——— Single trigger: button ——— */
      /* @__PURE__ */ jsxs(
        "button",
        {
          ref: triggerRef,
          id: triggerId,
          type: "button",
          className: triggerClasses,
          onClick: () => {
            if (disabled2) return;
            if (open2) closePopover();
            else openPopover();
          },
          onKeyDown: handleTriggerKeyDown,
          disabled: disabled2,
          "aria-haspopup": "dialog",
          "aria-expanded": open2,
          "aria-describedby": hasMessage ? messageId : void 0,
          "aria-invalid": isError || void 0,
          children: [
            /* @__PURE__ */ jsx(CalendarBlank, { size: iconSize2 }),
            /* @__PURE__ */ jsx("span", { className: triggerText ? s$r.value : s$r.placeholder, children: triggerText ?? placeholder2 ?? defaultPlaceholder }),
            /* @__PURE__ */ jsx(
              CaretDown,
              {
                size: iconSize2,
                className: `${s$r.caret} ${open2 ? s$r.caretOpen : ""}`
              }
            )
          ]
        }
      )
    ) }),
    open2 && createPortal(
      /* @__PURE__ */ jsxs(
        "div",
        {
          ref: popoverRef,
          className: s$r.popover,
          style: popoverStyle,
          role: "dialog",
          "aria-label": "Seletor de data",
          onMouseDown: (e) => e.preventDefault(),
          children: [
            !isRange && /* @__PURE__ */ jsx("div", { className: s$r.popoverHeader, children: /* @__PURE__ */ jsx(
              "input",
              {
                ref: startInputRef,
                type: "text",
                className: s$r.headerInput,
                placeholder: "DD/MM/AAAA",
                value: startText,
                onChange: handleStartTextChange,
                onKeyDown: handleInputKeyDown,
                onBlur: handleStartBlur,
                "aria-label": "Data"
              }
            ) }),
            /* @__PURE__ */ jsxs("div", { className: s$r.calendarNav, children: [
              /* @__PURE__ */ jsx(
                Button,
                {
                  variant: "tertiary",
                  size: "sm",
                  leftIcon: CaretLeft,
                  onClick: goPrevMonth,
                  "aria-label": "Mês anterior"
                }
              ),
              /* @__PURE__ */ jsxs("span", { className: s$r.monthYear, "aria-live": "polite", children: [
                MONTH_LABELS[viewMonth.month - 1],
                " ",
                viewMonth.year
              ] }),
              /* @__PURE__ */ jsx(
                Button,
                {
                  variant: "tertiary",
                  size: "sm",
                  leftIcon: CaretRight,
                  onClick: goNextMonth,
                  "aria-label": "Próximo mês"
                }
              )
            ] }),
            /* @__PURE__ */ jsx("div", { className: s$r.weekdays, children: WEEKDAY_LABELS.map((wd) => /* @__PURE__ */ jsx("div", { className: s$r.weekday, children: wd }, wd)) }),
            /* @__PURE__ */ jsx(
              "div",
              {
                className: s$r.grid,
                role: "grid",
                onKeyDown: handleGridKeyDown,
                children: grid2.map((cell2, i) => {
                  const d = cell2.date;
                  const key = `${d.year}-${d.month}-${d.day}`;
                  const dayDisabled2 = isDisabled(d, minDate, maxDate);
                  const isFocused = focusedDay !== null && isSameDay(d, focusedDay);
                  return /* @__PURE__ */ jsx(
                    "button",
                    {
                      type: "button",
                      role: "gridcell",
                      "data-date": key,
                      className: getDayClasses(cell2),
                      tabIndex: isFocused && !dayDisabled2 ? 0 : -1,
                      disabled: dayDisabled2,
                      onClick: () => handleDayClick(d),
                      onMouseEnter: () => {
                        if (!dayDisabled2) setHoverDate(d);
                      },
                      onMouseLeave: () => setHoverDate(null),
                      onFocus: () => {
                        if (!dayDisabled2) setFocusedDay(d);
                      },
                      "aria-label": `${d.day} de ${MONTH_LABELS[d.month - 1]} de ${d.year}`,
                      "aria-disabled": dayDisabled2 || void 0,
                      "aria-selected": !isRange && !!singleValue && isSameDay(d, singleValue) || isRange && rangeValue[0] !== null && isSameDay(d, rangeValue[0]) || isRange && rangeValue[1] !== null && isSameDay(d, rangeValue[1]) || void 0,
                      children: d.day
                    },
                    i
                  );
                })
              }
            ),
            /* @__PURE__ */ jsx("div", { className: s$r.popoverFooter, children: /* @__PURE__ */ jsx(
              Button,
              {
                variant: "tertiary",
                size: "sm",
                className: s$r.todayBtn,
                onClick: handleToday,
                children: "Hoje"
              }
            ) })
          ]
        }
      ),
      document.body
    ),
    hasMessage && /* @__PURE__ */ jsxs("div", { id: messageId, className: `${s$r.message} ${s$r[messageType]}`, children: [
      MsgIcon && /* @__PURE__ */ jsx(MsgIcon, { size: 14 }),
      /* @__PURE__ */ jsx("span", { children: message2 })
    ] })
  ] });
}
function addDays(d, n) {
  const date = new Date(d.year, d.month - 1, d.day + n);
  return {
    year: date.getFullYear(),
    month: date.getMonth() + 1,
    day: date.getDate()
  };
}
const trigger$1 = "_trigger_14a29_3";
const sm$7 = "_sm_14a29_19";
const md$5 = "_md_14a29_26";
const lg$4 = "_lg_14a29_33";
const hasLeftIcon = "_hasLeftIcon_14a29_41";
const primary = "_primary_14a29_48";
const open$1 = "_open_14a29_55";
const secondary = "_secondary_14a29_61";
const tertiary = "_tertiary_14a29_74";
const caret$1 = "_caret_14a29_105";
const caretOpen$1 = "_caretOpen_14a29_112";
const label$9 = "_label_14a29_118";
const menu = "_menu_14a29_125";
const menuIn = "_menuIn_14a29_1";
const searchWrapper = "_searchWrapper_14a29_149";
const searchIcon$1 = "_searchIcon_14a29_157";
const searchInput$2 = "_searchInput_14a29_162";
const list$3 = "_list_14a29_179";
const item$5 = "_item_14a29_187";
const focused$4 = "_focused_14a29_197";
const itemIcon$1 = "_itemIcon_14a29_201";
const itemText = "_itemText_14a29_206";
const itemLabel$1 = "_itemLabel_14a29_214";
const itemDescription = "_itemDescription_14a29_224";
const empty$4 = "_empty_14a29_235";
const s$q = {
  trigger: trigger$1,
  sm: sm$7,
  md: md$5,
  lg: lg$4,
  hasLeftIcon,
  primary,
  open: open$1,
  secondary,
  tertiary,
  caret: caret$1,
  caretOpen: caretOpen$1,
  label: label$9,
  menu,
  menuIn,
  searchWrapper,
  searchIcon: searchIcon$1,
  searchInput: searchInput$2,
  list: list$3,
  item: item$5,
  focused: focused$4,
  itemIcon: itemIcon$1,
  itemText,
  itemLabel: itemLabel$1,
  itemDescription,
  empty: empty$4
};
const iconSize = {
  sm: 14,
  md: 16,
  lg: 20
};
function DropdownButton({
  items: items2,
  onSelect,
  leftIcon: LeftIcon,
  variant = "secondary",
  size = "md",
  searchable = false,
  searchPlaceholder = "Buscar...",
  children,
  disabled: disabled2 = false,
  className
}) {
  const [open2, setOpen] = useState(false);
  const [search, setSearch] = useState("");
  const [focusedIndex, setFocusedIndex] = useState(-1);
  const triggerRef = useRef(null);
  const menuRef = useRef(null);
  const searchRef = useRef(null);
  const listRef = useRef(null);
  const triggerId = useId();
  const listId = `${triggerId}-list`;
  const iSize = iconSize[size];
  const filtered = searchable ? items2.filter(
    (item2) => item2.label.toLowerCase().includes(search.toLowerCase())
  ) : items2;
  const applyPosition = useCallback(() => {
    const trigger2 = triggerRef.current;
    const menu2 = menuRef.current;
    if (!trigger2 || !menu2) return;
    const tr = trigger2.getBoundingClientRect();
    const gap = 4;
    const margin = 8;
    menu2.style.position = "fixed";
    menu2.style.minWidth = `${tr.width}px`;
    const mr = menu2.getBoundingClientRect();
    const { top: top2, left: left2 } = resolveAnchoredOverlayPosition({
      anchorTop: tr.top,
      anchorBottom: tr.bottom,
      anchorLeft: tr.left,
      anchorRight: tr.right,
      overlayWidth: Math.max(mr.width, tr.width),
      overlayHeight: mr.height,
      viewportWidth: window.innerWidth,
      viewportHeight: window.innerHeight,
      gap,
      margin,
      horizontalAlign: "start",
      preferredVertical: "bottom"
    });
    menu2.style.left = `${left2}px`;
    menu2.style.top = `${top2}px`;
    menu2.style.bottom = "auto";
  }, []);
  useInitialReposition(open2, applyPosition);
  useViewportReposition(open2, applyPosition);
  const openMenu = useCallback(() => {
    setOpen(true);
    setSearch("");
    setFocusedIndex(-1);
  }, []);
  const closeMenu = useCallback(() => {
    setOpen(false);
    setFocusedIndex(-1);
  }, []);
  const handleSelect = useCallback(
    (item2) => {
      onSelect(item2);
      closeMenu();
    },
    [onSelect, closeMenu]
  );
  useDocumentEscape(open2, closeMenu);
  useOpenFocus({
    active: open2,
    containerRef: menuRef,
    initialFocusRef: searchable ? searchRef : void 0
  });
  useDocumentClickOutside({
    active: open2,
    refs: [triggerRef, menuRef],
    onOutside: closeMenu
  });
  const handleKeyDown = useCallback(
    (e) => {
      if (!open2) {
        if (e.key === "ArrowDown" || e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          openMenu();
        }
        return;
      }
      switch (e.key) {
        case "ArrowDown":
          e.preventDefault();
          setFocusedIndex((i) => i < filtered.length - 1 ? i + 1 : 0);
          break;
        case "ArrowUp":
          e.preventDefault();
          setFocusedIndex((i) => i > 0 ? i - 1 : filtered.length - 1);
          break;
        case "Home":
          e.preventDefault();
          setFocusedIndex(filtered.length > 0 ? 0 : -1);
          break;
        case "End":
          e.preventDefault();
          setFocusedIndex(filtered.length > 0 ? filtered.length - 1 : -1);
          break;
        case "Enter":
          e.preventDefault();
          if (focusedIndex >= 0 && focusedIndex < filtered.length) {
            handleSelect(filtered[focusedIndex]);
          }
          break;
        case "Escape":
          e.preventDefault();
          closeMenu();
          break;
        case "Tab":
          closeMenu();
          break;
      }
    },
    [open2, openMenu, closeMenu, handleSelect, filtered, focusedIndex]
  );
  useEffect(() => {
    if (!open2) return;
    setFocusedIndex(filtered.length > 0 ? 0 : -1);
  }, [open2, filtered.length]);
  useEffect(() => {
    if (focusedIndex < 0 || !listRef.current) return;
    const items22 = listRef.current.children;
    if (items22[focusedIndex]) {
      items22[focusedIndex].scrollIntoView({ block: "nearest" });
    }
  }, [focusedIndex]);
  const triggerClasses = [
    s$q.trigger,
    s$q[variant],
    s$q[size],
    LeftIcon ? s$q.hasLeftIcon : "",
    open2 ? s$q.open : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs(Fragment, { children: [
    /* @__PURE__ */ jsxs(
      "button",
      {
        ref: triggerRef,
        type: "button",
        id: triggerId,
        className: triggerClasses,
        disabled: disabled2,
        onClick: () => open2 ? closeMenu() : openMenu(),
        onKeyDown: handleKeyDown,
        "aria-haspopup": "menu",
        "aria-expanded": open2,
        "aria-controls": open2 ? listId : void 0,
        "aria-activedescendant": open2 && focusedIndex >= 0 ? `${triggerId}-opt-${focusedIndex}` : void 0,
        children: [
          LeftIcon && /* @__PURE__ */ jsx(LeftIcon, { size: iSize }),
          /* @__PURE__ */ jsx("span", { className: s$q.label, children }),
          /* @__PURE__ */ jsx(CaretDown, { size: iSize, className: `${s$q.caret} ${open2 ? s$q.caretOpen : ""}` })
        ]
      }
    ),
    open2 && createPortal(
      /* @__PURE__ */ jsxs(
        "div",
        {
          ref: menuRef,
          className: s$q.menu,
          role: "menu",
          "aria-label": "Opções",
          "aria-labelledby": triggerId,
          onKeyDown: handleKeyDown,
          children: [
            searchable && /* @__PURE__ */ jsxs("div", { className: s$q.searchWrapper, children: [
              /* @__PURE__ */ jsx(MagnifyingGlass, { size: 14, className: s$q.searchIcon }),
              /* @__PURE__ */ jsx(
                "input",
                {
                  ref: searchRef,
                  className: s$q.searchInput,
                  type: "text",
                  role: "combobox",
                  "aria-expanded": "true",
                  "aria-haspopup": "menu",
                  placeholder: searchPlaceholder,
                  value: search,
                  "aria-label": "Buscar itens",
                  "aria-controls": listId,
                  "aria-activedescendant": focusedIndex >= 0 ? `${triggerId}-opt-${focusedIndex}` : void 0,
                  onChange: (e) => {
                    setSearch(e.target.value);
                    setFocusedIndex(0);
                  },
                  onKeyDown: handleKeyDown
                }
              )
            ] }),
            /* @__PURE__ */ jsxs("ul", { ref: listRef, id: listId, className: s$q.list, role: "presentation", children: [
              filtered.map((item2, i) => /* @__PURE__ */ jsxs(
                "li",
                {
                  id: `${triggerId}-opt-${i}`,
                  className: `${s$q.item} ${i === focusedIndex ? s$q.focused : ""}`,
                  role: "menuitem",
                  onMouseEnter: () => setFocusedIndex(i),
                  onClick: () => handleSelect(item2),
                  children: [
                    item2.icon && /* @__PURE__ */ jsx(item2.icon, { size: 16, className: s$q.itemIcon }),
                    /* @__PURE__ */ jsxs("div", { className: s$q.itemText, children: [
                      /* @__PURE__ */ jsx("span", { className: s$q.itemLabel, children: item2.label }),
                      item2.description && /* @__PURE__ */ jsx("span", { className: s$q.itemDescription, children: item2.description })
                    ] })
                  ]
                },
                item2.id
              )),
              filtered.length === 0 && /* @__PURE__ */ jsx("li", { className: s$q.empty, role: "presentation", children: "Nenhum resultado encontrado" })
            ] })
          ]
        }
      ),
      document.body
    )
  ] });
}
const bar = "_bar_y0fo7_3";
const filters = "_filters_y0fo7_13";
const actions$1 = "_actions_y0fo7_23";
const addTrigger = "_addTrigger_y0fo7_32";
const addOpen = "_addOpen_y0fo7_58";
const actionBtn$1 = "_actionBtn_y0fo7_71";
const actionLabel = "_actionLabel_y0fo7_101";
const overlay$5 = "_overlay_y0fo7_107";
const sheetHandle$1 = "_sheetHandle_y0fo7_113";
const addPopover = "_addPopover_y0fo7_119";
const addSearchBox = "_addSearchBox_y0fo7_143";
const addSearchIcon = "_addSearchIcon_y0fo7_151";
const addSearchInput = "_addSearchInput_y0fo7_156";
const addList = "_addList_y0fo7_173";
const addItem = "_addItem_y0fo7_181";
const addItemFocused = "_addItemFocused_y0fo7_201";
const addItemIcon = "_addItemIcon_y0fo7_205";
const addEmpty = "_addEmpty_y0fo7_210";
const chip = "_chip_y0fo7_221";
const chipActive = "_chipActive_y0fo7_245";
const chipIcon = "_chipIcon_y0fo7_255";
const chipLabel = "_chipLabel_y0fo7_260";
const chipStatic = "_chipStatic_y0fo7_266";
const chipRemove = "_chipRemove_y0fo7_270";
const separator = "_separator_y0fo7_299";
const dropdown$1 = "_dropdown_y0fo7_308";
const s$p = {
  bar,
  filters,
  actions: actions$1,
  addTrigger,
  addOpen,
  actionBtn: actionBtn$1,
  actionLabel,
  overlay: overlay$5,
  sheetHandle: sheetHandle$1,
  addPopover,
  addSearchBox,
  addSearchIcon,
  addSearchInput,
  addList,
  addItem,
  addItemFocused,
  addItemIcon,
  addEmpty,
  chip,
  chipActive,
  chipIcon,
  chipLabel,
  chipStatic,
  chipRemove,
  separator,
  dropdown: dropdown$1
};
function FilterChip({
  label: label2,
  icon: Icon,
  onClick,
  onRemove,
  active: active2 = false,
  className
}) {
  const chipClasses = [
    s$p.chip,
    active2 ? s$p.chipActive : "",
    onClick ? "" : s$p.chipStatic,
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs(
    "div",
    {
      className: chipClasses,
      onClick,
      onKeyDown: onClick ? (e) => {
        if (e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          onClick();
        }
      } : void 0,
      tabIndex: onClick ? 0 : void 0,
      role: onClick ? "button" : void 0,
      children: [
        Icon && /* @__PURE__ */ jsx(Icon, { size: 14, className: s$p.chipIcon }),
        /* @__PURE__ */ jsx("span", { className: s$p.chipLabel, children: label2 }),
        onRemove && /* @__PURE__ */ jsx(
          "button",
          {
            type: "button",
            className: s$p.chipRemove,
            onClick: (e) => {
              e.stopPropagation();
              onRemove();
            },
            "aria-label": `Remover filtro ${label2}`,
            children: /* @__PURE__ */ jsx(X, { size: 10 })
          }
        )
      ]
    }
  );
}
function FilterDropdown({
  open: open2,
  onClose,
  anchorRef,
  children,
  className,
  placement = "bottom-start",
  ignoreRefs,
  noOverlay
}) {
  const dropdownRef = useRef(null);
  const outsideRefs = useMemo(
    () => [anchorRef, dropdownRef, ...ignoreRefs ?? []],
    [anchorRef, ignoreRefs]
  );
  useOpenFocus({ active: open2, containerRef: dropdownRef });
  const applyPosition = useCallback(() => {
    const anchor2 = anchorRef.current;
    const el = dropdownRef.current;
    if (!anchor2 || !el) return;
    if (window.innerWidth <= 480) return;
    const ar = anchor2.getBoundingClientRect();
    const gap = 4;
    const margin = 8;
    el.style.position = "fixed";
    const dr = el.getBoundingClientRect();
    if (placement === "right-start") {
      const { left: left2, top: top2 } = resolveSideStartOverlayPosition({
        anchorTop: ar.top,
        anchorLeft: ar.left,
        anchorRight: ar.right,
        overlayWidth: dr.width,
        overlayHeight: dr.height,
        viewportWidth: window.innerWidth,
        viewportHeight: window.innerHeight,
        gap,
        margin,
        preferredSide: "right"
      });
      el.style.left = `${left2}px`;
      el.style.top = `${top2}px`;
      el.style.bottom = "auto";
    } else {
      const { top: top2, left: left2 } = resolveAnchoredOverlayPosition({
        anchorTop: ar.top,
        anchorBottom: ar.bottom,
        anchorLeft: ar.left,
        anchorRight: ar.right,
        overlayWidth: dr.width,
        overlayHeight: dr.height,
        viewportWidth: window.innerWidth,
        viewportHeight: window.innerHeight,
        gap,
        margin,
        horizontalAlign: "start",
        preferredVertical: "bottom"
      });
      el.style.left = `${left2}px`;
      el.style.top = `${top2}px`;
      el.style.bottom = "auto";
    }
  }, [anchorRef, placement]);
  useInitialReposition(open2, applyPosition);
  useViewportReposition(open2, applyPosition);
  useDocumentClickOutside({
    active: open2,
    refs: outsideRefs,
    onOutside: onClose,
    relatedAnchorRef: anchorRef,
    relatedPortalSelectors: ["[data-filter-dropdown]"]
  });
  useDocumentEscape(open2, onClose);
  if (!open2) return null;
  const classes = [s$p.dropdown, className ?? ""].filter(Boolean).join(" ");
  return createPortal(
    /* @__PURE__ */ jsxs(Fragment, { children: [
      !noOverlay && /* @__PURE__ */ jsx("div", { className: s$p.overlay, role: "presentation", onClick: onClose }),
      /* @__PURE__ */ jsxs(
        "div",
        {
          ref: dropdownRef,
          className: classes,
          "data-filter-dropdown": true,
          onMouseDown: (e) => {
            const tag = e.target.tagName;
            if (tag !== "INPUT" && tag !== "TEXTAREA" && tag !== "SELECT") {
              e.preventDefault();
            }
          },
          children: [
            /* @__PURE__ */ jsx("div", { className: s$p.sheetHandle }),
            children
          ]
        }
      )
    ] }),
    document.body
  );
}
function FilterBar({
  filters: filters2,
  onAddFilter,
  onClearAll,
  onSaveView,
  saveViewLabel = "Salvar visualização",
  primaryAction,
  searchPlaceholder = "Buscar filtro...",
  defaultOpen = false,
  children,
  className
}) {
  const [open2, setOpen] = useState(defaultOpen);
  const [search, setSearch] = useState("");
  const [focusedIndex, setFocusedIndex] = useState(-1);
  const instanceId = useId();
  const listboxId = `${instanceId}-listbox`;
  const getOptionId = useCallback(
    (index) => `${instanceId}-option-${index}`,
    [instanceId]
  );
  const triggerRef = useRef(null);
  const popoverRef = useRef(null);
  const searchRef = useRef(null);
  const listRef = useRef(null);
  const hasChildren = children !== null && children !== void 0 && children !== false;
  const filtered = useMemo(() => {
    if (!search) return filters2;
    const q = search.toLowerCase();
    return filters2.filter((f) => f.label.toLowerCase().includes(q));
  }, [filters2, search]);
  const applyPosition = useCallback(() => {
    const trigger2 = triggerRef.current;
    const el = popoverRef.current;
    if (!trigger2 || !el) return;
    if (window.innerWidth <= 480) return;
    const tr = trigger2.getBoundingClientRect();
    const gap = 4;
    const margin = 8;
    el.style.position = "fixed";
    const pr = el.getBoundingClientRect();
    const { top: top2, left: left2 } = resolveAnchoredOverlayPosition({
      anchorTop: tr.top,
      anchorBottom: tr.bottom,
      anchorLeft: tr.left,
      anchorRight: tr.right,
      overlayWidth: pr.width,
      overlayHeight: pr.height,
      viewportWidth: window.innerWidth,
      viewportHeight: window.innerHeight,
      gap,
      margin,
      horizontalAlign: "start",
      preferredVertical: "bottom"
    });
    el.style.left = `${left2}px`;
    el.style.top = `${top2}px`;
    el.style.bottom = "auto";
  }, []);
  useInitialReposition(open2, applyPosition);
  useViewportReposition(open2, applyPosition);
  const openPopover = useCallback(() => {
    setOpen(true);
    setSearch("");
    setFocusedIndex(-1);
  }, []);
  const closePopover = useCallback(() => {
    setOpen(false);
    setSearch("");
    setFocusedIndex(-1);
  }, []);
  useDocumentEscape(open2, closePopover);
  const handleSelectFilter = useCallback(
    (filterId) => {
      onAddFilter(filterId);
      closePopover();
    },
    [onAddFilter, closePopover]
  );
  useOpenFocus({
    active: open2,
    containerRef: popoverRef,
    initialFocusRef: searchRef
  });
  useDocumentClickOutside({
    active: open2,
    refs: [triggerRef, popoverRef],
    onOutside: closePopover
  });
  const handleKeyDown = useCallback(
    (e) => {
      if (!open2) {
        if (e.key === "ArrowDown" || e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          openPopover();
        }
        return;
      }
      switch (e.key) {
        case "ArrowDown":
          e.preventDefault();
          setFocusedIndex((i) => i < filtered.length - 1 ? i + 1 : 0);
          break;
        case "ArrowUp":
          e.preventDefault();
          setFocusedIndex((i) => i > 0 ? i - 1 : filtered.length - 1);
          break;
        case "Home":
          e.preventDefault();
          setFocusedIndex(filtered.length > 0 ? 0 : -1);
          break;
        case "End":
          e.preventDefault();
          setFocusedIndex(filtered.length > 0 ? filtered.length - 1 : -1);
          break;
        case "Enter":
          e.preventDefault();
          if (focusedIndex >= 0 && focusedIndex < filtered.length) {
            handleSelectFilter(filtered[focusedIndex].id);
          }
          break;
        case "Escape":
          e.preventDefault();
          closePopover();
          break;
        case "Tab":
          closePopover();
          break;
      }
    },
    [open2, openPopover, closePopover, handleSelectFilter, filtered, focusedIndex]
  );
  useEffect(() => {
    if (open2) setFocusedIndex(filtered.length > 0 ? 0 : -1);
  }, [filtered, open2]);
  useEffect(() => {
    if (focusedIndex < 0 || !listRef.current) return;
    const el = listRef.current.querySelector(
      `#${CSS.escape(getOptionId(focusedIndex))}`
    );
    el?.scrollIntoView({ block: "nearest" });
  }, [focusedIndex, getOptionId]);
  const barClasses = [s$p.bar, className ?? ""].filter(Boolean).join(" ");
  const triggerClasses = [s$p.addTrigger, open2 ? s$p.addOpen : ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("div", { className: barClasses, children: [
    /* @__PURE__ */ jsxs("div", { className: s$p.filters, children: [
      /* @__PURE__ */ jsxs(
        "button",
        {
          ref: triggerRef,
          type: "button",
          className: triggerClasses,
          onClick: () => open2 ? closePopover() : openPopover(),
          onKeyDown: handleKeyDown,
          "aria-haspopup": "listbox",
          "aria-expanded": open2,
          "aria-controls": open2 ? listboxId : void 0,
          children: [
            /* @__PURE__ */ jsx(Plus, { size: 14 }),
            /* @__PURE__ */ jsx("span", { children: "Adicionar filtro" })
          ]
        }
      ),
      children,
      open2 && createPortal(
        /* @__PURE__ */ jsxs(Fragment, { children: [
          /* @__PURE__ */ jsx("div", { className: s$p.overlay, role: "presentation", onClick: closePopover }),
          /* @__PURE__ */ jsxs(
            "div",
            {
              ref: popoverRef,
              className: s$p.addPopover,
              onMouseDown: (e) => e.preventDefault(),
              children: [
                /* @__PURE__ */ jsx("div", { className: s$p.sheetHandle }),
                /* @__PURE__ */ jsxs("div", { className: s$p.addSearchBox, children: [
                  /* @__PURE__ */ jsx(MagnifyingGlass, { size: 14, className: s$p.addSearchIcon }),
                  /* @__PURE__ */ jsx(
                    "input",
                    {
                      ref: searchRef,
                      type: "text",
                      className: s$p.addSearchInput,
                      placeholder: searchPlaceholder,
                      value: search,
                      onChange: (e) => setSearch(e.target.value),
                      onKeyDown: handleKeyDown,
                      role: "combobox",
                      "aria-expanded": "true",
                      "aria-label": "Buscar filtro",
                      "aria-controls": listboxId,
                      "aria-activedescendant": focusedIndex >= 0 ? getOptionId(focusedIndex) : void 0
                    }
                  )
                ] }),
                /* @__PURE__ */ jsxs(
                  "div",
                  {
                    ref: listRef,
                    id: listboxId,
                    className: s$p.addList,
                    role: "listbox",
                    "aria-label": "Filtros disponíveis",
                    children: [
                      filtered.map((f, i) => /* @__PURE__ */ jsxs(
                        "div",
                        {
                          id: getOptionId(i),
                          className: `${s$p.addItem} ${i === focusedIndex ? s$p.addItemFocused : ""}`,
                          role: "option",
                          "aria-selected": i === focusedIndex,
                          onMouseEnter: () => setFocusedIndex(i),
                          onClick: () => handleSelectFilter(f.id),
                          children: [
                            f.icon && /* @__PURE__ */ jsx(f.icon, { size: 16, className: s$p.addItemIcon }),
                            /* @__PURE__ */ jsx("span", { children: f.label })
                          ]
                        },
                        f.id
                      )),
                      filtered.length === 0 && /* @__PURE__ */ jsx("div", { className: s$p.addEmpty, role: "presentation", children: "Nenhum filtro encontrado" })
                    ]
                  }
                )
              ]
            }
          )
        ] }),
        document.body
      )
    ] }),
    hasChildren && (onClearAll || onSaveView || primaryAction) && /* @__PURE__ */ jsxs("div", { className: s$p.actions, children: [
      onClearAll && /* @__PURE__ */ jsxs(
        "button",
        {
          type: "button",
          className: s$p.actionBtn,
          onClick: onClearAll,
          "aria-label": "Limpar filtros",
          children: [
            /* @__PURE__ */ jsx(Broom, { size: 14 }),
            /* @__PURE__ */ jsx("span", { className: s$p.actionLabel, children: "Limpar filtros" })
          ]
        }
      ),
      onSaveView && /* @__PURE__ */ jsxs(Fragment, { children: [
        onClearAll && /* @__PURE__ */ jsx("span", { className: s$p.separator }),
        /* @__PURE__ */ jsxs(
          "button",
          {
            type: "button",
            className: s$p.actionBtn,
            onClick: onSaveView,
            "aria-label": saveViewLabel,
            children: [
              /* @__PURE__ */ jsx(FloppyDisk, { size: 14 }),
              /* @__PURE__ */ jsx("span", { className: s$p.actionLabel, children: saveViewLabel })
            ]
          }
        )
      ] }),
      primaryAction
    ] })
  ] });
}
const root$7 = "_root_353ec_1";
const step = "_step_353ec_7";
const stepDimmed = "_stepDimmed_353ec_12";
const stepHovered = "_stepHovered_353ec_16";
const label$8 = "_label_353ec_20";
const valueText$2 = "_valueText_353ec_27";
const s$o = {
  root: root$7,
  step,
  stepDimmed,
  stepHovered,
  label: label$8,
  valueText: valueText$2
};
const CHART_COLORS = [
  "var(--color-chart-1)",
  "var(--color-chart-2)",
  "var(--color-chart-3)",
  "var(--color-chart-4)",
  "var(--color-chart-5)"
];
function getColor(index, override) {
  if (override) return override;
  return CHART_COLORS[index % CHART_COLORS.length];
}
const GAP$1 = 4;
const MIN_WIDTH_RATIO = 0.2;
const LABEL_AREA_WIDTH = 180;
const LABEL_GAP = 12;
function Funnel({
  data,
  height = 300,
  showValues = true,
  showPercentage = true,
  formatValue,
  className
}) {
  const [hoveredIndex, setHoveredIndex] = useState(null);
  if (!data.length) return null;
  const firstValue = data[0].value;
  const stepCount = data.length;
  const totalGap = GAP$1 * (stepCount - 1);
  const stepHeight = (height - totalGap) / stepCount;
  const trapezoidWidth = 400;
  const svgWidth = trapezoidWidth + LABEL_GAP + LABEL_AREA_WIDTH;
  function widthRatio(value2) {
    if (firstValue === 0) return MIN_WIDTH_RATIO;
    return Math.max(MIN_WIDTH_RATIO, value2 / firstValue);
  }
  function buildPath(index) {
    const currentRatio = widthRatio(data[index].value);
    const topW = index === 0 ? trapezoidWidth : currentRatio * trapezoidWidth;
    const bottomW = index === stepCount - 1 ? currentRatio * trapezoidWidth : Math.min(currentRatio, widthRatio(data[index + 1].value)) * trapezoidWidth;
    const y = index * (stepHeight + GAP$1);
    const topOffset = (trapezoidWidth - topW) / 2;
    const bottomOffset = (trapezoidWidth - bottomW) / 2;
    return [
      `M ${topOffset} ${y}`,
      `L ${topOffset + topW} ${y}`,
      `L ${bottomOffset + bottomW} ${y + stepHeight}`,
      `L ${bottomOffset} ${y + stepHeight}`,
      "Z"
    ].join(" ");
  }
  const ariaLabel = data.map((step2) => {
    const pct = firstValue > 0 ? Math.round(step2.value / firstValue * 100) : 0;
    return `${step2.label}: ${formatValue ? formatValue(step2.value) : step2.value} (${pct}%)`;
  }).join(", ");
  return /* @__PURE__ */ jsx("div", { className: [s$o.root, className].filter(Boolean).join(" "), children: /* @__PURE__ */ jsx(
    "svg",
    {
      role: "img",
      "aria-label": `Funnel: ${ariaLabel}`,
      viewBox: `0 0 ${svgWidth} ${height}`,
      width: "100%",
      height,
      preserveAspectRatio: "xMidYMid meet",
      children: data.map((step2, i) => {
        const y = i * (stepHeight + GAP$1);
        const pct = firstValue > 0 ? Math.round(step2.value / firstValue * 100) : 0;
        const isDimmed = hoveredIndex !== null && hoveredIndex !== i;
        const isHovered = hoveredIndex === i;
        const labelX = trapezoidWidth + LABEL_GAP;
        const labelY = y + stepHeight / 2;
        const formattedValue = formatValue ? formatValue(step2.value) : String(step2.value);
        const valueParts = [];
        if (showValues) valueParts.push(formattedValue);
        if (showPercentage) valueParts.push(`${pct}%`);
        const valueStr = valueParts.join(" · ");
        return /* @__PURE__ */ jsxs(
          "g",
          {
            className: `${s$o.step}${isDimmed ? ` ${s$o.stepDimmed}` : ""}${isHovered ? ` ${s$o.stepHovered}` : ""}`,
            onMouseEnter: () => setHoveredIndex(i),
            onMouseLeave: () => setHoveredIndex(null),
            children: [
              /* @__PURE__ */ jsx("path", { d: buildPath(i), fill: getColor(i, step2.color) }),
              /* @__PURE__ */ jsx(
                "text",
                {
                  className: s$o.label,
                  x: labelX,
                  y: labelY,
                  dy: "-0.15em",
                  dominantBaseline: "middle",
                  children: step2.label
                }
              ),
              (showValues || showPercentage) && /* @__PURE__ */ jsx(
                "text",
                {
                  className: s$o.valueText,
                  x: labelX,
                  y: labelY,
                  dy: "1.15em",
                  dominantBaseline: "middle",
                  children: valueStr
                }
              )
            ]
          },
          i
        );
      })
    }
  ) });
}
const wrapper$7 = "_wrapper_oyanq_6";
const label$7 = "_label_oyanq_13";
const trackContainer = "_trackContainer_oyanq_21";
const track$1 = "_track_oyanq_21";
const fill$1 = "_fill_oyanq_34";
const expectedMarker = "_expectedMarker_oyanq_41";
const safeZone = "_safeZone_oyanq_52";
const gaugeMarker = "_gaugeMarker_oyanq_60";
const thumb$1 = "_thumb_oyanq_74";
const interactive = "_interactive_oyanq_91";
const dragging$1 = "_dragging_oyanq_97";
const footer$6 = "_footer_oyanq_116";
const statusGroup = "_statusGroup_oyanq_123";
const statusDot = "_statusDot_oyanq_130";
const statusText = "_statusText_oyanq_137";
const valueText$1 = "_valueText_oyanq_143";
const onTrack = "_onTrack_oyanq_155";
const attention$3 = "_attention_oyanq_170";
const offTrack = "_offTrack_oyanq_185";
const s$n = {
  wrapper: wrapper$7,
  label: label$7,
  trackContainer,
  track: track$1,
  fill: fill$1,
  expectedMarker,
  safeZone,
  gaugeMarker,
  thumb: thumb$1,
  interactive,
  dragging: dragging$1,
  footer: footer$6,
  statusGroup,
  statusDot,
  statusText,
  valueText: valueText$1,
  onTrack,
  attention: attention$3,
  offTrack
};
const STATUS_LABELS = {
  "on-track": "Em dia",
  attention: "Atenção",
  "off-track": "Em risco"
};
const STATUS_CLASS = {
  "on-track": s$n.onTrack,
  attention: s$n.attention,
  "off-track": s$n.offTrack
};
function clamp$2(v, lo, hi) {
  return Math.min(hi, Math.max(lo, v));
}
function resolveProgressStatus(percent2, expected) {
  if (expected !== void 0) {
    if (percent2 >= expected) return "on-track";
    if (expected - percent2 <= 25) return "attention";
    return "off-track";
  }
  if (percent2 >= 70) return "on-track";
  if (percent2 >= 40) return "attention";
  return "off-track";
}
function GoalProgressBar({
  label: label2,
  value: value2,
  target = 100,
  min = 0,
  formattedValue,
  expected,
  status: statusOverride,
  onChange,
  className
}) {
  const range = target - min;
  const percent2 = range > 0 ? clamp$2((value2 - min) / range * 100, 0, 100) : 0;
  const status = statusOverride ?? resolveProgressStatus(percent2, expected);
  const display = formattedValue ?? String(value2);
  const interactive2 = !!onChange;
  const trackRef = useRef(null);
  const [dragging2, setDragging] = useState(false);
  function updateFromPointer(e) {
    if (!trackRef.current || !onChange) return;
    const rect = trackRef.current.getBoundingClientRect();
    const x = clamp$2(e.clientX - rect.left, 0, rect.width);
    const pct = x / rect.width;
    onChange(Math.round(min + pct * range));
  }
  function handlePointerDown(e) {
    if (!interactive2) return;
    e.preventDefault();
    setDragging(true);
    trackRef.current?.setPointerCapture(e.pointerId);
    updateFromPointer(e);
  }
  function handlePointerMove(e) {
    if (!dragging2) return;
    updateFromPointer(e);
  }
  function handlePointerUp() {
    setDragging(false);
  }
  return /* @__PURE__ */ jsxs("div", { className: `${s$n.wrapper} ${STATUS_CLASS[status]} ${className ?? ""}`, children: [
    /* @__PURE__ */ jsx("span", { className: s$n.label, children: label2 }),
    /* @__PURE__ */ jsxs(
      "div",
      {
        className: `${s$n.trackContainer} ${interactive2 ? s$n.interactive : ""} ${dragging2 ? s$n.dragging : ""}`,
        ref: trackRef,
        onPointerDown: interactive2 ? handlePointerDown : void 0,
        onPointerMove: interactive2 ? handlePointerMove : void 0,
        onPointerUp: interactive2 ? handlePointerUp : void 0,
        children: [
          /* @__PURE__ */ jsx(
            "div",
            {
              className: s$n.track,
              role: "progressbar",
              "aria-valuenow": value2,
              "aria-valuemin": min,
              "aria-valuemax": target,
              "aria-label": label2,
              children: /* @__PURE__ */ jsx("div", { className: s$n.fill, style: { width: `${percent2}%` } })
            }
          ),
          expected !== void 0 && /* @__PURE__ */ jsx(
            "div",
            {
              className: s$n.expectedMarker,
              style: { left: `${clamp$2(expected, 0, 100)}%` },
              "aria-hidden": "true"
            }
          ),
          interactive2 && /* @__PURE__ */ jsx(
            "div",
            {
              className: s$n.thumb,
              style: { left: `${percent2}%` },
              role: "slider",
              tabIndex: 0,
              "aria-valuemin": 0,
              "aria-valuemax": 100,
              "aria-valuenow": value2,
              "aria-label": "Ajustar progresso",
              onKeyDown: (e) => {
                if (!onChange) return;
                let next2 = value2;
                switch (e.key) {
                  case "ArrowRight":
                  case "ArrowUp":
                    next2 = clamp$2(value2 + 1, 0, 100);
                    break;
                  case "ArrowLeft":
                  case "ArrowDown":
                    next2 = clamp$2(value2 - 1, 0, 100);
                    break;
                  case "Home":
                    next2 = 0;
                    break;
                  case "End":
                    next2 = 100;
                    break;
                  case "PageUp":
                    next2 = clamp$2(value2 + 10, 0, 100);
                    break;
                  case "PageDown":
                    next2 = clamp$2(value2 - 10, 0, 100);
                    break;
                  default:
                    return;
                }
                e.preventDefault();
                onChange(next2);
              }
            }
          )
        ]
      }
    ),
    /* @__PURE__ */ jsxs("div", { className: s$n.footer, children: [
      /* @__PURE__ */ jsxs("span", { className: s$n.statusGroup, children: [
        /* @__PURE__ */ jsx("span", { className: s$n.statusDot }),
        /* @__PURE__ */ jsx("span", { className: s$n.statusText, children: STATUS_LABELS[status] })
      ] }),
      /* @__PURE__ */ jsx("span", { className: s$n.valueText, children: display })
    ] })
  ] });
}
const GAUGE_STATUS_LABELS = {
  "on-track": "Dentro do previsto",
  attention: "Atenção",
  "off-track": "Fora do previsto"
};
function resolveGaugeStatus(value2, goalType, low, high, min, max) {
  const totalRange = max - min;
  const margin = totalRange * 0.1;
  if (goalType === "between" && low !== void 0 && high !== void 0) {
    if (value2 >= low && value2 <= high) return "on-track";
    if (value2 >= low - margin && value2 <= high + margin) return "attention";
    return "off-track";
  }
  if (goalType === "above" && low !== void 0) {
    if (value2 >= low) return "on-track";
    if (value2 >= low - margin) return "attention";
    return "off-track";
  }
  if (goalType === "below" && high !== void 0) {
    if (value2 <= high) return "on-track";
    if (value2 <= high + margin) return "attention";
    return "off-track";
  }
  return "on-track";
}
function toPos(v, min, max) {
  const range = max - min;
  return range > 0 ? clamp$2((v - min) / range * 100, 0, 100) : 0;
}
function GoalGaugeBar({
  label: label2,
  value: value2,
  low,
  high,
  goalType,
  min = 0,
  max = 100,
  formattedValue,
  status: statusOverride,
  onChange,
  className
}) {
  const status = statusOverride ?? resolveGaugeStatus(value2, goalType, low, high, min, max);
  const display = formattedValue ?? String(value2);
  const valuePos = toPos(value2, min, max);
  const interactive2 = !!onChange;
  const trackRef = useRef(null);
  const [dragging2, setDragging] = useState(false);
  function updateFromPointer(e) {
    if (!trackRef.current || !onChange) return;
    const rect = trackRef.current.getBoundingClientRect();
    const x = clamp$2(e.clientX - rect.left, 0, rect.width);
    const pct = x / rect.width;
    onChange(Math.round(min + pct * (max - min)));
  }
  function handlePointerDown(e) {
    if (!interactive2) return;
    e.preventDefault();
    setDragging(true);
    trackRef.current?.setPointerCapture(e.pointerId);
    updateFromPointer(e);
  }
  function handlePointerMove(e) {
    if (!dragging2) return;
    updateFromPointer(e);
  }
  function handlePointerUp() {
    setDragging(false);
  }
  let safeLeft = 0;
  let safeWidth = 100;
  if (goalType === "between" && low !== void 0 && high !== void 0) {
    safeLeft = toPos(low, min, max);
    safeWidth = toPos(high, min, max) - safeLeft;
  } else if (goalType === "above" && low !== void 0) {
    safeLeft = toPos(low, min, max);
    safeWidth = 100 - safeLeft;
  } else if (goalType === "below" && high !== void 0) {
    safeLeft = 0;
    safeWidth = toPos(high, min, max);
  }
  const touchesLeft = safeLeft <= 0;
  const touchesRight = safeLeft + safeWidth >= 100;
  return /* @__PURE__ */ jsxs("div", { className: `${s$n.wrapper} ${STATUS_CLASS[status]} ${className ?? ""}`, children: [
    /* @__PURE__ */ jsx("span", { className: s$n.label, children: label2 }),
    /* @__PURE__ */ jsxs(
      "div",
      {
        className: `${s$n.trackContainer} ${interactive2 ? s$n.interactive : ""} ${dragging2 ? s$n.dragging : ""}`,
        ref: trackRef,
        onPointerDown: interactive2 ? handlePointerDown : void 0,
        onPointerMove: interactive2 ? handlePointerMove : void 0,
        onPointerUp: interactive2 ? handlePointerUp : void 0,
        children: [
          /* @__PURE__ */ jsx(
            "div",
            {
              className: s$n.track,
              role: "meter",
              "aria-valuenow": value2,
              "aria-valuemin": min,
              "aria-valuemax": max,
              "aria-label": label2,
              "aria-valuetext": `${display} — ${GAUGE_STATUS_LABELS[status]}`,
              children: /* @__PURE__ */ jsx(
                "div",
                {
                  className: s$n.safeZone,
                  style: {
                    left: `${safeLeft}%`,
                    width: `${safeWidth}%`,
                    borderRadius: `${touchesLeft ? "var(--radius-full)" : "0"} ${touchesRight ? "var(--radius-full)" : "0"} ${touchesRight ? "var(--radius-full)" : "0"} ${touchesLeft ? "var(--radius-full)" : "0"}`
                  }
                }
              )
            }
          ),
          /* @__PURE__ */ jsx(
            "div",
            {
              className: s$n.gaugeMarker,
              style: { left: `${valuePos}%` },
              "aria-hidden": "true"
            }
          ),
          interactive2 && /* @__PURE__ */ jsx(
            "div",
            {
              className: s$n.thumb,
              style: { left: `${valuePos}%` },
              role: "slider",
              tabIndex: 0,
              "aria-valuemin": 0,
              "aria-valuemax": 100,
              "aria-valuenow": value2,
              "aria-label": "Ajustar progresso",
              onKeyDown: (e) => {
                if (!onChange) return;
                let next2 = value2;
                switch (e.key) {
                  case "ArrowRight":
                  case "ArrowUp":
                    next2 = clamp$2(value2 + 1, 0, 100);
                    break;
                  case "ArrowLeft":
                  case "ArrowDown":
                    next2 = clamp$2(value2 - 1, 0, 100);
                    break;
                  case "Home":
                    next2 = 0;
                    break;
                  case "End":
                    next2 = 100;
                    break;
                  case "PageUp":
                    next2 = clamp$2(value2 + 10, 0, 100);
                    break;
                  case "PageDown":
                    next2 = clamp$2(value2 - 10, 0, 100);
                    break;
                  default:
                    return;
                }
                e.preventDefault();
                onChange(next2);
              }
            }
          )
        ]
      }
    ),
    /* @__PURE__ */ jsxs("div", { className: s$n.footer, children: [
      /* @__PURE__ */ jsxs("span", { className: s$n.statusGroup, children: [
        /* @__PURE__ */ jsx("span", { className: s$n.statusDot }),
        /* @__PURE__ */ jsx("span", { className: s$n.statusText, children: GAUGE_STATUS_LABELS[status] })
      ] }),
      /* @__PURE__ */ jsx("span", { className: s$n.valueText, children: display })
    ] })
  ] });
}
const root$6 = "_root_xw21k_5";
const headerRow = "_headerRow_xw21k_16";
const dataRow = "_dataRow_xw21k_17";
const cornerCell = "_cornerCell_xw21k_25";
const colLabel = "_colLabel_xw21k_33";
const rowLabel = "_rowLabel_xw21k_50";
const cell = "_cell_xw21k_67";
const cellDimmed = "_cellDimmed_xw21k_86";
const cellHighlight = "_cellHighlight_xw21k_90";
const cellActive = "_cellActive_xw21k_94";
const legend = "_legend_xw21k_102";
const legendCell = "_legendCell_xw21k_110";
const legendLabel = "_legendLabel_xw21k_116";
const tooltip$1 = "_tooltip_xw21k_126";
const tooltipRow = "_tooltipRow_xw21k_155";
const tooltipCol = "_tooltipCol_xw21k_162";
const tooltipValue = "_tooltipValue_xw21k_171";
const s$m = {
  root: root$6,
  headerRow,
  dataRow,
  cornerCell,
  colLabel,
  rowLabel,
  cell,
  cellDimmed,
  cellHighlight,
  cellActive,
  legend,
  legendCell,
  legendLabel,
  tooltip: tooltip$1,
  tooltipRow,
  tooltipCol,
  tooltipValue
};
function clamp$1(v, lo, hi) {
  return Math.min(hi, Math.max(lo, v));
}
const COLOR_MAP$1 = {
  orange: "--color-orange",
  green: "--color-green",
  red: "--color-red",
  yellow: "--color-yellow",
  wine: "--color-wine",
  neutral: "--color-neutral"
};
const STEPS = [50, 200, 400, 500, 700];
function Heatmap({
  data,
  rows,
  columns,
  min: minProp,
  max: maxProp,
  color = "orange",
  colorScale = "sequential",
  showValues = true,
  formatValue,
  cellSize = 40,
  labelWidth = 100,
  gap = 4,
  columnTooltips,
  className
}) {
  const [hoveredCell, setHoveredCell] = useState(null);
  const rootRef = useRef(null);
  const lookup = /* @__PURE__ */ new Map();
  for (const cell2 of data) {
    lookup.set(`${cell2.row}__${cell2.col}`, cell2.value);
  }
  const values = data.map((d) => d.value);
  const dataMin = minProp ?? Math.min(...values);
  const dataMax = maxProp ?? Math.max(...values);
  const range = dataMax - dataMin || 1;
  const prefix = COLOR_MAP$1[color];
  const isDivergent = colorScale === "divergent";
  const DIVERGENT_STOPS = [
    { at: 0, bg: "var(--color-red-200)", dark: false },
    { at: 0.25, bg: "var(--color-red-100)", dark: false },
    { at: 0.4, bg: "var(--color-yellow-100)", dark: false },
    { at: 0.5, bg: "var(--color-yellow-50)", dark: false },
    { at: 0.6, bg: "var(--color-yellow-100)", dark: false },
    { at: 0.75, bg: "var(--color-green-100)", dark: false },
    { at: 1, bg: "var(--color-green-300)", dark: false }
  ];
  function getDivergentStop(value2) {
    const normalized = clamp$1((value2 - dataMin) / range, 0, 1);
    let best = DIVERGENT_STOPS[0];
    let bestDist = Math.abs(normalized - best.at);
    for (const stop of DIVERGENT_STOPS) {
      const dist = Math.abs(normalized - stop.at);
      if (dist < bestDist) {
        best = stop;
        bestDist = dist;
      }
    }
    return best;
  }
  function getIntensity(value2) {
    const normalized = clamp$1((value2 - dataMin) / range, 0, 1);
    return Math.round(normalized * (STEPS.length - 1));
  }
  function getCellColor(value2) {
    if (isDivergent) return getDivergentStop(value2).bg;
    const step2 = STEPS[getIntensity(value2)];
    return `var(${prefix}-${step2})`;
  }
  function getTextColor(value2) {
    if (isDivergent) return getDivergentStop(value2).dark ? "var(--color-white)" : "var(--color-neutral-950)";
    const intensity = getIntensity(value2);
    return intensity >= 3 ? "var(--color-white)" : "var(--color-neutral-950)";
  }
  const format = formatValue ?? ((v) => String(v));
  const vars = {
    "--heatmap-cell": `${cellSize}px`,
    "--heatmap-label-width": `${labelWidth}px`,
    "--heatmap-gap": `${gap}px`
  };
  const [tooltip2, setTooltip] = useState(null);
  function handleCellEnter(row2, col, value2, e) {
    setHoveredCell({ row: row2, col });
    const cellRect = e.currentTarget.getBoundingClientRect();
    setTooltip({
      row: row2,
      col,
      value: value2,
      x: cellRect.left + cellRect.width / 2,
      y: cellRect.top
    });
  }
  function handleCellLeave() {
    setHoveredCell(null);
    setTooltip(null);
  }
  const colTooltipLabel = (col) => columnTooltips?.[col] ? `${col} — ${columnTooltips[col]}` : col;
  return /* @__PURE__ */ jsxs(Fragment, { children: [
    /* @__PURE__ */ jsxs(
      "div",
      {
        ref: rootRef,
        className: `${s$m.root} ${className ?? ""}`,
        style: vars,
        role: "table",
        "aria-label": "Mapa de calor",
        children: [
          /* @__PURE__ */ jsxs("div", { className: s$m.headerRow, role: "row", children: [
            /* @__PURE__ */ jsx("span", { className: s$m.cornerCell, role: "columnheader" }),
            columns.map((col) => /* @__PURE__ */ jsx(
              "span",
              {
                className: s$m.colLabel,
                role: "columnheader",
                title: columnTooltips?.[col],
                children: col
              },
              col
            ))
          ] }),
          rows.map((row2) => /* @__PURE__ */ jsxs("div", { className: s$m.dataRow, role: "row", children: [
            /* @__PURE__ */ jsx("span", { className: s$m.rowLabel, role: "rowheader", title: row2, children: row2 }),
            columns.map((col) => {
              const key = `${row2}__${col}`;
              const value2 = lookup.get(key);
              const hasValue = value2 !== void 0;
              const isExact = hoveredCell?.row === row2 && hoveredCell?.col === col;
              const isRowOrCol = hoveredCell?.row === row2 || hoveredCell?.col === col;
              const hoverClass = hoveredCell ? isExact ? s$m.cellActive : isRowOrCol ? s$m.cellHighlight : s$m.cellDimmed : "";
              return /* @__PURE__ */ jsx(
                "span",
                {
                  className: `${s$m.cell} ${hoverClass}`,
                  style: hasValue ? {
                    backgroundColor: getCellColor(value2),
                    color: getTextColor(value2)
                  } : void 0,
                  role: "cell",
                  "aria-label": hasValue ? `${row2}, ${col}: ${format(value2)}` : `${row2}, ${col}: sem dados`,
                  onMouseEnter: hasValue ? (e) => handleCellEnter(row2, col, value2, e) : void 0,
                  onMouseLeave: handleCellLeave,
                  children: hasValue && showValues ? format(value2) : ""
                },
                key
              );
            })
          ] }, row2)),
          /* @__PURE__ */ jsx("div", { className: s$m.legend, "aria-hidden": "true", children: isDivergent ? /* @__PURE__ */ jsxs(Fragment, { children: [
            /* @__PURE__ */ jsx("span", { className: s$m.legendLabel, children: "Crítico" }),
            DIVERGENT_STOPS.map((stop, i) => /* @__PURE__ */ jsx(
              "span",
              {
                className: s$m.legendCell,
                style: { backgroundColor: stop.bg }
              },
              i
            )),
            /* @__PURE__ */ jsx("span", { className: s$m.legendLabel, children: "Excelente" })
          ] }) : /* @__PURE__ */ jsxs(Fragment, { children: [
            /* @__PURE__ */ jsx("span", { className: s$m.legendLabel, children: "Menor" }),
            STEPS.map((step2) => /* @__PURE__ */ jsx(
              "span",
              {
                className: s$m.legendCell,
                style: { backgroundColor: `var(${prefix}-${step2})` }
              },
              step2
            )),
            /* @__PURE__ */ jsx("span", { className: s$m.legendLabel, children: "Maior" })
          ] }) })
        ]
      }
    ),
    tooltip2 && createPortal(
      /* @__PURE__ */ jsxs(
        "div",
        {
          className: s$m.tooltip,
          style: {
            left: `${tooltip2.x}px`,
            top: `${tooltip2.y}px`
          },
          "aria-hidden": "true",
          children: [
            /* @__PURE__ */ jsx("span", { className: s$m.tooltipRow, children: tooltip2.row }),
            /* @__PURE__ */ jsx("span", { className: s$m.tooltipCol, children: colTooltipLabel(tooltip2.col) }),
            /* @__PURE__ */ jsx("span", { className: s$m.tooltipValue, children: format(tooltip2.value) })
          ]
        }
      ),
      document.body
    )
  ] });
}
const panel$1 = "_panel_1ujao_3";
const header$5 = "_header_1ujao_28";
const headerTitle = "_headerTitle_1ujao_36";
const markAllBtn = "_markAllBtn_1ujao_44";
const list$2 = "_list_1ujao_67";
const item$4 = "_item_1ujao_74";
const itemUnread = "_itemUnread_1ujao_96";
const itemAvatar = "_itemAvatar_1ujao_106";
const itemContent = "_itemContent_1ujao_120";
const itemTitle = "_itemTitle_1ujao_128";
const itemDesc = "_itemDesc_1ujao_140";
const itemTime = "_itemTime_1ujao_151";
const itemDot = "_itemDot_1ujao_162";
const itemDotPlaceholder = "_itemDotPlaceholder_1ujao_171";
const empty$3 = "_empty_1ujao_178";
const emptyIcon = "_emptyIcon_1ujao_187";
const emptyText = "_emptyText_1ujao_191";
const footer$5 = "_footer_1ujao_199";
const footerBtn = "_footerBtn_1ujao_207";
const sheetHandle = "_sheetHandle_1ujao_230";
const overlay$4 = "_overlay_1ujao_236";
const s$l = {
  panel: panel$1,
  header: header$5,
  headerTitle,
  markAllBtn,
  list: list$2,
  item: item$4,
  itemUnread,
  itemAvatar,
  itemContent,
  itemTitle,
  itemDesc,
  itemTime,
  itemDot,
  itemDotPlaceholder,
  empty: empty$3,
  emptyIcon,
  emptyText,
  footer: footer$5,
  footerBtn,
  sheetHandle,
  overlay: overlay$4
};
function getTextContent(node) {
  if (node == null || typeof node === "boolean") return "";
  if (typeof node === "string" || typeof node === "number") return String(node);
  if (Array.isArray(node)) return node.map(getTextContent).join("");
  if (typeof node === "object" && "props" in node) {
    return getTextContent(node.props.children);
  }
  return "";
}
function NotificationPanel({
  open: open2,
  onClose,
  anchorRef,
  notifications,
  onClickItem,
  onMarkAllRead,
  onViewAll,
  title: title2 = "Notificações",
  viewAllLabel = "Ver todas as notificações",
  emptyMessage = "Nenhuma notificação",
  className
}) {
  const panelRef = useRef(null);
  const hasUnread = notifications.some((n) => n.unread);
  useDocumentEscape(open2, onClose);
  useOpenFocus({ active: open2, containerRef: panelRef });
  const applyPosition = useCallback(() => {
    const anchor2 = anchorRef.current;
    const el = panelRef.current;
    if (!anchor2 || !el) return;
    if (window.innerWidth <= 480) return;
    const ar = anchor2.getBoundingClientRect();
    const gap = 4;
    const margin = 8;
    el.style.position = "fixed";
    el.style.top = `${ar.bottom + gap}px`;
    el.style.bottom = "auto";
    const pr = el.getBoundingClientRect();
    const { top: top2, left: left2 } = resolveAnchoredOverlayPosition({
      anchorTop: ar.top,
      anchorBottom: ar.bottom,
      anchorLeft: ar.left,
      anchorRight: ar.right,
      overlayWidth: pr.width,
      overlayHeight: pr.height,
      viewportWidth: window.innerWidth,
      viewportHeight: window.innerHeight,
      gap,
      margin,
      horizontalAlign: "end",
      preferredVertical: "bottom"
    });
    el.style.left = `${left2}px`;
    el.style.top = `${top2}px`;
    el.style.right = "auto";
    el.style.bottom = "auto";
  }, [anchorRef]);
  useInitialReposition(open2, applyPosition);
  useViewportReposition(open2, applyPosition);
  useDocumentClickOutside({
    active: open2,
    refs: [anchorRef, panelRef],
    onOutside: onClose
  });
  const handleKeyDown = useCallback(
    (e) => {
      trapFocusWithin(panelRef.current, e);
    },
    []
  );
  if (!open2) return null;
  const classes = [s$l.panel, className ?? ""].filter(Boolean).join(" ");
  return createPortal(
    /* @__PURE__ */ jsxs(Fragment, { children: [
      /* @__PURE__ */ jsx("div", { className: s$l.overlay, role: "presentation", onClick: onClose }),
      /* @__PURE__ */ jsxs(
        "div",
        {
          ref: panelRef,
          className: classes,
          role: "dialog",
          "aria-modal": "true",
          "aria-label": title2,
          onKeyDown: handleKeyDown,
          onMouseDown: (e) => e.preventDefault(),
          children: [
            /* @__PURE__ */ jsx("div", { className: s$l.sheetHandle }),
            /* @__PURE__ */ jsxs("div", { className: s$l.header, children: [
              /* @__PURE__ */ jsx("span", { className: s$l.headerTitle, children: title2 }),
              hasUnread && onMarkAllRead && /* @__PURE__ */ jsx(
                "button",
                {
                  type: "button",
                  className: s$l.markAllBtn,
                  onClick: onMarkAllRead,
                  children: "Marcar todas como lidas"
                }
              )
            ] }),
            /* @__PURE__ */ jsx("div", { className: s$l.list, children: notifications.length === 0 ? /* @__PURE__ */ jsxs("div", { className: s$l.empty, children: [
              /* @__PURE__ */ jsx(BellSlash, { size: 32, className: s$l.emptyIcon }),
              /* @__PURE__ */ jsx("span", { className: s$l.emptyText, children: emptyMessage })
            ] }) : notifications.map((n) => /* @__PURE__ */ jsxs(
              "button",
              {
                type: "button",
                className: `${s$l.item} ${n.unread ? s$l.itemUnread : ""}`,
                "aria-label": `${getTextContent(n.title)}${n.time ? `, ${n.time}` : ""}${n.unread ? ", não lida" : ""}`,
                onClick: () => onClickItem?.(n.id),
                children: [
                  n.avatarUrl ? /* @__PURE__ */ jsx(
                    "img",
                    {
                      src: n.avatarUrl,
                      alt: "",
                      className: s$l.itemAvatar
                    }
                  ) : n.icon ? /* @__PURE__ */ jsx("div", { className: s$l.itemAvatar, children: /* @__PURE__ */ jsx(n.icon, { size: 16 }) }) : /* @__PURE__ */ jsx("div", { className: s$l.itemAvatar }),
                  /* @__PURE__ */ jsxs("div", { className: s$l.itemContent, children: [
                    /* @__PURE__ */ jsx("span", { className: s$l.itemTitle, children: n.title }),
                    n.description && /* @__PURE__ */ jsx("span", { className: s$l.itemDesc, children: n.description }),
                    /* @__PURE__ */ jsx("span", { className: s$l.itemTime, children: n.time })
                  ] }),
                  n.unread ? /* @__PURE__ */ jsx("span", { className: s$l.itemDot, "aria-label": "Não lida" }) : /* @__PURE__ */ jsx("span", { className: s$l.itemDotPlaceholder })
                ]
              },
              n.id
            )) }),
            onViewAll && notifications.length > 0 && /* @__PURE__ */ jsx("div", { className: s$l.footer, children: /* @__PURE__ */ jsx(
              "button",
              {
                type: "button",
                className: s$l.footerBtn,
                onClick: onViewAll,
                children: viewAllLabel
              }
            ) })
          ]
        }
      )
    ] }),
    document.body
  );
}
const header$4 = "_header_1i28v_3";
const title$5 = "_title_1i28v_16";
const actions = "_actions_1i28v_31";
const notificationWrapper = "_notificationWrapper_1i28v_40";
const badge = "_badge_1i28v_45";
const assistantBtn = "_assistantBtn_1i28v_59";
const assistantLabel = "_assistantLabel_1i28v_80";
const s$k = {
  header: header$4,
  title: title$5,
  actions,
  notificationWrapper,
  badge,
  assistantBtn,
  assistantLabel
};
const SearchButton = forwardRef(
  function SearchButton2({ className, ...rest }, ref) {
    return /* @__PURE__ */ jsx(
      Button,
      {
        ref,
        variant: "tertiary",
        size: "md",
        leftIcon: MagnifyingGlass,
        "aria-label": "Buscar",
        className,
        ...rest
      }
    );
  }
);
const NotificationButton = forwardRef(function NotificationButton2({ hasUnread = false, className, ...rest }, ref) {
  return /* @__PURE__ */ jsxs("span", { className: s$k.notificationWrapper, children: [
    /* @__PURE__ */ jsx(
      Button,
      {
        ref,
        variant: "tertiary",
        size: "md",
        leftIcon: Bell,
        "aria-label": hasUnread ? "Notificações (não lidas)" : "Notificações",
        className,
        ...rest
      }
    ),
    hasUnread && /* @__PURE__ */ jsx("span", { className: s$k.badge, "aria-hidden": "true" })
  ] });
});
const AssistantButton = forwardRef(function AssistantButton2({ label: label2 = "Assistente", active: active2 = false, className, ...rest }, ref) {
  const classes = [
    s$k.assistantBtn,
    active2 ? btnStyles.active : "",
    className ?? ""
  ].filter(Boolean).join(" ") || void 0;
  return /* @__PURE__ */ jsx(
    Button,
    {
      ref,
      variant: "tertiary",
      size: "md",
      leftIcon: active2 ? SidebarSimple : Lightning,
      "aria-label": label2,
      "aria-expanded": active2,
      className: classes,
      ...rest,
      children: /* @__PURE__ */ jsx("span", { className: s$k.assistantLabel, children: label2 })
    },
    active2 ? "active" : "default"
  );
});
function PageHeader({
  title: title2,
  children,
  className
}) {
  const classes = [s$k.header, className ?? ""].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("header", { className: classes, children: [
    /* @__PURE__ */ jsx("h1", { className: s$k.title, children: title2 }),
    children && /* @__PURE__ */ jsx("div", { className: s$k.actions, children })
  ] });
}
const overlay$3 = "_overlay_1k9fb_3";
const container$1 = "_container_1k9fb_22";
const inputRow = "_inputRow_1k9fb_48";
const inputIcon = "_inputIcon_1k9fb_56";
const input$3 = "_input_1k9fb_48";
const list$1 = "_list_1k9fb_80";
const groupLabel$1 = "_groupLabel_1k9fb_88";
const item$3 = "_item_1k9fb_98";
const itemSelected = "_itemSelected_1k9fb_117";
const itemIcon = "_itemIcon_1k9fb_127";
const itemHint = "_itemHint_1k9fb_136";
const empty$2 = "_empty_1k9fb_144";
const footer$4 = "_footer_1k9fb_154";
const kbd = "_kbd_1k9fb_165";
const s$j = {
  overlay: overlay$3,
  container: container$1,
  inputRow,
  inputIcon,
  input: input$3,
  list: list$1,
  groupLabel: groupLabel$1,
  item: item$3,
  itemSelected,
  itemIcon,
  itemHint,
  empty: empty$2,
  footer: footer$4,
  kbd
};
function matchesQuery(item2, q) {
  if (item2.label.toLowerCase().includes(q)) return true;
  if (item2.keywords?.some((kw) => kw.toLowerCase().includes(q))) return true;
  return false;
}
function CommandPalette({
  open: open2,
  onClose,
  onSelect,
  groups,
  placeholder: placeholder2 = "Buscar...",
  emptyMessage = "Nenhum resultado encontrado",
  footer: footer2
}) {
  const [query, setQuery] = useState("");
  const [selectedIndex, setSelectedIndex] = useState(0);
  const inputRef = useRef(null);
  const listRef = useRef(null);
  const containerRef = useRef(null);
  const uid2 = useId();
  const listboxId = `${uid2}-listbox`;
  useDocumentEscape(open2, onClose);
  useOpenFocus({
    active: open2,
    containerRef,
    initialFocusRef: inputRef
  });
  const getItemId = useCallback(
    (flatIdx) => `${uid2}-item-${flatIdx}`,
    [uid2]
  );
  const filteredGroups = useMemo(() => {
    const q = query.trim().toLowerCase();
    if (!q) return groups;
    return groups.map((g) => ({
      ...g,
      items: g.items.filter((item2) => matchesQuery(item2, q))
    })).filter((g) => g.items.length > 0);
  }, [groups, query]);
  const flatItems = useMemo(
    () => filteredGroups.flatMap((g) => g.items),
    [filteredGroups]
  );
  useEffect(() => {
    if (open2) {
      setQuery("");
      setSelectedIndex(0);
    }
  }, [open2]);
  useEffect(() => {
    setSelectedIndex(0);
  }, [flatItems.length]);
  useEffect(() => {
    if (!listRef.current) return;
    const els = listRef.current.querySelectorAll(
      "[data-command-item]"
    );
    els[selectedIndex]?.scrollIntoView({ block: "nearest" });
  }, [selectedIndex]);
  function handleKeyDown(e) {
    trapFocusWithin(containerRef.current, e);
    switch (e.key) {
      case "ArrowDown":
        e.preventDefault();
        setSelectedIndex((i) => i < flatItems.length - 1 ? i + 1 : 0);
        break;
      case "ArrowUp":
        e.preventDefault();
        setSelectedIndex((i) => i > 0 ? i - 1 : flatItems.length - 1);
        break;
      case "Enter":
        if (flatItems.length > 0) {
          e.preventDefault();
          onSelect(flatItems[selectedIndex].id);
        }
        break;
      case "Escape":
        e.preventDefault();
        break;
    }
  }
  if (!open2) return null;
  let flatIndex = 0;
  const renderGroups = filteredGroups.map((g) => {
    const itemsWithIndex = g.items.map((item2) => ({
      item: item2,
      flatIdx: flatIndex++
    }));
    return { label: g.label, items: itemsWithIndex };
  });
  return createPortal(
    /* @__PURE__ */ jsx(
      "div",
      {
        className: s$j.overlay,
        role: "presentation",
        onMouseDown: (e) => {
          if (e.target === e.currentTarget) onClose();
        },
        onKeyDown: handleKeyDown,
        children: /* @__PURE__ */ jsxs("div", { className: s$j.container, ref: containerRef, role: "dialog", "aria-modal": "true", "aria-label": "Busca", children: [
          /* @__PURE__ */ jsxs("div", { className: s$j.inputRow, children: [
            /* @__PURE__ */ jsx(MagnifyingGlass, { size: 16, className: s$j.inputIcon }),
            /* @__PURE__ */ jsx(
              "input",
              {
                ref: inputRef,
                className: s$j.input,
                type: "text",
                role: "combobox",
                "aria-expanded": "true",
                "aria-haspopup": "listbox",
                "aria-controls": listboxId,
                "aria-activedescendant": flatItems.length > 0 ? getItemId(selectedIndex) : void 0,
                placeholder: placeholder2,
                value: query,
                onChange: (e) => setQuery(e.target.value)
              }
            )
          ] }),
          /* @__PURE__ */ jsx("div", { className: s$j.list, ref: listRef, id: listboxId, role: "listbox", children: flatItems.length === 0 ? /* @__PURE__ */ jsx("div", { className: s$j.empty, children: emptyMessage }) : renderGroups.map((g) => /* @__PURE__ */ jsxs("div", { children: [
            g.label && /* @__PURE__ */ jsx("div", { className: s$j.groupLabel, children: g.label }),
            g.items.map(({ item: item2, flatIdx }) => /* @__PURE__ */ jsxs(
              "button",
              {
                id: getItemId(flatIdx),
                type: "button",
                role: "option",
                "aria-selected": flatIdx === selectedIndex,
                "data-command-item": true,
                className: `${s$j.item} ${flatIdx === selectedIndex ? s$j.itemSelected : ""}`,
                onMouseEnter: () => setSelectedIndex(flatIdx),
                onClick: () => onSelect(item2.id),
                children: [
                  item2.icon && /* @__PURE__ */ jsx(item2.icon, { size: 16, className: s$j.itemIcon }),
                  /* @__PURE__ */ jsx("span", { children: item2.label }),
                  item2.hint && /* @__PURE__ */ jsx("span", { className: s$j.itemHint, children: item2.hint })
                ]
              },
              item2.id
            ))
          ] }, g.label)) }),
          footer2 !== void 0 ? footer2 : /* @__PURE__ */ jsxs("div", { className: s$j.footer, children: [
            /* @__PURE__ */ jsxs("span", { children: [
              /* @__PURE__ */ jsx("kbd", { className: s$j.kbd, children: "↑" }),
              " ",
              /* @__PURE__ */ jsx("kbd", { className: s$j.kbd, children: "↓" }),
              " ",
              "navegar"
            ] }),
            /* @__PURE__ */ jsxs("span", { children: [
              /* @__PURE__ */ jsx("kbd", { className: s$j.kbd, children: "↵" }),
              " selecionar"
            ] }),
            /* @__PURE__ */ jsxs("span", { children: [
              /* @__PURE__ */ jsx("kbd", { className: s$j.kbd, children: "esc" }),
              " fechar"
            ] })
          ] })
        ] })
      }
    ),
    document.body
  );
}
const wrapper$6 = "_wrapper_1tpaf_2";
const label$6 = "_label_1tpaf_10";
const inputBox = "_inputBox_1tpaf_19";
const sm$6 = "_sm_1tpaf_32";
const md$4 = "_md_1tpaf_37";
const lg$3 = "_lg_1tpaf_42";
const disabled$7 = "_disabled_1tpaf_47";
const hovered$3 = "_hovered_1tpaf_48";
const focused$3 = "_focused_1tpaf_53";
const error$3 = "_error_1tpaf_59";
const input$2 = "_input_1tpaf_19";
const message$2 = "_message_1tpaf_118";
const attention$2 = "_attention_1tpaf_132";
const success$3 = "_success_1tpaf_136";
const s$i = {
  wrapper: wrapper$6,
  label: label$6,
  inputBox,
  sm: sm$6,
  md: md$4,
  lg: lg$3,
  disabled: disabled$7,
  hovered: hovered$3,
  focused: focused$3,
  error: error$3,
  input: input$2,
  message: message$2,
  attention: attention$2,
  success: success$3
};
const ICON_SIZES$1 = { sm: 14, md: 16, lg: 20 };
const messageIconMap$2 = {
  error: WarningCircle,
  attention: WarningCircle,
  success: CheckCircle
};
const Input = forwardRef(
  ({
    size = "md",
    label: label2,
    leftIcon: LeftIcon,
    rightIcon: RightIcon,
    message: message2,
    messageType,
    disabled: disabled2 = false,
    className,
    ...rest
  }, ref) => {
    const autoId = useId();
    const inputId = rest.id ?? autoId;
    const messageId = `${inputId}-msg`;
    const hasMessage = !!message2 && !!messageType;
    const isError = messageType === "error";
    const iconSize2 = ICON_SIZES$1[size];
    const wrapperClasses = [s$i.wrapper, className ?? ""].filter(Boolean).join(" ");
    const inputBoxClasses = [
      s$i.inputBox,
      s$i[size],
      isError ? s$i.error : "",
      disabled2 ? s$i.disabled : ""
    ].filter(Boolean).join(" ");
    const MsgIcon = messageType ? messageIconMap$2[messageType] : null;
    return /* @__PURE__ */ jsxs("div", { className: wrapperClasses, children: [
      label2 && /* @__PURE__ */ jsx("label", { className: s$i.label, htmlFor: inputId, children: label2 }),
      /* @__PURE__ */ jsxs("div", { className: inputBoxClasses, children: [
        LeftIcon && /* @__PURE__ */ jsx(LeftIcon, { size: iconSize2, "aria-hidden": "true" }),
        /* @__PURE__ */ jsx(
          "input",
          {
            ref,
            id: inputId,
            className: s$i.input,
            disabled: disabled2,
            "aria-invalid": isError || void 0,
            "aria-describedby": hasMessage ? messageId : void 0,
            ...rest
          }
        ),
        RightIcon && /* @__PURE__ */ jsx(RightIcon, { size: iconSize2, "aria-hidden": "true" })
      ] }),
      hasMessage && /* @__PURE__ */ jsxs("div", { id: messageId, className: `${s$i.message} ${s$i[messageType]}`, children: [
        MsgIcon && /* @__PURE__ */ jsx(MsgIcon, { size: 14, "aria-hidden": "true" }),
        /* @__PURE__ */ jsx("span", { children: message2 })
      ] })
    ] });
  }
);
Input.displayName = "Input";
const screen = "_screen_gizqi_3";
const logoContainer = "_logoContainer_gizqi_14";
const logoBackground = "_logoBackground_gizqi_20";
const logoFilled = "_logoFilled_gizqi_26";
const label$5 = "_label_gizqi_46";
const s$h = {
  screen,
  logoContainer,
  logoBackground,
  logoFilled,
  label: label$5
};
const SYMBOL_PATH = "M239.986 44.8463C239.406 21.0338 220.559 1.4407 196.862 0.0776031C189.718 -0.333 182.901 0.900098 176.753 3.43709V3.43451L155.54 12.2361C153.849 12.9376 151.913 12.6686 150.476 11.532L148.329 9.83555C140.532 3.67521 130.693 0.000373357 119.999 0.000373357C113.82 0.000373357 107.926 1.22832 102.545 3.45511V3.45253L6.85696 43.3789L6.85439 43.3802L6.50423 43.5269L6.47217 43.5398V43.5411C5.14977 44.1216 3.9659 44.9634 2.98469 45.9983C1.13641 47.947 0 50.5831 0 53.4869C0 56.3882 1.13385 59.0243 2.9834 60.9756C4.41226 62.4816 6.26695 63.5757 8.35507 64.0648L8.37816 64.0712L107.653 90.31C111.58 91.4105 115.722 91.9987 119.999 91.9987C130.693 91.9987 140.53 88.3239 148.329 82.1635C148.969 81.6577 149.595 81.1364 150.207 80.5984C151.668 79.3138 153.718 78.9521 155.513 79.7012L176.325 88.3844L176.755 88.5646C182.124 90.7785 188.003 92 194.166 92C214.255 92 231.327 79.0306 237.515 60.9782C239.243 55.9364 240.124 50.5046 239.986 44.8463ZM148.33 46.2827V46.3342C148.33 46.3754 148.33 46.4166 148.329 46.4565C148.245 51.7789 146.704 56.7447 144.087 60.9743C141.591 65.0083 138.118 68.3678 133.998 70.7207C129.87 73.0788 125.094 74.4264 120.004 74.4264C115.121 74.4264 110.528 73.1869 106.52 71.0052C102.178 68.6419 98.5204 65.1731 95.9217 60.9743C94.498 58.6729 93.3924 56.1539 92.6664 53.4754C92.02 51.0915 91.6762 48.5854 91.6762 45.997C91.6762 43.3287 92.0418 40.7441 92.7293 38.2946C94.5083 31.9335 98.4383 26.4734 103.683 22.7586C108.297 19.4879 113.927 17.5675 120.006 17.5675C125.053 17.5675 129.792 18.892 133.897 21.2153C142.39 26.019 148.167 35.0947 148.33 45.5387C148.331 45.5799 148.331 45.6198 148.331 45.6597V45.7112C148.334 45.8065 148.334 45.9017 148.334 45.997C148.333 46.0922 148.333 46.1875 148.33 46.2827ZM218.252 60.9743C213.256 69.05 204.338 74.4264 194.169 74.4264C187.602 74.4264 181.558 72.1829 176.755 68.4205C174.119 66.3572 171.857 63.8344 170.087 60.9743C167.396 56.625 165.842 51.4931 165.842 45.9983C165.842 36.8878 170.113 28.7774 176.755 23.576C181.558 19.8123 187.602 17.5701 194.169 17.5701C209.814 17.5701 222.498 30.2975 222.498 45.9995C222.497 51.4931 220.943 56.6237 218.252 60.9743Z";
const VB_W = 240;
const VB_H = 92;
function LoadingScreen({
  message: message2 = "Carregando...",
  className
}) {
  const cls = [s$h.screen, className].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs("div", { className: cls, role: "status", "aria-label": message2, children: [
    /* @__PURE__ */ jsxs("div", { className: s$h.logoContainer, children: [
      /* @__PURE__ */ jsx(
        "svg",
        {
          className: s$h.logoBackground,
          viewBox: `0 0 ${VB_W} ${VB_H}`,
          width: 96,
          height: 37,
          "aria-hidden": true,
          children: /* @__PURE__ */ jsx("path", { d: SYMBOL_PATH, fill: "currentColor" })
        }
      ),
      /* @__PURE__ */ jsxs(
        "svg",
        {
          className: s$h.logoFilled,
          viewBox: `0 0 ${VB_W} ${VB_H}`,
          width: 96,
          height: 37,
          "aria-hidden": true,
          children: [
            /* @__PURE__ */ jsx("defs", { children: /* @__PURE__ */ jsx("clipPath", { id: "bud-liquid", children: /* @__PURE__ */ jsx(
              "path",
              {
                className: s$h.liquidRect,
                d: `
                  M -${VB_W} -10
                  L ${VB_W - 4} -10
                  C ${VB_W + 4} 10, ${VB_W - 4} 20, ${VB_W} 30
                  C ${VB_W + 4} 40, ${VB_W - 4} 52, ${VB_W} 62
                  C ${VB_W + 4} 72, ${VB_W - 4} 82, ${VB_W} 102
                  L -${VB_W} 102
                  Z
                `
              }
            ) }) }),
            /* @__PURE__ */ jsx(
              "path",
              {
                d: SYMBOL_PATH,
                fill: "currentColor",
                clipPath: "url(#bud-liquid)"
              }
            )
          ]
        }
      )
    ] }),
    /* @__PURE__ */ jsx("span", { className: s$h.label, children: message2 })
  ] });
}
const overlay$2 = "_overlay_v47mb_2";
const hidden$1 = "_hidden_v47mb_17";
const panel = "_panel_v47mb_24";
const right$1 = "_right_v47mb_38";
const left$1 = "_left_v47mb_46";
const panelOpen = "_panelOpen_v47mb_54";
const sm$5 = "_sm_v47mb_59";
const md$3 = "_md_v47mb_63";
const lg$2 = "_lg_v47mb_67";
const header$3 = "_header_v47mb_72";
const headerTop$1 = "_headerTop_v47mb_81";
const headerActions$1 = "_headerActions_v47mb_87";
const headerText$1 = "_headerText_v47mb_94";
const title$4 = "_title_v47mb_102";
const description$4 = "_description_v47mb_111";
const body$3 = "_body_v47mb_121";
const footer$3 = "_footer_v47mb_128";
const footerBetween$1 = "_footerBetween_v47mb_138";
const s$g = {
  overlay: overlay$2,
  hidden: hidden$1,
  panel,
  right: right$1,
  left: left$1,
  panelOpen,
  sm: sm$5,
  md: md$3,
  lg: lg$2,
  header: header$3,
  headerTop: headerTop$1,
  headerActions: headerActions$1,
  headerText: headerText$1,
  title: title$4,
  description: description$4,
  body: body$3,
  footer: footer$3,
  footerBetween: footerBetween$1
};
const DrawerTitleIdContext = createContext(void 0);
function Drawer({
  open: open2,
  onClose,
  side = "right",
  size = "md",
  width,
  className,
  "aria-label": ariaLabel,
  children
}) {
  const overlayRef = useRef(null);
  const panelRef = useRef(null);
  const everOpened = useHasOpened(open2);
  const titleId = useId();
  useDocumentEscape(open2, onClose);
  useBodyScrollLock(open2);
  useOpenFocus({ active: open2, containerRef: panelRef });
  const handleKeyDown = useCallback(
    (e) => {
      trapFocusWithin(panelRef.current, e);
    },
    []
  );
  if (!everOpened) return null;
  const sideClass = side === "left" ? s$g.left : s$g.right;
  const sizeClass2 = width ? "" : s$g[size];
  const panelStyle = width ? { width } : void 0;
  const panelCls = [s$g.panel, sizeClass2, sideClass, open2 && s$g.panelOpen, className].filter(Boolean).join(" ");
  return createPortal(
    /* @__PURE__ */ jsx(
      "div",
      {
        ref: overlayRef,
        className: `${s$g.overlay} ${open2 ? "" : s$g.hidden}`,
        onMouseDown: (e) => {
          if (open2 && e.target === overlayRef.current) onClose();
        },
        onKeyDown: open2 ? handleKeyDown : void 0,
        children: /* @__PURE__ */ jsx(DrawerTitleIdContext.Provider, { value: titleId, children: /* @__PURE__ */ jsx(
          "aside",
          {
            ref: panelRef,
            className: panelCls,
            style: panelStyle,
            role: "dialog",
            "aria-modal": "true",
            "aria-label": ariaLabel,
            "aria-labelledby": ariaLabel ? void 0 : titleId,
            children
          }
        ) })
      }
    ),
    document.body
  );
}
function DrawerHeader({
  title: title2,
  description: description2,
  onClose,
  children,
  afterTitle
}) {
  const titleId = useContext(DrawerTitleIdContext);
  return /* @__PURE__ */ jsxs("div", { className: s$g.header, children: [
    /* @__PURE__ */ jsxs("div", { className: s$g.headerTop, children: [
      /* @__PURE__ */ jsxs("div", { className: s$g.headerText, children: [
        /* @__PURE__ */ jsx("h2", { id: titleId, className: s$g.title, children: title2 }),
        description2 && /* @__PURE__ */ jsx("p", { className: s$g.description, children: description2 })
      ] }),
      (children || onClose) && /* @__PURE__ */ jsxs("div", { className: s$g.headerActions, children: [
        children,
        onClose && /* @__PURE__ */ jsx(
          Button,
          {
            variant: "tertiary",
            size: "md",
            leftIcon: X,
            onClick: onClose,
            "aria-label": "Fechar"
          }
        )
      ] })
    ] }),
    afterTitle
  ] });
}
function DrawerBody({ children }) {
  return /* @__PURE__ */ jsx("div", { className: s$g.body, children });
}
function DrawerFooter({ children, align = "end" }) {
  return /* @__PURE__ */ jsx(
    "div",
    {
      className: `${s$g.footer} ${align === "between" ? s$g.footerBetween : ""}`,
      children
    }
  );
}
function DragToCloseDrawer({
  children,
  open: open2,
  onClose,
  dragToCloseEnabled = true,
  dragThreshold = 80,
  velocityThreshold = 0.5,
  dragZoneHeight = 60,
  ...drawerProps
}) {
  const touchStartY = useRef(0);
  const touchStartTime = useRef(0);
  const isDragging = useRef(false);
  const panelRef = useRef(null);
  const cleanupRef = useRef(null);
  useEffect(() => {
    if (!dragToCloseEnabled || !open2) {
      panelRef.current = null;
      return;
    }
    const findPanel = () => {
      const panel2 = document.querySelector(
        '[role="dialog"][aria-modal="true"]:not([data-drag-handled])'
      );
      if (panel2) {
        panel2.setAttribute("data-drag-handled", "true");
        panelRef.current = panel2;
      }
    };
    findPanel();
    const timeout = setTimeout(findPanel, 50);
    return () => {
      clearTimeout(timeout);
      if (panelRef.current) {
        panelRef.current.removeAttribute("data-drag-handled");
      }
    };
  }, [open2, dragToCloseEnabled]);
  useEffect(() => {
    if (!dragToCloseEnabled || !open2) return;
    function handleTouchStart(e) {
      const panel2 = panelRef.current;
      if (!panel2) return;
      const touch = e.touches[0];
      if (!touch) return;
      const rect = panel2.getBoundingClientRect();
      const relativeY = touch.clientY - rect.top;
      if (relativeY <= dragZoneHeight) {
        touchStartY.current = touch.clientY;
        touchStartTime.current = Date.now();
        isDragging.current = true;
      }
    }
    function handleTouchMove(e) {
      if (!isDragging.current) return;
      const panel2 = panelRef.current;
      if (!panel2) return;
      const touch = e.touches[0];
      if (!touch) return;
      const deltaY = touch.clientY - touchStartY.current;
      if (deltaY > 0) {
        panel2.style.transform = `translateY(${deltaY}px)`;
        const opacity = Math.max(0.3, 1 - deltaY / 400);
        panel2.style.opacity = String(opacity);
        e.preventDefault();
      }
    }
    function handleTouchEnd(e) {
      if (!isDragging.current) return;
      const panel2 = panelRef.current;
      if (!panel2) return;
      const touch = e.changedTouches[0];
      if (!touch) return;
      const deltaY = touch.clientY - touchStartY.current;
      const deltaTime = Date.now() - touchStartTime.current;
      const velocity = deltaY / deltaTime;
      if (deltaY > dragThreshold || velocity > velocityThreshold) {
        onClose();
      } else {
        const cleanupTransition = () => {
          panel2.style.transition = "";
          panel2.removeEventListener("transitionend", cleanupTransition);
        };
        panel2.addEventListener("transitionend", cleanupTransition);
        panel2.style.transition = "transform 250ms ease, opacity 250ms ease";
        panel2.style.transform = "";
        panel2.style.opacity = "";
      }
      isDragging.current = false;
    }
    document.addEventListener("touchstart", handleTouchStart, { passive: true });
    document.addEventListener("touchmove", handleTouchMove, { passive: false });
    document.addEventListener("touchend", handleTouchEnd, { passive: true });
    cleanupRef.current = () => {
      document.removeEventListener("touchstart", handleTouchStart);
      document.removeEventListener("touchmove", handleTouchMove);
      document.removeEventListener("touchend", handleTouchEnd);
    };
    return cleanupRef.current;
  }, [open2, dragToCloseEnabled, dragThreshold, velocityThreshold, dragZoneHeight, onClose]);
  return /* @__PURE__ */ jsx(Drawer, { open: open2, onClose, ...drawerProps, children });
}
const overlay$1 = "_overlay_nyg3l_2";
const hidden = "_hidden_nyg3l_21";
const container = "_container_nyg3l_28";
const sm$4 = "_sm_nyg3l_55";
const md$2 = "_md_nyg3l_59";
const lg$1 = "_lg_nyg3l_63";
const doubleLayout = "_doubleLayout_nyg3l_68";
const sidePanel = "_sidePanel_nyg3l_92";
const sidePanelOpen = "_sidePanelOpen_nyg3l_102";
const header$2 = "_header_nyg3l_115";
const headerTop = "_headerTop_nyg3l_123";
const headerActions = "_headerActions_nyg3l_129";
const headerText = "_headerText_nyg3l_136";
const title$3 = "_title_nyg3l_144";
const description$3 = "_description_nyg3l_153";
const body$2 = "_body_nyg3l_163";
const footer$2 = "_footer_nyg3l_170";
const footerBetween = "_footerBetween_nyg3l_179";
const s$f = {
  overlay: overlay$1,
  hidden,
  container,
  sm: sm$4,
  md: md$2,
  lg: lg$1,
  doubleLayout,
  sidePanel,
  sidePanelOpen,
  header: header$2,
  headerTop,
  headerActions,
  headerText,
  title: title$3,
  description: description$3,
  body: body$2,
  footer: footer$2,
  footerBetween
};
const ModalTitleIdContext = createContext(void 0);
function Modal({
  open: open2,
  onClose,
  size = "md",
  children,
  sidePanel: sidePanel2,
  width,
  className,
  "aria-label": ariaLabel
}) {
  const overlayRef = useRef(null);
  const containerRef = useRef(null);
  const everOpened = useHasOpened(open2);
  const titleId = useId();
  useDocumentEscape(open2, onClose);
  useBodyScrollLock(open2);
  useOpenFocus({ active: open2, containerRef });
  const handleKeyDown = useCallback(
    (e) => {
      trapFocusWithin(containerRef.current, e);
    },
    []
  );
  if (!everOpened) return null;
  const sizeClass2 = width ? "" : s$f[size];
  const containerStyle = width ? { width } : void 0;
  const containerClasses = [s$f.container, sizeClass2, className].filter(Boolean).join(" ");
  const content2 = /* @__PURE__ */ jsx("div", { className: containerClasses, style: containerStyle, children });
  return createPortal(
    /* @__PURE__ */ jsx(
      "div",
      {
        ref: overlayRef,
        className: `${s$f.overlay} ${open2 ? "" : s$f.hidden}`,
        onMouseDown: (e) => {
          if (open2 && e.target === overlayRef.current) onClose();
        },
        onKeyDown: open2 ? handleKeyDown : void 0,
        children: /* @__PURE__ */ jsx(ModalTitleIdContext.Provider, { value: titleId, children: sidePanel2 !== void 0 ? /* @__PURE__ */ jsxs(
          "div",
          {
            ref: containerRef,
            className: s$f.doubleLayout,
            role: "dialog",
            "aria-modal": "true",
            "aria-label": ariaLabel,
            "aria-labelledby": ariaLabel ? void 0 : titleId,
            children: [
              content2,
              /* @__PURE__ */ jsx("div", { className: `${s$f.sidePanel} ${sidePanel2 ? s$f.sidePanelOpen : ""}`, children: sidePanel2 })
            ]
          }
        ) : /* @__PURE__ */ jsx(
          "div",
          {
            ref: containerRef,
            className: containerClasses,
            style: containerStyle,
            role: "dialog",
            "aria-modal": "true",
            "aria-label": ariaLabel,
            "aria-labelledby": ariaLabel ? void 0 : titleId,
            children
          }
        ) })
      }
    ),
    document.body
  );
}
function ModalHeader({
  title: title2,
  description: description2,
  onClose,
  children,
  afterTitle
}) {
  const titleId = useContext(ModalTitleIdContext);
  return /* @__PURE__ */ jsxs("div", { className: s$f.header, children: [
    /* @__PURE__ */ jsxs("div", { className: s$f.headerTop, children: [
      /* @__PURE__ */ jsxs("div", { className: s$f.headerText, children: [
        /* @__PURE__ */ jsx("h2", { id: titleId, className: s$f.title, children: title2 }),
        description2 && /* @__PURE__ */ jsx("p", { className: s$f.description, children: description2 })
      ] }),
      (children || onClose) && /* @__PURE__ */ jsxs("div", { className: s$f.headerActions, children: [
        children,
        onClose && /* @__PURE__ */ jsx(
          Button,
          {
            variant: "tertiary",
            size: "md",
            leftIcon: X,
            onClick: onClose,
            "aria-label": "Fechar"
          }
        )
      ] })
    ] }),
    afterTitle
  ] });
}
function ModalBody({ children }) {
  return /* @__PURE__ */ jsx("div", { className: s$f.body, children });
}
function ModalFooter({ children, align = "end" }) {
  return /* @__PURE__ */ jsx("div", { className: `${s$f.footer} ${align === "between" ? s$f.footerBetween : ""}`, children });
}
const wrapper$5 = "_wrapper_8mn0t_2";
const hasLabel$1 = "_hasLabel_8mn0t_10";
const disabled$6 = "_disabled_8mn0t_14";
const input$1 = "_input_8mn0t_19";
const box = "_box_8mn0t_31";
const dot$1 = "_dot_8mn0t_43";
const sm$3 = "_sm_8mn0t_50";
const md$1 = "_md_8mn0t_55";
const hovered$2 = "_hovered_8mn0t_109";
const focused$2 = "_focused_8mn0t_118";
const text$2 = "_text_8mn0t_130";
const text_sm = "_text_sm_8mn0t_139";
const title$2 = "_title_8mn0t_143";
const text_md = "_text_md_8mn0t_149";
const description$2 = "_description_8mn0t_163";
const s$e = {
  wrapper: wrapper$5,
  hasLabel: hasLabel$1,
  disabled: disabled$6,
  input: input$1,
  box,
  dot: dot$1,
  sm: sm$3,
  md: md$1,
  hovered: hovered$2,
  focused: focused$2,
  text: text$2,
  text_sm,
  title: title$2,
  text_md,
  description: description$2
};
const boxSize = { sm: 16, md: 20 };
const Radio = forwardRef(
  ({
    size = "md",
    label: label2,
    description: description2,
    disabled: disabled2 = false,
    className,
    ...rest
  }, ref) => {
    const wrapperClasses = [
      s$e.wrapper,
      label2 ? s$e.hasLabel : "",
      disabled2 ? s$e.disabled : "",
      className ?? ""
    ].filter(Boolean).join(" ");
    return /* @__PURE__ */ jsxs("label", { className: wrapperClasses, children: [
      /* @__PURE__ */ jsx(
        "input",
        {
          ref,
          type: "radio",
          className: s$e.input,
          disabled: disabled2,
          ...rest
        }
      ),
      /* @__PURE__ */ jsx(
        "span",
        {
          className: `${s$e.box} ${s$e[size]}`,
          style: {
            width: boxSize[size],
            height: boxSize[size]
          },
          children: /* @__PURE__ */ jsx("span", { className: s$e.dot })
        }
      ),
      label2 && /* @__PURE__ */ jsxs("span", { className: `${s$e.text} ${s$e[`text_${size}`]}`, children: [
        /* @__PURE__ */ jsx("span", { className: s$e.title, children: label2 }),
        description2 && /* @__PURE__ */ jsx("span", { className: s$e.description, children: description2 })
      ] })
    ] });
  }
);
Radio.displayName = "Radio";
const body$1 = "_body_2u6e7_3";
const item$2 = "_item_2u6e7_13";
const itemActive$1 = "_itemActive_2u6e7_36";
const searchRow = "_searchRow_2u6e7_50";
const searchIcon = "_searchIcon_2u6e7_59";
const searchInput$1 = "_searchInput_2u6e7_64";
const createRow = "_createRow_2u6e7_82";
const createInput = "_createInput_2u6e7_91";
const footer$1 = "_footer_2u6e7_109";
const empty$1 = "_empty_2u6e7_116";
const styles = {
  body: body$1,
  item: item$2,
  itemActive: itemActive$1,
  searchRow,
  searchIcon,
  searchInput: searchInput$1,
  createRow,
  createInput,
  footer: footer$1,
  empty: empty$1
};
function formatMultiLabel(ids, options, fallback) {
  if (ids.length === 0) return fallback;
  const labels = ids.map((id) => options.find((o) => o.id === id)?.label ?? id);
  if (labels.length === 1) return labels[0];
  return `${labels[0]} +${labels.length - 1}`;
}
function PopoverSelect(props) {
  const {
    open: open2,
    onClose,
    anchorRef,
    placement,
    noOverlay,
    ignoreRefs,
    options,
    renderOptionPrefix,
    searchable,
    searchPlaceholder = "Buscar...",
    footer: footer2,
    emptyText: emptyText2 = "Nenhum resultado encontrado"
  } = props;
  const [search, setSearch] = useState("");
  const [createValue, setCreateValue] = useState("");
  const filtered = useMemo(() => {
    if (!searchable || !search.trim()) return options;
    const term = search.toLowerCase();
    return options.filter((o) => o.label.toLowerCase().includes(term));
  }, [options, search, searchable]);
  function isSelected(id) {
    if (props.mode === "single") return props.value === id;
    return props.value.includes(id);
  }
  function handleSelect(id) {
    if (props.mode === "single") {
      props.onChange(id);
      if (props.closeOnSelect !== false) onClose();
    } else {
      const prev2 = props.value;
      props.onChange(
        prev2.includes(id) ? prev2.filter((v) => v !== id) : [...prev2, id]
      );
    }
  }
  function handleCreate() {
    if (props.mode !== "multiple" || !props.creatable || !props.onCreateOption) return;
    const label2 = createValue.trim();
    if (!label2) return;
    const newOpt = props.onCreateOption(label2);
    props.onChange([...props.value, newOpt.id]);
    setCreateValue("");
  }
  const Indicator = props.mode === "single" ? Radio : Checkbox;
  const CreatableIcon = props.mode === "multiple" && props.creatable ? props.creatableIcon : void 0;
  return /* @__PURE__ */ jsxs(
    FilterDropdown,
    {
      open: open2,
      onClose: () => {
        setSearch("");
        onClose();
      },
      anchorRef,
      placement,
      noOverlay,
      ignoreRefs,
      children: [
        /* @__PURE__ */ jsxs("div", { className: styles.body, children: [
          searchable && /* @__PURE__ */ jsxs("div", { className: styles.searchRow, children: [
            /* @__PURE__ */ jsx(MagnifyingGlass, { size: 14, className: styles.searchIcon }),
            /* @__PURE__ */ jsx(
              "input",
              {
                type: "text",
                className: styles.searchInput,
                placeholder: searchPlaceholder,
                value: search,
                onChange: (e) => setSearch(e.target.value)
              }
            )
          ] }),
          filtered.length === 0 && /* @__PURE__ */ jsx("div", { className: styles.empty, children: emptyText2 }),
          filtered.map((opt) => {
            const active2 = isSelected(opt.id);
            return /* @__PURE__ */ jsxs(
              "button",
              {
                type: "button",
                className: `${styles.item} ${active2 ? styles.itemActive : ""}`,
                onClick: () => handleSelect(opt.id),
                children: [
                  /* @__PURE__ */ jsx(Indicator, { checked: active2, readOnly: true }),
                  renderOptionPrefix?.(opt),
                  !renderOptionPrefix && opt.initials && /* @__PURE__ */ jsx(Avatar, { initials: opt.initials, src: opt.avatarSrc, size: "xs" }),
                  !renderOptionPrefix && opt.icon && /* @__PURE__ */ jsx(opt.icon, { size: 14 }),
                  /* @__PURE__ */ jsx("span", { children: opt.label })
                ]
              },
              opt.id
            );
          }),
          props.mode === "multiple" && props.creatable && props.onCreateOption && /* @__PURE__ */ jsxs("div", { className: styles.createRow, children: [
            /* @__PURE__ */ jsx(
              "input",
              {
                type: "text",
                className: styles.createInput,
                placeholder: props.createPlaceholder ?? "Criar novo...",
                value: createValue,
                onChange: (e) => setCreateValue(e.target.value),
                onKeyDown: (e) => {
                  if (e.key === "Enter") handleCreate();
                }
              }
            ),
            /* @__PURE__ */ jsx(
              Button,
              {
                variant: "tertiary",
                size: "sm",
                leftIcon: CreatableIcon ?? Plus,
                "aria-label": "Criar",
                disabled: !createValue.trim(),
                onClick: handleCreate
              }
            )
          ] })
        ] }),
        footer2 && /* @__PURE__ */ jsx("div", { className: styles.footer, children: footer2 })
      ]
    }
  );
}
const root$5 = "_root_7o3zr_1";
const grid = "_grid_7o3zr_6";
const axis = "_axis_7o3zr_12";
const area = "_area_7o3zr_17";
const outline = "_outline_7o3zr_22";
const dot = "_dot_7o3zr_29";
const label$4 = "_label_7o3zr_33";
const valueText = "_valueText_7o3zr_40";
const s$d = {
  root: root$5,
  grid,
  axis,
  area,
  outline,
  dot,
  label: label$4,
  valueText
};
const COLOR_MAP = {
  orange: "var(--color-orange-500)",
  green: "var(--color-green-500)",
  red: "var(--color-red-500)",
  wine: "var(--color-wine-500)",
  neutral: "var(--color-neutral-500)"
};
function vertexPosition(index, total, radius, cx, cy) {
  const angle = 2 * Math.PI * index / total - Math.PI / 2;
  return [cx + radius * Math.cos(angle), cy + radius * Math.sin(angle)];
}
function polygonPoints(total, radius, cx, cy) {
  return Array.from({ length: total }).map((_, i) => vertexPosition(i, total, radius, cx, cy).join(",")).join(" ");
}
function Radar({
  data,
  maxValue,
  size = 200,
  color = "orange",
  showValues = false,
  levels = 4,
  className
}) {
  const n = data.length;
  if (n < 3) return null;
  const resolvedMax = maxValue ?? Math.max(...data.map((d) => d.value));
  const effectiveMax = resolvedMax > 0 ? resolvedMax : 1;
  const padding = 40;
  const svgSize = size + padding * 2;
  const cx = svgSize / 2;
  const cy = svgSize / 2;
  const radius = size / 2;
  const dotRadius = 3;
  const fillColor = COLOR_MAP[color] ?? COLOR_MAP.orange;
  const strokeColor = fillColor;
  const dataPoints = data.map((d, i) => {
    const r = Math.min(d.value, effectiveMax) / effectiveMax * radius;
    return vertexPosition(i, n, r, cx, cy);
  });
  const dataPolygon = dataPoints.map((p) => p.join(",")).join(" ");
  const ariaLabel = `Gráfico radar com ${n} eixos: ${data.map((d) => `${d.label} ${d.value}`).join(", ")}`;
  return /* @__PURE__ */ jsxs(
    "svg",
    {
      role: "img",
      "aria-label": ariaLabel,
      width: svgSize,
      height: svgSize,
      viewBox: `0 0 ${svgSize} ${svgSize}`,
      className: [s$d.root, className].filter(Boolean).join(" "),
      children: [
        Array.from({ length: levels }).map((_, l) => {
          const r = radius / levels * (l + 1);
          return /* @__PURE__ */ jsx(
            "polygon",
            {
              className: s$d.grid,
              points: polygonPoints(n, r, cx, cy)
            },
            `grid-${l}`
          );
        }),
        Array.from({ length: n }).map((_, i) => {
          const [x, y] = vertexPosition(i, n, radius, cx, cy);
          return /* @__PURE__ */ jsx(
            "line",
            {
              className: s$d.axis,
              x1: cx,
              y1: cy,
              x2: x,
              y2: y
            },
            `axis-${i}`
          );
        }),
        /* @__PURE__ */ jsx(
          "polygon",
          {
            className: s$d.area,
            points: dataPolygon,
            fill: fillColor
          }
        ),
        /* @__PURE__ */ jsx(
          "polygon",
          {
            className: s$d.outline,
            points: dataPolygon,
            stroke: strokeColor
          }
        ),
        dataPoints.map(([x, y], i) => /* @__PURE__ */ jsx(
          "circle",
          {
            className: s$d.dot,
            cx: x,
            cy: y,
            r: dotRadius,
            fill: fillColor
          },
          `dot-${i}`
        )),
        data.map((d, i) => {
          const labelOffset = 16;
          const [lx, ly] = vertexPosition(i, n, radius + labelOffset, cx, cy);
          const angle = 2 * Math.PI * i / n - Math.PI / 2;
          const cos = Math.cos(angle);
          let anchor2 = "middle";
          if (cos > 0.01) anchor2 = "start";
          if (cos < -0.01) anchor2 = "end";
          const sin = Math.sin(angle);
          let dy = "0.35em";
          if (sin < -0.5) dy = "0em";
          if (sin > 0.5) dy = "0.7em";
          return /* @__PURE__ */ jsxs("g", { children: [
            /* @__PURE__ */ jsx(
              "text",
              {
                className: s$d.label,
                x: lx,
                y: ly,
                textAnchor: anchor2,
                dy,
                children: d.label
              }
            ),
            showValues && /* @__PURE__ */ jsx(
              "text",
              {
                className: s$d.valueText,
                x: lx,
                y: ly + 14,
                textAnchor: anchor2,
                dy,
                children: d.value
              }
            )
          ] }, `label-${i}`);
        })
      ]
    }
  );
}
function RowActionsPopover({
  items: items2,
  open: open2,
  onToggle,
  onClose,
  className,
  buttonAriaLabel = "Abrir ações"
}) {
  const triggerRef = useRef(null);
  return /* @__PURE__ */ jsxs("div", { className, children: [
    /* @__PURE__ */ jsx(
      Button,
      {
        ref: triggerRef,
        variant: "secondary",
        size: "md",
        leftIcon: DotsThreeVertical,
        "aria-label": buttonAriaLabel,
        onClick: onToggle
      }
    ),
    /* @__PURE__ */ jsx(
      Popover,
      {
        items: items2,
        open: open2,
        onClose,
        anchorRef: triggerRef
      }
    )
  ] });
}
const wrapper$4 = "_wrapper_j97f8_2";
const label$3 = "_label_j97f8_10";
const anchor = "_anchor_j97f8_19";
const trigger = "_trigger_j97f8_24";
const sm$2 = "_sm_j97f8_40";
const md = "_md_j97f8_45";
const lg = "_lg_j97f8_50";
const disabled$5 = "_disabled_j97f8_55";
const hovered$1 = "_hovered_j97f8_56";
const open = "_open_j97f8_61";
const focused$1 = "_focused_j97f8_62";
const error$2 = "_error_j97f8_68";
const value = "_value_j97f8_93";
const placeholder = "_placeholder_j97f8_106";
const caret = "_caret_j97f8_134";
const caretOpen = "_caretOpen_j97f8_139";
const dropdown = "_dropdown_j97f8_144";
const searchBox = "_searchBox_j97f8_156";
const searchInput = "_searchInput_j97f8_165";
const optionList = "_optionList_j97f8_183";
const empty = "_empty_j97f8_194";
const option = "_option_j97f8_183";
const selected$2 = "_selected_j97f8_228";
const message$1 = "_message_j97f8_233";
const attention$1 = "_attention_j97f8_247";
const success$2 = "_success_j97f8_251";
const s$c = {
  wrapper: wrapper$4,
  label: label$3,
  anchor,
  trigger,
  sm: sm$2,
  md,
  lg,
  disabled: disabled$5,
  hovered: hovered$1,
  open,
  focused: focused$1,
  error: error$2,
  value,
  placeholder,
  caret,
  caretOpen,
  dropdown,
  searchBox,
  searchInput,
  optionList,
  empty,
  option,
  selected: selected$2,
  message: message$1,
  attention: attention$1,
  success: success$2
};
const ICON_SIZES = { sm: 14, md: 16, lg: 20 };
const messageIconMap$1 = {
  error: WarningCircle,
  attention: WarningCircle,
  success: CheckCircle
};
function Select(props) {
  const {
    size = "md",
    label: label2,
    leftIcon: LeftIcon,
    placeholder: placeholder2 = "Selecione...",
    options,
    searchable = false,
    searchPlaceholder = "Buscar...",
    multiple = false,
    message: message2,
    messageType,
    disabled: disabled2 = false,
    className
  } = props;
  const iconSize2 = ICON_SIZES[size];
  const isSingleControlled = !multiple && props.value !== void 0;
  const [singleInternal, setSingleInternal] = useState(
    (!multiple ? props.defaultValue : "") ?? ""
  );
  const singleValue = isSingleControlled ? props.value : singleInternal;
  const isMultiControlled = multiple && props.value !== void 0;
  const [multiInternal, setMultiInternal] = useState(
    (multiple ? props.defaultValue : []) ?? []
  );
  const multiValue = isMultiControlled ? props.value : multiInternal;
  const [open2, setOpen] = useState(false);
  const [focusedIndex, setFocusedIndex] = useState(-1);
  const [search, setSearch] = useState("");
  const [dropdownStyle, setDropdownStyle] = useState({});
  const wrapperRef = useRef(null);
  const dropdownRef = useRef(null);
  const listRef = useRef(null);
  const searchRef = useRef(null);
  const triggerRef = useRef(null);
  const triggerId = useId();
  const listId = useId();
  const messageId = useId();
  const filtered = useMemo(() => {
    if (!search) return options;
    const q = search.toLowerCase();
    return options.filter((o) => o.label.toLowerCase().includes(q));
  }, [options, search]);
  const triggerText = useMemo(() => {
    if (multiple) {
      if (multiValue.length === 0) return null;
      if (multiValue.length === 1) {
        return options.find((o) => o.value === multiValue[0])?.label;
      }
      return `${multiValue.length} selecionados`;
    }
    return options.find((o) => o.value === singleValue)?.label ?? null;
  }, [multiple, multiValue, singleValue, options]);
  const hasMessage = !!message2 && !!messageType;
  const isError = messageType === "error";
  const MsgIcon = messageType ? messageIconMap$1[messageType] : null;
  const wrapperClasses = [s$c.wrapper, className ?? ""].filter(Boolean).join(" ");
  const triggerClasses = [
    s$c.trigger,
    s$c[size],
    isError ? s$c.error : "",
    disabled2 ? s$c.disabled : "",
    open2 ? s$c.open : ""
  ].filter(Boolean).join(" ");
  const selectSingle = useCallback(
    (val) => {
      if (!isSingleControlled) setSingleInternal(val);
      props.onChange?.(val);
      setOpen(false);
      setSearch("");
    },
    [isSingleControlled, props]
  );
  const toggleMulti = useCallback(
    (val) => {
      const next2 = multiValue.includes(val) ? multiValue.filter((v) => v !== val) : [...multiValue, val];
      if (!isMultiControlled) setMultiInternal(next2);
      props.onChange?.(next2);
    },
    [multiValue, isMultiControlled, props]
  );
  const updatePosition = useCallback(() => {
    const trigger2 = triggerRef.current;
    if (!trigger2) return;
    const rect = trigger2.getBoundingClientRect();
    const margin = 8;
    const gap = 4;
    const spaceBelow = window.innerHeight - rect.bottom - 8;
    setDropdownStyle({
      position: "fixed",
      top: rect.bottom + gap,
      left: clampToViewport({
        value: rect.left,
        size: rect.width,
        viewportSize: window.innerWidth,
        margin
      }),
      width: rect.width,
      maxHeight: Math.max(spaceBelow, 120)
    });
  }, []);
  const openDropdown = useCallback(() => {
    updatePosition();
    setOpen(true);
    setSearch("");
    if (!multiple) {
      const idx = filtered.findIndex((o) => o.value === singleValue);
      setFocusedIndex(idx >= 0 ? idx : 0);
    } else {
      setFocusedIndex(0);
    }
  }, [multiple, filtered, singleValue, updatePosition]);
  const closeDropdown = useCallback(() => {
    setOpen(false);
    setSearch("");
  }, []);
  useOpenFocus({
    active: open2,
    containerRef: dropdownRef,
    initialFocusRef: searchable ? searchRef : void 0
  });
  useInitialReposition(open2, updatePosition);
  useViewportReposition(open2, updatePosition);
  useDocumentClickOutside({
    active: open2,
    refs: [wrapperRef, dropdownRef],
    onOutside: closeDropdown
  });
  useDocumentEscape(open2, closeDropdown);
  useEffect(() => {
    if (!open2 || focusedIndex < 0) return;
    const list2 = listRef.current;
    if (!list2) return;
    const items2 = list2.querySelectorAll("[role='option']");
    const item2 = items2[focusedIndex];
    if (!item2) return;
    const listTop = list2.scrollTop;
    const listBottom = listTop + list2.clientHeight;
    const itemTop = item2.offsetTop;
    const itemBottom = itemTop + item2.offsetHeight;
    if (itemTop < listTop) {
      list2.scrollTop = itemTop;
    } else if (itemBottom > listBottom) {
      list2.scrollTop = itemBottom - list2.clientHeight;
    }
  }, [open2, focusedIndex]);
  useEffect(() => {
    if (open2) setFocusedIndex(filtered.length > 0 ? 0 : -1);
  }, [filtered, open2]);
  const handleListKeyDown = (e) => {
    switch (e.key) {
      case "ArrowDown":
        e.preventDefault();
        if (!open2) {
          openDropdown();
        } else {
          setFocusedIndex((i) => i < filtered.length - 1 ? i + 1 : i);
        }
        break;
      case "ArrowUp":
        e.preventDefault();
        if (open2) {
          setFocusedIndex((i) => i > 0 ? i - 1 : i);
        }
        break;
      case "Enter":
        e.preventDefault();
        if (open2 && focusedIndex >= 0 && filtered[focusedIndex]) {
          if (multiple) {
            toggleMulti(filtered[focusedIndex].value);
          } else {
            selectSingle(filtered[focusedIndex].value);
          }
        } else if (!open2) {
          openDropdown();
        }
        break;
      case "Home":
        if (open2) {
          e.preventDefault();
          setFocusedIndex(filtered.length > 0 ? 0 : -1);
        }
        break;
      case "End":
        if (open2) {
          e.preventDefault();
          setFocusedIndex(filtered.length > 0 ? filtered.length - 1 : -1);
        }
        break;
      case " ":
        if (!searchable || !open2) {
          e.preventDefault();
          if (!open2) openDropdown();
          else if (focusedIndex >= 0 && filtered[focusedIndex]) {
            if (multiple) {
              toggleMulti(filtered[focusedIndex].value);
            } else {
              selectSingle(filtered[focusedIndex].value);
            }
          }
        }
        break;
      case "Escape":
        e.preventDefault();
        closeDropdown();
        break;
      case "Tab":
        if (open2) {
          setOpen(false);
          setSearch("");
          triggerRef.current?.focus();
        }
        break;
    }
  };
  const handleTriggerClick = () => {
    if (disabled2) return;
    if (open2) {
      closeDropdown();
    } else {
      openDropdown();
    }
  };
  const handleOptionClick = (opt, e) => {
    e.preventDefault();
    if (multiple) {
      toggleMulti(opt.value);
    } else {
      selectSingle(opt.value);
    }
  };
  return /* @__PURE__ */ jsxs("div", { className: wrapperClasses, ref: wrapperRef, children: [
    label2 && /* @__PURE__ */ jsx("label", { className: s$c.label, htmlFor: triggerId, children: label2 }),
    /* @__PURE__ */ jsx("div", { className: s$c.anchor, children: /* @__PURE__ */ jsxs(
      "button",
      {
        ref: triggerRef,
        type: "button",
        id: triggerId,
        className: triggerClasses,
        onClick: handleTriggerClick,
        onKeyDown: !searchable || !open2 ? handleListKeyDown : void 0,
        disabled: disabled2,
        "aria-haspopup": "listbox",
        "aria-expanded": open2,
        "aria-controls": open2 ? listId : void 0,
        "aria-activedescendant": open2 && focusedIndex >= 0 ? `${listId}-opt-${focusedIndex}` : void 0,
        "aria-describedby": hasMessage ? messageId : void 0,
        "aria-invalid": isError || void 0,
        children: [
          LeftIcon && /* @__PURE__ */ jsx(LeftIcon, { size: iconSize2 }),
          /* @__PURE__ */ jsx("span", { className: triggerText ? s$c.value : s$c.placeholder, children: triggerText ?? placeholder2 }),
          /* @__PURE__ */ jsx(
            CaretDown,
            {
              size: iconSize2,
              className: `${s$c.caret} ${open2 ? s$c.caretOpen : ""}`
            }
          )
        ]
      }
    ) }),
    open2 && createPortal(
      /* @__PURE__ */ jsxs(
        "div",
        {
          ref: dropdownRef,
          className: s$c.dropdown,
          style: dropdownStyle,
          onMouseDown: (e) => e.preventDefault(),
          children: [
            searchable && /* @__PURE__ */ jsxs("div", { className: s$c.searchBox, children: [
              /* @__PURE__ */ jsx(MagnifyingGlass, { size: 16 }),
              /* @__PURE__ */ jsx(
                "input",
                {
                  ref: searchRef,
                  type: "text",
                  className: s$c.searchInput,
                  placeholder: searchPlaceholder,
                  value: search,
                  onChange: (e) => setSearch(e.target.value),
                  onKeyDown: handleListKeyDown,
                  "aria-label": "Buscar opções",
                  "aria-activedescendant": focusedIndex >= 0 ? `${listId}-opt-${focusedIndex}` : void 0
                }
              )
            ] }),
            /* @__PURE__ */ jsxs(
              "ul",
              {
                id: listId,
                ref: listRef,
                className: s$c.optionList,
                role: "listbox",
                "aria-labelledby": triggerId,
                "aria-multiselectable": multiple || void 0,
                children: [
                  filtered.length === 0 && /* @__PURE__ */ jsx("li", { className: s$c.empty, role: "status", "aria-live": "polite", children: "Nenhum resultado" }),
                  filtered.map((opt, i) => {
                    const isSelected = multiple ? multiValue.includes(opt.value) : opt.value === singleValue;
                    return /* @__PURE__ */ jsx(
                      "li",
                      {
                        id: `${listId}-opt-${i}`,
                        className: `${s$c.option} ${isSelected && !multiple ? s$c.selected : ""} ${i === focusedIndex ? s$c.focused : ""}`,
                        role: "option",
                        "aria-selected": isSelected,
                        onMouseEnter: () => setFocusedIndex(i),
                        onMouseDown: (e) => handleOptionClick(opt, e),
                        children: multiple ? /* @__PURE__ */ jsx(
                          Checkbox,
                          {
                            size: "sm",
                            checked: isSelected,
                            label: opt.label,
                            readOnly: true,
                            tabIndex: -1
                          }
                        ) : opt.label
                      },
                      opt.value
                    );
                  })
                ]
              }
            )
          ]
        }
      ),
      document.body
    ),
    hasMessage && /* @__PURE__ */ jsxs("div", { id: messageId, className: `${s$c.message} ${s$c[messageType]}`, children: [
      MsgIcon && /* @__PURE__ */ jsx(MsgIcon, { size: 14 }),
      /* @__PURE__ */ jsx("span", { children: message2 })
    ] })
  ] });
}
const root$4 = "_root_1adx1_2";
const disabled$4 = "_disabled_1adx1_8";
const buttons = "_buttons_1adx1_14";
const btn = "_btn_1adx1_20";
const selected$1 = "_selected_1adx1_50";
const btnNumber = "_btnNumber_1adx1_62";
const btnLabel = "_btnLabel_1adx1_66";
const labelsDesktop = "_labelsDesktop_1adx1_72";
const label$2 = "_label_1adx1_72";
const sm$1 = "_sm_1adx1_86";
const s$b = {
  root: root$4,
  disabled: disabled$4,
  buttons,
  btn,
  selected: selected$1,
  btnNumber,
  btnLabel,
  labelsDesktop,
  label: label$2,
  sm: sm$1
};
function ScaleInput({
  min = 0,
  max = 10,
  value: value2,
  onChange,
  minLabel,
  maxLabel,
  disabled: disabled2 = false,
  size = "md",
  className
}) {
  const uid2 = useId();
  const containerRef = useRef(null);
  const count = max - min + 1;
  const values = Array.from({ length: count }, (_, i) => min + i);
  const handleSelect = useCallback(
    (v) => {
      if (disabled2) return;
      onChange?.(v);
    },
    [disabled2, onChange]
  );
  const handleKeyDown = useCallback(
    (e, currentValue) => {
      if (disabled2) return;
      let nextValue = null;
      switch (e.key) {
        case "ArrowRight":
        case "ArrowDown":
          e.preventDefault();
          nextValue = currentValue < max ? currentValue + 1 : min;
          break;
        case "ArrowLeft":
        case "ArrowUp":
          e.preventDefault();
          nextValue = currentValue > min ? currentValue - 1 : max;
          break;
        case "Home":
          e.preventDefault();
          nextValue = min;
          break;
        case "End":
          e.preventDefault();
          nextValue = max;
          break;
      }
      if (nextValue !== null) {
        onChange?.(nextValue);
        const btn2 = containerRef.current?.querySelector(
          `[data-value="${nextValue}"]`
        );
        btn2?.focus();
      }
    },
    [disabled2, onChange, min, max]
  );
  const rootClasses = [
    s$b.root,
    size === "sm" ? s$b.sm : "",
    disabled2 ? s$b.disabled : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  const hasLabels = !!(minLabel || maxLabel);
  return /* @__PURE__ */ jsxs("div", { className: rootClasses, children: [
    hasLabels && /* @__PURE__ */ jsxs("div", { className: s$b.labelsDesktop, children: [
      /* @__PURE__ */ jsx("span", { className: s$b.label, children: minLabel }),
      /* @__PURE__ */ jsx("span", { className: s$b.label, children: maxLabel })
    ] }),
    /* @__PURE__ */ jsx(
      "div",
      {
        ref: containerRef,
        className: s$b.buttons,
        role: "radiogroup",
        "aria-label": "Escala",
        children: values.map((v, i) => {
          const isSelected = value2 === v;
          const isFirst = i === 0;
          const isLast = i === values.length - 1;
          const tabIndex = value2 !== void 0 ? isSelected ? 0 : -1 : i === 0 ? 0 : -1;
          const inlineLabel = isFirst && minLabel ? minLabel : isLast && maxLabel ? maxLabel : null;
          return /* @__PURE__ */ jsxs(
            "button",
            {
              type: "button",
              role: "radio",
              "aria-checked": isSelected,
              "aria-label": inlineLabel ? `${v} — ${inlineLabel}` : String(v),
              id: `${uid2}-${v}`,
              "data-value": v,
              tabIndex,
              className: `${s$b.btn} ${isSelected ? s$b.selected : ""}`,
              disabled: disabled2,
              onClick: () => handleSelect(v),
              onKeyDown: (e) => handleKeyDown(e, v),
              children: [
                /* @__PURE__ */ jsx("span", { className: s$b.btnNumber, children: v }),
                inlineLabel && /* @__PURE__ */ jsx("span", { className: s$b.btnLabel, children: inlineLabel })
              ]
            },
            v
          );
        })
      }
    )
  ] });
}
const skeleton = "_skeleton_p77d1_3";
const shimmer = "_shimmer_p77d1_1";
const text$1 = "_text_p77d1_30";
const circular = "_circular_p77d1_35";
const rectangular = "_rectangular_p77d1_39";
const rounded = "_rounded_p77d1_43";
const noAnimation = "_noAnimation_p77d1_49";
const imagePlaceholder = "_imagePlaceholder_p77d1_63";
const videoPlaceholder = "_videoPlaceholder_p77d1_64";
const chartPlaceholder = "_chartPlaceholder_p77d1_92";
const chartBar = "_chartBar_p77d1_105";
const testimonialPlaceholder = "_testimonialPlaceholder_p77d1_123";
const testimonialAvatar = "_testimonialAvatar_p77d1_131";
const s$a = {
  skeleton,
  shimmer,
  text: text$1,
  circular,
  rectangular,
  rounded,
  noAnimation,
  imagePlaceholder,
  videoPlaceholder,
  chartPlaceholder,
  chartBar,
  testimonialPlaceholder,
  testimonialAvatar
};
const SKELETON_HEIGHTS = {
  text: 14,
  heading: 24,
  subheading: 18,
  button: 40,
  input: 40,
  avatar: 40,
  avatarLg: 48
};
function SkeletonContainer({
  children,
  loadingText = "Carregando...",
  className
}) {
  return /* @__PURE__ */ jsxs("div", { role: "status", className, children: [
    children,
    /* @__PURE__ */ jsx("span", { className: "sr-only", children: loadingText })
  ] });
}
function Skeleton({
  variant = "rectangular",
  width,
  height,
  animation = true,
  className
}) {
  const style = {};
  if (width !== void 0) {
    style.width = typeof width === "number" ? `${width}px` : width;
  }
  if (height !== void 0) {
    style.height = typeof height === "number" ? `${height}px` : height;
  }
  if (variant === "circular") {
    if (!style.width) style.width = "40px";
    if (!style.height) style.height = style.width;
  }
  const cls = [
    s$a.skeleton,
    s$a[variant],
    !animation && s$a.noAnimation,
    className
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx(
    "span",
    {
      className: cls,
      style,
      "aria-hidden": "true"
    }
  );
}
const list = "_list_pimja_3";
const disabled$3 = "_disabled_pimja_9";
const item$1 = "_item_pimja_16";
const dragging = "_dragging_pimja_30";
const over = "_over_pimja_53";
const handle = "_handle_pimja_61";
const number$1 = "_number_pimja_70";
const label$1 = "_label_pimja_82";
const moveButtons = "_moveButtons_pimja_92";
const moveBtn = "_moveBtn_pimja_98";
const sm = "_sm_pimja_128";
const s$9 = {
  list,
  disabled: disabled$3,
  item: item$1,
  dragging,
  over,
  handle,
  number: number$1,
  label: label$1,
  moveButtons,
  moveBtn,
  sm
};
function reorder(items2, from, to) {
  const arr = [...items2];
  const [moved] = arr.splice(from, 1);
  arr.splice(to, 0, moved);
  return arr;
}
function SortableList({
  items: items2,
  onChange,
  disabled: disabled2 = false,
  size = "md",
  className
}) {
  const listRef = useRef(null);
  const [dragIndex, setDragIndex] = useState(null);
  const [overIndex, setOverIndex] = useState(null);
  const dragCounterRef = useRef(0);
  const handleDragStart = useCallback(
    (e, index) => {
      if (disabled2) {
        e.preventDefault();
        return;
      }
      e.dataTransfer.effectAllowed = "move";
      e.dataTransfer.setData("text/plain", String(index));
      setDragIndex(index);
    },
    [disabled2]
  );
  const handleDragEnd = useCallback(() => {
    setDragIndex(null);
    setOverIndex(null);
    dragCounterRef.current = 0;
  }, []);
  const handleDragEnter = useCallback(
    (e, index) => {
      e.preventDefault();
      dragCounterRef.current++;
      setOverIndex(index);
    },
    []
  );
  const handleDragLeave = useCallback(() => {
    dragCounterRef.current--;
    if (dragCounterRef.current <= 0) {
      setOverIndex(null);
      dragCounterRef.current = 0;
    }
  }, []);
  const handleDragOver = useCallback((e) => {
    e.preventDefault();
    e.dataTransfer.dropEffect = "move";
  }, []);
  const handleDrop = useCallback(
    (e, dropIndex) => {
      e.preventDefault();
      const fromIndex = Number(e.dataTransfer.getData("text/plain"));
      if (isNaN(fromIndex) || fromIndex === dropIndex) {
        handleDragEnd();
        return;
      }
      onChange?.(reorder(items2, fromIndex, dropIndex));
      handleDragEnd();
    },
    [items2, onChange, handleDragEnd]
  );
  const touchState = useRef(null);
  const [touchDragIndex, setTouchDragIndex] = useState(null);
  const [touchOverIndex, setTouchOverIndex] = useState(null);
  const handleTouchStart = useCallback(
    (e, index) => {
      if (disabled2) return;
      const touch = e.touches[0];
      const itemEls = listRef.current?.querySelectorAll(`.${s$9.item}`);
      const rects = itemEls ? Array.from(itemEls).map((el) => el.getBoundingClientRect()) : [];
      touchState.current = {
        index,
        startY: touch.clientY,
        currentY: touch.clientY,
        itemRects: rects
      };
      setTouchDragIndex(index);
      setTouchOverIndex(index);
    },
    [disabled2]
  );
  useEffect(() => {
    function handleTouchMove(e) {
      if (!touchState.current) return;
      e.preventDefault();
      const touch = e.touches[0];
      touchState.current.currentY = touch.clientY;
      const y = touch.clientY;
      const { itemRects } = touchState.current;
      let target = touchState.current.index;
      for (let i = 0; i < itemRects.length; i++) {
        const rect = itemRects[i];
        const midY = rect.top + rect.height / 2;
        if (y < midY) {
          target = i;
          break;
        }
        target = i;
      }
      setTouchOverIndex(target);
    }
    function handleTouchEnd() {
      if (!touchState.current) return;
      const fromIndex = touchState.current.index;
      const toIndex = touchOverIndexRef.current !== null ? touchOverIndexRef.current : fromIndex;
      if (fromIndex !== toIndex) {
        onChange?.(reorder(items2, fromIndex, toIndex));
      }
      touchState.current = null;
      setTouchDragIndex(null);
      setTouchOverIndex(null);
    }
    document.addEventListener("touchmove", handleTouchMove, { passive: false });
    document.addEventListener("touchend", handleTouchEnd);
    document.addEventListener("touchcancel", handleTouchEnd);
    return () => {
      document.removeEventListener("touchmove", handleTouchMove);
      document.removeEventListener("touchend", handleTouchEnd);
      document.removeEventListener("touchcancel", handleTouchEnd);
    };
  }, [items2, onChange]);
  const touchOverIndexRef = useRef(null);
  touchOverIndexRef.current = touchOverIndex;
  const handleMoveUp = useCallback(
    (index) => {
      if (disabled2 || index <= 0) return;
      onChange?.(reorder(items2, index, index - 1));
    },
    [disabled2, items2, onChange]
  );
  const handleMoveDown = useCallback(
    (index) => {
      if (disabled2 || index >= items2.length - 1) return;
      onChange?.(reorder(items2, index, index + 1));
    },
    [disabled2, items2, onChange]
  );
  const handleKeyDown = useCallback(
    (e, index) => {
      if (disabled2) return;
      let targetIndex = null;
      if (e.altKey && e.key === "ArrowUp" && index > 0) {
        e.preventDefault();
        targetIndex = index - 1;
      } else if (e.altKey && e.key === "ArrowDown" && index < items2.length - 1) {
        e.preventDefault();
        targetIndex = index + 1;
      }
      if (targetIndex !== null) {
        onChange?.(reorder(items2, index, targetIndex));
        requestAnimationFrame(() => {
          const target = listRef.current?.querySelectorAll(
            `.${s$9.item}`
          )[targetIndex];
          target?.focus();
        });
      }
    },
    [disabled2, items2, onChange]
  );
  const rootClasses = [
    s$9.list,
    size === "sm" ? s$9.sm : "",
    disabled2 ? s$9.disabled : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  const activeDragIndex = dragIndex ?? touchDragIndex;
  const activeOverIndex = overIndex ?? touchOverIndex;
  return /* @__PURE__ */ jsx("div", { ref: listRef, className: rootClasses, role: "list", "aria-label": "Lista reordenável", children: items2.map((item2, index) => {
    const isDragging = activeDragIndex === index;
    const isOver = activeOverIndex === index && activeDragIndex !== index;
    const itemClasses = [
      s$9.item,
      isDragging ? s$9.dragging : "",
      isOver ? s$9.over : ""
    ].filter(Boolean).join(" ");
    return /* @__PURE__ */ jsxs(
      "div",
      {
        role: "listitem",
        className: itemClasses,
        draggable: !disabled2,
        tabIndex: 0,
        "aria-label": `${index + 1}. ${item2.label}. Use Alt+Setas para reordenar.`,
        onDragStart: (e) => handleDragStart(e, index),
        onDragEnd: handleDragEnd,
        onDragEnter: (e) => handleDragEnter(e, index),
        onDragLeave: handleDragLeave,
        onDragOver: handleDragOver,
        onDrop: (e) => handleDrop(e, index),
        onTouchStart: (e) => handleTouchStart(e, index),
        onKeyDown: (e) => handleKeyDown(e, index),
        children: [
          /* @__PURE__ */ jsx("span", { className: s$9.handle, "aria-hidden": true, children: /* @__PURE__ */ jsx(DotsSixVertical, { size: size === "sm" ? 16 : 20 }) }),
          /* @__PURE__ */ jsx("span", { className: s$9.number, children: index + 1 }),
          /* @__PURE__ */ jsx("span", { className: s$9.label, children: item2.label }),
          /* @__PURE__ */ jsxs("span", { className: s$9.moveButtons, children: [
            /* @__PURE__ */ jsx(
              "button",
              {
                type: "button",
                className: s$9.moveBtn,
                "aria-label": "Mover para cima",
                disabled: disabled2 || index === 0,
                onClick: (e) => {
                  e.stopPropagation();
                  handleMoveUp(index);
                },
                tabIndex: -1,
                children: /* @__PURE__ */ jsx(ArrowUp, { size: 14 })
              }
            ),
            /* @__PURE__ */ jsx(
              "button",
              {
                type: "button",
                className: s$9.moveBtn,
                "aria-label": "Mover para baixo",
                disabled: disabled2 || index === items2.length - 1,
                onClick: (e) => {
                  e.stopPropagation();
                  handleMoveDown(index);
                },
                tabIndex: -1,
                children: /* @__PURE__ */ jsx(ArrowDown, { size: 14 })
              }
            )
          ] })
        ]
      },
      item2.id
    );
  }) });
}
const root$3 = "_root_10dxi_1";
const line = "_line_10dxi_6";
const fill = "_fill_10dxi_13";
const orange = "_orange_10dxi_20";
const green = "_green_10dxi_27";
const red = "_red_10dxi_34";
const neutral$1 = "_neutral_10dxi_41";
const s$8 = {
  root: root$3,
  line,
  fill,
  orange,
  green,
  red,
  neutral: neutral$1
};
function getTrendLabel(data) {
  if (data.length < 2) return "sem dados suficientes";
  const first = data[0];
  const last = data[data.length - 1];
  if (last > first) return "tendência de alta";
  if (last < first) return "tendência de queda";
  return "tendência estável";
}
function Sparkline({
  data,
  width = 80,
  height = 24,
  color = "orange",
  filled = false,
  className
}) {
  if (data.length < 2) return null;
  const min = Math.min(...data);
  const max = Math.max(...data);
  const range = max - min || 1;
  const padY = 1.5;
  const innerH = height - padY * 2;
  const points = data.map((value2, i) => {
    const x = i / (data.length - 1) * width;
    const y = padY + innerH - (value2 - min) / range * innerH;
    return `${x},${y}`;
  });
  const polylinePoints = points.join(" ");
  const polygonPoints2 = filled ? `${polylinePoints} ${width},${height} 0,${height}` : void 0;
  const classes = [s$8.root, s$8[color], className].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsxs(
    "svg",
    {
      className: classes,
      width,
      height,
      viewBox: `0 0 ${width} ${height}`,
      role: "img",
      "aria-label": `Sparkline: ${getTrendLabel(data)}`,
      children: [
        filled && polygonPoints2 && /* @__PURE__ */ jsx("polygon", { className: s$8.fill, points: polygonPoints2 }),
        /* @__PURE__ */ jsx(
          "polyline",
          {
            className: s$8.line,
            points: polylinePoints,
            strokeWidth: 1.5
          }
        )
      ]
    }
  );
}
const wrapper$3 = "_wrapper_73bor_3";
const tooltip = "_tooltip_73bor_9";
const tooltipIn = "_tooltipIn_73bor_1";
const arrow = "_arrow_73bor_38";
const top = "_top_73bor_46";
const bottom = "_bottom_73bor_50";
const left = "_left_73bor_54";
const right = "_right_73bor_58";
const s$7 = {
  wrapper: wrapper$3,
  tooltip,
  tooltipIn,
  arrow,
  top,
  bottom,
  left,
  right
};
const ARROW_SIZE = 6;
const GAP = 8;
const VIEWPORT_MARGIN = 8;
function Tooltip({
  content: content2,
  placement = "top",
  delay = 200,
  disabled: disabled2 = false,
  children
}) {
  const id = useId();
  const tooltipId = `tooltip-${id}`;
  const wrapperRef = useRef(null);
  const tooltipRef = useRef(null);
  const arrowRef = useRef(null);
  const timeoutRef = useRef(void 0);
  const [visible, setVisible] = useState(false);
  const [resolvedPlacement, setResolvedPlacement] = useState(placement);
  const show = useCallback(() => {
    if (disabled2) return;
    timeoutRef.current = setTimeout(() => setVisible(true), delay);
  }, [disabled2, delay]);
  const hide = useCallback(() => {
    clearTimeout(timeoutRef.current);
    setVisible(false);
  }, []);
  useEffect(() => {
    return () => clearTimeout(timeoutRef.current);
  }, []);
  const applyPosition = useCallback(() => {
    const anchor2 = wrapperRef.current;
    const tip = tooltipRef.current;
    const arrow2 = arrowRef.current;
    if (!anchor2 || !tip || !arrow2) return;
    const ar = anchor2.getBoundingClientRect();
    const tr = tip.getBoundingClientRect();
    const vw = window.innerWidth;
    const vh = window.innerHeight;
    const placements = [placement, "top", "bottom", "right", "left"];
    let chosen = placement;
    for (const p of placements) {
      if (checkFit(p, ar, tr)) {
        chosen = p;
        break;
      }
    }
    setResolvedPlacement(chosen);
    const anchorCx = ar.left + ar.width / 2;
    const anchorCy = ar.top + ar.height / 2;
    let top2;
    let left2;
    if (chosen === "top" || chosen === "bottom") {
      left2 = anchorCx - tr.width / 2;
      left2 = clampToViewport({
        value: left2,
        size: tr.width,
        viewportSize: vw,
        margin: VIEWPORT_MARGIN
      });
      top2 = chosen === "top" ? ar.top - tr.height - GAP : ar.bottom + GAP;
      const arrowLeft = clamp(anchorCx - left2, ARROW_SIZE + 4, tr.width - ARROW_SIZE - 4);
      arrow2.style.left = `${arrowLeft}px`;
      arrow2.style.top = "";
      arrow2.style.transform = `translateX(-50%) rotate(45deg)`;
    } else {
      top2 = anchorCy - tr.height / 2;
      top2 = clampToViewport({
        value: top2,
        size: tr.height,
        viewportSize: vh,
        margin: VIEWPORT_MARGIN
      });
      left2 = chosen === "left" ? ar.left - tr.width - GAP : ar.right + GAP;
      const arrowTop = clamp(anchorCy - top2, ARROW_SIZE + 4, tr.height - ARROW_SIZE - 4);
      arrow2.style.top = `${arrowTop}px`;
      arrow2.style.left = "";
      arrow2.style.transform = `translateY(-50%) rotate(45deg)`;
    }
    tip.style.top = `${top2}px`;
    tip.style.left = `${left2}px`;
  }, [placement]);
  useInitialReposition(visible, applyPosition);
  useViewportReposition(visible, applyPosition);
  useDocumentEscape(visible, hide);
  return /* @__PURE__ */ jsxs(Fragment, { children: [
    /* @__PURE__ */ jsx(
      "span",
      {
        ref: wrapperRef,
        className: s$7.wrapper,
        onMouseEnter: show,
        onMouseLeave: hide,
        onFocus: show,
        onBlur: hide,
        "aria-describedby": visible ? tooltipId : void 0,
        children
      }
    ),
    visible && createPortal(
      /* @__PURE__ */ jsxs(
        "div",
        {
          ref: tooltipRef,
          id: tooltipId,
          role: "tooltip",
          className: [s$7.tooltip, s$7[resolvedPlacement]].join(" "),
          children: [
            content2,
            /* @__PURE__ */ jsx("span", { ref: arrowRef, className: s$7.arrow })
          ]
        }
      ),
      document.body
    )
  ] });
}
function clamp(value2, min, max) {
  return Math.max(min, Math.min(value2, max));
}
function checkFit(p, ar, tr) {
  switch (p) {
    case "top":
      return ar.top - tr.height - GAP >= 0;
    case "bottom":
      return ar.bottom + tr.height + GAP <= window.innerHeight;
    case "left":
      return ar.left - tr.width - GAP >= 0;
    case "right":
      return ar.right + tr.width + GAP <= window.innerWidth;
  }
}
const root$2 = "_root_zyflz_3";
const rootCollapsed = "_rootCollapsed_zyflz_16";
const collapseHitArea = "_collapseHitArea_zyflz_22";
const collapseBtn = "_collapseBtn_zyflz_38";
const header$1 = "_header_zyflz_64";
const headerBrand = "_headerBrand_zyflz_73";
const orgWrapper = "_orgWrapper_zyflz_85";
const orgSwitcher = "_orgSwitcher_zyflz_90";
const orgIcon = "_orgIcon_zyflz_113";
const orgImage = "_orgImage_zyflz_125";
const orgLabel = "_orgLabel_zyflz_134";
const orgCaret = "_orgCaret_zyflz_148";
const divider = "_divider_zyflz_155";
const nav = "_nav_zyflz_163";
const navList = "_navList_zyflz_170";
const group = "_group_zyflz_178";
const groupLabel = "_groupLabel_zyflz_184";
const itemWrapper = "_itemWrapper_zyflz_196";
const item = "_item_zyflz_196";
const itemActive = "_itemActive_zyflz_237";
const itemLabel = "_itemLabel_zyflz_242";
const itemCaret = "_itemCaret_zyflz_250";
const itemCaretOpen = "_itemCaretOpen_zyflz_256";
const subItems = "_subItems_zyflz_262";
const subItemsOpen = "_subItemsOpen_zyflz_268";
const subItemsInner = "_subItemsInner_zyflz_272";
const subItem = "_subItem_zyflz_262";
const subItemActive = "_subItemActive_zyflz_317";
const flyout = "_flyout_zyflz_324";
const flyoutInner = "_flyoutInner_zyflz_335";
const flyoutLabel = "_flyoutLabel_zyflz_350";
const flyoutItems = "_flyoutItems_zyflz_360";
const footer = "_footer_zyflz_368";
const user = "_user_zyflz_377";
const userText = "_userText_zyflz_400";
const userName = "_userName_zyflz_408";
const userRole = "_userRole_zyflz_419";
const userCaret = "_userCaret_zyflz_430";
const overlay = "_overlay_zyflz_503";
const mobileClose = "_mobileClose_zyflz_509";
const rootMobileOpen = "_rootMobileOpen_zyflz_529";
const overlayVisible = "_overlayVisible_zyflz_627";
const s$6 = {
  root: root$2,
  rootCollapsed,
  collapseHitArea,
  collapseBtn,
  header: header$1,
  headerBrand,
  orgWrapper,
  orgSwitcher,
  orgIcon,
  orgImage,
  orgLabel,
  orgCaret,
  divider,
  nav,
  navList,
  group,
  groupLabel,
  itemWrapper,
  item,
  itemActive,
  itemLabel,
  itemCaret,
  itemCaretOpen,
  subItems,
  subItemsOpen,
  subItemsInner,
  subItem,
  subItemActive,
  flyout,
  flyoutInner,
  flyoutLabel,
  flyoutItems,
  footer,
  user,
  userText,
  userName,
  userRole,
  userCaret,
  overlay,
  mobileClose,
  rootMobileOpen,
  overlayVisible
};
const CollapsedContext = createContext(false);
const CollapseCallbackContext = createContext(void 0);
function Sidebar({
  children,
  className,
  collapsed = false,
  onCollapse,
  mobileOpen = false,
  onMobileClose,
  ...rest
}) {
  const asideRef = useRef(null);
  const closeRef = useRef(null);
  useBodyScrollLock(mobileOpen);
  const restoreFocus = useOpenFocus({
    active: mobileOpen,
    containerRef: asideRef,
    initialFocusRef: closeRef
  });
  const handleMobileClose = useCallback(() => {
    onMobileClose?.();
    restoreFocus();
  }, [onMobileClose, restoreFocus]);
  useDocumentEscape(mobileOpen && !!onMobileClose, handleMobileClose);
  const handleAsideKeyDown = useCallback(
    (e) => {
      if (!mobileOpen) return;
      trapFocusWithin(asideRef.current, e);
    },
    [mobileOpen]
  );
  const cls = [
    s$6.root,
    collapsed && s$6.rootCollapsed,
    mobileOpen && s$6.rootMobileOpen,
    className
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx(CollapsedContext.Provider, { value: collapsed, children: /* @__PURE__ */ jsxs(CollapseCallbackContext.Provider, { value: onCollapse, children: [
    /* @__PURE__ */ jsx(
      "div",
      {
        className: `${s$6.overlay}${mobileOpen ? ` ${s$6.overlayVisible}` : ""}`,
        onClick: handleMobileClose,
        "aria-hidden": "true"
      }
    ),
    /* @__PURE__ */ jsxs(
      "aside",
      {
        ref: asideRef,
        className: cls,
        "aria-label": "Menu lateral",
        role: mobileOpen ? "dialog" : void 0,
        "aria-modal": mobileOpen || void 0,
        onKeyDown: handleAsideKeyDown,
        ...rest,
        children: [
          onMobileClose && /* @__PURE__ */ jsx(
            "button",
            {
              ref: closeRef,
              type: "button",
              className: s$6.mobileClose,
              onClick: handleMobileClose,
              "aria-label": "Fechar menu",
              children: /* @__PURE__ */ jsx(X, { size: 20 })
            }
          ),
          children
        ]
      }
    )
  ] }) });
}
function SidebarHeader({ children, collapsedContent }) {
  const collapsed = useContext(CollapsedContext);
  const onCollapse = useContext(CollapseCallbackContext);
  function handleCollapseClick(e) {
    onCollapse?.();
    if (e.detail > 0) {
      const target = e.currentTarget;
      requestAnimationFrame(() => {
        target.blur();
      });
    }
  }
  return /* @__PURE__ */ jsxs("div", { className: s$6.header, children: [
    /* @__PURE__ */ jsx("div", { className: s$6.headerBrand, children: collapsed && collapsedContent ? collapsedContent : children }),
    onCollapse && /* @__PURE__ */ jsx("div", { className: s$6.collapseHitArea, children: /* @__PURE__ */ jsx(
      Tooltip,
      {
        content: collapsed ? "Expandir" : "Recolher",
        placement: "right",
        children: /* @__PURE__ */ jsx(
          "button",
          {
            type: "button",
            className: s$6.collapseBtn,
            onClick: handleCollapseClick,
            "aria-label": collapsed ? "Expandir sidebar" : "Recolher sidebar",
            children: collapsed ? /* @__PURE__ */ jsx(CaretLineRight, { size: 14 }) : /* @__PURE__ */ jsx(CaretLineLeft, { size: 14 })
          }
        )
      }
    ) })
  ] });
}
const SidebarOrgSwitcher = forwardRef(function SidebarOrgSwitcher2({ icon: Icon, image: image2, label: label2, onClick }, ref) {
  return /* @__PURE__ */ jsx("div", { className: s$6.orgWrapper, children: /* @__PURE__ */ jsxs(
    "button",
    {
      ref,
      type: "button",
      className: s$6.orgSwitcher,
      onClick,
      "aria-haspopup": "listbox",
      "aria-label": `Organização: ${label2}`,
      children: [
        image2 ? /* @__PURE__ */ jsx("img", { src: image2, alt: "", className: s$6.orgImage }) : Icon ? /* @__PURE__ */ jsx("span", { className: s$6.orgIcon, children: /* @__PURE__ */ jsx(Icon, { size: 16 }) }) : null,
        /* @__PURE__ */ jsx("span", { className: s$6.orgLabel, children: label2 }),
        /* @__PURE__ */ jsx(CaretDown, { size: 16, className: s$6.orgCaret })
      ]
    }
  ) });
});
function SidebarDivider() {
  return /* @__PURE__ */ jsx("hr", { className: s$6.divider });
}
function SidebarNav({
  children,
  "aria-label": ariaLabel = "Menu principal"
}) {
  return /* @__PURE__ */ jsx("nav", { className: s$6.nav, "aria-label": ariaLabel, children: /* @__PURE__ */ jsx("div", { className: s$6.navList, role: "list", children }) });
}
function SidebarGroup({ label: label2, children }) {
  return /* @__PURE__ */ jsxs("div", { className: s$6.group, role: "group", "aria-label": label2, children: [
    /* @__PURE__ */ jsx("span", { className: s$6.groupLabel, children: label2 }),
    children
  ] });
}
function SidebarItem({
  icon: Icon,
  label: label2,
  href,
  active: active2 = false,
  defaultExpanded = false,
  onClick,
  children
}) {
  const [expanded, setExpanded] = useState(defaultExpanded);
  const collapsed = useContext(CollapsedContext);
  const hasChildren = !!children;
  const panelId = useId();
  function handleClick() {
    if (hasChildren) {
      setExpanded((v) => !v);
    }
    onClick?.();
  }
  const cls = `${s$6.item}${active2 ? ` ${s$6.itemActive}` : ""}`;
  const trigger2 = href && !hasChildren ? /* @__PURE__ */ jsxs(
    "a",
    {
      href,
      className: cls,
      onClick,
      "aria-current": active2 ? "page" : void 0,
      "aria-label": collapsed ? label2 : void 0,
      children: [
        /* @__PURE__ */ jsx(Icon, { size: 16 }),
        /* @__PURE__ */ jsx("span", { className: s$6.itemLabel, children: label2 })
      ]
    }
  ) : /* @__PURE__ */ jsxs(
    "button",
    {
      type: "button",
      className: cls,
      onClick: collapsed ? void 0 : handleClick,
      "aria-current": active2 && !hasChildren ? "page" : void 0,
      "aria-expanded": hasChildren && !collapsed ? expanded : void 0,
      "aria-controls": hasChildren && !collapsed ? panelId : void 0,
      "aria-label": collapsed ? label2 : void 0,
      children: [
        /* @__PURE__ */ jsx(Icon, { size: 16 }),
        /* @__PURE__ */ jsx("span", { className: s$6.itemLabel, children: label2 }),
        hasChildren && /* @__PURE__ */ jsx(
          CaretDown,
          {
            size: 16,
            className: `${s$6.itemCaret}${expanded ? ` ${s$6.itemCaretOpen}` : ""}`
          }
        )
      ]
    }
  );
  return /* @__PURE__ */ jsxs("div", { className: s$6.itemWrapper, children: [
    trigger2,
    hasChildren && !collapsed && /* @__PURE__ */ jsx(
      "div",
      {
        id: panelId,
        className: `${s$6.subItems}${expanded ? ` ${s$6.subItemsOpen}` : ""}`,
        role: "group",
        "aria-label": label2,
        children: /* @__PURE__ */ jsx("div", { className: s$6.subItemsInner, children })
      }
    ),
    collapsed && hasChildren && /* @__PURE__ */ jsx("div", { className: s$6.flyout, children: /* @__PURE__ */ jsxs("div", { className: s$6.flyoutInner, role: "group", "aria-label": label2, children: [
      /* @__PURE__ */ jsx("span", { className: s$6.flyoutLabel, children: label2 }),
      /* @__PURE__ */ jsx("div", { className: s$6.flyoutItems, children })
    ] }) })
  ] });
}
function SidebarSubItem({
  href,
  active: active2 = false,
  onClick,
  children
}) {
  const cls = `${s$6.subItem}${active2 ? ` ${s$6.subItemActive}` : ""}`;
  return href ? /* @__PURE__ */ jsx(
    "a",
    {
      href,
      className: cls,
      onClick,
      "aria-current": active2 ? "page" : void 0,
      children
    }
  ) : /* @__PURE__ */ jsx(
    "button",
    {
      type: "button",
      className: cls,
      onClick,
      "aria-current": active2 ? "page" : void 0,
      children
    }
  );
}
function SidebarFooter({ children }) {
  return /* @__PURE__ */ jsx("div", { className: s$6.footer, children });
}
const SidebarUser = forwardRef(
  function SidebarUser2({ name: name2, role, avatar: avatar2, onClick }, ref) {
    return /* @__PURE__ */ jsxs(
      "button",
      {
        ref,
        type: "button",
        className: s$6.user,
        onClick,
        "aria-label": `Menu do usuário: ${name2}`,
        "aria-haspopup": "menu",
        children: [
          avatar2,
          /* @__PURE__ */ jsxs("div", { className: s$6.userText, children: [
            /* @__PURE__ */ jsx("span", { className: s$6.userName, children: name2 }),
            role && /* @__PURE__ */ jsx("span", { className: s$6.userRole, children: role })
          ] }),
          /* @__PURE__ */ jsx(DotsThree, { size: 24, className: s$6.userCaret })
        ]
      }
    );
  }
);
const wrapper$2 = "_wrapper_1n4be_3";
const overflowLeft = "_overflowLeft_1n4be_30";
const overflowRight = "_overflowRight_1n4be_34";
const tabList = "_tabList_1n4be_40";
const tab = "_tab_1n4be_40";
const disabled$2 = "_disabled_1n4be_90";
const active = "_active_1n4be_111";
const s$5 = {
  wrapper: wrapper$2,
  overflowLeft,
  overflowRight,
  tabList,
  tab,
  disabled: disabled$2,
  active
};
function getTabId(baseId, value2) {
  return `${baseId}-tab-${value2}`;
}
function getPanelId(baseId, value2) {
  return `${baseId}-panel-${value2}`;
}
function TabBar({
  tabs,
  activeTab,
  onTabChange,
  ariaLabel = "Abas",
  id,
  className
}) {
  const tabListRef = useRef(null);
  const autoId = useId();
  const baseId = id ?? autoId;
  const [canScrollLeft, setCanScrollLeft] = useState(false);
  const [canScrollRight, setCanScrollRight] = useState(false);
  const updateScrollHints = useCallback(() => {
    const el = tabListRef.current;
    if (!el) return;
    const tolerance = 2;
    setCanScrollLeft(el.scrollLeft > tolerance);
    setCanScrollRight(
      el.scrollLeft + el.clientWidth < el.scrollWidth - tolerance
    );
  }, []);
  useEffect(() => {
    const el = tabListRef.current;
    if (!el) return;
    updateScrollHints();
    el.addEventListener("scroll", updateScrollHints, { passive: true });
    const ro = new ResizeObserver(updateScrollHints);
    ro.observe(el);
    return () => {
      el.removeEventListener("scroll", updateScrollHints);
      ro.disconnect();
    };
  }, [updateScrollHints, tabs]);
  const getEnabledTabs = useCallback(() => {
    return tabs.filter((t) => !t.disabled);
  }, [tabs]);
  const handleKeyDown = useCallback(
    (e) => {
      const enabled = getEnabledTabs();
      if (enabled.length === 0) return;
      const currentIdx = enabled.findIndex((t) => t.value === activeTab);
      let nextIdx = null;
      switch (e.key) {
        case "ArrowRight":
          e.preventDefault();
          nextIdx = currentIdx < enabled.length - 1 ? currentIdx + 1 : 0;
          break;
        case "ArrowLeft":
          e.preventDefault();
          nextIdx = currentIdx > 0 ? currentIdx - 1 : enabled.length - 1;
          break;
        case "Home":
          e.preventDefault();
          nextIdx = 0;
          break;
        case "End":
          e.preventDefault();
          nextIdx = enabled.length - 1;
          break;
      }
      if (nextIdx !== null) {
        const nextTab = enabled[nextIdx];
        onTabChange(nextTab.value);
        const btn2 = tabListRef.current?.querySelector(
          `[data-tab-value="${nextTab.value}"]`
        );
        btn2?.focus();
      }
    },
    [activeTab, getEnabledTabs, onTabChange]
  );
  const wrapperCls = [
    s$5.wrapper,
    canScrollLeft ? s$5.overflowLeft : "",
    canScrollRight ? s$5.overflowRight : ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx("div", { className: wrapperCls, children: /* @__PURE__ */ jsx(
    "div",
    {
      ref: tabListRef,
      role: "tablist",
      "aria-label": ariaLabel,
      className: [s$5.tabList, className ?? ""].filter(Boolean).join(" "),
      onKeyDown: handleKeyDown,
      children: tabs.map((tab2) => {
        const isActive = tab2.value === activeTab;
        return /* @__PURE__ */ jsxs(
          "button",
          {
            type: "button",
            role: "tab",
            id: `${baseId}-tab-${tab2.value}`,
            "data-tab-value": tab2.value,
            "aria-selected": isActive,
            "aria-controls": `${baseId}-panel-${tab2.value}`,
            "aria-disabled": tab2.disabled || void 0,
            tabIndex: isActive ? 0 : -1,
            className: [
              s$5.tab,
              isActive ? s$5.active : "",
              tab2.disabled ? s$5.disabled : ""
            ].filter(Boolean).join(" "),
            onClick: () => {
              if (!tab2.disabled) onTabChange(tab2.value);
            },
            children: [
              tab2.label,
              tab2.badge
            ]
          },
          tab2.value
        );
      })
    }
  ) });
}
const root$1 = "_root_cu52d_3";
const prev = "_prev_cu52d_12";
const next = "_next_cu52d_18";
const buttonLabel = "_buttonLabel_cu52d_27";
const numbers = "_numbers_cu52d_33";
const number = "_number_cu52d_33";
const numberActive = "_numberActive_cu52d_67";
const ellipsis = "_ellipsis_cu52d_74";
const mobileLabel = "_mobileLabel_cu52d_88";
const s$4 = {
  root: root$1,
  prev,
  next,
  buttonLabel,
  numbers,
  number,
  numberActive,
  ellipsis,
  mobileLabel
};
function Pagination({
  currentPage,
  totalPages,
  onPageChange,
  className,
  ...rest
}) {
  const pages = useMemo(() => {
    const items2 = [];
    if (totalPages <= 7) {
      for (let i = 1; i <= totalPages; i++) items2.push(i);
    } else {
      items2.push(1);
      if (currentPage > 3) items2.push("ellipsis");
      const start = Math.max(2, currentPage - 1);
      const end = Math.min(totalPages - 1, currentPage + 1);
      for (let i = start; i <= end; i++) items2.push(i);
      if (currentPage < totalPages - 2) items2.push("ellipsis");
      items2.push(totalPages);
    }
    return items2;
  }, [currentPage, totalPages]);
  return /* @__PURE__ */ jsxs(
    "nav",
    {
      className: [s$4.root, className ?? ""].filter(Boolean).join(" "),
      "aria-label": "Paginação",
      ...rest,
      children: [
        /* @__PURE__ */ jsx("div", { className: s$4.prev, children: /* @__PURE__ */ jsx(
          Button,
          {
            variant: "secondary",
            size: "md",
            leftIcon: CaretLeft,
            disabled: currentPage === 1,
            onClick: () => onPageChange(currentPage - 1),
            "aria-label": "Página anterior",
            children: /* @__PURE__ */ jsx("span", { className: s$4.buttonLabel, children: "Anterior" })
          }
        ) }),
        /* @__PURE__ */ jsx("div", { className: s$4.numbers, children: pages.map(
          (item2, i) => item2 === "ellipsis" ? /* @__PURE__ */ jsx("span", { className: s$4.ellipsis, "aria-hidden": true, children: "..." }, `e${i}`) : /* @__PURE__ */ jsx(
            "button",
            {
              type: "button",
              className: [
                s$4.number,
                item2 === currentPage ? s$4.numberActive : ""
              ].filter(Boolean).join(" "),
              onClick: () => onPageChange(item2),
              "aria-label": `Página ${item2}`,
              "aria-current": item2 === currentPage ? "page" : void 0,
              children: item2
            },
            item2
          )
        ) }),
        /* @__PURE__ */ jsxs("div", { className: s$4.mobileLabel, children: [
          "Página ",
          currentPage,
          " de ",
          totalPages
        ] }),
        /* @__PURE__ */ jsx("div", { className: s$4.next, children: /* @__PURE__ */ jsx(
          Button,
          {
            variant: "secondary",
            size: "md",
            rightIcon: CaretRight,
            disabled: currentPage === totalPages,
            onClick: () => onPageChange(currentPage + 1),
            "aria-label": "Próxima página",
            children: /* @__PURE__ */ jsx("span", { className: s$4.buttonLabel, children: "Próximo" })
          }
        ) })
      ]
    }
  );
}
const root = "_root_1o7ep_3";
const elevated = "_elevated_1o7ep_10";
const borderless = "_borderless_1o7ep_14";
const cardHeader = "_cardHeader_1o7ep_21";
const cardHeaderContent = "_cardHeaderContent_1o7ep_26";
const cardHeaderTextBadge = "_cardHeaderTextBadge_1o7ep_34";
const cardHeaderTitle = "_cardHeaderTitle_1o7ep_40";
const cardHeaderActions = "_cardHeaderActions_1o7ep_49";
const cardHeaderDivider = "_cardHeaderDivider_1o7ep_55";
const bulkBar = "_bulkBar_1o7ep_62";
const bulkLeft = "_bulkLeft_1o7ep_90";
const bulkClose = "_bulkClose_1o7ep_96";
const bulkCount = "_bulkCount_1o7ep_120";
const bulkDivider = "_bulkDivider_1o7ep_128";
const bulkActions = "_bulkActions_1o7ep_135";
const scrollWrapper = "_scrollWrapper_1o7ep_143";
const table = "_table_1o7ep_155";
const th = "_th_1o7ep_165";
const thCheckbox = "_thCheckbox_1o7ep_180";
const sortButton = "_sortButton_1o7ep_187";
const sortIconIdle = "_sortIconIdle_1o7ep_207";
const row = "_row_1o7ep_213";
const td = "_td_1o7ep_218";
const striped = "_striped_1o7ep_223";
const selected = "_selected_1o7ep_228";
const tdCheckbox = "_tdCheckbox_1o7ep_251";
const pagination = "_pagination_1o7ep_258";
const s$3 = {
  root,
  elevated,
  borderless,
  cardHeader,
  cardHeaderContent,
  cardHeaderTextBadge,
  cardHeaderTitle,
  cardHeaderActions,
  cardHeaderDivider,
  bulkBar,
  bulkLeft,
  bulkClose,
  bulkCount,
  bulkDivider,
  bulkActions,
  scrollWrapper,
  table,
  th,
  thCheckbox,
  sortButton,
  sortIconIdle,
  row,
  td,
  striped,
  selected,
  tdCheckbox,
  pagination
};
const TableContext = createContext({
  variant: "divider",
  selectable: false,
  selectedRows: /* @__PURE__ */ new Set(),
  allSelected: false,
  someSelected: false
});
function Table({
  variant = "divider",
  elevated: elevated2 = true,
  bordered = true,
  selectable = false,
  selectedRows,
  rowIds = [],
  onSelectRow,
  onSelectAll,
  children,
  className,
  ...rest
}) {
  const selected2 = useMemo(() => selectedRows ?? /* @__PURE__ */ new Set(), [selectedRows]);
  const allSelected = rowIds.length > 0 && rowIds.every((id) => selected2.has(id));
  const someSelected = !allSelected && rowIds.some((id) => selected2.has(id));
  const ctx = useMemo(
    () => ({
      variant,
      selectable,
      selectedRows: selected2,
      onSelectRow,
      onSelectAll,
      allSelected,
      someSelected
    }),
    [variant, selectable, selected2, onSelectRow, onSelectAll, allSelected, someSelected]
  );
  return /* @__PURE__ */ jsx(TableContext.Provider, { value: ctx, children: /* @__PURE__ */ jsx(
    "div",
    {
      className: [s$3.root, elevated2 && s$3.elevated, !bordered && s$3.borderless, className].filter(Boolean).join(" "),
      ...rest,
      children
    }
  ) });
}
function TableCardHeader({ title: title2, badge: badge2, actions: actions2 }) {
  return /* @__PURE__ */ jsxs("div", { className: s$3.cardHeader, children: [
    /* @__PURE__ */ jsxs("div", { className: s$3.cardHeaderContent, children: [
      /* @__PURE__ */ jsxs("div", { className: s$3.cardHeaderTextBadge, children: [
        /* @__PURE__ */ jsx("h3", { className: s$3.cardHeaderTitle, children: title2 }),
        badge2
      ] }),
      actions2 && /* @__PURE__ */ jsx("div", { className: s$3.cardHeaderActions, children: actions2 })
    ] }),
    /* @__PURE__ */ jsx("div", { className: s$3.cardHeaderDivider })
  ] });
}
function TableBulkActions({ count, onClear, children }) {
  if (count === 0) return null;
  return createPortal(
    /* @__PURE__ */ jsxs("div", { className: s$3.bulkBar, role: "toolbar", "aria-label": "Ações em lote", children: [
      /* @__PURE__ */ jsxs("div", { className: s$3.bulkLeft, children: [
        /* @__PURE__ */ jsx(
          "button",
          {
            type: "button",
            className: s$3.bulkClose,
            onClick: onClear,
            "aria-label": "Desmarcar todos",
            children: /* @__PURE__ */ jsx(X, { size: 16 })
          }
        ),
        /* @__PURE__ */ jsxs("span", { className: s$3.bulkCount, "aria-live": "polite", role: "status", children: [
          count,
          " ",
          count === 1 ? "item selecionado" : "itens selecionados"
        ] })
      ] }),
      /* @__PURE__ */ jsx("span", { className: s$3.bulkDivider }),
      /* @__PURE__ */ jsx("div", { className: s$3.bulkActions, children })
    ] }),
    document.body
  );
}
function TableContent({
  children,
  className,
  ...rest
}) {
  return /* @__PURE__ */ jsx(
    "div",
    {
      className: [s$3.scrollWrapper, className ?? ""].filter(Boolean).join(" "),
      tabIndex: 0,
      role: "region",
      "aria-label": "Tabela com rolagem horizontal",
      ...rest,
      children: /* @__PURE__ */ jsx("table", { className: s$3.table, children })
    }
  );
}
function TableHead({
  children,
  ...rest
}) {
  return /* @__PURE__ */ jsx("thead", { ...rest, children });
}
function TableBody({
  children,
  ...rest
}) {
  return /* @__PURE__ */ jsx("tbody", { ...rest, children });
}
function TableRow({ rowId, children, className, ...rest }) {
  const { variant, selectable, selectedRows } = useContext(TableContext);
  const isSelected = rowId ? selectedRows.has(rowId) : false;
  const classes = [
    s$3.row,
    variant === "striped" ? s$3.striped : "",
    isSelected ? s$3.selected : "",
    className ?? ""
  ].filter(Boolean).join(" ");
  return /* @__PURE__ */ jsx(
    "tr",
    {
      className: classes,
      "aria-selected": selectable ? isSelected || void 0 : void 0,
      ...rest,
      children
    }
  );
}
function TableHeaderCell({
  sortable = false,
  sortDirection,
  onSort,
  isCheckbox = false,
  selectAllLabel,
  sortLabel,
  children,
  className,
  ...rest
}) {
  const { selectable, allSelected, someSelected, onSelectAll } = useContext(TableContext);
  if (isCheckbox && selectable) {
    return /* @__PURE__ */ jsx(
      "th",
      {
        className: [s$3.th, s$3.thCheckbox, className ?? ""].filter(Boolean).join(" "),
        scope: "col",
        ...rest,
        children: /* @__PURE__ */ jsx(
          Checkbox,
          {
            size: "md",
            checked: allSelected,
            indeterminate: someSelected,
            onChange: (e) => onSelectAll?.(e.currentTarget.checked),
            "aria-label": selectAllLabel ?? "Selecionar todas as linhas"
          }
        )
      }
    );
  }
  const columnLabel = sortLabel ?? (typeof children === "string" ? children : void 0);
  const sortAriaLabel = sortable ? sortDirection === "asc" ? columnLabel ? `Ordenado crescente por ${columnLabel}` : "Ordenado crescente" : sortDirection === "desc" ? columnLabel ? `Ordenado decrescente por ${columnLabel}` : "Ordenado decrescente" : columnLabel ? `Ordenar por ${columnLabel}` : "Clique para ordenar" : void 0;
  return /* @__PURE__ */ jsx(
    "th",
    {
      className: [s$3.th, className ?? ""].filter(Boolean).join(" "),
      scope: "col",
      "aria-sort": sortDirection === "asc" ? "ascending" : sortDirection === "desc" ? "descending" : sortable ? "none" : void 0,
      ...rest,
      children: sortable ? /* @__PURE__ */ jsxs(
        "button",
        {
          type: "button",
          className: s$3.sortButton,
          onClick: onSort,
          "aria-label": sortAriaLabel,
          children: [
            /* @__PURE__ */ jsx("span", { children }),
            sortDirection === "asc" && /* @__PURE__ */ jsx(ArrowUp, { size: 12 }),
            sortDirection === "desc" && /* @__PURE__ */ jsx(ArrowDown, { size: 12 }),
            !sortDirection && /* @__PURE__ */ jsx(ArrowDown, { size: 12, className: s$3.sortIconIdle })
          ]
        }
      ) : children
    }
  );
}
function TableCell({
  isCheckbox = false,
  rowId,
  selectionLabel,
  children,
  className,
  ...rest
}) {
  const { selectable, selectedRows, onSelectRow } = useContext(TableContext);
  if (isCheckbox && selectable && rowId) {
    return /* @__PURE__ */ jsx(
      "td",
      {
        className: [s$3.td, s$3.tdCheckbox, className ?? ""].filter(Boolean).join(" "),
        ...rest,
        children: /* @__PURE__ */ jsx(
          Checkbox,
          {
            size: "md",
            checked: selectedRows.has(rowId),
            onChange: (e) => onSelectRow?.(rowId, e.currentTarget.checked),
            "aria-label": selectionLabel ?? `Selecionar ${rowId}`
          }
        )
      }
    );
  }
  return /* @__PURE__ */ jsx(
    "td",
    {
      className: [s$3.td, className ?? ""].filter(Boolean).join(" "),
      ...rest,
      children
    }
  );
}
function TablePagination({
  currentPage,
  totalPages,
  onPageChange
}) {
  return /* @__PURE__ */ jsx(
    Pagination,
    {
      className: s$3.pagination,
      currentPage,
      totalPages,
      onPageChange
    }
  );
}
const wrapper$1 = "_wrapper_1ythi_2";
const label = "_label_1ythi_10";
const textareaBox = "_textareaBox_1ythi_19";
const disabled$1 = "_disabled_1ythi_29";
const hovered = "_hovered_1ythi_30";
const focused = "_focused_1ythi_35";
const error$1 = "_error_1ythi_41";
const textarea = "_textarea_1ythi_19";
const message = "_message_1ythi_94";
const attention = "_attention_1ythi_108";
const success$1 = "_success_1ythi_112";
const s$2 = {
  wrapper: wrapper$1,
  label,
  textareaBox,
  disabled: disabled$1,
  hovered,
  focused,
  error: error$1,
  textarea,
  message,
  attention,
  success: success$1
};
const messageIconMap = {
  error: WarningCircle,
  attention: WarningCircle,
  success: CheckCircle
};
const Textarea = forwardRef(
  ({
    label: label2,
    message: message2,
    messageType,
    disabled: disabled2 = false,
    className,
    rows = 4,
    ...rest
  }, ref) => {
    const autoId = useId();
    const textareaId = rest.id ?? autoId;
    const messageId = `${textareaId}-message`;
    const hasMessage = !!message2 && !!messageType;
    const isError = messageType === "error";
    const wrapperClasses = [s$2.wrapper, className ?? ""].filter(Boolean).join(" ");
    const boxClasses = [
      s$2.textareaBox,
      isError ? s$2.error : "",
      disabled2 ? s$2.disabled : ""
    ].filter(Boolean).join(" ");
    const MsgIcon = messageType ? messageIconMap[messageType] : null;
    return /* @__PURE__ */ jsxs("div", { className: wrapperClasses, children: [
      label2 && /* @__PURE__ */ jsx("label", { className: s$2.label, htmlFor: textareaId, children: label2 }),
      /* @__PURE__ */ jsx("div", { className: boxClasses, children: /* @__PURE__ */ jsx(
        "textarea",
        {
          ref,
          id: textareaId,
          className: s$2.textarea,
          disabled: disabled2,
          rows,
          "aria-invalid": isError || void 0,
          "aria-describedby": hasMessage ? messageId : void 0,
          ...rest
        }
      ) }),
      hasMessage && /* @__PURE__ */ jsxs("div", { id: messageId, className: `${s$2.message} ${s$2[messageType]}`, children: [
        MsgIcon && /* @__PURE__ */ jsx(MsgIcon, { size: 14 }),
        /* @__PURE__ */ jsx("span", { children: message2 })
      ] })
    ] });
  }
);
Textarea.displayName = "Textarea";
const toaster = "_toaster_1nq09_2";
const toast$1 = "_toast_1nq09_2";
const header = "_header_1nq09_27";
const headerIcon = "_headerIcon_1nq09_36";
const title$1 = "_title_1nq09_41";
const closeBtn = "_closeBtn_1nq09_51";
const body = "_body_1nq09_62";
const description$1 = "_description_1nq09_69";
const actionBtn = "_actionBtn_1nq09_79";
const success = "_success_1nq09_100";
const error = "_error_1nq09_106";
const warning = "_warning_1nq09_112";
const neutral = "_neutral_1nq09_118";
const black = "_black_1nq09_124";
const s$1 = {
  toaster,
  toast: toast$1,
  header,
  headerIcon,
  title: title$1,
  closeBtn,
  body,
  description: description$1,
  actionBtn,
  success,
  error,
  warning,
  neutral,
  black
};
let items = [];
const listeners = /* @__PURE__ */ new Set();
let counter = 0;
function emit() {
  items = [...items];
  listeners.forEach((fn) => fn());
}
function add(variant, title2, opts) {
  const id = String(++counter);
  items.unshift({
    id,
    variant,
    title: title2,
    description: opts?.description,
    duration: opts?.duration ?? 4e3,
    action: opts?.action
  });
  emit();
  return id;
}
function remove(id) {
  items = items.filter((t) => t.id !== id);
  emit();
}
function toast(title2, options) {
  return add("neutral", title2, options);
}
toast.success = (title2, options) => add("success", title2, options);
toast.error = (title2, options) => add("error", title2, options);
toast.warning = (title2, options) => add("warning", title2, options);
toast.black = (title2, options) => add("black", title2, options);
toast.dismiss = (id) => {
  if (id) remove(id);
  else {
    items = [];
    emit();
  }
};
const VISIBLE = 3;
const STACK_GAP = 8;
const EXPAND_GAP = 12;
const ANIM_MS = 400;
const variantIcon = {
  success: CheckCircle,
  error: WarningCircle,
  warning: WarningCircle,
  neutral: Info,
  black: CheckCircle
};
function Toaster() {
  const toasts = useSyncExternalStore(
    useCallback((cb) => {
      listeners.add(cb);
      return () => {
        listeners.delete(cb);
      };
    }, []),
    () => items,
    () => items
  );
  const [mounted, setMounted] = useState(() => /* @__PURE__ */ new Set());
  const [removing, setRemoving] = useState(() => /* @__PURE__ */ new Set());
  const [expanded, setExpanded] = useState(false);
  const heights = useRef(/* @__PURE__ */ new Map());
  const timers = useRef(/* @__PURE__ */ new Map());
  useEffect(
    () => () => {
      timers.current.forEach((t) => clearTimeout(t));
    },
    []
  );
  useEffect(() => {
    const newIds = toasts.slice(0, VISIBLE).map((t) => t.id).filter((id) => !mounted.has(id));
    if (newIds.length === 0) return;
    const raf = requestAnimationFrame(() => {
      setMounted((prev2) => {
        const next2 = new Set(prev2);
        newIds.forEach((id) => next2.add(id));
        return next2;
      });
    });
    return () => cancelAnimationFrame(raf);
  }, [toasts, mounted]);
  const dismiss2 = useCallback((id) => {
    timers.current.delete(id);
    setRemoving((prev2) => new Set(prev2).add(id));
    setTimeout(() => {
      remove(id);
      setRemoving((prev2) => {
        const next2 = new Set(prev2);
        next2.delete(id);
        return next2;
      });
      setMounted((prev2) => {
        const next2 = new Set(prev2);
        next2.delete(id);
        return next2;
      });
      heights.current.delete(id);
    }, ANIM_MS);
  }, []);
  useEffect(() => {
    if (expanded) {
      timers.current.forEach((t) => clearTimeout(t));
      timers.current.clear();
      return;
    }
    toasts.slice(0, VISIBLE).forEach((item2) => {
      if (removing.has(item2.id) || timers.current.has(item2.id)) return;
      if (item2.duration === Infinity) return;
      const timer = setTimeout(() => {
        timers.current.delete(item2.id);
        dismiss2(item2.id);
      }, item2.duration);
      timers.current.set(item2.id, timer);
    });
  }, [toasts, expanded, removing, dismiss2]);
  const visible = toasts.slice(0, VISIBLE);
  if (visible.length === 0 && removing.size === 0) return null;
  const expandedOffsets = [];
  let offset = 0;
  for (let i = 0; i < visible.length; i++) {
    expandedOffsets.push(offset);
    offset += (heights.current.get(visible[i].id) ?? 0) + EXPAND_GAP;
  }
  return /* @__PURE__ */ jsx(
    "section",
    {
      className: s$1.toaster,
      "aria-label": "Notificações",
      onMouseEnter: () => setExpanded(true),
      onMouseLeave: () => setExpanded(false),
      children: visible.map((item2, index) => {
        const isMounted = mounted.has(item2.id);
        const isRemoving = removing.has(item2.id);
        let transform;
        let opacity;
        if (!isMounted || isRemoving) {
          transform = "translateY(100%) scale(1)";
          opacity = 0;
        } else if (expanded) {
          transform = `translateY(-${expandedOffsets[index]}px) scale(1)`;
          opacity = 1;
        } else {
          transform = `translateY(${index * STACK_GAP}px) scale(${1 - index * 0.05})`;
          opacity = 1 - index * 0.15;
        }
        const style = {
          transform,
          opacity,
          zIndex: VISIBLE - index,
          transition: `all ${ANIM_MS}ms ease`,
          transformOrigin: "bottom center"
        };
        return /* @__PURE__ */ jsx(
          ToastCard,
          {
            data: item2,
            style,
            onDismiss: () => dismiss2(item2.id),
            onHeight: (h) => heights.current.set(item2.id, h)
          },
          item2.id
        );
      })
    }
  );
}
function ToastCard({ data, style, onDismiss, onHeight }) {
  const ref = useRef(null);
  const Icon = variantIcon[data.variant];
  useEffect(() => {
    if (ref.current) onHeight(ref.current.offsetHeight);
  });
  const handleAction = () => {
    data.action?.onClick();
    onDismiss();
  };
  return /* @__PURE__ */ jsxs(
    "div",
    {
      ref,
      className: `${s$1.toast} ${s$1[data.variant]}`,
      style,
      role: "status",
      "aria-live": "polite",
      children: [
        /* @__PURE__ */ jsxs("div", { className: s$1.header, children: [
          /* @__PURE__ */ jsx("span", { className: s$1.headerIcon, children: /* @__PURE__ */ jsx(Icon, { size: 20, "aria-hidden": true }) }),
          /* @__PURE__ */ jsx("p", { className: s$1.title, children: data.title }),
          /* @__PURE__ */ jsx(
            Button,
            {
              variant: "tertiary",
              size: "sm",
              leftIcon: X,
              onClick: onDismiss,
              "aria-label": "Fechar notificação",
              className: s$1.closeBtn
            }
          )
        ] }),
        (data.description || data.action) && /* @__PURE__ */ jsxs("div", { className: s$1.body, children: [
          data.description && /* @__PURE__ */ jsx("p", { className: s$1.description, children: data.description }),
          data.action && /* @__PURE__ */ jsx("button", { type: "button", className: s$1.actionBtn, onClick: handleAction, children: data.action.label })
        ] })
      ]
    }
  );
}
const wrapper = "_wrapper_w3vu9_2";
const hasLabel = "_hasLabel_w3vu9_10";
const disabled = "_disabled_w3vu9_14";
const input = "_input_w3vu9_19";
const track = "_track_w3vu9_31";
const thumb = "_thumb_w3vu9_44";
const text = "_text_w3vu9_110";
const title = "_title_w3vu9_120";
const description = "_description_w3vu9_127";
const s = {
  wrapper,
  hasLabel,
  disabled,
  input,
  track,
  thumb,
  text,
  title,
  description
};
const Toggle = forwardRef(
  ({
    label: label2,
    description: description2,
    disabled: disabled2 = false,
    className,
    ...rest
  }, ref) => {
    const wrapperClasses = [
      s.wrapper,
      label2 ? s.hasLabel : "",
      disabled2 ? s.disabled : "",
      className ?? ""
    ].filter(Boolean).join(" ");
    return /* @__PURE__ */ jsxs("label", { className: wrapperClasses, children: [
      /* @__PURE__ */ jsx(
        "input",
        {
          ref,
          type: "checkbox",
          role: "switch",
          className: s.input,
          disabled: disabled2,
          ...rest
        }
      ),
      /* @__PURE__ */ jsx("span", { className: s.track, children: /* @__PURE__ */ jsx("span", { className: s.thumb }) }),
      label2 && /* @__PURE__ */ jsxs("span", { className: s.text, children: [
        /* @__PURE__ */ jsx("span", { className: s.title, children: label2 }),
        description2 && /* @__PURE__ */ jsx("span", { className: s.description, children: description2 })
      ] })
    ] });
  }
);
Toggle.displayName = "Toggle";
function useDataTable() {
  const [selectedRows, setSelectedRows] = useState(/* @__PURE__ */ new Set());
  const [sortKey, setSortKey] = useState(null);
  const [sortDir, setSortDir] = useState("asc");
  function handleSort(key) {
    if (sortKey === key) {
      setSortDir((prev2) => prev2 === "asc" ? "desc" : "asc");
      return;
    }
    setSortKey(key);
    setSortDir("asc");
  }
  function getSortDirection(key) {
    return sortKey === key ? sortDir : void 0;
  }
  function handleSelectRow(id, checked) {
    setSelectedRows((prev2) => {
      const next2 = new Set(prev2);
      if (checked) next2.add(id);
      else next2.delete(id);
      return next2;
    });
  }
  function handleSelectAll(checked, rowIds) {
    setSelectedRows(checked ? new Set(rowIds) : /* @__PURE__ */ new Set());
  }
  function clearSelection() {
    setSelectedRows(/* @__PURE__ */ new Set());
  }
  return {
    selectedRows,
    setSelectedRows,
    clearSelection,
    sortKey,
    sortDir,
    handleSort,
    getSortDirection,
    handleSelectRow,
    handleSelectAll
  };
}
function useFilterChips({ chipRefs, onResetFilter } = {}) {
  const [activeFilters, setActiveFilters] = useState([]);
  const [openFilter, setOpenFilter] = useState(null);
  const ignoreChipRefs = useMemo(() => {
    if (!chipRefs) return void 0;
    return Object.values(chipRefs);
  }, [chipRefs]);
  function addFilter(id) {
    setActiveFilters((prev2) => prev2.includes(id) ? prev2 : [...prev2, id]);
  }
  function addFilterAndOpen(id) {
    addFilter(id);
    requestAnimationFrame(() => {
      requestAnimationFrame(() => {
        setOpenFilter(id);
      });
    });
  }
  function removeFilter(id) {
    setActiveFilters((prev2) => prev2.filter((filterId) => filterId !== id));
    if (openFilter === id) setOpenFilter(null);
    onResetFilter?.(id);
  }
  function clearAllFilters() {
    setActiveFilters((prev2) => {
      prev2.forEach((id) => onResetFilter?.(id));
      return [];
    });
    setOpenFilter(null);
  }
  function toggleFilterDropdown(id) {
    setOpenFilter((prev2) => prev2 === id ? null : id);
  }
  function getAvailableFilters(allFilters) {
    return allFilters.filter((filter) => !activeFilters.includes(filter.id));
  }
  return {
    activeFilters,
    setActiveFilters,
    openFilter,
    setOpenFilter,
    addFilter,
    addFilterAndOpen,
    removeFilter,
    clearAllFilters,
    toggleFilterDropdown,
    getAvailableFilters,
    ignoreChipRefs
  };
}
export {
  Accordion,
  AccordionItem,
  AiAssistant,
  Alert,
  AssistantButton,
  Avatar,
  AvatarGroup,
  AvatarLabelGroup,
  Badge,
  Breadcrumb,
  Button,
  Card,
  CardBody,
  CardDivider,
  CardFooter,
  CardHeader,
  Chart,
  ChartTooltipContent,
  Checkbox,
  ChoiceBox,
  ChoiceBoxGroup,
  CommandPalette,
  DatePicker,
  DragToCloseDrawer,
  Drawer,
  DrawerBody,
  DrawerFooter,
  DrawerHeader,
  DropdownButton,
  FilterBar,
  FilterChip,
  FilterDropdown,
  Funnel,
  GoalGaugeBar,
  GoalProgressBar,
  Heatmap,
  Input,
  LoadingScreen,
  MONTH_LABELS,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  NotificationButton,
  NotificationPanel,
  PageHeader,
  Popover,
  PopoverSelect,
  Radar,
  Radio,
  RowActionsPopover,
  SKELETON_HEIGHTS,
  ScaleInput,
  SearchButton,
  Select,
  Sidebar,
  SidebarDivider,
  SidebarFooter,
  SidebarGroup,
  SidebarHeader,
  SidebarItem,
  SidebarNav,
  SidebarOrgSwitcher,
  SidebarSubItem,
  SidebarUser,
  Skeleton,
  SkeletonContainer,
  SortableList,
  Sparkline,
  TabBar,
  Table,
  TableBody,
  TableBulkActions,
  TableCardHeader,
  TableCell,
  TableContent,
  TableHead,
  TableHeaderCell,
  TablePagination,
  TableRow,
  Textarea,
  Toaster,
  Toggle,
  Tooltip,
  WEEKDAY_LABELS,
  compareDates,
  daysInMonth,
  firstDayOfWeek,
  formatDate,
  formatMultiLabel,
  getPanelId,
  getTabId,
  isDisabled,
  isInRange,
  isSameDay,
  isToday,
  isValidDate,
  nextMonth,
  parseDate,
  prevMonth,
  toast,
  today,
  useDataTable,
  useFilterChips
};
//# sourceMappingURL=index.js.map
