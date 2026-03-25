export interface FunnelStep {
    /** Step label */
    label: string;
    /** Numeric value for this step */
    value: number;
    /** Optional color override (CSS variable or valid CSS color) */
    color?: string;
}
export interface FunnelProps {
    /** Array of funnel steps from top to bottom */
    data: FunnelStep[];
    /** Total height in px (default: 300) */
    height?: number;
    /** Show numeric values (default: true) */
    showValues?: boolean;
    /** Show percentage relative to first step (default: true) */
    showPercentage?: boolean;
    /** Custom value formatter */
    formatValue?: (value: number) => string;
    /** Additional CSS class */
    className?: string;
}
export declare function Funnel({ data, height, showValues, showPercentage, formatValue, className, }: FunnelProps): import("react/jsx-runtime").JSX.Element | null;
