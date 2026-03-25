import { ReactNode } from 'react';
import { DrawerProps } from './Drawer';
interface DragToCloseDrawerProps extends DrawerProps {
    children: ReactNode;
    /** Enable drag-to-close gesture (default: true) */
    dragToCloseEnabled?: boolean;
    /** Minimum distance (px) to trigger close (default: 80) */
    dragThreshold?: number;
    /** Minimum velocity (px/ms) to trigger close (default: 0.5) */
    velocityThreshold?: number;
    /** Height of draggable zone at top of drawer in px (default: 60) */
    dragZoneHeight?: number;
}
/**
 * Wrapper do DS Drawer que adiciona funcionalidade drag-to-close no mobile.
 *
 * Detecta gestos de arraste vertical para baixo no topo do drawer e fecha quando:
 * - O arraste ultrapassa o `dragThreshold`, OU
 * - A velocidade do arraste ultrapassa o `velocityThreshold`
 *
 * @example
 * ```tsx
 * <DragToCloseDrawer
 *   open={open}
 *   onClose={() => setOpen(false)}
 *   dragThreshold={100}
 *   dragZoneHeight={80}
 * >
 *   <DrawerHeader title="Detalhes" onClose={() => setOpen(false)} />
 *   <DrawerBody>...</DrawerBody>
 * </DragToCloseDrawer>
 * ```
 */
export declare function DragToCloseDrawer({ children, open, onClose, dragToCloseEnabled, dragThreshold, velocityThreshold, dragZoneHeight, ...drawerProps }: DragToCloseDrawerProps): import("react/jsx-runtime").JSX.Element;
export default DragToCloseDrawer;
