export interface RadarDataPoint {
    /** Axis label */
    label: string;
    /** Numeric value for this axis */
    value: number;
}
export interface RadarProps {
    /** Data points — one per axis */
    data: RadarDataPoint[];
    /** Maximum value for the scale (default: auto from data) */
    maxValue?: number;
    /** Diameter of the chart in px (default: 200) */
    size?: number;
    /** Color theme (default: "orange") */
    color?: "orange" | "green" | "red" | "wine" | "neutral";
    /** Show numeric values next to each axis (default: false) */
    showValues?: boolean;
    /** Number of concentric grid rings (default: 4) */
    levels?: number;
    className?: string;
}
export declare function Radar({ data, maxValue, size, color, showValues, levels, className, }: RadarProps): import("react/jsx-runtime").JSX.Element | null;
